package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

const (
	authorizationHeader = "Authorization"
)

func (h *Handler) JWTAuthAdminMiddleware() gin.HandlerFunc {
	//header := c.GetHeader(authorizationHeader)
	//if header == "" {
	//	NewErrorResponse(c, http.StatusUnauthorized, "provide authorization credentials")
	//	return
	//}
	//
	//headerParts := strings.Split(header, " ")
	//if len(headerParts) != 2 {
	//	NewErrorResponse(c, http.StatusUnauthorized, "invalid authorization credentials")
	//	return
	//}
	//
	//adminRole, err := h.services.Authorization.ParseAdminToken(headerParts[1])
	//if err != nil {
	//	NewErrorResponse(c, http.StatusUnauthorized, err.Error())
	//	return
	//}
	//
	//c.Set(adminCtx, adminRole)

	return func(c *gin.Context) {
		var token string
		tQ := c.Query("token")
		if tQ != "" {
			token = tQ
		}
		tB := strings.Split(c.Request.Header.Get(authorizationHeader), " ")
		if len(tB) == 2 {
			token = tB[1]
		}

		if token == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, basicResponse{"No token provided"})
			return
		}

		err := h.services.Authorization.AdminTokenValid(token)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized,
				basicResponse{
					"Unauthorized",
				})
			return
		}
		c.Next()
	}
}
