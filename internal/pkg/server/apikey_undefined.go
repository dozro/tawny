package server

import (
	"time"

	apiError2 "github.com/dozro/tawny/pkg/apiError"
	"github.com/gin-gonic/gin"
)

func apikeyUndefined(apikey string, c *gin.Context) bool {
	if apikey == "" {
		render(c, 401, apiError2.ApiError{
			HttpCode:          401,
			InternalErrorCode: apiError2.MissingApiKeyInRequest,
			InternalErrorMsg:  apiError2.MissingApiKeyInRequest.String(),
			Message:           "apikey is required",
			Success:           false,
			Date:              time.Now().String(),
			Data:              c.Request,
		})
		return true
	}
	return false
}
