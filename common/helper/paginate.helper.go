package helper

import (
	"fmt"
	"strconv"
)

type Paginate struct {
	Page       int   `json:"page"`
	Limit      int   `json:"limit"`
	TotalPages int   `json:"total_pages"`
	TotalData  int64 `json:"total_data"`
}

func Pagination(pageParam, limitParam string) (int, int, int, error) {

	if pageParam == "0" || pageParam == "" {
		pageParam = "1"
	}
	if limitParam == "0" || limitParam == "" {
		limitParam = "10"
	}

	page, err := strconv.Atoi(pageParam)
	if err != nil {
		err = fmt.Errorf("invalid value for 'page': %s", err.Error())
		return 0, 0, 0, err
	}

	limit, err := strconv.Atoi(limitParam)
	if err != nil {
		err = fmt.Errorf("invalid value for 'limit': %s", err.Error())
		return 0, 0, 0, err
	}

	offset := (page - 1) * limit

	return page, limit, offset, nil
}
