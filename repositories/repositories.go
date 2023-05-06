package repositories

import "perpustakaan/models"

type BookRepository interface {
	GetAll() ([]models.Book, error)
	GetByID(id string) (models.Book, error)
	Create(bookInput models.BookInput) (models.Book, error)
	Update(bookInput models.BookInput, id string) (models.Book, error)
	Delete(id string) error
	Restore(id string) (models.Book, error)
	ForceDelete(id string) error
}

type MemberRepository interface {
	GetAll() ([]models.Member, error)
	GetByID(id string) (models.Member, error)
	Create(memberInput models.MemberInput) (models.Member, error)
	Update(memberInput models.MemberInput, id string) (models.Member, error)
	Delete(id string) error
	Restore(id string) (models.Member, error)
	ForceDelete(id string) error
}

type AdminRepository interface {
	Register(adminInput models.AdminInput) (models.Admin, error)
	GetByEmail(adminInput models.AdminInput) (models.Admin, error)
}

type TransactionRepository interface {
	GetAll() ([]models.Transaction, error)
	GetByID(id string) (models.Transaction, error)
	Create(transactionInput models.TransactionInput) (models.Transaction, error)
	Update(transactionInput models.TransactionInput, id string) (models.Transaction, error)
	Delete(id string) error
	Restore(id string) (models.Transaction, error)
	ForceDelete(id string) error
}
