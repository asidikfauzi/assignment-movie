package movies

import (
	"assignment-movie/common/helper"
	"assignment-movie/common/validator"
	"assignment-movie/models"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"log"
	"net/http"
	"strconv"
	"time"
)

func (m *MasterMovies) GetAll(c *gin.Context) {
	startTime := time.Now()

	search := c.Query("search")
	pageParam := c.Query("page")
	limitParam := c.Query("limit")
	orderByParam := c.Query("orderBy")

	dataCustomer, paginate, err := m.MoviesService.GetAll(c, pageParam, limitParam, orderByParam, search, startTime)
	if err != nil {
		log.Printf("error movies controller GetAll :%s", err)
		return
	}

	helper.ResponseDataPaginationAPI(c, true, http.StatusOK, http.StatusText(http.StatusOK), []string{helper.SuccessGetData}, dataCustomer, paginate, startTime)
	return
}

func (m *MasterMovies) GetByID(c *gin.Context) {
	startTime := time.Now()

	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Printf("Error parsing id: %s", err)
		return
	}

	dataCustomer, err := m.MoviesService.GetByID(c, id, startTime)
	if err != nil {
		log.Printf("error movies controller GetByID :%s", err)
		return
	}

	helper.ResponseDataAPI(c, true, http.StatusOK, http.StatusText(http.StatusOK), []string{helper.SuccessGetData}, dataCustomer, startTime)
	return
}

func (m *MasterMovies) Create(c *gin.Context) {
	startTime := time.Now()

	file, err := c.FormFile("image")
	if err != nil {
		if err == http.ErrMissingFile {
			helper.ResponseAPI(c, false, http.StatusUnprocessableEntity, http.StatusText(http.StatusUnprocessableEntity), []string{"Field 'image'" + helper.RequiredMessage}, startTime)
			return
		} else {
			helper.ResponseAPI(c, false, http.StatusBadRequest, http.StatusText(http.StatusBadRequest), []string{err.Error()}, startTime)
			return
		}
	}

	if file.Size > helper.MaxFileSize {
		helper.ResponseAPI(c, false, http.StatusUnprocessableEntity, http.StatusText(http.StatusUnprocessableEntity), []string{helper.MaxFileSizeMessage}, startTime)
		return
	}

	var req models.ReqMovie
	if err = c.ShouldBindWith(&req, binding.Form); err != nil {
		helper.ResponseAPI(c, false, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), []string{err.Error()}, startTime)
		return
	}

	validate := validator.ValidatorMessage(req)
	if len(validate) > 0 {
		helper.ResponseAPI(c, false, http.StatusUnprocessableEntity, http.StatusText(http.StatusUnprocessableEntity), validate, startTime)
		return
	}

	err = m.MoviesService.Create(c, file, req, startTime)
	if err != nil {
		log.Printf("error movies controller Create :%s", err)
		return
	}

	helper.ResponseAPI(c, true, http.StatusCreated, http.StatusText(http.StatusCreated), []string{helper.SuccessCreatedData}, startTime)
	return
}

func (m *MasterMovies) Update(c *gin.Context) {
	startTime := time.Now()

	id := c.Param("id")
	file, _ := c.FormFile("image")

	if file != nil {
		if file.Size > helper.MaxFileSize {
			helper.ResponseAPI(c, false, http.StatusUnprocessableEntity, http.StatusText(http.StatusUnprocessableEntity), []string{helper.MaxFileSizeMessage}, startTime)
			return
		}
	}

	var req models.ReqMovie
	if err := c.ShouldBindWith(&req, binding.Form); err != nil {
		helper.ResponseAPI(c, false, http.StatusUnprocessableEntity, http.StatusText(http.StatusUnprocessableEntity), []string{err.Error()}, startTime)
		return
	}

	validate := validator.ValidatorMessage(req)
	if len(validate) > 0 {
		helper.ResponseAPI(c, false, http.StatusUnprocessableEntity, http.StatusText(http.StatusUnprocessableEntity), validate, startTime)
		return
	}

	err := m.MoviesService.Update(c, id, file, req, startTime)
	if err != nil {
		log.Printf("error movies controller Update :%s", err)
		return
	}

	helper.ResponseAPI(c, true, http.StatusOK, http.StatusText(http.StatusOK), []string{helper.SuccessUpdatedData}, startTime)
	return
}

func (m *MasterMovies) Delete(c *gin.Context) {
	startTime := time.Now()

	id := c.Param("id")
	err := m.MoviesService.Delete(c, id, startTime)
	if err != nil {
		log.Printf("error movies controller Delete :%s", err)
		return
	}

	helper.ResponseAPI(c, true, http.StatusOK, http.StatusText(http.StatusOK), []string{helper.SuccessDeletedData}, startTime)
	return
}
