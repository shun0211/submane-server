package utils

import (
	"api/dto"
	"math"
	"strconv"

	"github.com/labstack/echo/v4"
)

func ConvertToPage(c echo.Context, totalCount int) dto.Page {
	page, _ := strconv.Atoi(c.QueryParam("page"))
	perPage, _ := strconv.Atoi(c.QueryParam("perPage"))

	totalPages := int(math.Ceil(float64(totalCount) / float64(perPage)))
	return dto.Page{Page: page, PerPage: perPage, TotalCounts: totalCount, TotalPages: totalPages}
}
