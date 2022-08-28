package main

import (
	"log"
	"time"

	"github.com/Vikash123abc/Fampay-Assignment.git/config"
	"github.com/Vikash123abc/Fampay-Assignment.git/controller"
	"github.com/Vikash123abc/Fampay-Assignment.git/datastore"
	"github.com/Vikash123abc/Fampay-Assignment.git/service"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func main() {

	r := gin.New()
	log.Println("Backend Service starting...")

	// Intialising Logger
	config.InitializeLogger()

	// Loading Config values
	config1, err := config.Load(".")
	if err != nil {
		log.Fatal(err)
	}

	// Connecting to MomgoDB database
	collection, err := datastore.ConnectMongo(config1)
	if err != nil {
		log.Println("Error connecting to Database", err)
		return
	} else {
		log.Println("connected to Database")
	}

	youtubeService := service.Service{
		Config: &config1,
	}

	// creating youtube-api to access APIs
	youtubeApi := controller.YoutubeAPI{
		Config:          &config1,
		MongoCollection: collection,
		YoutubeService:  youtubeService,
	}

	// go routine for getting data from video for every 10 secs
	go getNewVideos(&config1, collection)
	errorHandler := config.WrapperAPI()

	// APIs

	r.GET("/", errorHandler(youtubeApi.LoadStoredVideos))
	r.GET("/search", errorHandler(youtubeApi.LoadStoredVideosByQuery))

	// Running the server on port number 2000
	PORT := ":2000"
	r.Run(PORT)
}

// Function for calling youtube apis every second for query = cricket
func getNewVideos(config *config.Config, collection *mongo.Collection) {
	ticker := time.NewTicker(10 * time.Second)
	done := make(chan bool)
	go func() {
		for {
			select {
			case <-ticker.C:
				service.GetVideosByQuery(config, "cricket", collection)

			case <-done:
				return
			}
		}
	}()
}
