package services

import (
	"perpustakaan/models"
	"perpustakaan/repositories"
)

type BookService struct {
	repository repositories.BookRepository
}

func InitBookService() BookService {
	return BookService{
		repository: &repositories.BookRepositoryImpl{},
	}
}

func (bs *BookService) GetAll() ([]models.Book, error) {
	return bs.repository.GetAll()
}
func (bs *BookService) GetByID(id string) (models.Book, error) {
	return bs.repository.GetByID(id)
}
func (bs *BookService) Create(bookInput models.BookInput) (models.Book, error) {
	return bs.repository.Create(bookInput)
}
func (bs *BookService) Update(bookInput models.BookInput, id string) (models.Book, error) {
	return bs.repository.Update(bookInput, id)
}
func (bs *BookService) Delete(id string) error {
	return bs.repository.Delete(id)
}
func (bs *BookService) Restore(id string) (models.Book, error) {
	return bs.repository.Restore(id)
}
func (bs *BookService) ForceDelete(id string) error {
	return bs.repository.ForceDelete(id)
}
