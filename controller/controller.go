package controller

import (
	"errors"
	"log"
	"net/http"
	"strconv"

	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"

	"github.com/Vikash123abc/Fampay-Assignment.git/config"
	"github.com/Vikash123abc/Fampay-Assignment.git/model"
	"github.com/Vikash123abc/Fampay-Assignment.git/service"
	"github.com/gin-gonic/gin"
)

type YoutubeAPI struct {
	Config          *config.Config
	YoutubeService  service.Service
	MongoCollection *mongo.Collection
}

func NewYouTubeAPI(logger *zap.SugaredLogger, config *config.Config, youtubeService service.Service, mongoCollection *mongo.Collection) *YoutubeAPI {
	return &YoutubeAPI{
		Config:          config,
		YoutubeService:  youtubeService,
		MongoCollection: mongoCollection,
	}
}

func (yt YoutubeAPI) LoadStoredVideos(ctx *gin.Context) (int, interface{}, error) {

	limit, err := strconv.Atoi(ctx.Query("limit"))
	if err != nil {
		limit = 10
	}

	offset, err := strconv.Atoi(ctx.Query("offset"))
	if err != nil {
		offset = 1
	}

	page := model.Page{
		Offset: offset,
		Limit:  limit,
	}

	videos, err := yt.YoutubeService.LoadStoredVideos(ctx, page, "", yt.MongoCollection)
	if err != nil {
		log.Println("error getting videos list", "error", err)
		return http.StatusInternalServerError, gin.H{"success": false, "error": "error loading videos"}, err
	}

	return http.StatusOK, gin.H{"success": true, "videos": videos, "error": ""}, nil
}

func (yt YoutubeAPI) LoadStoredVideosByQuery(ctx *gin.Context) (int, interface{}, error) {

	offset, err := strconv.Atoi(ctx.Query("offset"))
	if err != nil {
		offset = 1
	}
	limit, err := strconv.Atoi(ctx.Query("limit"))
	if err != nil {
		limit = 10
	}

	searchQuery := ctx.Query("search")
	if searchQuery == "" {
		log.Panicln("search query empty")
		return http.StatusInternalServerError, gin.H{"success": false, "error": "search query empty"}, errors.New("search query empty")
	}

	page := model.Page{
		Offset: offset,
		Limit:  limit,
	}

	videos, err := yt.YoutubeService.LoadStoredVideos(ctx, page, searchQuery, yt.MongoCollection)
	if err != nil {
		log.Println("error getting videos list", "error", err)
		return http.StatusInternalServerError, gin.H{"success": false, "error": "error loading videos"}, err
	}

	return http.StatusOK, gin.H{"success": true, "videos": videos, "error": ""}, nil
}
