package movies

import (
	"assignment-movie/common/helper"
	"github.com/gin-gonic/gin"
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
		log.Printf("error movies controller GetAllMovies :%s", err)
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
		log.Printf("error movies controller GetMovieByID :%s", err)
		return
	}

	helper.ResponseDataAPI(c, true, http.StatusOK, http.StatusText(http.StatusOK), []string{helper.SuccessGetData}, dataCustomer, startTime)
	return
}
