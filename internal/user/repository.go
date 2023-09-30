package user

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var UserColl *mongo.Collection

var db *mongo.Client

type repository struct {
	db *mongo.Client
}

type Repository interface {
	CreateUser(c *gin.Context)
}

func NewRepository(db *mongo.Client) Repository {
	return &repository{
		db: db,
	}
}

// CreateUser implements Respository.
func (r *repository) CreateUser(c *gin.Context) {
	panic("unimplemented")
}

func ConnectDB() {

	fmt.Println("Iniciando conexion a base de datos")

	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		log.Fatal("You must set your 'MONGODB_URI' environmental variable. See\n\t https://www.mongodb.com/docs/drivers/go/current/usage-examples/#environment-variable")
	}

	db, _ = mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb+srv://nacasas:D960bi8NAzg3hKIN@clusteranelka.vkgfe8i.mongodb.net/GoCafe"))

	UserColl = db.Database("GoCafe").Collection("users")
	fmt.Printf("Base de datos en linea")

}
