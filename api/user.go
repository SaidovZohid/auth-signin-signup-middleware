package api

import (
	"net/http"
	"time"

	"github.com/SaidovZohid/auth-signin-signup-middleware/api/models"
	"github.com/SaidovZohid/auth-signin-signup-middleware/pkg/utils"
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
	user := models.User{
		FirstName: req.FirstName,
		LastName: req.LastName,
		Email: req.Email,
		Password: hashedPassword,
	}
	result := sr.Cfg.Create(user)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": result.Error.Error(),
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
		user models.SignInUser
	)
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return 
	}
	result := sr.Cfg.Table("users").Select("password", "email").Where("email = ?", req.Email).Scan(&user)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid email or password!",
		})
		return 
	}
	
	err := utils.CheckPassword(req.Password, user.Password)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid email or password!",
		})
		return 
	}

	token, _, err := utils.CreateToken(req.Email, time.Hour * 24 * 30)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": result.Error.Error(),
		})
		return 
	}

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", token, 3600 * 24 * 30, "", "", false, true)

	c.JSON(http.StatusOK, gin.H{
		"message": "succesfully sign in",
	})
}

func (sr *RouteOptions) Validate(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "You are Logged in",
	})
}