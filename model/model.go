package model

type VideoSchema struct {
	Title        string         `json:"offset"`
	ChannelId    string         `json:"channelId"`
	ChannelTitle string         `json:"channelTitle"`
	VideoId      string         `json:"videoId "`
	Description  string         `json:"description"`
	Thumbnails   ThumbnailsData `json:"thumbnails"`
	PublishedAt  uint64         `json:"publishedAt"`
}

type ThumbnailsData struct {
	Url    string `json:"url"`
	Width  uint64 `json:"width"`
	Height uint64 `json:"hight"`
}

type Page struct {
	Offset int `json:"offset"`
	Limit  int `json:"limit"`
}
