package server

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GithubWebhookPullRequestOpened struct {
	Action string `json:"action" binding:"required"`
}

func Webhook(c *gin.Context) {
	var body GithubWebhookPullRequestOpened
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"ok":    false,
			"error": err.Error(),
		})
		return

	}

	fmt.Println("✅ ", body)

	c.Status(http.StatusOK)
}
