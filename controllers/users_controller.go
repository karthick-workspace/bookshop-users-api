package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/karthick-workspace/bookshop-users-api/domain/users"
	"github.com/karthick-workspace/bookshop-users-api/services"
	"github.com/karthick-workspace/bookshop-users-api/utils/errors"
	"net/http"
	"strconv"
)

func GetUser(c *gin.Context) {
	userId, userErr := strconv.ParseInt(c.Param("user_id"), 10, 64)

	if userErr != nil {
		err := errors.NewBadRequestError("User id should be number")
		c.JSON(err.Status, err)
		return
	}

	user, getErr := services.GetUser(userId)

	if getErr != nil {
		c.JSON(getErr.Status, getErr)
	}

	c.JSON(http.StatusOK, user)
}

func CreateUser(c *gin.Context) {
	var user users.User
	//bytes, err := ioutil.ReadAll(c.Request.Body)
	//if err != nil {
	//	// TODO: Handle error
	//	return
	//}
	//if err := json.Unmarshal(bytes, &user); err != nil {
	//	fmt.Println(err.Error())
	//	//TODO:Handle json error
	//	return
	//}

	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}

	result, saveErr := services.CreateUser(user)
	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		return
	}
	fmt.Println(result)
	c.JSON(http.StatusCreated, result)
}

func SearchUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "SearchUser:Implement me")
}
