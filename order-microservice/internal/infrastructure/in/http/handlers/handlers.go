package handlers

import (
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/trng-tr/order-microservice/internal/application/in"
	"github.com/trng-tr/order-microservice/internal/infrastructure/in/http/dtos"
	"github.com/trng-tr/order-microservice/internal/infrastructure/in/http/mappers"
)

// OrderHandlerImpl implement OrderHandler
type OrderHandlerImpl struct {
	inOrderSvc          in.InOrderService
	inOrderLineSvc      in.InOrderLineService
	inRemoteCustomerSvc in.RemoteCustomerService
	inRemoteProdSvc     in.RemoteProductService
}

// NewOrderHandlerImpl DI by constuctor
func NewOrderHandlerImpl(inOrderSvc in.InOrderService, inOrderLineSvc in.InOrderLineService,
	inRemoteCustomerSvc in.RemoteCustomerService, inRemoteProdSvc in.RemoteProductService) *OrderHandlerImpl {
	return &OrderHandlerImpl{
		inOrderSvc:          inOrderSvc,
		inOrderLineSvc:      inOrderLineSvc,
		inRemoteCustomerSvc: inRemoteCustomerSvc,
		inRemoteProdSvc:     inRemoteProdSvc,
	}
}

// HandleCreateOrder implement interface
func (o *OrderHandlerImpl) HandleCreateOrder(ctx *gin.Context) {
	var request dtos.OrderRequest

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, dtos.NewResponse(fail, err.Error()))
		return
	}

	var OrderRequest = dtos.OrderRequest{
		CustomerID: request.CustomerID,
	}
	order, err := o.inOrderSvc.CreateOrder(ctx.Request.Context(), OrderRequest.CustomerID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, dtos.NewResponse(fail, err.Error()))
		return
	}
	var orderLinesResponses []dtos.OrderLineResponse
	for _, line := range request.OrderLines {
		ordeLine, err := o.inOrderLineSvc.CreateOrderLine(ctx.Request.Context(), order.ID, mappers.ToBusinessOrderLine(line))
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, dtos.NewResponse(fail, err.Error()))
			return
		}
		product, err := o.inRemoteProdSvc.GetRemoteProductByID(ctx.Request.Context(), ordeLine.ProductID)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, dtos.NewResponse(fail, err.Error()))
			return
		}
		orderLinesResponses = append(orderLinesResponses, mappers.ToOrderLineResponse(ordeLine, product))
	}
	customer, err := o.inRemoteCustomerSvc.GetRemoteCustomerByID(ctx.Request.Context(), order.CustomerID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, dtos.NewResponse(fail, err.Error()))
		return
	}

	var orderResponse dtos.OrderResponse = dtos.OrderResponse{
		ID:                  order.ID,
		CustomerResponse:    mappers.ToCustomerResponse(customer),
		Status:              string(order.Status),
		OrderLinesResponses: orderLinesResponses,
		CreatedAt:           order.CreatedAt.Format(time.RFC3339),
	}

	ctx.JSON(http.StatusOK, orderResponse)
}

// HandleGetOrderByID implement interface
func (o *OrderHandlerImpl) HandleGetOrderByID(ctx *gin.Context) {
	var idStr string = ctx.Param("id")
	if strings.TrimSpace(idStr) == "" {
		ctx.JSON(http.StatusBadRequest, dtos.NewResponse(fail, errInvalidParams))
		return
	}

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, dtos.NewResponse(fail, err.Error()))
		return
	}

	order, err := o.inOrderSvc.GetOrderByID(ctx.Request.Context(), id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, dtos.NewResponse(fail, err.Error()))
		return
	}

	customer, err := o.inRemoteCustomerSvc.GetRemoteCustomerByID(ctx, order.CustomerID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, dtos.NewResponse(fail, err.Error()))
		return
	}
	orderLines, err := o.inOrderLineSvc.GetOrderLinesByOrderID(ctx, order.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, dtos.NewResponse(fail, err.Error()))
		return
	}

	var orderLinesResponses []dtos.OrderLineResponse
	for _, line := range orderLines {
		product, err := o.inRemoteProdSvc.GetRemoteProductByID(ctx, line.ProductID)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, dtos.NewResponse(fail, err.Error()))
			return
		}
		orderLinesResponses = append(orderLinesResponses, mappers.ToOrderLineResponse(line, product))
	}
	var orderResponse = dtos.OrderResponse{
		ID:                  order.ID,
		CustomerResponse:    mappers.ToCustomerResponse(customer),
		Status:              string(order.Status),
		OrderLinesResponses: orderLinesResponses,
		CreatedAt:           order.CreatedAt.Format(time.RFC3339),
	}

	ctx.JSON(http.StatusOK, orderResponse)
}
