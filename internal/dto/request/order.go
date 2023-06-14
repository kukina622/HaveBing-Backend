package request

type AddOrderRequestDTO struct {
	UserId         uint `json:"userId"`
	ProductList    []AddOrderProductListRequestDTO `json:"productList"`
	Note           string `json:"note"`
	RecipientName  string `json:"recipientName"`
	RecipientPhone string `json:"recipientPhone"`
	Address        string `json:"address"`
	ShippingMethod string `json:"shippingMethod"`
}

type AddOrderProductListRequestDTO struct {
	ProductId uint `json:"productId"`
	Quantity  uint `json:"quantity"`
}
