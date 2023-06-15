package request

type AddOrderRequestDTO struct {
	UserId         uint                            `json:"userId"`
	ProductList    []AddOrderProductListRequestDTO `json:"productList"`
	Note           string                          `json:"note"`
	RecipientName  string                          `json:"recipientName"`
	RecipientPhone string                          `json:"recipientPhone"`
	Email          string                          `json:"email"`
	Address        string                          `json:"address"`
	ShippingMethod string                          `json:"shippingMethod"`
	InvoiceType    string                          `json:"invoiceType"`
}

type AddOrderProductListRequestDTO struct {
	ProductId uint `json:"productId"`
	Quantity  uint `json:"quantity"`
}
