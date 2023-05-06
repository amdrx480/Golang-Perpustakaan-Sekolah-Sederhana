package repositories

import (
	"perpustakaan/database"
	"perpustakaan/models"

	"gorm.io/gorm"
)

type TransactionRepositoryImpl struct {
}

func InitTransactionRepository() TransactionRepository {
	return &TransactionRepositoryImpl{}
}

func (tr *TransactionRepositoryImpl) GetAll() ([]models.Transaction, error) {
	var transactions []models.Transaction

	//Preload digunakan untuk mendapatkan data realtionnya
	err := database.DB.Preload("Book").Preload("Member").Find(&transactions).Error
	if err != nil {
		return nil, err
	}

	return transactions, nil
}

func (tr *TransactionRepositoryImpl) GetByID(id string) (models.Transaction, error) {
	var transaction models.Transaction

	//penyebab gagal saat get by id adalah tidak menambahkan parameter
	err := database.DB.Preload("Book").Preload("Member").First(&transaction, "id = ?", id).Error

	if err != nil {
		return models.Transaction{}, err
	}

	return transaction, nil

}

func (tr *TransactionRepositoryImpl) Create(transactionInput models.TransactionInput) (models.Transaction, error) {
	//membuat terlebih dahulu variable untuk input data ke database dan mengecek hasilnya
	var transaction models.Transaction = models.Transaction{
		BookID:    transactionInput.BookId,
		MemberID:  transactionInput.MemberID,
		CreatedAt: transactionInput.CreatedAt,
	}

	//menambahkakan data baru menggunakan create
	result := database.DB.Create(&transaction)

	if err := result.Error; err != nil {
		return models.Transaction{}, err
	}

	//last untuk mengecek data yang terakhir kali ditambahkan
	err := result.Preload("Book").Preload("Member").Last(&transaction).Error

	if err != nil {
		return models.Transaction{}, err
	}

	return transaction, nil

}

func (tr *TransactionRepositoryImpl) Update(transactionInput models.TransactionInput, id string) (models.Transaction, error) {
	transaction, err := tr.GetByID(id)

	if err != nil {
		return models.Transaction{}, err
	}

	transaction.BookID = transactionInput.BookId
	transaction.MemberID = transactionInput.MemberID

	err = database.DB.Save(&transaction).Error

	if err != nil {
		return models.Transaction{}, err
	}

	return transaction, nil
}

func (tr *TransactionRepositoryImpl) Delete(id string) error {
	transaction, err := tr.GetByID(id)

	if err != nil {
		return err
	}

	err = database.DB.Delete(&transaction).Error

	if err != nil {
		return err
	}

	return nil

}

func (tr *TransactionRepositoryImpl) Restore(id string) (models.Transaction, error) {
	var trashedtransaction models.Transaction

	//method unscoped untuk mengakses semua data yang ada di database, baik dengan deleted_at ada maupun null
	err := database.DB.Unscoped().First(&trashedtransaction, "id = ?", id).Error

	if err != nil {
		return models.Transaction{}, err
	}

	trashedtransaction.DeletedAt = gorm.DeletedAt{}

	err = database.DB.Unscoped().Save(&trashedtransaction).Error

	if err != nil {
		return models.Transaction{}, err
	}

	return trashedtransaction, nil
}

func (tr *TransactionRepositoryImpl) ForceDelete(id string) error {
	transaction, err := tr.GetByID(id)

	if err != nil {
		return err
	}

	err = database.DB.Unscoped().Delete(&transaction).Error

	if err != nil {
		return err
	}

	return nil
}
