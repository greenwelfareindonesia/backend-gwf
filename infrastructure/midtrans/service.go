package midtrans

import (
	"fmt"
	"github.com/midtrans/midtrans-go"
	orders "greenwelfare/order"
	"greenwelfare/payment"
)
import "github.com/midtrans/midtrans-go/coreapi"

var (
	listOfBank = []midtrans.Bank{
		midtrans.BankBca,
		midtrans.BankBri,
		midtrans.BankCimb,
	}
)

type midtransGateway struct {
	client *coreapi.Client
	order  orders.RepositoryOrder
}

type Config struct {
	ServerKey string
}

func NewMidtransGateway(cfg *Config) (payment.Gateway, error) {
	var client *coreapi.Client
	client.New(cfg.ServerKey, midtrans.Sandbox)
	return &midtransGateway{
		client: client,
	}, nil
}

// SubmitPayment call charge api
func (m *midtransGateway) SubmitPayment(req *payment.SubmitPaymentRequest) (*payment.Response, error) {
	var chosenBank midtrans.Bank
	for _, bank := range listOfBank {
		if string(bank) == req.Dest {
			chosenBank = bank
		}
	}

	if chosenBank == "" {
		return nil, fmt.Errorf("unsupported bank")
	}

	resp, err := m.client.ChargeTransaction(&coreapi.ChargeReq{
		PaymentType: coreapi.PaymentTypeBankTransfer,
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  req.OrderID,
			GrossAmt: req.Amount,
		},
		BankTransfer: &coreapi.BankTransferDetails{
			Bank: chosenBank,
		},
	})

	if err != nil {
		return nil, err
	}

	return mapChargeToResponse(resp), nil
}

func (m *midtransGateway) CheckPaymentStatus(req *payment.CheckPaymentStatusRequest) (*payment.Response, error) {
	resp, err := m.client.CheckTransaction(req.OrderID)

	if err != nil {
		return nil, err
	}

	return mapStatusResponseToResponse(resp), nil
}

func (m *midtransGateway) CompletePayment(req *payment.CompletePaymentRequest) (*payment.Response, error) {
	//update db status based on request status
	return nil, nil
}

func mapChargeToResponse(resp *coreapi.ChargeResponse) *payment.Response {
	return nil
}

func mapStatusResponseToResponse(resp *coreapi.TransactionStatusResponse) *payment.Response {
	return nil
}
