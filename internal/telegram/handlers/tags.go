package handlers

import (
	"fmt"
	"github.com/guisecreator/pintebot/internal/config"
	"github.com/guisecreator/pintebot/internal/state"
	"github.com/guisecreator/pintebot/internal/telegram/types"
	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
	"github.com/sirupsen/logrus"
	"io"
	"log"
	"net/http"
	"os"
)

type UserImageStore struct {
	ImageLists     map[int64][]string
	CurrentIndices map[int64]int
	UserStates     map[int64]state.UserState
}

type TagsCommand struct {
	*types.CommandsOptions
	logger         *logrus.Logger
	UserImageStore *UserImageStore
}

func (tags *TagsCommand) GetImageList(
	update telego.Update,
) ([]string, error) {
	chatID := update.Message.Chat.ID

	messageRequest, err := tags.handleUserMessage(update)
	if err != nil {
		return nil, err
	}

	if tags.UserImageStore == nil {
		tags.UserImageStore = &UserImageStore{
			ImageLists:     make(map[int64][]string),
			CurrentIndices: make(map[int64]int),
		}
	}

	pins, err := tags.Services.
		PinterestAPI.
		GetPinsBySearch(messageRequest)
	if err != nil {
		return nil, err
	}

	userImageList, exists := tags.UserImageStore.ImageLists[chatID]
	if !exists {
		tags.logger.Errorf("Empty image list for chatID: %d", chatID)
		userImageList = make([]string, 0)
	}

	for _, pin := range *pins {
		if pin.Id == "" && pin.Url == "" {
			log.Fatal("Empty id or url")
			return nil, err
		}

		imageData, imgErr := http.NewRequest("GET", pin.Url, nil)
		if err != nil {
			return nil, imgErr
		}
		defer imageData.Body.Close()

		imageName := fmt.Sprintf("media/%s/%s.jpg", messageRequest, pin.Note)
		imageFile, createErr := os.Create(imageName)
		if createErr != nil {
			return nil, createErr
		}
		defer imageFile.Close()

		_, copyErr := io.Copy(imageFile, imageData.Body)
		if copyErr != nil {
			return nil, copyErr
		}

		userImageList = append(userImageList, imageName)
	}

	tags.UserImageStore.ImageLists[chatID] = userImageList

	return userImageList, nil
}

func (tags *TagsCommand) handleUserMessage(
	update telego.Update,
) (string, error) {
	var messageError = fmt.Sprintf(
		"Empty message or wrong message: %s",
		update.Message.Text,
	)

	user := update.Message.Text
	if user == "" {
		return messageError, nil
	}

	if update.Message.ReplyToMessage != nil {
		user = update.Message.ReplyToMessage.Text
		if user == "" {
			return messageError, nil
		}
	}

	tags.UserImageStore.CurrentIndices[update.Message.Chat.ID] = 0

	return user, nil
}

// func (tags *TagsCommand) handleNextImageQuery(
// 	chatID telego.ChatID,
// 	update telego.Update,
// ) *telego.SendPhotoParams {
// 	if chatID.ID == 0 {
// 		tags.logger.Error("Empty chatID")
// 		return nil
// 	}

// 	imageList, exist := tags.UserImageStore.ImageLists[chatID.ID]
// 	if !exist {
// 		tags.logger.Error("Empty image list")
// 		return nil
// 	}

// 	currentIndex, exists := tags.UserImageStore.CurrentIndices[chatID.ID]
// 	if !exists {
// 		currentIndex = 0
// 	}

// 	// increment the index or reset if the end of the list is reached
// 	currentIndex = (currentIndex + 1) % len(imageList)
// 	tags.UserImageStore.CurrentIndices[chatID.ID] = currentIndex

// 	imageData := []byte(imageList[currentIndex])
// 	reader := bytes.NewReader(imageData)

// 	photo := telego.InputFile{
// 		File: tu.NameReader(reader, "image"),
// 	}

// 	return &telego.SendPhotoParams{
// 		ChatID:              tu.ID(chatID.ID),
// 		Photo:               photo,
// 		DisableNotification: false,
// 		ParseMode:           telego.ModeHTML,
// 		Caption:             "",
// 	}
// }

func (tags *TagsCommand) handleNextImageQuery(
	chatID telego.ChatID,
	update telego.Update,
) *telego.SendPhotoParams {
	filePath := "media/test.jpg"

	file, err := os.Open(filePath)
	if err != nil {
		tags.logger.Errorf("Ошибка при открытии файла: %v\n", err)
	}

	photo := telego.InputFile{
		File: file,
	}

	return &telego.SendPhotoParams{
		ChatID: chatID,
		Photo:  photo,
	}
}

func (tags *TagsCommand) MessageTag(bot *telego.Bot, update telego.Update) {
	userId := tu.ID(update.Message.Chat.ID)

	messages, err := config.InitCommandsText("locales/en.yaml")
	if err != nil {
		log.Fatal(err)
	}

	// keyboard := tu.Keyboard(
	// 	tu.KeyboardRow(
	// 		tu.KeyboardButton(
	// 			"find_pins",
	// 		),
	// 	),
	// ).WithResizeKeyboard().
	// WithInputFieldPlaceholder("Enter pin name")

	_, botErr := bot.SendMessage(
		tu.Message(userId, messages.AnyTagText).
			WithReplyMarkup(tu.ForceReply()),
	)
	if botErr != nil {
		log.Fatalf("send message error: %v\n", botErr)
	}
}

func (tags *TagsCommand) NewTagsCommand(bot *telego.Bot, update telego.Update) {
	userId := tu.ID(update.Message.From.ID)

	messages, err := config.InitCommandsText("locales/en.yaml")
	if err != nil {
		log.Fatal(err)
	}

	pinsName := update.Message.Text
	if pinsName == "" {
		_, msgErr := bot.SendMessage(
			MessageError(userId, 1, "Message Error", true),
		)
		if msgErr != nil {
			tags.logger.Errorf("send message to %v user: %v", userId, msgErr)
		}
		return
	}

	if update.Message.From.ID == 0 {
			return
	}

	//remove it
	if tags.UserImageStore == nil {
		tags.UserImageStore = &UserImageStore{
			ImageLists:     make(map[int64][]string),
			CurrentIndices: make(map[int64]int),
			UserStates:     make(map[int64]state.UserState),
		}
	}
	if tags.UserImageStore.UserStates == nil {
		tags.UserImageStore.UserStates = make(map[int64]state.UserState)
	}

	// get state of user
	userState := state.GetUserState(
		userId.ID,
		tags.UserImageStore.UserStates,
	)

	switch userState {
	case state.StateInitial:
		buttonMessage := tu.Message(
			userId,
			messages.SuccessfulSearchByTags+update.Message.Text,
		).WithReplyMarkup(tu.ForceReply())
		_, buttonErr := bot.SendMessage(buttonMessage)
		if buttonErr != nil {
			tags.logger.Errorf("send button error: %v\n", buttonErr)
		}

		sendPhotoParams := tags.handleNextImageQuery(
			userId,
			update,
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

		// update the user's state.
		state.SetUserState(userId.ID, state.StateImageSent, tags.UserImageStore.UserStates)

	case state.StateImageSent:
		if pinsName == "Next ⬇️" {
			sendPhotoParams := tags.handleNextImageQuery(
				userId,
				update,
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
		}
	}
}
