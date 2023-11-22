package pinterest_service

type PinterestElement struct {
	Step string
}

type PinterestService struct {
	PinterestMap map[int64]*PinterestElement
}

func NewPinterestService() *PinterestService {
	return &PinterestService{
		PinterestMap: map[int64]*PinterestElement{},
	}
}

func (p *PinterestService) CreatePinterestService(userId int64) {
	p.PinterestMap[userId] = &PinterestElement{}
}

func (p *PinterestService) DeleteService(chatId int64) {
	delete(p.PinterestMap, chatId)
}
