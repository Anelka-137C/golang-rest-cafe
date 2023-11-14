package routes

import (
	"github.com/Anelka-137C/cafe-app/internal/card"
	"github.com/Anelka-137C/cafe-app/src/handlers"
	"github.com/Anelka-137C/cafe-app/src/middlewares"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/mongo"
)

type cardRouter struct {
	db  *mongo.Client
	eng *gin.Engine
	rg  *gin.RouterGroup
}

type CardRouter interface {
	MapCardRoutes()
}

func NewCardtRouter(eng *gin.Engine, db *mongo.Client) CardRouter {
	return &cardRouter{eng: eng, db: db}
}

// MapRoutes implements Router.
func (r *cardRouter) MapCardRoutes() {
	r.setCardGroup()
	r.card()
}

func (r *cardRouter) setCardGroup() {
	r.rg = r.eng.Group("/card")
}

func (r *cardRouter) card() {

	group := r.rg.Group("/card")
	repo := card.NewRepository(r.db)
	service := card.NewService(repo)
	handler := handlers.NewCard(service)

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("ValidateIsEmptyProducts", middlewares.ValidateIsEmptyProducts())
		v.RegisterValidation("ValidateProducts", middlewares.ValidateProducts(repo))
	}
	group.POST("/create" /*middlewares.ValidateJwt()*/, handler.CreateCard())
	group.GET("/get/:_id" /*middlewares.ValidateJwtForUsers()*/, handler.GetCard())
	group.PUT("/addToCard/:_id" /*middlewares.ValidateJwtForUsers()*/, handler.AddToCard())
	// group.DELETE("/delete/:_id", middlewares.ValidateJwt(), handler.DeleteProduct())
	// group.GET("/get/:_id", middlewares.ValidateJwtForUsers(), handler.GetProduct())
	// group.PUT("/update/:_id", middlewares.ValidateJwt(), handler.UpdateProduct())
	// group.GET("/getAll", middlewares.ValidateJwtForUsers(), handler.GetAllProduct())
	// group.GET("/getByName", middlewares.ValidateJwtForUsers(), handler.GetProductByName())

}
