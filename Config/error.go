package Config

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Controller func(*gin.Context) (int, interface{}, error)
type Wrapper func(Controller) gin.HandlerFunc

// error handler
func ProvideAPIWrap() Wrapper {
	return func(controller Controller) gin.HandlerFunc {
		return func(ctx *gin.Context) {
			status, body, err := controller(ctx)
			if err != nil {
				//	logger.Errorw("error in handler", "error", err)
				ctx.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
				return
			}

			ctx.JSON(status, body)
		}
	}
}
