package transactions

import (
	"errors"

	"greenwelfare/external"
	"greenwelfare/product"
	"greenwelfare/user"
)

type ServiceOrder interface {
	CreateOrder(getIDProduct GetIDProduct, GetUserID int, input OrderInput) (Order, error)
	GetOrder(productID int, userID int) ([]Order, error)
	// MarkAsPaid()
	Pay(req *PaymentRequest) (*PaymentResponse, error)
	// SaveToOrder(userID int, OrderID int) (Orderdetail.OrderDetails, error)
}

type PaymentRequest struct {
}

type PaymentResponse struct {
}

type serviceOrder struct {
	repository        RepositoryOrder
	repositoryProduct product.RepositoryProduct
	repositoryUser    user.Repository
	paymentGateway    external.PaymentGateway
}

func NewServiceOrder(repository RepositoryOrder, repositoryProduct product.RepositoryProduct, repositoryUser user.Repository, paymentGateway external.PaymentGateway) *serviceOrder {
	return &serviceOrder{repository, repositoryProduct, repositoryUser, paymentGateway}
}

func (s *serviceOrder) CreateOrder(getIDProduct GetIDProduct, GetUserID int, input OrderInput) (Order, error) {
	order := Order{}

	cek, err := s.repositoryProduct.FindById(getIDProduct.ID)
	if err != nil {
		print("erorrrrrr")
		return Order{}, errors.New("item not found")
	}

	if cek.Stock < input.Quantity {
		return Order{}, errors.New("quantity not enough")
	}

	cekSaldo, err := s.repositoryUser.FindById(GetUserID)
	if err != nil {
		return Order{}, err
	}

	totalBelanja, err := s.repositoryProduct.FindById(getIDProduct.ID)
	if err != nil {
		return Order{}, err
	}
	cek.Price = input.Quantity * totalBelanja.Price

	cek.Stock = cek.Stock - input.Quantity

	order.Quantity = input.Quantity
	order.TotalPrice = cek.Price
	order.ProductID = cek.ID
	order.UserID = cekSaldo.ID

	newProduct, err := s.repository.Save(order)
	if err != nil {
		return newProduct, err
	}

	return newProduct, nil

}

func (s *serviceOrder) GetOrder(productID int, userID int) ([]Order, error) {

	order, err := s.repository.FindByUserId(productID, userID)
	if err != nil {
		return order, err
	}
	return order, nil

}
