package handlers

import (
	"net/http"

	"github.com/Anelka-137C/cafe-app/internal/card"
	"github.com/Anelka-137C/cafe-app/src/util"
	"github.com/gin-gonic/gin"
)

const (
	cardCreationMessage = "Card successfully created"
	cardGetMessage      = "Card successfully obtained"
	cardAddMessage      = "Card successfully added"
)

type Card struct {
	CardService card.Service
}

func NewCard(s card.Service) *Card {
	return &Card{
		CardService: s,
	}
}

func (ca *Card) CreateCard() gin.HandlerFunc {
	return func(c *gin.Context) {
		card, err := ca.CardService.CreateCard(c)
		if err != nil {
			util.BuildBadResponse(http.StatusBadRequest, err, c)
			return
		}
		util.BuildResponse(http.StatusOK, card, cardCreationMessage, c, "Product")
	}
}

func (ca *Card) GetCard() gin.HandlerFunc {
	return func(c *gin.Context) {
		card, err := ca.CardService.GetCard(c)
		if err != nil {
			util.BuildBadResponse(http.StatusBadRequest, err, c)
			return
		}
		util.BuildResponse(http.StatusOK, card, cardGetMessage, c, "Card")
	}
}

func (ca *Card) AddToCard() gin.HandlerFunc {
	return func(c *gin.Context) {
		card, err := ca.CardService.AddToCard(c)
		if err != nil {
			util.BuildBadResponse(http.StatusBadRequest, err, c)
			return
		}
		util.BuildResponse(http.StatusOK, card, cardAddMessage, c, "Card")
	}
}

func (ca *Card) RestToCard() gin.HandlerFunc {
	return func(c *gin.Context) {
		card, err := ca.CardService.RestToCard(c)
		if err != nil {
			util.BuildBadResponse(http.StatusBadRequest, err, c)
			return
		}
		util.BuildResponse(http.StatusOK, card, cardAddMessage, c, "Card")
	}
}

func (ca *Card) DeleteCard() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := ca.CardService.DeleteCard(c)
		if err != nil {
			util.BuildBadResponse(http.StatusBadRequest, err, c)
			return
		}
		util.BuildResponse(http.StatusOK, "", cardAddMessage, c, "Card")
	}
}
