package domain

import (
	"assignment-movie/common/helper"
	"assignment-movie/models"
	"github.com/gin-gonic/gin"
	"time"
)

type (
	MovieServices interface {
		GetAll(c *gin.Context, pageParam, limitParam, orderByParam, search string, startTime time.Time) ([]models.GetMovies, helper.Paginate, error)
		GetByID(c *gin.Context, id int, startTime time.Time) (models.GetMovies, error)
		Create(c *gin.Context, req models.ReqMovie, startTime time.Time) error
	}

	MoviePostgres interface {
		GetAll(limit, offset int, orderBy, search string) ([]models.GetMovies, int64, error)
		GetByID(id int) (models.GetMovies, error)
		GetByTitle(title string) (models.GetMovies, error)
		Create(movie models.Movies) error
	}
)
