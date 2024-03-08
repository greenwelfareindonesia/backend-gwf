package payment

import (
	"fmt"
	orders "greenwelfare/order"
)

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

type Service interface {
	DoPayment(req *SubmitPaymentRequest) error
}

type service struct {
	gw    Gateway
	order orders.RepositoryOrder
}

func NewService(repo orders.RepositoryOrder, gateway Gateway) Service {
	return &service{
		gw:    gateway,
		order: repo,
	}
}

func (s *service) DoPayment(req *SubmitPaymentRequest) error {
	or, err := s.order.FindById()
	if err != nil {
		return err
	}

	or.Status = "PROCESSING"
	s.order.Update(or)

	resp, err := s.gw.SubmitPayment(nil)
	if err != nil {
		return err
	}

	if resp.Status != "SUCCESS" {
		or.Status == "FAILED"
		s.order.Update(or)
		return fmt.Errorf("failed to do payment")
	}

	or.Status = "PENDING"
	s.order.Update(or)

	return nil
}
