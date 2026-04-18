package handlers

import (
	"context"
	"time"
	"tutorial/repository"

	"github.com/gin-gonic/gin"
)

type RoleHandler struct {
	repo *repository.RoleRepository
}

func NewRoleHandler(repo *repository.RoleRepository) *RoleHandler {
	return &RoleHandler{repo: repo}
}

func (controller *RoleHandler) GetRole(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)

	defer cancel()

	role, err := controller.repo.GetRoles(ctx)

	if err != nil {
		c.JSON(404, gin.H{"error": "Something Wrong..."})
		return
	}

	c.JSON(200, role)

}
