package routes

import (
	"assignment-movie/common/helper"
	"assignment-movie/controllers/movies"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type InitRoutes interface {
	InitRouter()
}

type RouteService struct {
	MovieService movies.MovieController `inject:"controller_movie_master"`
}

func InitPackage() *RouteService {
	return &RouteService{
		MovieService: &movies.MasterMovies{},
	}
}

func (r *RouteService) InitRouter() {
	router := gin.Default()

	api := router.Group("/api")
	{
		prefix := api.Group("/v1")
		{
			movie := prefix.Group("/movies")
			{
				movie.GET("", r.MovieService.GetAll)
				movie.GET(":id", r.MovieService.GetByID)
				movie.POST("", r.MovieService.Create)
			}
		}

	}

	router.POST("/upload", func(c *gin.Context) {
		file, err := c.FormFile("image")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		nameImage, err := helper.UploadImageMovies(c, file)

		fmt.Println(nameImage)

		c.JSON(http.StatusOK, gin.H{"message": "Gambar berhasil diunggah"})
	})

	router.Static("/common/public/assets/img", "./common/public/assets/img")

	err := router.Run(":" + helper.GetEnv("APP_PORT"))
	if err != nil {
		return
	}

}
