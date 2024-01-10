package app

import "producer/internal/domain"

type BuyService interface {
	CreateBuy(buy *domain.Buy) error
}

type BuyApp struct {
	BuyService BuyService
}

func NewBuyApplication(buyService BuyService) *BuyApp {
	return &BuyApp{
		BuyService: buyService,
	}
}

func (a *BuyApp) CreateBuy(buy *domain.Buy) error {
	return a.BuyService.CreateBuy(buy)
}
