package middleware

import (
	"context"
	"fmt"
	"test_dealls/src/business/usecase/user"
	"test_dealls/src/utils/appcontext"
	"test_dealls/src/utils/config"
	"test_dealls/src/utils/log"
	"time"

	// "github.com/anekapay/go-sdk/codes"
	// "github.com/anekapay/go-sdk/errors"
	// "github.com/anekapay/go-sdk/codes"
	// "github.com/anekapay/go-sdk/errors"
	// "github.com/anekapay/kios-svc/src/utils/constants"
	// "github.com/anekapay/kios-svc/src/utils/keys"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

const (
	KeyRequestID string = "x-request-id"
)

func DeallsMiddleware(ucUser user.Interface, proxyCfg config.Application, log log.Interface) gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("General")
		getTime := func() time.Time {
			location, err := time.LoadLocation("Asia/Jakarta")
			if err != nil {
				log.Error(c, fmt.Sprintf("Failed Load Location Time, %v", err))
				return time.Now().UTC()
			}
			return time.Now().UTC().In(location)
		}

		ctx := context.WithValue(c.Request.Context(), "ucUser", ucUser)

		headers := c.Request.Header.Clone()
		reqid := headers.Get(KeyRequestID)
		if reqid == "" {
			reqid = uuid.New().String()
		}

		// Add request ID and start time to the context
		ctx = appcontext.SetRequestId(ctx, reqid)
		ctx = appcontext.SetRequestStartTime(ctx, getTime())

		// Update the context in the request
		c.Request = c.Request.WithContext(ctx)

		c.Next()
	}
}
