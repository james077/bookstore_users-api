package users

import (
	//"encoding/json"
	//"fmt"
	//"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/james077/bookstore_users-api/domain/users"
	"github.com/james077/bookstore_users-api/services"
	"github.com/james077/bookstore_users-api/utils/errors"
)

// CreateUser permite crear un usuario
func CreateUser(c *gin.Context){
	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil{
		restErr := errors.NewBadRequestError("Payload Json Inválido")
		c.JSON(restErr.Status, restErr)
		return
	}
	result, saveErr := services.CreateUser(user)
	if saveErr != nil{
		c.JSON(saveErr.Status, saveErr)
		return
	}
	//fmt.Println(err)
	c.JSON(http.StatusCreated, result)
}

// GetUser permite obtener un usuario
func GetUser(c *gin.Context){
	userId, userErr := strconv.ParseInt(c.Param("user_id"),10,64)
	if userErr != nil{
		err:= errors.NewNotFoundError("Id de usuario inválido")
		c.JSON(err.Status,err)
	}

	user, getErr := services.GetUser(userId)
	if getErr != nil{
		c.JSON(getErr.Status, getErr)
		return
	}
	c.JSON(http.StatusOK, user)
}

// SearchUser permite buscar un usuario
/*func SearchUser(c *gin.Context){
	c.String(http.StatusNotImplemented, "busqueda por id por implementar...")
}*/