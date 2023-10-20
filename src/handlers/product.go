package handlers

import (
	"net/http"

	"github.com/Anelka-137C/cafe-app/internal/product"
	"github.com/Anelka-137C/cafe-app/src/util"
	"github.com/gin-gonic/gin"
)

const (
	productCreationMessage = "product successfully created"
	productDeleteMessage   = "product successfully deleted"
	productGetMessage      = "Product successfully obtained"
)

type Product struct {
	productService product.Service
}

func NewProduct(s product.Service) *Product {
	return &Product{
		productService: s,
	}
}

func (p *Product) CreateProduct() gin.HandlerFunc {
	return func(c *gin.Context) {
		product, err := p.productService.CreateProduct(c)
		if err != nil {
			util.BuildBadResponse(http.StatusBadRequest, err, c)
			return
		}
		util.BuildResponse(http.StatusOK, product, productCreationMessage, c)
	}
}

func (p *Product) DeleteProduct() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := p.productService.DeleteProduct(c)
		if err != nil {
			util.BuildBadResponse(http.StatusBadRequest, err, c)
			return
		}
		util.BuildResponse(http.StatusOK, nil, productDeleteMessage, c)
	}
}

func (p *Product) GetProduct() gin.HandlerFunc {
	return func(c *gin.Context) {
		product, err := p.productService.GetProduct(c)
		if err != nil {
			util.BuildBadResponse(http.StatusBadRequest, err, c)
			return
		}
		util.BuildResponse(http.StatusOK, product, productGetMessage, c)
	}
}
