package payment

type SubmitPaymentRequest struct {
	OrderID string
	Amount  int64
	Dest    string
}

type Response struct {
}

type CompletePaymentRequest struct {
}

type CheckPaymentStatusRequest struct {
	OrderID string
}

type Gateway interface {
	SubmitPayment(req *SubmitPaymentRequest) (*Response, error)
	CompletePayment(req *CompletePaymentRequest) (*Response, error)
	CheckPaymentStatus(req *CheckPaymentStatusRequest) (*Response, error)
}
