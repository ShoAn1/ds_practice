package handler

import "github.com/labstack/echo/v4"

func SetRoutes(e *echo.Echo) *echo.Echo {
	e.GET("/photos", GetPhotosbyAlbumsAndId)
	e.POST("/photos", AddPhotos)
	e.POST("/albums", AddAlbums)
	e.GET("/albums", GetAlbums)
	e.GET("/search", Search)
	return e
}
