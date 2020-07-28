package users

import (
	"github.com/alvarezcarlos/bookstore_users-api/domain/users"
	"github.com/alvarezcarlos/bookstore_users-api/services"
	"github.com/alvarezcarlos/bookstore_users-api/utils/date"
	"github.com/alvarezcarlos/bookstore_users-api/utils/errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func CreateUser(c *gin.Context) {
	var user users.User
	user.DateCreated = date.GetNowString()
	user.Status = "active"
	user.Password = "1234"
	//bytes, err := ioutil.ReadAll(c.Request.Body)
	//fmt.Println(err)
	//fmt.Println("this is body " + string(bytes))
	//if err := json.Unmarshal(bytes, &user); err != nil {
	//	fmt.Println(err)
	//	return
	//}

	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.BadRequestError("invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}


	result, saveErr := services.CreateUser(user)
	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		return
	}

	c.JSON(http.StatusCreated, result)
	return
}

func GetUser(c *gin.Context) {
	userId, userErr := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if userErr != nil{
		err := errors.BadRequestError("invalid Id")
		c.JSON(err.Status, err)
		return
	}

	user, err := services.GetUser(userId)

	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusOK, user)
}

func FindByStatus(c *gin.Context) {
	status := c.Query("status")
	user := users.User{}
	result, err := user.FindByStatus(status)
	if err != nil {
		res := errors.InternalServerError(err.Message)
		c.JSON(res.Status, res)
	}
	c.JSON(http.StatusOK, result)
}
