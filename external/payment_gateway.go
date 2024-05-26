package external

type PaymentGateway interface {
	SubmitPayment(req *PaymentRequest) (*PaymentResponse, error)
}

type PaymentRequest struct{}

type PaymentResponse struct{}
