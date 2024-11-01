package middleware

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"test_dealls/src/business/usecase/user"
	"test_dealls/src/entity"
	"test_dealls/src/utils/appcontext"

	"github.com/gin-gonic/gin"
)

func CheckToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Retrieve ucUser from context
		ucUser, ok := c.Request.Context().Value("ucUser").(user.Interface)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized access - user context missing"})
			c.Abort()
			return
		}

		fmt.Println("Check Token")
		token := c.GetHeader("Authorization")
		arr := strings.Split(token, " ")
		if len(arr) != 2 || arr[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token format"})
			c.Abort()
			return
		}

		// Decode the token payload
		userIdAndEmail, err := base64.StdEncoding.DecodeString(arr[1])
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token encoding"})
			c.Abort()
			return
		}

		// Extract userId and email
		values := strings.Split(string(userIdAndEmail), ":")
		if len(values) != 2 {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token payload invalid"})
			c.Abort()
			return
		}
		userId, email := values[0], values[1]

		// Convert userId to integer
		idInt, err := strconv.Atoi(userId)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "User ID conversion failed"})
			c.Abort()
			return
		}

		// Retrieve user details
		_, err = ucUser.GetDetail(c.Request.Context(), entity.User{ID: int64(idInt), Email: email})
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "User validation failed"})
			c.Abort()
			return
		}

		// Continue to the next middleware/handler
		fmt.Println(appcontext.GetRequestId(c.Request.Context()))

		// Add request ID and start time to the context
		ctx := appcontext.SetUserIDAgent(c.Request.Context(), userId)

		// Update the context in the request
		c.Request = c.Request.WithContext(ctx)

		c.Next()
	}
}
