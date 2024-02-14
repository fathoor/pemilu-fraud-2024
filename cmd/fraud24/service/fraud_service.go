package service

type FraudService interface {
	FraudCheck(k string) string
}
