package app

import (
	"fmt"
	"producer/internal/domain"
)

type BuyHandler struct {
	BuyApp *BuyApp
}

func NewBuyHandler(buyApp *BuyApp) *BuyHandler {
	return &BuyHandler{
		BuyApp: buyApp,
	}
}

func (h *BuyHandler) HandleBuyCreation(buy *domain.Buy) error {
	err := h.BuyApp.CreateBuy(buy)
	if err != nil {
		return fmt.Errorf("failed to create buy: %w", err)
	}

	return nil
}
