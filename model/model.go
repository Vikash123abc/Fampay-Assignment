package model

type Video struct {
	Title        string `json:"title" bson:"title"`
	ChannelId    string `json:"channelId" bson:"channel_id"`
	ChannelTitle string `json:"channelTitle" bson:"channel_title"`
	VideoId      string `json:"videoId " bson:"video_id"`
	Description  string `json:"description" bson:"description"`
	PublishedAt  string `json:"publishedAt" bson:"published_at"`
	ThumbnailUrl string `json:"thumbnail_url" bson:"thumbnail_url"`
	VideoETag    string `json:"video_etag" bson:"video_etag"`
}

type Page struct {
	Offset int `json:"offset" bson:"offset"`
	Limit  int `json:"limit" bson:"limit"`
}
