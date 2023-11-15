package handlers

import (
	"github.com/carrot/go-pinterest"
	"github.com/guisecreator/pintebot/internal/config"
	"github.com/guisecreator/pintebot/internal/telegram/types"
	"github.com/mymmrac/telego"
	th "github.com/mymmrac/telego/telegohandler"
	tu "github.com/mymmrac/telego/telegoutil"
	"github.com/sirupsen/logrus"
	"log"
)

type TagsCommand struct {
	*types.CommandsOptions
	logger *logrus.Logger
}

func (tags *TagsCommand) GetImageList() ([]string, error) {
	panic("implement me")
}

// юзер что либо вводит, это обрабатывается. сообщение от юзера уходит в поиск пинов
func (tags *TagsCommand) handleUserMessage(update *telego.Update, message string) error {
	if message == "" {
		return nil
	}

	userId := update.Message.From.ID
	if userId == 0 {
		return nil
	}

	return nil
}

// здесь сообщения должны читаться, обрабатываться и отправляться юзеру
func (tags *TagsCommand) handleImage(
	id telego.ChatID,
	photo telego.InputFile,
) *telego.SendPhotoParams {

	////imageList, err := tags.GetImageList()
	////if err != nil {
	////	return nil
	////}
	//
	//files := []string{
	//	"media/test.jpg",
	//	"media/test2.jpg",
	//	"media/test3.jpg",
	//	"media/test4.jpg",
	//}
	//
	//var photos []telego.InputFile
	//
	//for _, file := range files {
	//	if file == "" {
	//		return nil
	//	}
	//
	//	uncovered, err := os.Open(file)
	//	if err != nil {
	//		tags.logger.Errorf("Error opening file: %v\n", err)
	//		return nil
	//	}
	//	defer uncovered.Close()
	//
	//	formattedImage, err := img_processing.FormatImages(file)
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//
	//	photos = append(photos, telego.InputFile{
	//		File: formattedImage.GetReader(),
	//	})
	//
	//}
	//
	//photo = telego.InputFile{
	//	File: photos,
	//}
	//
	//return &telego.SendPhotoParams{
	//	ChatID: id,
	//	Photo:  photo,
	//}
}

func (tags *TagsCommand) handleDownloadImage(update *telego.Update, client pinterest.Client) {
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

		//pinsName, isPinsName := GetUserMessage(update)
		//if !isPinsName {
		//	tags.logger.Printf("get pins name error: %v\n", pinsName)
		//	return
		//}
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
				userId, telego.InputFile{},
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
