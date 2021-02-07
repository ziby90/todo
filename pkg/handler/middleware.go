package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

const (
	authorizationHeader = "Authorization"
	userCtx             = "userId"
)

func (h *Handler) userIdentity(c *gin.Context) {
	header := c.GetHeader(authorizationHeader)
	if header == "" {
		newErrorResponse(c, http.StatusUnauthorized, `пустой токен авторизации`)
		return
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		newErrorResponse(c, http.StatusUnauthorized, `невалидный токен авторизации`)
		return
	}

	// parse token
	userId, err := h.services.Authorization.ParseToken(headerParts[1])
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
	}

	c.Set(userCtx, userId)
}

func getUserId(c *gin.Context) (uint, error) {
	id, ok := c.Get(userCtx)
	if !ok {
		newErrorResponse(c, http.StatusInternalServerError, `не найден пользователь`)
		return 0, errors.New(`не найден пользователь`)
	}
	idUint, ok := id.(uint)
	if !ok {
		newErrorResponse(c, http.StatusInternalServerError, `не найден пользователь`)
		return 0, errors.New(`не найден пользователь`)
	}
	return idUint, nil
}
