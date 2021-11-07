package handler

import (
	"Kreditbee/db_create"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

const alb = "albums"

func GetAlbums(c echo.Context) error {
	res := []Albums{}
	statemnt := "select * from albums"
	err := db_create.GetDb_instance().Raw(statemnt).Scan(&res)
	if err.Error != nil {
		return c.JSON(http.StatusInternalServerError, "error in db opearations")
	}
	return c.JSON(http.StatusOK, res)
}

func GetPhotosbyAlbumsAndId(c echo.Context) error {
	resp := []Photos{}
	id := c.QueryParam("id")
	fmt.Println("**********", id)
	statemnt := fmt.Sprintf(`select * from photos where albums_id = %s`, id)
	err := db_create.GetDb_instance().Raw(statemnt).Scan(&resp)
	if err.Error != nil {
		return c.JSON(http.StatusInternalServerError, "error in db opearations")
	}
	if err.RowsAffected == 0 {
		return c.JSON(http.StatusNotFound, "value not in db id->"+id)
	}
	return c.JSON(http.StatusAccepted, resp)
}

func AddAlbums(c echo.Context) error {
	req := Albums{}
	er := c.Bind(&req)
	if er != nil {
		fmt.Println("BIND ERROR")
	}
	statement := fmt.Sprintf(`INSERT INTO albums(
	id, user_id, title)
	VALUES (%d, %d, '%s');`, req.Id, req.UserId, req.Title)
	err := db_create.GetDb_instance().Exec(statement)
	if err.Error != nil {
		return c.JSON(http.StatusInternalServerError, "error in db opearations")
	}
	return c.String(http.StatusCreated, "")
}

func AddPhotos(c echo.Context) error {
	req := Photos{}
	c.Bind(&req)
	statemnt := fmt.Sprintf(`INSERT INTO photos(
	albums_id, id, title, url, thumbnail_url)
	VALUES (%d, %d, '%s', '%s', '%s');`, req.AlbumsId, req.Id, req.Title, c.Request().RequestURI, req.ThumbnailUrl)
	err := db_create.GetDb_instance().Exec(statemnt)
	if err.Error != nil {
		return c.JSON(http.StatusInternalServerError, "error in db opearations")
	}
	return c.JSON(http.StatusCreated, nil)
}

func Search(c echo.Context) error {
	types := c.QueryParam("types")
	id := c.QueryParam("id")
	albums := c.QueryParam("albums")
	fmt.Println("**********", types, id, albums)
	var statement string
	if types == alb {
		statement = fmt.Sprintf(`select * from albums where id=%s`, id)
		res := []Albums{}
		err := db_create.GetDb_instance().Raw(statement).Scan(&res)
		if err.Error == nil {
			return c.JSON(http.StatusOK, res)
		}
	} else {
		statement = fmt.Sprintf(`select * from photos where id=%s and albums_id=%s`, id, albums)
		res := []Albums_photos_resp{}
		err := db_create.GetDb_instance().Raw(statement).Scan(&res)
		if err.Error == nil {
			return c.JSON(http.StatusOK, res)
		}
	}

	return c.JSON(http.StatusInternalServerError, "")
}
