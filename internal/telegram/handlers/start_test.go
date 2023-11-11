package handlers

import (
	"github.com/guisecreator/pintebot/internal/telegram/types"
	"github.com/mymmrac/telego"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type mockBot struct {
	mock.Mock
	telego.Bot
}

type mockCommandsOptions struct {
	mock.Mock
	*types.CommandsOptions
}

type mockTypesServices struct {
	mock.Mock
}

func TestMessageError(t *testing.T) {
	t.Parallel()

	testCase := struct {
		UserID           telego.ChatID
		ReplyToMessageID int
		Message          string
		IsReply          bool
	}{
		UserID: telego.ChatID{
			ID:       123,
			Username: "test_user",
		},
		ReplyToMessageID: 456,
		Message:          "test_message",
		IsReply:          true,
	}

	result := MessageError(
		testCase.UserID,
		testCase.ReplyToMessageID,
		testCase.Message,
		testCase.IsReply,
	)

	assert.NotNil(t, result)
	assert.Equal(t, testCase.UserID, result.ChatID)
	assert.Equal(t, testCase.Message, result.Text)
	assert.Equal(t, testCase.IsReply, result.DisableNotification)

	for _, tt := range testCase.Message {
		t.Run(string(tt), func(t *testing.T) {
			t.Parallel()
		})
	}
}

func TestBuildKeyboard(t *testing.T) {
	t.Parallel()

	result := BuildKeyboard()

	const numberOfRows = 5
	assert.Equal(t, numberOfRows, len(result.InlineKeyboard))

	assert.Equal(
		t,
		"find_pin_via_tag",
		result.InlineKeyboard[0][0].CallbackData,
	)

	assert.Equal(
		t,
		"boards",
		result.InlineKeyboard[1][0].CallbackData,
	)

	assert.Equal(
		t,
		"settings",
		result.InlineKeyboard[2][0].CallbackData,
	)

	assert.Equal(
		t,
		"help_info",
		result.InlineKeyboard[3][0].CallbackData,
	)

	assert.Equal(t, "GitHub", result.InlineKeyboard[5][0].Text)
	assert.Equal(t, "https://github.com/guisecreator/pintebot", result.InlineKeyboard[5][0].URL)

}

func TestStartCommand_HandleStartCallback(t *testing.T) {
	t.Parallel()

	mockBotInstance := new(mockBot)
	mockCommandsOptionsInstance := new(mockCommandsOptions)
	mockTypesServicesInstance := new(mockTypesServices)

	startCommand := &StartCommand{
		CommandsOptions: mockCommandsOptionsInstance.CommandsOptions,
		logger:          &logrus.Logger{},
	}

	mockBotInstance.On("SendMessage", mock.Anything).Return(&telego.Message{}, nil)
	mockBotInstance.On("AnswerCallbackQuery", mock.Anything).Return(nil)

	mockCommandsOptionsInstance.On("InitCommandsText", mock.Anything)
	mockTypesServicesInstance.On("InitCommandsText", mock.Anything)

	startCommand.HandleStartCallback()(&mockBotInstance.Bot, telego.Update{})

	mockBotInstance.AssertExpectations(t)
	mockCommandsOptionsInstance.AssertExpectations(t)
	mockTypesServicesInstance.AssertExpectations(t)
}
