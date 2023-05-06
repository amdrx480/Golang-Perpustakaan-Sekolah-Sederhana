package services

import (
	"perpustakaan/models"
	"perpustakaan/repositories"
)

type TransactionService struct {
	repository repositories.TransactionRepository
}

func InitTransactionService() TransactionService {
	return TransactionService{
		repository: &repositories.TransactionRepositoryImpl{},
	}
}

func (ts *TransactionService) GetAll() ([]models.Transaction, error) {
	return ts.repository.GetAll()
}

func (ts *TransactionService) GetByID(id string) (models.Transaction, error) {
	return ts.repository.GetByID(id)

}

func (ts *TransactionService) Create(transactionInput models.TransactionInput) (models.Transaction, error) {
	return ts.repository.Create(transactionInput)

}

func (ts *TransactionService) Update(transactionInput models.TransactionInput, id string) (models.Transaction, error) {
	return ts.repository.Update(transactionInput, id)

}

func (ts *TransactionService) Delete(id string) error {
	return ts.repository.Delete(id)

}

func (ts *TransactionService) Restore(id string) (models.Transaction, error) {
	return ts.repository.Restore(id)

}

func (ts *TransactionService) ForceDelete(id string) error {
	return ts.repository.ForceDelete(id)
}
