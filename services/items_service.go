package services

var (
	ItemsService itemServiceInterface = &itemService{}
)

type itemServiceInterface interface {
	GetItem()
	SaveItem()
}
type itemService struct {
}

func (s *itemService) GetItem() {

}
func (s *itemService) SaveItem() {

}
