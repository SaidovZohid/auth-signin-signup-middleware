package api

import (
	"net/http"
	"time"

	"github.com/SaidovZohid/auth-signin-signup-middleware/api/models"
	"github.com/SaidovZohid/auth-signin-signup-middleware/pkg/utils"
	"github.com/SaidovZohid/auth-signin-signup-middleware/storage/repo"
	"github.com/gin-gonic/gin"
)

func (sr *RouteOptions) SignUp(c *gin.Context) {
	var (
		req models.SignUpUser
	)
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return 
	}

	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	err = sr.Storage.User().Create(&repo.User{
		FirstName: req.FirstName,
		LastName: req.LastName,
		Email: req.Email,
		Password: hashedPassword,
	})
	if err!= nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return 
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Succesfully created",
	})
}

func (sr *RouteOptions) SignIn(c *gin.Context) {
	var (
		req models.SignInUser
	)
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, models.ResponseError{
			Error: err.Error(),
		})
		return 
	}
	user, err := sr.Storage.User().Get(req.Email)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ResponseError{
			Error: err.Error(),
		})
		return 
	}
	if user.ID == 0 {
		c.JSON(http.StatusBadRequest, models.ResponseError{
			Error:  err.Error(),
		})
		return 
	}

	err = utils.CheckPassword(req.Password, user.Password)

	if err != nil {
		c.JSON(http.StatusBadRequest, models.ResponseError{
			Error:  err.Error(),
		})
		return 
	}

	token, _, err := utils.CreateToken(user.FirstName, user.LastName, req.Email, time.Hour * 24 * 30)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ResponseError{
			Error:  err.Error(),
		})
		return 
	}

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", token, 3600 * 24 * 30, "", "", false, true)

	c.JSON(http.StatusOK, models.ResponseOK{
		Message: "Succesfully created.",
	})
}

func (sr *RouteOptions) Validate(c *gin.Context) {
	user, _ := c.Get("user")
	c.JSON(http.StatusOK, gin.H{
		"message": user,
	})
}