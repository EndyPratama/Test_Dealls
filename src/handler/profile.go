package handler

import (
	"net/http"
	"test_dealls/src/entity"

	"github.com/gin-gonic/gin"
)

// Untuk cari yg match
func (r *rest) SearchPeople(ctx *gin.Context) {
	var body entity.Profile
	if err := ctx.ShouldBind(&body); err != nil {
		// Return 400 Bad Request if JSON binding fails
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	// Call use case method to create profile
	user, err := r.uc.Profile.GetListPeople(ctx.Request.Context(), body)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Return 200 OK with user data
	ctx.JSON(http.StatusOK, user)
}

func (r *rest) detailProfile(ctx *gin.Context) {
	var body entity.Profile
	if err := ctx.ShouldBind(&body); err != nil {
		// Return 400 Bad Request if JSON binding fails
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	// Call use case method to create profile
	user, err := r.uc.Profile.GetDetail(ctx.Request.Context(), body)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Return 200 OK with user data
	ctx.JSON(http.StatusOK, user)
}

func (r *rest) createProfile(ctx *gin.Context) {
	var body entity.Profile
	if err := ctx.ShouldBind(&body); err != nil {
		// Return 400 Bad Request if JSON binding fails
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	// Call use case method to create profile
	user, err := r.uc.Profile.Create(ctx.Request.Context(), body)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Return 200 OK with user data
	ctx.JSON(http.StatusOK, user)
}

func (r *rest) UpdateProfile(ctx *gin.Context) {
	var body entity.Profile
	if err := ctx.ShouldBind(&body); err != nil {
		// Return 400 Bad Request if JSON binding fails
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	// Call use case method to create profile
	user, err := r.uc.Profile.Update(ctx.Request.Context(), body)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Return 200 OK with user data
	ctx.JSON(http.StatusOK, user)
}

func (r *rest) UpgradeProfile(ctx *gin.Context) {
	var body entity.Profile
	if err := ctx.ShouldBind(&body); err != nil {
		// Return 400 Bad Request if JSON binding fails
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	// Call use case method to create profile
	user, err := r.uc.Profile.Upgrade(ctx.Request.Context(), body)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Return 200 OK with user data
	ctx.JSON(http.StatusOK, user)
}
