package handler

import (
	"net/http"
	"test_dealls/src/entity"

	"github.com/gin-gonic/gin"
)

func (r *rest) createPhoto(ctx *gin.Context) {
	var body entity.Photo
	if err := ctx.ShouldBind(&body); err != nil {
		// Return 400 Bad Request if JSON binding fails
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	// Call use case method to create photo
	user, err := r.uc.Photo.Create(ctx.Request.Context(), body)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Return 200 OK with user data
	ctx.JSON(http.StatusOK, user)
}

func (r *rest) UpdatePhoto(ctx *gin.Context) {
	var body entity.Photo
	if err := ctx.ShouldBind(&body); err != nil {
		// Return 400 Bad Request if JSON binding fails
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	// Call use case method to create photo
	user, err := r.uc.Photo.Update(ctx.Request.Context(), body)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Return 200 OK with user data
	ctx.JSON(http.StatusOK, user)
}
