package handlers

import (
	"github.com/guisecreator/pintebot/internal/config"
	"github.com/guisecreator/pintebot/internal/telegram/types"
	"github.com/mymmrac/telego"
	th "github.com/mymmrac/telego/telegohandler"
	tu "github.com/mymmrac/telego/telegoutil"
	"github.com/sirupsen/logrus"
	"log"
	"os"
)

type TagsCommand struct {
	*types.CommandsOptions
	logger *logrus.Logger
}

// берет картинки по запросу июзера и складывает их в папку изображений
func (tags *TagsCommand) GetImageList(update telego.Update) ([]string, error) {
	messageRequest, err := tags.handleUserMessage(update)
	if err != nil {
		return nil, err
	}

	if messageRequest == "" {
		return nil, nil
	}

	//fix:
	pins, err := tags.Services.PinterestAPI.GetPinsBySearch(messageRequest)
	if err != nil {
		return nil, err
	}

	var imageList []string

	for _, pin := range *pins {
		if pin.Id == "" {
			tags.logger.Errorf("Pin %s has no image\n", pin.Id)
			continue
		}

		imageList = append(imageList, pin.Image.Original.Url)
	}

	return imageList, nil
}

// юзер что либо вводит, это обрабатывается. сообщение от юзера уходит в поиск пинов
func (tags *TagsCommand) handleUserMessage(update telego.Update) (string, error) {

	return "", nil
}

// здесь сообщения должны читаться, обрабатываться и отправляться юзеру
func (tags *TagsCommand) handleImage(
	id telego.ChatID,
	finishedImage telego.InputFile,
	update telego.Update,
) *telego.SendPhotoParams {

	imageList, err := tags.GetImageList(update)
	if err != nil {
		return nil
	}

	var images []telego.InputFile

	for _, image := range imageList {
		if image == "" {
			return nil
		}

		uncovered, expansionErr := os.Open(image)
		if expansionErr != nil {
			tags.logger.Errorf("Error opening image: %v\n", expansionErr)
			return nil
		}
		defer uncovered.Close()

		//formattedImage, formattingErr := img_processing.FormatImage(image)
		//if formattingErr != nil {
		//	log.Fatal(formattingErr)
		//}

		images = append(images, telego.InputFile{
			File: uncovered,
		})

	}

	finishedImage = telego.InputFile{
		//fix: File: images,
		File: images[0].File,
	}

	return &telego.SendPhotoParams{
		ChatID: id,
		Photo:  finishedImage,
	}
}

func (tags *TagsCommand) handleDownloadImage(update *telego.Update) {
	panic("implement me")
}

// TODO: implement me
func (tags *TagsCommand) NextImageQuery() {
	panic("implement me")
}

func (tags *TagsCommand) MessageTag() th.Handler {
	return func(bot *telego.Bot, update telego.Update) {
		messages, err := config.InitCommandsText("locales/en.yaml")
		if err != nil {
			log.Fatal(err)
		}

		userId := tu.ID(update.CallbackQuery.From.ID)

		messageText := messages.AnyTagText
		message := tu.Message(
			userId,
			messageText,
		).WithParseMode(telego.ModeHTML)

		_, botErr := bot.SendMessage(message)
		if botErr != nil {
			log.Printf("send message error: %v\n", botErr)
		}

	}
}

func (tags *TagsCommand) NewTagsCommand() th.Handler {
	return func(bot *telego.Bot, update telego.Update) {
		messages, err := config.InitCommandsText("locales/en.yaml")
		if err != nil {
			log.Fatal(err)
		}

		userId := tu.ID(update.Message.From.ID)

		pinsName := update.Message.Text
		if pinsName == "" {
			_, msgerr := bot.SendMessage(
				MessageError(userId, 1, "Message Error", true),
			)
			if msgerr != nil {
				tags.logger.Errorf("send message to %v user: %v", userId, msgerr)
			}
			return
		}

		//TODO: сделать обработку команды(и кнопки соответсвенно) Next, при нажатии на которую бот будет переходить на следующий пин
		if pinsName != "" {
			buttonMessage := tu.Message(
				userId,
				messages.SuccessfulSearchByTags+update.Message.Text,
			)
			_, buttonErr := bot.SendMessage(buttonMessage)
			if buttonErr != nil {
				tags.logger.Errorf("send button error: %v\n", buttonErr)
			}

			sendPhotoParams := tags.handleImage(
				userId,
				telego.InputFile{},
				update,
			).WithReplyMarkup(
				tu.InlineKeyboard(
					tu.InlineKeyboardRow(
						tu.InlineKeyboardButton("Download this image").
							WithCallbackData("download"),
					),
					tu.InlineKeyboardRow(
						tu.InlineKeyboardButton("Cancel").
							WithCallbackData("cancel"),
					),
				),
			).WithReplyMarkup(tu.Keyboard(
				tu.KeyboardRow(
					tu.KeyboardButton(
						messages.TagsCommand.InlineKeyboard.
							KeyboardRow1.NextButton,
					),
				),
			).WithResizeKeyboard())

			_, sendPhotoErr := bot.SendPhoto(sendPhotoParams)
			if sendPhotoErr != nil {
				tags.logger.Errorf("send photo error: %v\n", sendPhotoErr)
				return
			}

			//bot.FileDownloadURL(sendPhotoParams.Photo.FileID)
		}
	}
}
