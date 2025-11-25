package server

import (
	"github.com/dozro/tawny/internal/pkg/apiError"
	"github.com/gin-gonic/gin"
)

func apikeyUndefined(apikey string, c *gin.Context) bool {
	if apikey == "" {
		c.JSON(401, apiError.ApiError{
			HttpCode:          401,
			InternalErrorCode: apiError.MissingApiKeyInRequest,
			Message:           "apikey is required",
			Success:           false,
		})
		return true
	}
	return false
}
