package middleware

import (
	"fmt"
	"log"

	"github.com/golimix/go-learn/framework/gin/casbin/component"
	"github.com/casbin/casbin"
	gormadapter "github.com/casbin/gorm-adapter"
	"github.com/gin-gonic/gin"
)

func Authorize(obj string, act string, adapter *gormadapter.Adapter) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get current user/subject
		val, existed := c.Get("current_subject")
		if !existed {
			c.AbortWithStatusJSON(401, component.RestResponse{Message: "user hasn't logged in yet"})
			return
		}
		// Casbin enforces policy
		ok, err := enforce(val.(string), obj, act, adapter)
		if err != nil {
			log.Println(err)
			c.AbortWithStatusJSON(500, component.RestResponse{Message: "error occurred when authorizing user"})
			return
		}
		if !ok {
			c.AbortWithStatusJSON(403, component.RestResponse{Message: "forbidden"})
			return
		}
		c.Next()
	}
}

func enforce(sub string, obj string, act string, adapter *gormadapter.Adapter) (bool, error) {
	// Load model configuration file and policy store adapter
	enforcer, err := casbin.NewEnforcer("config/rbac_model.conf", adapter)
	if err != nil {
		return false, fmt.Errorf("failed to create casbin enforcer: %w", err)
	}
	// Load policies from DB dynamically
	err = enforcer.LoadPolicy()
	if err != nil {
		return false, fmt.Errorf("failed to load policy from DB: %w", err)
	}
	// Verify
	ok, err := enforcer.Enforce(sub, obj, act)
	return ok, err
}
