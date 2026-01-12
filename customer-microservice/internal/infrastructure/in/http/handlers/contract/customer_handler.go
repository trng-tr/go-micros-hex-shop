package contract

import "github.com/gin-gonic/gin"

// CustomerHandlderService interface that use gin-gonic/gin
type CustomerHandlerService interface {
	CustomerHandleCreate(c *gin.Context)
	CustomerHandleGetById(c *gin.Context)
	CustomerHandleGetAll(c *gin.Context)
	CustomerHandlePatch(c *gin.Context)
	CustomerHandleDelete(c *gin.Context)
}
