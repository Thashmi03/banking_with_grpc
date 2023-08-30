package netxddalmodels

import "time"

type Customer struct {
	CustomerId         int64     `json:"customer_id" bson:"customer_id" binding:"required"`
	Firstname          string    `json:"firstname" bson:"firstname" binding:"required"`
	Lastname           string    `json:"lastname" bson:"lastname" binding:"required,min=8"`
	BankId   		   string    `json:"bank_id" bson:"bank_id,omitempty" binding:"required"`
	Balance            string    `json:"balance" bson:"balance"`
	IsActive		   bool		 `json:"is_active" bson:"is_active"`
	CreatedAt          time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt          time.Time `json:"updated_at" bson:"updated_at"`
}

type DbResponse struct{
	CustomerId         int64     `json:"customer_id" bson:"customer_id" binding:"required"`
	CreatedAt          time.Time `json:"created_at" bson:"created_at"`
}