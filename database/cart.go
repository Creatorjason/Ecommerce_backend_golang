package database

import (
	"errors"
)

var(
	ErrCantFindProduct = errors.New("can't find the product")
	ErrCantDecodeProducts= errors.New("can't find the product")
	ErrUserIdNotValid = errors.New("this user is not valid")
	ErrCantUpdateUser = errors.New("can't add this product to cart")
	ErrCantRemoveItemCart = errors.New("can't remove item from cart")
	ErrCantGetItem = errors.New("unable to get item from cart")
	ErrCantBuyCartItem = errors.New("unable to update purchase")
)

func AddProductToCart(){}


func RemoveCartItem(){}

func BuyItemFromCart(){}

func InstantBuyer(){}

