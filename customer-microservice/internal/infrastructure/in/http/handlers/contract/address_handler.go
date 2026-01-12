package contract

import "github.com/gin-gonic/gin"

type AddressHandlerService interface {
	AddressHandleCreate(c *gin.Context)
	AddressHandleGetById(c *gin.Context)
	AddressHandleGetAll(c *gin.Context)
	AddressHandleDelete(c *gin.Context)
}
