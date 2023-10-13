package routes

import (
	"github.com/Anelka-137C/cafe-app/internal/user"
	"github.com/Anelka-137C/cafe-app/src/handlers"
	"github.com/Anelka-137C/cafe-app/src/middlewares"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/mongo"
)

type router struct {
	db  *mongo.Client
	eng *gin.Engine
	rg  *gin.RouterGroup
}

type Router interface {
	MapRoutes()
}

func NewRouter(eng *gin.Engine, db *mongo.Client) Router {
	return &router{eng: eng, db: db}
}

// MapRoutes implements Router.
func (r *router) MapRoutes() {
	r.setGroup()
	r.user()
}

func (r *router) setGroup() {
	r.rg = r.eng.Group("/users")
}

func (r *router) user() {

	group := r.rg.Group("/user")
	repo := user.NewRepository(r.db)
	service := user.NewService(repo)
	handler := handlers.NewUser(service)
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("validateEmail", middlewares.ValidateEmail(repo))
		v.RegisterValidation("validateRole", middlewares.ValidateRole(repo))

	}
	group.POST("/create", handler.CreateUser())
	group.GET("/get/:_id", handler.GetUser())
	group.DELETE("/delete/:_id", handler.DeleteUser())
	group.PUT("/update/:_id", handler.UpdateUser())
}
