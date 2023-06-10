package auth

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type Controller struct {
	UC UseCaseInterface
}

type ControllerInterface interface {
	Login(body *AuthRequest) (*ProfileItemResponse, error)
	ShowProfile(tokenSigned string) (*ProfileItem, error)
	VerifyToken(tokenSigned string) (*JWTClaim, error)
}

func NewController(uc UseCaseInterface) ControllerInterface {
	return &Controller{
		UC: uc,
	}
}

type ProfileItem struct {
	ID       uint   `json:"id"`
	FullName string `json:"full_name"`
	Username string `json:"username"`
}

type ProfileItemWithToken struct {
	ProfileItem
	Token string `json:"token"`
}

type ProfileItemResponse struct {
	Message string                `json:"message"`
	Data    *ProfileItemWithToken `json:"data"`
}

func (c *Controller) Login(body *AuthRequest) (*ProfileItemResponse, error) {
	payload := Payload{
		Username: body.Username,
		Password: body.Password,
	}
	user, err := c.UC.Login(&payload)
	if err != nil {
		return nil, err
	}

	res := &ProfileItemResponse{
		Message: "login successfully",
		Data: &ProfileItemWithToken{
			ProfileItem: ProfileItem{
				ID:       user.ID,
				FullName: user.FullName,
				Username: user.Username,
			},
			Token: user.Token,
		},
	}

	return res, nil
}

func (c *Controller) ShowProfile(tokenSigned string) (*ProfileItem, error) {
	user, err := c.UC.ShowProfile(tokenSigned)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, fmt.Errorf("user not found")
	}

	return &ProfileItem{
		ID:       user.ID,
		FullName: user.FullName,
		Username: user.Username,
	}, nil
}

func (c *Controller) VerifyToken(tokenSigned string) (*JWTClaim, error) {
	token, err := jwt.Parse(tokenSigned, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method")
		}

		return []byte(JWT_KEY), nil
	})
	if err != nil {
		return nil, err
	}

	if token.Valid {
		claims, _ := token.Claims.(jwt.MapClaims)

		return &JWTClaim{
			Username: claims["Username"].(string),
			RegisteredClaims: jwt.RegisteredClaims{
				Issuer:    "jwt-token",
				ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * 3600)),
			},
		}, nil
	} else {
		return nil, fmt.Errorf("invalid authorization token")
	}
}
