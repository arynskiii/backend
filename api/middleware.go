package api

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"simple_bank/token"
	"strings"
)

const (
	authorizationHeaderKey  = "authorization"
	authorizationTypeBearer = "bearer"
	authorizationPayloadKey = "authorization_payload"
)

func authMiddleware(tokenMaker token.Maker) gin.HandlerFunc {
	return func(context *gin.Context) {
		authorizationHeader := context.GetHeader(authorizationHeaderKey)
		fmt.Println(authorizationHeader)
		if len(authorizationHeader) == 0 {
			err := errors.New("authorization head is not provided")
			context.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(err))
			return
		}
		fields := strings.Fields(authorizationHeader)
		fmt.Println(fields)
		if len(fields) < 2 {
			err := errors.New("invalid authorization header format")
			context.AbortWithStatusJSON(401, errorResponse(err))
			return
		}
		authorizationType := strings.ToLower(fields[0])
		if authorizationType != authorizationTypeBearer {
			err := fmt.Errorf("unsupported authorization typs %s", authorizationType)
			context.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(err))
			return
		}
		accessToken := fields[1]

		payload, err := tokenMaker.VerifyToken(accessToken)
		fmt.Println(payload)
		if err != nil {
			context.AbortWithStatusJSON(401, errorResponse(err))
			return
		}
		context.Set(authorizationPayloadKey, payload)
		context.Next()

	}
}
