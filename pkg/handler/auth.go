package handler

import (
	"fmt"
	"net/http"

	"github.com/Hanqur/todo_app"
	"github.com/gin-gonic/gin"
)

func (h *Handler) signUp(c *gin.Context) {
	var input todo_app.User

	fmt.Println(input)

	// ловим JSON-структуру и упаковываем в структуру
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	fmt.Println(input)

	// запускаем метод CreateUser у сервиса, получаем id созданного юзера
	id, err := h.services.Authorization.CreateUser(input)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	// подготавливаем JSON с ответом
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

type signInInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (h *Handler) signIn(c *gin.Context) {
	var input signInInput

	fmt.Println(input)

	// ловим JSON-структуру и упаковываем в структуру
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	fmt.Println(input)

	
	token, err := h.services.Authorization.GenerateToken(input.Username, input.Password)
	if err != nil {
		
		newErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	// подготавливаем JSON с ответом
	c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})
}
