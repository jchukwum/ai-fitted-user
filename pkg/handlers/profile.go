package handlers

import (
	"log"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/jchukwum/ai-fitted-user/pkg/model"
	"github.com/jchukwum/ai-fitted-user/pkg/repository"
	uuid "github.com/satori/go.uuid"
)

// GetProfile returns UserProfile
func GetProfile(c *gin.Context) {
	userID := getUserIDFromSession(c)
	p, err := repository.GetProfile(userID)
	if err != nil {
		newProfile := model.Profile{}
		newProfile.UserID = userID
		c.JSON(http.StatusOK, newProfile)
		return
	}
	c.JSON(http.StatusOK, p)
}

// SaveProfile saves UserProfile
func SaveProfile(c *gin.Context) {
	var p model.Profile
	c.BindJSON(&p)
	if p.UserID == "" {
		log.Println("ERROR: Cannot save profile as was unable to find UserID from session")
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorised",
		})
		c.Abort()
		return
	}
	repository.SaveProfile(&p)
	c.JSON(http.StatusOK, &p)
}

func createGuestSession(c *gin.Context) string {
	session := sessions.Default(c)
	userID := uuid.NewV4().String()
	session.Set("userID", userID)
	session.Save()
	log.Println("INFO: Created guest session:", userID)
	return userID
}

func getUserIDFromSession(c *gin.Context) string {
	session := sessions.Default(c)
	userID := session.Get("userID")
	if userID != nil {
		return userID.(string)
	}
	log.Println("INFO: Unable to find an active session. Creating one")
	return createGuestSession(c)
}
