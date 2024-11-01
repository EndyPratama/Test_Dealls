package handler

import (
	"net/http"
	"test_dealls/src/entity"

	"github.com/gin-gonic/gin"
)

func (r *rest) Login(ctx *gin.Context) {
	var body entity.User
	if err := ctx.ShouldBind(&body); err != nil {
		// Return 400 Bad Request if JSON binding fails
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	// Call use case method to create profile
	user, token, err := r.uc.User.Login(ctx.Request.Context(), body)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Set the token in the response header
	ctx.Header("secret", token)
	ctx.JSON(http.StatusOK, user)
}

func (r *rest) RegisterUser(ctx *gin.Context) {
	var body entity.User
	if err := ctx.ShouldBind(&body); err != nil {
		// Return 400 Bad Request if JSON binding fails
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	// Call use case method to create profile
	user, err := r.uc.User.Register(ctx.Request.Context(), body)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, user)
}
