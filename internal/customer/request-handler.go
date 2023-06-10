package customer

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
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name" binding:"required"`
	Email     string `json:"email" binding:"required"`
	Avatar    string `json:"avatar" binding:"required"`
}

func (rh RequestHandler) ShowAll(c *gin.Context) {
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

func (rh RequestHandler) Create(c *gin.Context) {
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

func (rh RequestHandler) Show(c *gin.Context) {
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

func (rh RequestHandler) Update(c *gin.Context) {
	customer := Customer{}
	ID := c.Param("ID")
	if err := c.ShouldBindJSON(&customer); err != nil {
		c.JSON(http.StatusInternalServerError, dto.Response{
			Meta: dto.MetaResponse{
				Success: false,
				Message: err.Error(),
			},
			Data: nil,
		})
		return
	}

	res, err := rh.C.Update(ID, customer)
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

func (rh RequestHandler) Destroy(c *gin.Context) {
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
