package routes

import (
	"net/http"
	"strconv"

	"example.com/anon-project/models"
	"example.com/anon-project/utils"
	"github.com/gin-gonic/gin"
)

func getUsers(c *gin.Context) {
	users, err := models.GetAllUsers()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch data."})
		return
	}

	c.JSON(http.StatusOK, users)
}

func getUser(c *gin.Context) {
	userID, err := strconv.ParseInt(c.Param("id"), 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request"})
		return
	}

	user, err := models.GetUserByID(userID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch data"})
		return
	}

	c.JSON(http.StatusOK, user)
}

func getUserUsername(c *gin.Context) {
	username := c.Param("username")

	user, err := models.GetUserByUsername(username)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch data"})
		return
	}

	c.JSON(http.StatusOK, user)
}

func createUser(c *gin.Context) {
	var user models.User
	err := c.ShouldBindJSON(&user)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Could not bind JSON", "error": err.Error()})
		return
	}

	err = user.Save()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not store data"})
		return
	}

	c.JSON(http.StatusCreated, user)
}

func updateUser(c *gin.Context) {
	userID, err := strconv.ParseInt(c.Param("id"), 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request"})
		return
	}

	_, err = models.GetUserByID(userID)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request"})
		return
	}

	var user models.User
	err = c.ShouldBindJSON(&user)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request"})
		return
	}

	user.ID = userID
	err = user.Update()

	if err != nil {
		match := utils.GetDuplicateError(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not store data", "error": match, "field": match[2], "value": match[1]})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data updated successfully"})
}

func deleteUser(c *gin.Context) {
	userID, err := strconv.ParseInt(c.Param("id"), 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request"})
		return
	}

	user, err := models.GetUserByID(userID)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request"})
		return
	}

	err = user.Delete()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not store data"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data deleted succesfully."})
}

func login(c *gin.Context) {
	var user models.User
	err := c.ShouldBindJSON(&user)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request", "errors": "Wrong password."})
		return
	}

	err = user.ValidateCredentials()

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Could not authorized account", "errors": "Wrong password."})
		return
	}

	token, err := utils.GenerateJWTToken(user.ID, user.Username.String)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Could not authorized account", "errors": "Internal server error. Please wait and try again."})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Login success", "token": token})
}

func signup(c *gin.Context) {
	var user models.User
	err := c.ShouldBindJSON(&user)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Could not bind JSON", "error": err.Error()})
		return
	}

	err = user.Save()

	if err != nil {
		match := utils.GetDuplicateError(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not store data", "error": err.Error(), "field": match[2], "value": match[1]})
		return
	}

	c.JSON(http.StatusCreated, user)
}

func userChangePassword(c *gin.Context) {
	userID, err := strconv.ParseInt(c.Param("id"), 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request"})
		return
	}

	_, err = models.GetUserByID(userID)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request"})
		return
	}

	var user models.User
	err = c.ShouldBindJSON(&user)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request"})
		return
	}

	user.ID = userID
	err = user.UpdatePassword()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not store data", "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data updated successfully"})
}
