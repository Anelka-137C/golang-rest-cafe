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

type UserRouter interface {
	MapUserRoutes()
}

func NewUserRouter(eng *gin.Engine, db *mongo.Client) UserRouter {
	return &router{eng: eng, db: db}
}

// MapRoutes implements Router.
func (r *router) MapUserRoutes() {
	r.setUserGroup()
	r.user()
}

func (r *router) setUserGroup() {
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
		v.RegisterValidation("validateIfExistEmail", middlewares.ValidateIfExistEmail(repo))
	}
	group.POST("/create", middlewares.ValidateJwt(), handler.CreateUser())
	group.GET("/get/:_id", middlewares.ValidateJwt(), handler.GetUser())
	group.GET("/getEmail", middlewares.ValidateJwt(), handler.GetUserByEmail())
	group.GET("/getByName", middlewares.ValidateJwt(), handler.GetUsersByName())
	group.DELETE("/delete/:_id", middlewares.ValidateJwt(), handler.DeleteUser())
	group.PUT("/update/:_id", middlewares.ValidateJwt(), handler.UpdateUser())
	group.POST("/login", handler.Login())
}
