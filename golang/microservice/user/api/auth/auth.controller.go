package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthController interface {
	RegisterHandler(*gin.Context)
	LoginHandler(*gin.Context)
	ReadAllHandler(*gin.Context)
}

type authController struct {
	AuthService AuthService
}

func NewAuthController(authService AuthService) AuthController {
	return &authController{
		AuthService: authService,
	}
}

func (ac authController) RegisterHandler(ctx *gin.Context) {
	var authDTO AuthDTO
	if err := ctx.Bind(&authDTO); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	auth, err := ac.AuthService.Register(ctx.Request.Context(), authDTO)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"response": auth,
	})
}
func (ac authController) LoginHandler(ctx *gin.Context) {
	var authDTO AuthDTO
	if err := ctx.Bind(&authDTO); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	auth, err := ac.AuthService.Login(ctx.Request.Context(), authDTO)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"response": auth,
	})
}
func (ac authController) ReadAllHandler(ctx *gin.Context) {

	auths, err := ac.AuthService.ReadAll(ctx.Request.Context())
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"response": auths,
	})
}
