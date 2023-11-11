package config

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
)

type CommandsText struct {
	Description string `yaml:"description"`
	AnyTagText  string `yaml:"any_tag_text"`

	TagsCommand struct {
		Phrase string `yaml:"phrase"`

		InlineKeyboard struct {
			KeyboardRow1 struct {
				NextButton   string `yaml:"next_button"`
				CancelButton string `yaml:"cancel_button"`
			} `yaml:"keyboard_row_1"`
		} `yaml:"inline_keyboard"`
	} `yaml:"tags_command"`

	StartCommand struct {
		InlineKeyboard struct {
			KeyboardRow1 struct {
				FindPinViaTagButton string `yaml:"find_pin_via_tag_button"`
			} `yaml:"keyboard_row_1"`

			KeyboardRow2 struct {
				BoardsButton string `yaml:"boards_button"`
			} `yaml:"keyboard_row_2"`

			KeyboardRow3 struct {
				SettingsButton string `yaml:"settings_button"`
			} `yaml:"keyboard_row_3"`

			KeyboardRow4 struct {
				HelpButton string `yaml:"help_button"`
			} `yaml:"keyboard_row_4"`

			KeyboardRow5 struct {
				ProjectButton string `yaml:"project_button"`
			} `yaml:"keyboard_row_5"`
		} `yaml:"inline_keyboard"`
	} `yaml:"start_command"`

	BoardsCommand struct {
		InlineKeyboard struct {
			KeyboardRow1 struct {
				SpecificUsersBoardsButton string `yaml:"specific_users_boards_btn"`
			} `yaml:"keyboard_row_1"`

			KeyboardRow2 struct {
				BoardsByTitleButton string `yaml:"boards_by_title_btn"`
			} `yaml:"keyboard_row_2"`

			KeyboardRow3 struct {
				HistoryOfBoardsButton string `yaml:"history_of_boards_btn"`
			} `yaml:"keyboard_row_3"`

			KeyboardRow4 struct {
				HelpButton string `yaml:"help_button"`
			} `yaml:"keyboard_row_4"`

			KeyboardRow5 struct {
				BackToStartButton string `yaml:"back_to_start_btn"`
			} `yaml:"keyboard_row_5"`
		} `yaml:"inline_keyboard"`
	} `yaml:"boards_command"`
}

func InitCommandsText(path string) (CommandsText, error) {
	var command CommandsText

	err := cleanenv.ReadConfig(path, &command)
	if err != nil {
		return command, fmt.Errorf("ConfigErr: %v", err)
	}

	return command, nil
}
