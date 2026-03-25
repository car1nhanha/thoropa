package handler

import (
	"context"
	"fmt"
	"net/http"
	"thoropa/internal/model"
	"thoropa/internal/service"
	"time"

	"github.com/gin-gonic/gin"
	gonanoid "github.com/matoous/go-nanoid/v2"
)

type LinkHandler struct {
	service *service.LinkService
}

func NewLinkHandler(s *service.LinkService) *LinkHandler {
	return &LinkHandler{service: s}
}

func (h *LinkHandler) Create(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), 3*time.Second)
	defer cancel()
	id, _ := gonanoid.New()

	link := model.Link{
		Id:        id,
		Ip:        c.ClientIP(),
		CreatedAt: time.Now().Unix(),
		Accesses:  0,
		Original:  "aa",
	}

	fmt.Println("ANTES DO DYNAMO")
	err := h.service.Create(ctx, &link)
	fmt.Println("DEPOIS DO DYNAMO")

	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, link)
}
