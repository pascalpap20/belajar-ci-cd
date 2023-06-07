package user

import (
	"crud/dto"
	"crud/repository"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

type RequestHandlerUser struct {
	ctr ControllerUser
}

func NewUserRequestHandler(
	dbCrud *gorm.DB,
) RequestHandlerUser {
	return RequestHandlerUser{
		ctr: controllerUser{
			userUseCase: useCaseUser{
				userRepo: repository.NewUser(dbCrud),
			},
		}}
}

func (h RequestHandlerUser) CreateUser(c *gin.Context) {
	request := UserParam{}
	err := c.Bind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.DefaultBadRequestResponse())
		return
	}
	res, err := h.ctr.CreateUser(request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.DefaultErrorResponse())
		return
	}
	c.JSON(http.StatusOK, res)
}

func (h RequestHandlerUser) GetUserById(c *gin.Context) {
	request := UserParam{}
	err := c.BindQuery(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.DefaultBadRequestResponse())
		return
	}
	userId, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.DefaultBadRequestResponse())
		return
	}

	res, err := h.ctr.GetUserById(uint(userId))
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.DefaultErrorResponse())
		return
	}
	c.JSON(http.StatusOK, res)
}

func (h RequestHandlerUser) UpdateUser(c *gin.Context) {
	request := UserParam{}
	err := c.BindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.DefaultBadRequestResponse())
		return
	}
	userId, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.DefaultBadRequestResponse())
		return
	}

	res, err := h.ctr.UpdateUser(request, uint(userId))
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.DefaultErrorResponse())
		return
	}
	c.JSON(http.StatusOK, res)
}

func (h RequestHandlerUser) DeleteUser(c *gin.Context) {
	email := c.Param("email")
	res, err := h.ctr.DeleteUser(email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.DefaultErrorResponse())
		return
	}
	c.JSON(http.StatusOK, res)
}
