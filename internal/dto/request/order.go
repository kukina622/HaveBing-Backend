package request

type AddOrderRequestDTO struct {
	UserId               uint                            `json:"userId"`
	ProductList          []AddOrderProductListRequestDTO `json:"productList"`
	Note                 string                          `json:"note"`
	RecipientName        string                          `json:"recipientName"`
	RecipientPhone       string                          `json:"recipientPhone"`
	ExpectedDeliveryDate string                          `json:"expectedDeliveryDate"`
	Email                string                          `json:"email"`
	Address              string                          `json:"address"`
	ShippingMethod       string                          `json:"shippingMethod"`
	PaymentMethod        string                          `json:"paymentMethod"`
	InvoiceType          string                          `json:"invoiceType"`
}

type AddOrderProductListRequestDTO struct {
	ProductId uint `json:"productId"`
	Quantity  uint `json:"quantity"`
}
