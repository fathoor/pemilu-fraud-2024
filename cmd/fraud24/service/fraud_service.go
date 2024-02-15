package service

import "github.com/fathoor/fraud24/cmd/fraud24/entity"

type FraudService interface {
	FraudCheck() []entity.Fraud
	FraudCheckCache() []entity.Fraud
}
