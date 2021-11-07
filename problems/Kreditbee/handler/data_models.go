package handler

type Albums struct {
	Id     int    `json:"id"`
	UserId int    `json:"userId"`
	Title  string `json:"title"`
}

type Albums_photos_resp struct {
	AlbumsId     int    `json:"albumId"`
	Id           int    `json:"id"`
	Title        string `json:"title"`
	Url          string `json:"url"`
	ThumbnailUrl string `json:"thumbnailUrl"`
}

type Photos struct {
	AlbumsId     int
	Id           int
	Title        string
	Url          string
	ThumbnailUrl string
}
