package main

import (
	"RestApi-Golang/repository/mongodb"
	"RestApi-Golang/usecase"
	"context"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	slog.Info("Loaded .env file successfully")
}

func mongoConnection() *mongo.Client {
	// Use the SetServerAPIOptions() method to set the version of the Stable API on the client
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(os.Getenv("MONGO_URI")).SetServerAPIOptions(serverAPI)
	// Create a new client and connect to the server
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		panic(err)
	}

	// Send a ping to confirm a successful connection
	if err := client.Database(os.Getenv("MONGO_DBNAME")).RunCommand(context.TODO(), bson.D{{"ping", 1}}).Err(); err != nil {
		panic(err)
	}
	fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")

	return client
}

func main() {
	mongoClient := mongoConnection()
	defer mongoClient.Disconnect(context.Background())

	// Mongo Collection
	database := mongoClient.Database(os.Getenv("MONGO_DBNAME"))
	collection := database.Collection(os.Getenv("MONGO_COLLECTION_NAME"))

	repo := mongodb.New(*collection)

	// UserService instance
	userService := usecase.New(repo)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.SetHeader("Content-Type", "application/json"))

	r.Route("/api/v1/users", func(r chi.Router) {
		r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("server is healty...."))
		})

		r.Post("/", userService.CreateUser)
		r.Get("/{id}", userService.GetUserByID)
		r.Get("/", userService.GetAllUsers)
		r.Put("/{id}", userService.UpdateUserAgeByID)
		r.Delete("/{id}", userService.DeleteUserByID)
		r.Delete("/", userService.DeleteAllUsers)
	})

	slog.Info("Server started at :4444")
	http.ListenAndServe(":4444", r)
}
