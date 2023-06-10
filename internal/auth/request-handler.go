package auth

import (
	"github.com/gin-gonic/gin"
	"mini-project-3-azaz/dto"
	"net/http"
	"strings"
)

type RequestHandler struct {
	C ControllerInterface
}

type RequestHandlerInterface interface {
	Login(c *gin.Context)
	ShowProfile(c *gin.Context)
}

func NewRequestHandler(c ControllerInterface) RequestHandlerInterface {
	return &RequestHandler{
		C: c,
	}
}

type AuthRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (rh *RequestHandler) Login(c *gin.Context) {
	var req AuthRequest

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusInternalServerError, dto.Response{
			Meta: dto.MetaResponse{
				Success: false,
				Message: err.Error(),
			},
			Data: nil,
		})
		return
	}

	res, err := rh.C.Login(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.Response{
			Meta: dto.MetaResponse{
				Success: false,
				Message: err.Error(),
			},
			Data: nil,
		})
		return
	}

	c.JSON(http.StatusOK, dto.Response{
		Meta: dto.MetaResponse{
			Success: true,
			Message: res.Message,
		},
		Data: res.Data,
	})
}

func (rh *RequestHandler) ShowProfile(c *gin.Context) {
	authorization := c.Request.Header["Authorization"]
	if authorization == nil {
		c.JSON(http.StatusUnauthorized, dto.Response{
			Meta: dto.MetaResponse{
				Success: false,
				Message: "Unauthorized",
			},
			Data: nil,
		})
		return
	}

	tokenSigned := strings.Split(authorization[0], " ")[1]
	res, err := rh.C.ShowProfile(tokenSigned)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.Response{
			Meta: dto.MetaResponse{
				Success: false,
				Message: err.Error(),
			},
			Data: nil,
		})
		return
	}

	c.JSON(http.StatusOK, dto.Response{
		Meta: dto.MetaResponse{
			Success: true,
			Message: "",
		},
		Data: res,
	})
}
