package datastore

import (
	"context"
	"errors"
	"fmt"
	"time"

	config "github.com/Vikash123abc/Fampay-Assignment.git/config"
	"github.com/Vikash123abc/Fampay-Assignment.git/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
)

const (
	DatabaseName   = "youtubeapi"
	CollectionName = "youtubeapi"
)

// function to connect to mongodb and returns the collection

func ConnectMongo(config config.Config) (*mongo.Collection, error) {
	serverAPIOptions := options.ServerAPI(options.ServerAPIVersion1)

	if config.MongoURI == "" {
		return nil, errors.New("MongoURI is empty")
	}

	clientOptions := options.Client().
		ApplyURI(config.MongoURI).
		SetServerAPIOptions(serverAPIOptions)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)

	collection := client.Database(DatabaseName).Collection(CollectionName)
	if err != nil {
		return nil, err
	}

	return collection, nil
}

// Function for doing bulk write to db, if video already exists then update the existing data
func BulkUpsertVideos(ctx context.Context, videos []model.Video, mongoCollection *mongo.Collection) error {
	var operations []mongo.WriteModel

	for _, video := range videos {
		operation := mongo.NewUpdateOneModel()
		operation.SetUpsert(true)
		operation.SetFilter(bson.M{"video_etag": video.VideoETag})
		operation.SetUpdate(bson.M{"$set": bson.M{"title": video.Title, "description": video.Description, "thumbnail_url": video.ThumbnailUrl, "published_at": video.PublishedAt}})

		operations = append(operations, operation)
	}

	bulkWrite := options.BulkWriteOptions{}
	bulkWrite.SetOrdered(true)

	_, err := mongoCollection.BulkWrite(ctx, operations, &bulkWrite)
	if err != nil {
		return err
	}
	return nil
}

/*
   fetches video list from mongodb.
   supports pagination and search query.
*/
func GetVideosList(ctx context.Context, page model.Page, searchString string, mongoCollection *mongo.Collection) ([]*model.Video, error) {
	videoList := make([]*model.Video, 0)

	filter := bson.M{}

	// Regex search on title and description
	if searchString != "" {
		filter = bson.M{"$or": []bson.M{
			{"title": primitive.Regex{Pattern: searchString, Options: "i"}},
			{"description": primitive.Regex{Pattern: searchString, Options: "i"}},
		},
		}
	}

	// sort in descending order of published_at field
	findOptions := &options.FindOptions{}
	findOptions.SetSort(bson.M{"published_at": -1})

	// Pagination Implementation
	findOptions.SetSkip(int64((page.Offset - 1) * page.Limit))
	findOptions.SetLimit(int64(page.Limit))

	fmt.Println("test-vikash")
	cursor, err := mongoCollection.Find(ctx, filter, findOptions)
	if err != nil {
		return nil, err
	}
	fmt.Println("test-vikash1")

	for cursor.Next(ctx) {
		video := &model.Video{}
		err := cursor.Decode(video)
		if err != nil {
			return nil, err
		} else {
			videoList = append(videoList, video)
		}
	}
	fmt.Println("test-vikash2")

	return videoList, nil
}
