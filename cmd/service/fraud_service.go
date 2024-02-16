package service

import (
	"github.com/fathoor/pemilu-fraud-2024/cmd/entity"
)

type FraudService interface {
	FraudCheck(kota string) []entity.Fraud
	FraudCache(kota string) []entity.Fraud
}
