package config

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Controller func(*gin.Context) (int, interface{}, error)
type Wrapper func(Controller) gin.HandlerFunc

// error handler
func WrapperAPI() Wrapper {
	return func(controller Controller) gin.HandlerFunc {
		return func(ctx *gin.Context) {
			status, body, err := controller(ctx)
			if err != nil {
				log.Println("error in handler", err)
				ctx.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
				return
			}

			ctx.JSON(status, body)
		}
	}
}
