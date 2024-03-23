package services

import (
	"assignment-movie/common/helper"
	"assignment-movie/domain"
	"assignment-movie/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"math"
	"mime/multipart"
	"net/http"
	"strconv"
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

func (s *Movies) Create(c *gin.Context, file *multipart.FileHeader, req models.ReqMovie, startTime time.Time) error {
	var (
		movie models.Movies
		err   error
	)

	getMovie, _ := s.moviesPostgres.GetByTitle(req.Title)
	if getMovie.Title != "" {
		err = fmt.Errorf("title '%s' already exists", req.Title)
		helper.ResponseAPI(c, false, http.StatusConflict, http.StatusText(http.StatusConflict), []string{err.Error()}, startTime)
		return err
	}

	newImage, err := helper.UploadImageMovies(c, file)
	if err != nil {
		helper.ResponseAPI(c, false, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), []string{err.Error()}, startTime)
		return err
	}

	getImage := helper.GetImageMovies(c, newImage)

	movie = models.Movies{
		Title:       req.Title,
		Description: req.Description,
		Rating:      req.Rating,
		Image:       getImage,
		CreatedAt:   time.Now(),
	}

	err = s.moviesPostgres.Create(movie)
	if err != nil {
		log.Printf("error movie service Create: %s", err)
		helper.ResponseAPI(c, false, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), []string{err.Error()}, startTime)
		return err
	}

	return nil
}

func (s *Movies) Update(c *gin.Context, id string, file *multipart.FileHeader, req models.ReqMovie, startTime time.Time) error {
	var (
		movie models.Movies
		err   error
	)

	idInt, err := strconv.Atoi(id)
	if err != nil {
		helper.ResponseAPI(c, false, http.StatusBadRequest, http.StatusText(http.StatusBadRequest), []string{err.Error()}, startTime)
		return err

	}

	_, err = s.moviesPostgres.GetByID(idInt)
	if err != nil {
		err = fmt.Errorf("Movies ID '%s' not found.", req.Title)
		helper.ResponseAPI(c, false, http.StatusNotFound, http.StatusText(http.StatusNotFound), []string{err.Error()}, startTime)
		return err
	}

	timeNow := time.Now()

	movie = models.Movies{
		ID:          idInt,
		Title:       req.Title,
		Description: req.Description,
		Rating:      req.Rating,
		UpdatedAt:   &timeNow,
	}

	exists, _ := s.moviesPostgres.CheckIfExists(movie)
	if exists == true {
		err = fmt.Errorf("title '%s' already exists", req.Title)
		helper.ResponseAPI(c, false, http.StatusConflict, http.StatusText(http.StatusConflict), []string{err.Error()}, startTime)
		return err
	}

	if file != nil {
		newImage, err := helper.UploadImageMovies(c, file)
		if err != nil {
			helper.ResponseAPI(c, false, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), []string{err.Error()}, startTime)
			return err
		}

		getImage := helper.GetImageMovies(c, newImage)
		movie.Image = getImage
	}

	err = s.moviesPostgres.Update(movie)
	if err != nil {
		log.Printf("error movie service Update: %s", err)
		helper.ResponseAPI(c, false, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), []string{err.Error()}, startTime)
		return err
	}

	return nil
}

func (s *Movies) Delete(c *gin.Context, id string, startTime time.Time) error {
	var (
		movie models.Movies
		err   error
	)

	idInt, err := strconv.Atoi(id)
	if err != nil {
		helper.ResponseAPI(c, false, http.StatusBadRequest, http.StatusText(http.StatusBadRequest), []string{err.Error()}, startTime)
		return err

	}

	_, err = s.moviesPostgres.GetByID(idInt)
	if err != nil {
		err = fmt.Errorf("Movies ID '%s' not found.", id)
		helper.ResponseAPI(c, false, http.StatusNotFound, http.StatusText(http.StatusNotFound), []string{err.Error()}, startTime)
		return err
	}

	timeNow := time.Now()
	movie = models.Movies{
		ID:        idInt,
		DeletedAt: &timeNow,
	}

	err = s.moviesPostgres.Delete(movie)
	if err != nil {
		log.Printf("error movie service Delete: %s", err)
		helper.ResponseAPI(c, false, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), []string{err.Error()}, startTime)
		return err
	}

	return nil
}
