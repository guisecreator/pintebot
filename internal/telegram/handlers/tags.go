package handlers

import (
	"bytes"
	"github.com/carrot/go-pinterest"
	"github.com/guisecreator/pintebot/internal/config"
	"github.com/guisecreator/pintebot/internal/pinterest/pinterest_api"
	"github.com/guisecreator/pintebot/internal/telegram/types"
	"github.com/h2non/bimg"
	"github.com/mymmrac/telego"
	th "github.com/mymmrac/telego/telegohandler"
	tu "github.com/mymmrac/telego/telegoutil"
	"github.com/sirupsen/logrus"
	"io"
	"log"
)

type TagsCommand struct {
	*types.CommandsOptions
	logger *logrus.Logger
}

type NamedReaderImpl struct {
	io.Reader
	fileName string
}

func (n NamedReaderImpl) Name() string {
	return n.fileName
}

const (
	widthTgPhoto  = 512
	heightTgPhoto = 512
)

func (tags *TagsCommand) SendUserMessageToPinterest() *telego.SendMessageParams {
	panic("implement me")
}

func (tags *TagsCommand) SetImagesFromPinterest() ([]string, error) {
	return nil, nil
}

// the message is processed, the line goes to Pinterest search
func HandleUserMessage(update *telego.Update, client pinterest.Client) {
	userMessage := update.Message.Text

	_, err := pinterest_api.GetPinsBySearch(client, userMessage)
	if err != nil {
		log.Fatalf("get pins by search error: %v\n", err)
	}
}

func (tags *TagsCommand) formatImageForTg(images []string) (*NamedReaderImpl, error) {
	var (
		processedBuffer bytes.Buffer
		formattingErr   error
	)

	for _, photo := range images {
		buf, err := bimg.Read(photo)
		if err != nil {
			tags.logger.Fatalf("read photo error: %v\n", formattingErr)
			return nil, err
		}

		newImage, err := bimg.NewImage(buf).
			Process(bimg.Options{
				Width:   widthTgPhoto,
				Height:  heightTgPhoto,
				Crop:    true,
				Quality: 95,
			})
		if err != nil {
			tags.logger.Fatalf("process photo error: %v\n", formattingErr)
			return nil, err
		}

		_, err = processedBuffer.Write(newImage)
		if err != nil {
			tags.logger.Fatalf("write processed photo to buffer error: %v\n", formattingErr)
			return nil, err
		}
	}

	processed := processedBuffer.Bytes()
	if processed == nil {
		return nil, formattingErr
	}

	//mb remove newReader?
	newReader := bytes.NewReader(processed)

	return &NamedReaderImpl{
		Reader:   newReader,
		fileName: string(processed),
	}, nil
}

func (tags *TagsCommand) SendImageUser(id telego.ChatID) *telego.SendPhotoParams {
	images, err := tags.SetImagesFromPinterest()
	if err != nil {
		log.Fatal(err)
	}

	formattedImages, err := tags.formatImageForTg(images)
	if err != nil {
		log.Fatal(err)
	}

	receivedFile := tu.File(formattedImages)
	if receivedFile.String() == "" {
		tags.logger.Error(err)
		return nil
	}

	return &telego.SendPhotoParams{
		ChatID: id,
		Photo:  receivedFile,
	}
}

func (tags *TagsCommand) NewTagsCommand() th.Handler {
	return func(bot *telego.Bot, update telego.Update) {
		messages, err := config.InitCommandsText("locales/en.yaml")
		if err != nil {
			log.Fatal(err)
		}

		userId := tu.ID(update.CallbackQuery.From.ID)
		if update.Message != nil && update.Message.Text != "" {
			inlineKeyboard := tu.Keyboard(
				tu.KeyboardRow(
					tu.KeyboardButton(
						messages.TagsCommand.InlineKeyboard.
							KeyboardRow1.NextButton,
					),
				),
			).WithResizeKeyboard()

			buttonMessage := tu.Message(
				userId,
				"",
			).WithReplyMarkup(inlineKeyboard)
			_, buttonErr := bot.SendMessage(buttonMessage)
			if buttonErr != nil {
				tags.logger.Errorf("send button error: %v\n", buttonErr)
			}

			//add a cancel command button to the photo that was sent to the user
			_, sendPhotoErr := bot.SendPhoto(tags.SendImageUser(userId))
			if sendPhotoErr != nil {
				tags.logger.Errorf("send photo error: %v\n", sendPhotoErr)
				return
			}

		} else {
			messageText := messages.AnyTagText
			message := tu.Message(
				userId,
				messageText,
			).WithParseMode(telego.ModeHTML)

			_, botErr := bot.SendMessage(message.WithProtectContent())
			if botErr != nil {
				log.Printf("send message error: %v\n", botErr)
			}
		}
	}
}
