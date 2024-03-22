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

	api := router.Group("/api")
	{
		prefix := api.Group("/v1")
		{
			movie := prefix.Group("/movie")
			{
				movie.GET("", r.MovieService.GetAll)
				movie.GET(":id", r.MovieService.GetByID)
			}
		}

	}

	err := router.Run(":" + helper.GetEnv("APP_PORT"))
	if err != nil {
		return
	}

}
