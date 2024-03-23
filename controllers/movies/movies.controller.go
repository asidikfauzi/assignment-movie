package movies

import (
	"assignment-movie/domain"
	"github.com/gin-gonic/gin"
)

type MovieController interface {
	GetAll(c *gin.Context)
	GetByID(c *gin.Context)
	Create(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}

type MasterMovies struct {
	MoviesService domain.MovieServices `inject:"movies_service"`
}
