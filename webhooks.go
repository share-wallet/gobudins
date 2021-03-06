package gobudins

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type WebhooksListeners struct {
	OnUserCreated       func(User)
	OnAccountSynced     func(SyncedAccount)
	OnAccountDisabled   func(SyncedAccount)
	OnConnectionDeleted func(Connection)
}

func (ctrl *Controller) SetupRoutesGin(e *gin.RouterGroup) {
	e.POST("/user/created", ctrl.whUserCreated)
	e.POST("/accounts/synced", ctrl.whAccountSynced)
	e.POST("/accounts/disabled", ctrl.whAccountDisabled)
	e.POST("/connection/deleted", ctrl.whConnectionDeleted)

}

func (ctrl *Controller) whUserCreated(c *gin.Context) {
	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		fmt.Printf("Failed to unmarshal created user: %s\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	ctrl.listeners.OnUserCreated(user)
}

func (ctrl *Controller) whAccountSynced(c *gin.Context) {
	var account SyncedAccount
	if err := c.ShouldBindJSON(&account); err != nil {
		fmt.Printf("Failed to unmarshal synced account: %s\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	ctrl.listeners.OnAccountSynced(account)
}

func (ctrl *Controller) whAccountDisabled(c *gin.Context) {
	var account SyncedAccount
	if err := c.ShouldBindJSON(&account); err != nil {
		fmt.Printf("Failed to unmarshal synced account: %s\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	ctrl.listeners.OnAccountDisabled(account)
}

func (ctrl *Controller) whConnectionDeleted(c *gin.Context) {
	var connection Connection
	if err := c.ShouldBindJSON(&connection); err != nil {
		fmt.Printf("Failed to unmarshal connection: %s\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	ctrl.listeners.OnConnectionDeleted(connection)
}
