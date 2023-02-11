package service

import (
	"a21hc3NpZ25tZW50/database"
	"a21hc3NpZ25tZW50/entity"
	"errors"
)

// Service is package for any logic needed in this program

type ServiceInterface interface {
	AddCart(productName string, quantity int) error
	RemoveCart(productName string) error
	ShowCart() ([]entity.CartItem, error)
	ResetCart() error
	GetAllProduct() ([]entity.Product, error)
	Pay(money int) (entity.PaymentInformation, error)
}

type Service struct {
	database database.DatabaseInterface
}

func NewService(database database.DatabaseInterface) *Service {
	return &Service{
		database: database,
	}
}

func (s *Service) AddCart(productName string, quantity int) error {
	product, err := s.database.GetProductByName(productName)
	if err != nil {
		return err
	}

	if quantity <= 0 {
		return errors.New("invalid quantity")
	}

	cart, err := s.database.GetCartItems()
	if err != nil {
		return err
	}

	found := false
	for i := 0; i < len(cart); i++ {
		if cart[i].ProductName == productName {
			cart[i].Quantity += quantity
			found = true
		}
	}

	if !found {
		cart = append(cart, entity.CartItem{
			ProductName: productName,
			Price: product.Price,
			Quantity: quantity,
		})
	}

	err = s.database.SaveCartItems(cart)
	if err != nil {
		return err
	}
	
	return nil // TODO: replace this
}

func (s *Service) RemoveCart(productName string) error {
	cart, err := s.database.GetCartItems()
	if err != nil {
		return err
	}

	found := false
	newCart := []entity.CartItem{}

	for i := 0; i < len(cart); i++ {
		if cart[i].ProductName == productName {
			found = true
			continue
		} else {
			newCart = append(newCart, cart[i])
		}
	}

	if !found {
		return errors.New("product not found")
	}

	err = s.database.SaveCartItems(newCart)
	if err != nil {
		return err
	}
	
	return nil // TODO: replace this
}

func (s *Service) ShowCart() ([]entity.CartItem, error) {
	cart, err := s.database.GetCartItems()
	if err != nil {
		return nil, err
	}

	return cart, nil
}

func (s *Service) ResetCart() error {
	err := s.database.SaveCartItems([]entity.CartItem{})
	if err != nil {
		return err
	}
	
	return nil // TODO: replace this
}

func (s *Service) GetAllProduct() ([]entity.Product, error) {
	return s.database.GetProductData(), nil // TODO: replace this
}

func (s *Service) Pay(money int) (entity.PaymentInformation, error) {
	cart, err := s.database.GetCartItems()
	if err != nil {
		return entity.PaymentInformation{}, err
	}

	var total int
	for i := 0; i < len(cart); i++ {
		total += (cart[i].Price * cart[i].Quantity)
	}

	if money < total {
		return entity.PaymentInformation{}, errors.New("money is not enough")
	}

	var paymentInfo = entity.PaymentInformation {
		ProductList: cart,
		TotalPrice: total,
		MoneyPaid: money,
		Change: money - total,
	}
	
	err = s.ResetCart()
	if err != nil {
		return entity.PaymentInformation{}, err
	}

	return paymentInfo, nil // TODO: replace this
}
