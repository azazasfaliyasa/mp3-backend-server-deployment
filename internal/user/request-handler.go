package user

import (
	"github.com/gin-gonic/gin"
	"mini-project-3-azaz/dto"
	"net/http"
)

type RequestHandler struct {
	C ControllerInterface
}

type RequestHandlerInterface interface {
	ShowAll(c *gin.Context)
	Create(c *gin.Context)
	Show(c *gin.Context)
	Update(c *gin.Context)
	Destroy(c *gin.Context)
}

func NewRequestHandler(c ControllerInterface) RequestHandlerInterface {
	return &RequestHandler{
		C: c,
	}
}

type CreateRequest struct {
	FullName string `json:"full_name" binding:"required"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (rh *RequestHandler) ShowAll(c *gin.Context) {
	res, err := rh.C.ShowAll()
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

func (rh *RequestHandler) Create(c *gin.Context) {
	var req CreateRequest
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

	res, err := rh.C.Create(&req)
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

func (rh *RequestHandler) Show(c *gin.Context) {
	ID := c.Param("ID")
	res, err := rh.C.Show(ID)
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

func (rh *RequestHandler) Update(c *gin.Context) {
	user := User{}
	ID := c.Param("ID")
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusInternalServerError, dto.Response{
			Meta: dto.MetaResponse{
				Success: false,
				Message: err.Error(),
			},
			Data: nil,
		})
		return
	}

	res, err := rh.C.Update(ID, user)
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

func (rh *RequestHandler) Destroy(c *gin.Context) {
	ID := c.Param("ID")
	res, err := rh.C.Destroy(ID)
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
