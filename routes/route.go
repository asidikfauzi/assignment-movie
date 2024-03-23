package routes

import (
	"assignment-movie/common/helper"
	"assignment-movie/controllers/movies"
	"github.com/gin-gonic/gin"
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

	router.Static("/common/public/assets/img", "./common/public/assets/img")

	api := router.Group("/api")
	{
		prefix := api.Group("/v1")
		{
			movie := prefix.Group("/movies")
			{
				movie.GET("", r.MovieService.GetAll)
				movie.GET(":id", r.MovieService.GetByID)
				movie.POST("", r.MovieService.Create)
				movie.PATCH(":id", r.MovieService.Update)
			}
		}

	}

	err := router.Run(":" + helper.GetEnv("APP_PORT"))
	if err != nil {
		return
	}

}
