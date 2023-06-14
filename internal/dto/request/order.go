package request

type AddOrderRequestDTO struct {
	UserId      uint `json:"userId"`
	ProductList []struct {
		ProductId uint `json:"productId"`
		Quantity  uint `json:"quantity"`
	}
	Note           string `json:"note"`
	RecipientName  string `json:"recipientName"`
	RecipientPhone string `json:"recipientPhone"`
	Address        string `json:"address"`
	ShippingMethod string `json:"shippingMethod"`
}
