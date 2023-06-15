package response

import (
	"HaveBing-Backend/internal/domain"

	"github.com/shopspring/decimal"
)

type OrderResponseDTO struct {
	OrderId     uint                   `json:"orderId"`
	OrderNumber string                 `json:"orderNumber"`
	Status      string                 `json:"status"`
	UserId      uint                   `json:"userId"`
	Note        string                 `json:"note"`
	OrderDate   string                 `json:"orderDate"`
	Payment     domain.Payment         `json:"payment"`
	Shipping    domain.Shipping        `json:"shipping"`
	OrderItem   []OrderItemResponseDTO `json:"orderItem"`
}

type OrderItemResponseDTO struct {
	Quality uint            `json:"quality"`
	Price   decimal.Decimal `json:"price"`
	Product struct {
		ID              uint     `json:"productId"`
		ProductName     string   `json:"productName"`
		ProductImage    []string `json:"productImage"`
		ProductCategory string   `json:"productCategory"`
	} `json:"product"`
}

func NewOrderResponseDTO[in *domain.Order | []domain.Order](_order in) any {
	switch order := any(_order).(type) {
	case *domain.Order:
		result := OrderResponseDTO{
			OrderId:     order.ID,
			OrderNumber: order.OrderNumber,
			Status:      string(order.Status),
			UserId:      order.UserId,
			Note:        order.Note,
			OrderDate:   order.OrderDate.Format("2006-01-02"),
			Payment:     order.Payment,
			Shipping:    order.Shipping,
			OrderItem:   NewOrderItemDTO(order.OrderItem),
		}
		return result
	case []domain.Order:
		result := []OrderResponseDTO{}
		for _, _order := range order {
			result = append(result, OrderResponseDTO{
				OrderId:     _order.ID,
				OrderNumber: _order.OrderNumber,
				Status:      string(_order.Status),
				UserId:      _order.UserId,
				Note:        _order.Note,
				OrderDate:   _order.OrderDate.Format("2006-01-02"),
				Payment:     _order.Payment,
				Shipping:    _order.Shipping,
				OrderItem:   NewOrderItemDTO(_order.OrderItem),
			})
		}
		return result
	}
	return nil
}

func NewOrderItemDTO(_orderItem []domain.OrderItem) []OrderItemResponseDTO {
	var result []OrderItemResponseDTO
	for _, orderItem := range _orderItem {
		result = append(result, OrderItemResponseDTO{
			Quality: orderItem.Quality,
			Price:   orderItem.Price,
			Product: struct {
				ID              uint     `json:"productId"`
				ProductName     string   `json:"productName"`
				ProductImage    []string `json:"productImage"`
				ProductCategory string   `json:"productCategory"`
			}{
				ID:              orderItem.Product.ID,
				ProductName:     orderItem.Product.ProductName,
				ProductImage:    getProductImagePath(orderItem.Product.ProductImage),
				ProductCategory: orderItem.Product.ProductCategory.CategoryName,
			},
		})
	}
	return result
}
