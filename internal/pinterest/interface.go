package pinterest

type Interface interface {
	GetPinsById(pinId string) ([]PinData, error)
	GetBoard(bookmark string) (*BoardsData, error)
	GetBoards() ([]BoardData, error)
}
