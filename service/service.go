package service

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"log"

	"github.com/Vikash123abc/Fampay-Assignment.git/config"
	"github.com/Vikash123abc/Fampay-Assignment.git/datastore"
	"github.com/Vikash123abc/Fampay-Assignment.git/model"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"
)

type Service struct {
	Config *config.Config
}

var (
	maxResults = flag.Int64("max-results", 25, "Max YouTube results")
)

func YoutubeMetaData(config *config.Config, query string) (*youtube.SearchListResponse, error) {
	keys := config.YoutubeApiKeys
	fmt.Println("keys", keys)
	for _, key := range keys {
		youtubeClient, err := youtube.NewService(context.TODO(), option.WithAPIKey(key))
		if err != nil {
			continue
		}
		call := youtubeClient.Search.List([]string{"id,snippet"}).
			Q(query).
			MaxResults(*maxResults).
			Order("date").
			Type("video").
			PublishedAfter("2021-01-01T00:00:00Z")

		response, err := call.Do()
		if err != nil {
			continue
		}

		return response, nil
	}
	return nil, errors.New("invalid api key(s)")
}

func GetVideosByQuery(config *config.Config, query string, mongoCollection *mongo.Collection) error {
	response, err := YoutubeMetaData(config, query)
	if err != nil {
		log.Println("error getting video from youtube", err)
		return err
	}

	videosList := []model.Video{}

	// Create a list of videos metadata for upserting
	for _, item := range response.Items {
		newVideo := model.Video{
			Title:        item.Snippet.Title,
			ChannelId:    item.Snippet.ChannelId,
			VideoId:      item.Id.PlaylistId,
			ChannelTitle: item.Snippet.ChannelTitle,
			Description:  item.Snippet.Description,
			PublishedAt:  item.Snippet.PublishedAt,
			ThumbnailUrl: item.Snippet.Thumbnails.Default.Url,
			VideoETag:    item.Etag,
		}
		videosList = append(videosList, newVideo)
	}

	err = datastore.BulkUpsertVideos(context.TODO(), videosList, mongoCollection)
	if err != nil {
		log.Println("error while bulk upserting videos", err)
		return err
	}

	// Printing data to the terminal
	fmt.Println(videosList)
	return nil
}

// Loading the videos from database
func (svc Service) LoadStoredVideos(ctx context.Context, page model.Page, searchString string, mongoCollection *mongo.Collection) ([]*model.Video, error) {
	videos, err := datastore.GetVideosList(ctx, page, searchString, mongoCollection)
	if err != nil {
		return nil, err
	}
	return videos, nil
}
