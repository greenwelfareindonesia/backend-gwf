package payment

import (
	"fmt"
	"greenwelfare/order"
	"strconv"
)

type SubmitPaymentRequest struct {
	OrderID      string
	Amount       int64
	BankTransfer string
	PaymentType  string
}

type Response struct {
	Status string
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
	order order.RepositoryOrder
}

func NewService(repo order.RepositoryOrder, gateway Gateway) Service {
	return &service{
		gw:    gateway,
		order: repo,
	}
}

func (s *service) DoPayment(req *SubmitPaymentRequest) error {
	orderIDConv, _ := strconv.Atoi(req.OrderID)

	findOrderID, err := s.order.FindById(orderIDConv)
	if err != nil {
		return err
	}

	findOrderID.Status = "PROCESSING"
	req.OrderID = strconv.Itoa(findOrderID.ID)
	req.Amount = int64(findOrderID.TotalPrice)
	// req.BankTransfer =

	_, err = s.order.Update(findOrderID)
	if err != nil {
		return err
	}

	resp, err := s.gw.SubmitPayment(req)
	if err != nil {
		return err
	}

	if resp.Status != "SUCCESS" {
		findOrderID.Status = "FAILED"
		s.order.Update(findOrderID)
		return fmt.Errorf("failed to do payment")
	}

	findOrderID.Status = "PENDING"
	s.order.Update(findOrderID)

	return nil
}
