package routes

import (
	"github.com/Anelka-137C/cafe-app/internal/product"
	"github.com/Anelka-137C/cafe-app/src/handlers"
	"github.com/Anelka-137C/cafe-app/src/middlewares"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/mongo"
)

type productRouter struct {
	db  *mongo.Client
	eng *gin.Engine
	rg  *gin.RouterGroup
}

type ProductRouter interface {
	MapProductRoutes()
}

func NewProductRouter(eng *gin.Engine, db *mongo.Client) ProductRouter {
	return &productRouter{eng: eng, db: db}
}

// MapRoutes implements Router.
func (r *productRouter) MapProductRoutes() {
	r.setProductGroup()
	r.product()
}

func (r *productRouter) setProductGroup() {
	r.rg = r.eng.Group("/products")
}

func (r *productRouter) product() {

	group := r.rg.Group("/product")
	repo := product.NewRepository(r.db)
	service := product.NewService(repo)
	handler := handlers.NewProduct(service)
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("ValidateCategory", middlewares.ValidateCategory(repo))
	}
	group.POST("/create", middlewares.ValidateJwt(), handler.CreateProduct())
	group.DELETE("/delete/:_id", middlewares.ValidateJwt(), handler.DeleteProduct())
	group.GET("/get/:_id", middlewares.ValidateJwtForUsers(), handler.GetProduct())
	group.PUT("/update/:_id", middlewares.ValidateJwt(), handler.UpdateProduct())
	group.GET("/getAll", middlewares.ValidateJwtForUsers(), handler.GetAllProduct())
	group.GET("/getByName", middlewares.ValidateJwtForUsers(), handler.GetProductByName())

}
