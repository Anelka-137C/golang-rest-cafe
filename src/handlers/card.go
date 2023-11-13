package handlers

import (
	"net/http"

	"github.com/Anelka-137C/cafe-app/internal/card"
	"github.com/Anelka-137C/cafe-app/src/util"
	"github.com/gin-gonic/gin"
)

const (
	cardCreationMessage = "Card successfully created"
)

type Card struct {
	CardService card.Service
}

func NewCard(s card.Service) *Card {
	return &Card{
		CardService: s,
	}
}

func (p *Card) CreateCard() gin.HandlerFunc {
	return func(c *gin.Context) {
		card, err := p.CardService.CreateCard(c)
		if err != nil {
			util.BuildBadResponse(http.StatusBadRequest, err, c)
			return
		}
		util.BuildResponse(http.StatusOK, card, cardCreationMessage, c, "Product")
	}
}
