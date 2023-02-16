package models

import (
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
)


type User struct{
	ID  			primitive.ObjectID   `json:"_id" bson:"_id"`
	FirstName 		*string				`json:"first_name"  validate:"required, min=2, max=20"`
	LastName 		*string				`json:"last_name"   validate:"required, min=2, max=20"`
	Password		*string				`json:"password"    validate:"required, min=6"`
	Email 			*string				`json:"email"       validate:"email, required"`	
	Phone 			*string				`json:"phone"       validate:"required"`
	Token 			*string				`json:"token"`
	Refresh_Token	*string				`json:"refresh_token"`	
	Created_At		time.Time			`json:"created_at"`	
	Updated_At		time.Time			`json:"updated_at"`
	User_ID			string				`json:"user_id"`	
	UserCart		[]ProductUser		`json:"product_user" bson:"product_user"`
	Address_Details []Address			`json:"address" bson:"address"`
	Order_Status	[] Order			`json:"order" bson:"order"`
}


type Product struct{
	Product_ID			primitive.ObjectID			`bson:"_id"`
	Product_Name		*string						`json:"product_name"`	
	Price				*uint64						`json:"price"`
	Rating				*uint8						`json:"rating"`	
	Image				*string						`json:"image"`

}


type ProductUser struct{ 
	Product_ID			primitive.ObjectID			`bson:"_id"`
	Product_Name		*string						`json:"product_name" bson:"product_name"`
	Price				*uint64						`json:"price" bson:"price"`
	Rating				*uint8						`json:"rating" bson"rating"`
	Image				*string						`json:"image" bson:"image"`

}

type Address struct{
	Address_id         primitive.ObjectID          	`bson:"_id"`
	House			   *string						`json:"house_name" bson:"house_name"`
	Street			   *string						`json:"street" bson:"street"`	
	City			   *string						`json:"city" bson:"city"`
	PinCode            *string						`json:"pin_code" bson:"pin_code"`	

}

type Order struct{
	Order_ID		primitive.ObjectID			`bson:"_id`
	Order_Cart		[]ProductUser				`json:"order_list" bson:"order_list"`	
	Ordered_At		time.Time					`json:"ordered_at" bson:"ordered_at"`
	Price			*uint64						`json:"total_price" bson:"total_price"`
	Discount		*uint8						`json:"discount" bson:"discount"`
	Payment_Method  Payment						`json:"payment_method" bson:"payment_method"`

}

type Payment struct{	
	Digital 		bool						`json:"digital" bson:"digital"`
	CashOnDelivery  bool						`json:"COD" bson"COD"`	
}