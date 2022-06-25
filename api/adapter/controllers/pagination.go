package controllers

import (
	"api/domain"
	"math"
	"strconv"

	"github.com/labstack/echo/v4"
)

// NOTE: pageを定義しているところに移す
func CreateToPage(c echo.Context, totalCount int) domain.Page {
	page, _ := strconv.Atoi(c.QueryParam("page"))
	perPage, _ := strconv.Atoi(c.QueryParam("perPage"))

	totalPages := int(math.Ceil(float64(totalCount) / float64(perPage)))
	return domain.Page{Page: page, PerPage: perPage, TotalCounts: totalCount, TotalPages: totalPages}
}
