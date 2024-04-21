package http

import (
	"runtime"

	"github.com/labstack/echo/v4"
)

func NewRecoverMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) (returnErr error) {
		defer func() {
			if p := recover(); p != nil {
				stack := make([]byte, 1<<16)
				stack = stack[:runtime.Stack(stack, false)]
				returnErr = &PanicError{
					Panic: p,
					Stack: stack,
				}
			}
		}()
		return next(c)
	}
}
