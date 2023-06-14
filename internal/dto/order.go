package dto

type AddOrderDTO struct {
	UserId      uint
	ProductList []struct {
		ProductId uint
		Quantity  uint
	}
	Note           string
	RecipientName  string
	RecipientPhone string
	Address        string
	ShippingMethod string
}