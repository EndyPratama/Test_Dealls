package handler

import (
	"net/http"
	"test_dealls/src/entity"

	"github.com/gin-gonic/gin"
)

func (r *rest) listLikes(ctx *gin.Context) {
	var body entity.Likes
	if err := ctx.ShouldBind(&body); err != nil {
		// Return 400 Bad Request if JSON binding fails
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	// Call use case method to create photo
	user, err := r.uc.Likes.GetList(ctx.Request.Context(), body)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Return 200 OK with user data
	ctx.JSON(http.StatusOK, user)
}

func (r *rest) approve(ctx *gin.Context) {
	var body entity.Likes
	if err := ctx.ShouldBind(&body); err != nil {
		// Return 400 Bad Request if JSON binding fails
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	// Call use case method to create photo
	user, err := r.uc.Likes.Approve(ctx.Request.Context(), body)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Return 200 OK with user data
	ctx.JSON(http.StatusOK, user)
}

func (r *rest) skip(ctx *gin.Context) {
	var body entity.Likes
	if err := ctx.ShouldBind(&body); err != nil {
		// Return 400 Bad Request if JSON binding fails
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	// Call use case method to create photo
	user, err := r.uc.Likes.Skip(ctx.Request.Context(), body)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Return 200 OK with user data
	ctx.JSON(http.StatusOK, user)
}

func (r *rest) DeleteLikes(ctx *gin.Context) {
	var body entity.Likes
	if err := ctx.ShouldBind(&body); err != nil {
		// Return 400 Bad Request if JSON binding fails
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	// Call use case method to create photo
	user, err := r.uc.Likes.Delete(ctx.Request.Context(), body)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Return 200 OK with user data
	ctx.JSON(http.StatusOK, user)
}
