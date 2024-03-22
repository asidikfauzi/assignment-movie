package movies

import (
	"assignment-movie/domain"
	"github.com/gin-gonic/gin"
)

type MovieController interface {
	GetAll(c *gin.Context)
	GetByID(c *gin.Context)
}

type MasterMovies struct {
	MoviesService domain.MovieServices `inject:"movies_service"`
}
