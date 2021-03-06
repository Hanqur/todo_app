package handler

import (
	"net/http"
	"strconv"

	"github.com/Hanqur/todo_app"
	"github.com/gin-gonic/gin"
)

func (h *Handler) createTag(c *gin.Context) {
	
	userId, err := getUserId(c)
	if err != nil {
		return
	}
	
	itemId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
	}

	var input todo_app.Tag
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Tag.CreateTag(userId, itemId, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

type getAllTagsResponse struct {
	Data []todo_app.Tag `json:"data"`
}

func (h *Handler) getAllTags(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	itemId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
	}

	tags, err := h.services.Tag.GetAllTags(userId, itemId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAllTagsResponse{
		Data: tags,
	})

}