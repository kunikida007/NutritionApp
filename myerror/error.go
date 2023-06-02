package myerror

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

// BadRequestError HTTP Status Code: 400
type BadRequestError struct {
	Err error
	Msg string
}

func (e *BadRequestError) Error() string {
	// Errを返さない、もしくは補足が必要な場合Msgを利用
	if e.Msg == "" {
		return fmt.Sprintf("%s", e.Err)
	}
	if e.Err == nil {
		return e.Msg
	}
	return fmt.Sprintf("%s, %s", e.Err, e.Msg)
}

// UnauthorizedError HTTP Status Code: 401
type UnauthorizedError struct {
	Err error
	Msg string
}

func (e *UnauthorizedError) Error() string {
	if e.Msg == "" {
		return fmt.Sprintf("%s", e.Err)
	}
	if e.Err == nil {
		return e.Msg
	}
	return fmt.Sprintf("%s, %s", e.Err, e.Msg)
}

// NotFoundError HTTP Status Code: 404
type NotFoundError struct {
	Err error
	Msg string
}

func (e *NotFoundError) Error() string {
	if e.Msg == "" {
		return fmt.Sprintf("%s", e.Err)
	}
	if e.Err == nil {
		return e.Msg
	}
	return fmt.Sprintf("%s, %s", e.Err, e.Msg)
}

// InternalServerError HTTP Status Code: 500
type InternalServerError struct {
	Msg string
	Err error
}

func (e *InternalServerError) Error() string {
	if e.Msg == "" {
		return fmt.Sprintf("%s", e.Err)
	}
	if e.Err == nil {
		return e.Msg
	}
	return fmt.Sprintf("%s, %s", e.Err, e.Msg)
}

func JSONErrorHandler(err error, c echo.Context) {
	if err == nil {
		return
	}
	switch e := err.(type) {
	case *BadRequestError:
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": e.Error(),
		})
	case *UnauthorizedError:
		c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"error": e.Error(),
		})
	case *NotFoundError:
		c.JSON(http.StatusNotFound, map[string]interface{}{
			"error": e.Error(),
		})
	case *InternalServerError:
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": e.Error(),
		})
	default:
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
	}
}
