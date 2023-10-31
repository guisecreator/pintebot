package handlers

type CommandsHandler struct {
	StartCommand  *StartCommand
	BoardsCommand *BoardsCommand
	TagsCommand   *TagsCommand
}

func NewCommandsHandler() (*CommandsHandler, error) {
	return &CommandsHandler{
		StartCommand:  &StartCommand{},
		BoardsCommand: &BoardsCommand{},
	}, nil
}
