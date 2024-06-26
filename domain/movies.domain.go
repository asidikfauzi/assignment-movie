package domain

import (
	"assignment-movie/common/helper"
	"assignment-movie/models"
	"github.com/gin-gonic/gin"
	"mime/multipart"
	"time"
)

type (
	MovieServices interface {
		GetAll(c *gin.Context, pageParam, limitParam, orderByParam, search string, startTime time.Time) ([]models.GetMovies, helper.Paginate, error)
		GetByID(c *gin.Context, id int, startTime time.Time) (models.GetMovies, error)
		Create(c *gin.Context, file *multipart.FileHeader, req models.ReqMovie, startTime time.Time) error
		Update(c *gin.Context, id string, file *multipart.FileHeader, req models.ReqMovie, startTime time.Time) error
		Delete(c *gin.Context, id string, startTime time.Time) error
	}

	MoviePostgres interface {
		GetAll(limit, offset int, orderBy, search string) ([]models.GetMovies, int64, error)
		GetByID(id int) (models.GetMovies, error)
		GetByTitle(title string) (models.GetMovies, error)
		Create(movie models.Movies) error
		Update(movie models.Movies) error
		CheckIfExists(movie models.Movies) (bool, error)
		Delete(movie models.Movies) error
	}
)
