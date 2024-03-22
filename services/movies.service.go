package services

import (
	"assignment-movie/common/helper"
	"assignment-movie/domain"
	"assignment-movie/models"
	"github.com/gin-gonic/gin"
	"log"
	"math"
	"net/http"
	"time"
)

type Movies struct {
	moviesPostgres domain.MoviePostgres
}

func NewMoviesService(mp domain.MoviePostgres) domain.MovieServices {
	return &Movies{
		moviesPostgres: mp,
	}
}

func (s *Movies) GetAll(c *gin.Context, pageParam, limitParam, orderByParam, search string, startTime time.Time) ([]models.GetMovies, helper.Paginate, error) {
	var (
		dataMovies []models.GetMovies
		paginate   helper.Paginate
		totalData  int64
		err        error
	)

	page, limit, offset, err := helper.Pagination(pageParam, limitParam)
	if err != nil {
		helper.ResponseAPI(c, false, http.StatusBadRequest, http.StatusText(http.StatusBadRequest), []string{err.Error()}, startTime)
		return dataMovies, paginate, err
	}

	dataMovies, totalData, err = s.moviesPostgres.GetAll(limit, offset, orderByParam, search)
	if err != nil {
		log.Printf("error movie service GetAll: %s", err)
		helper.ResponseAPI(c, false, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), []string{err.Error()}, startTime)
		return dataMovies, paginate, err
	}

	for i, movie := range dataMovies {
		createdAtStr := movie.CreatedAt.Format("2006-01-02 15:04:05")
		createdAt, err := time.Parse("2006-01-02 15:04:05", createdAtStr)
		if err != nil {
			log.Printf("Error parsing created_at: %s", err)
			return dataMovies, paginate, err
		} else {
			dataMovies[i].CreatedAt = createdAt
		}

		if movie.UpdatedAt != nil {
			updatedAtStr := movie.UpdatedAt.Format("2006-01-02 15:04:05")
			updatedAt, err := time.Parse("2006-01-02 15:04:05", updatedAtStr)
			if err != nil {
				log.Printf("Error parsing updated_at: %s", err)
				return dataMovies, paginate, err
			} else {
				dataMovies[i].UpdatedAt = &updatedAt
			}
		}
	}

	totalPages := int(math.Ceil(float64(totalData) / float64(limit)))

	paginate = helper.Paginate{
		Page:       page,
		Limit:      limit,
		TotalPages: totalPages,
		TotalData:  totalData,
	}

	return dataMovies, paginate, nil
}

func (s *Movies) GetByID(c *gin.Context, id int, startTime time.Time) (models.GetMovies, error) {
	var (
		dataMovies models.GetMovies
		err        error
	)

	dataMovies, err = s.moviesPostgres.GetByID(id)
	if err != nil {
		log.Printf("error movie service GetByID: %s", err)
		helper.ResponseAPI(c, false, http.StatusNotFound, http.StatusText(http.StatusNotFound), []string{err.Error()}, startTime)
		return dataMovies, err
	}

	createdAtStr := dataMovies.CreatedAt.Format("2006-01-02 15:04:05")
	createdAt, err := time.Parse("2006-01-02 15:04:05", createdAtStr)
	if err != nil {
		log.Printf("Error parsing created_at: %s", err)
		return dataMovies, err
	} else {
		dataMovies.CreatedAt = createdAt
	}

	return dataMovies, nil
}
