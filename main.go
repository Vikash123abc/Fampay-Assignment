package main

import (
	"fmt"
	"log"
	"time"

	config1 "github.com/Vikash123abc/Fampay-Assignment.git/Config"
	"github.com/Vikash123abc/Fampay-Assignment.git/controller"
	"github.com/Vikash123abc/Fampay-Assignment.git/datastore"
	ytservice "github.com/Vikash123abc/Fampay-Assignment.git/service"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func refreshVideoList(config *config1.Config, collection *mongo.Collection) {
	ticker := time.NewTicker(10 * time.Second)
	quit := make(chan struct{})
	go func() {
		for {
			select {
			case <-ticker.C:
				//logger.Infow("loading video list")
				ytservice.FetchVideosByQuery(config, "football", collection)

			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()
}

func main() {

	r := gin.New()
	fmt.Println("Backend Service starting...")
	config1.InitializeLogger()

	config, err := config1.Load(".")
	if err != nil {
		log.Fatal(err)
	}
	collection, err := datastore.ConnectMongo(config)
	if err != nil {
		fmt.Println("Error connecting to Database", err)
		return
	} else {
		fmt.Println("connected to Database")
	}

	youtubeService := ytservice.Service{
		Config: &config,
	}

	// create youtube-api to access APIs
	youtubeApi := controller.YoutubeAPI{
		Config:          &config,
		MongoCollection: collection,
		YoutubeService:  youtubeService,
	}
	fmt.Println(youtubeApi)

	// go routine to pull videos metadata from youtube
	go refreshVideoList(&config, collection)
	errorHandler := config1.ProvideAPIWrap()

	// handler functions
	r.GET("/", errorHandler(youtubeApi.LoadStoredVideos))

	r.GET("/search", errorHandler(youtubeApi.LoadStoredVideosByQuery))
	//http.ListenAndServe(":1000", nil)
	r.Run(":2000")
}
