package response

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type DefaultResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

type Pagination struct {
	CurrentPage int     `json:"current_page"`
	PerPage     int     `json:"per_page"`
	TotalPage   float64 `json:"total_page"`
	TotalData   int64   `json:"total_data"`
}

func ResponseSuccess(c echo.Context, data interface{}) error {
	resp := DefaultResponse{
		Success: true,
		Message: "Success",
		Data:    data,
	}

	return c.JSON(http.StatusOK, resp)
}

func ResponseFail(c echo.Context, data interface{}) error {
	resp := DefaultResponse{
		Success: false,
		Message: "Fail",
		Data:    data,
	}

	return c.JSON(http.StatusOK, resp)
}
