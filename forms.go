package main

type PaymentForm struct {
	UserID       uint   `form:"SubscriptionReferrer"`
	Plan         string `form:"SubscriptionPath"`
	SecurityData string `form:"security_data"`
	SecurityHash string `form:"security_hash"`
}
