package services

import (
	"github.com/tv2169145/golang-microservices/mvc/domain"
	"github.com/tv2169145/golang-microservices/mvc/untils"
	"net/http"
)

type itemService struct {

}

var ItemService itemService

func (i *itemService) GetItem(itemId string) (*domain.Item, *untils.ApplicationError) {
	return nil, &untils.ApplicationError{
		Message: "implement me",
		StatusCode: http.StatusInternalServerError,
	}
}
