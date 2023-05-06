package repositories

//database lokalan tutorial
import (
	"perpustakaan/database"
	"perpustakaan/models"

	"gorm.io/gorm"
)

type BookRepositoryImpl struct{}

func InitBookRepository() BookRepository {
	return &BookRepositoryImpl{}
}

// method terdapat receiver
func (br *BookRepositoryImpl) GetAll() ([]models.Book, error) {
	//membuat var dengan slice untuk mengambil semua data book
	var books []models.Book

	//operasi query find untuk mengambil semua data dari database dan menyimpannya ke variable books
	err := database.DB.Find(&books).Error

	//melakukan pengecekan
	if err != nil {
		return nil, err
	}

	return books, nil
}

func (br *BookRepositoryImpl) GetByID(id string) (models.Book, error) {
	//karena getbyid mengambil 1 data maka tidak menggunakan slice
	var book models.Book

	//method first untuk mencari record pertama kali berdasarkan kondisinya
	err := database.DB.First(&book, "id = ?", id).Error

	if err != nil {
		return models.Book{}, err
	}

	return book, nil
}

func (br *BookRepositoryImpl) Create(bookInput models.BookInput) (models.Book, error) {
	var createdBook models.Book = models.Book{
		Author:     bookInput.Author,
		Title:      bookInput.Title,
		Publisher:  bookInput.Publisher,
		FiscalYear: bookInput.FiscalYear,
		Isbn:       bookInput.Isbn,
		Qty:        bookInput.Qty,
		Rack:       bookInput.Rack,
	}

	result := database.DB.Create(&createdBook)

	if err := result.Error; err != nil {
		return models.Book{}, err
	}

	//mendapatkan data yang baru ditambahkan
	err := database.DB.Last(&createdBook).Error

	if err != nil {
		return models.Book{}, err
	}

	return createdBook, nil
}

func (br *BookRepositoryImpl) Update(bookInput models.BookInput, id string) (models.Book, error) {

	// validasi.. ongoing tutorial

	book, err := br.GetByID(id)

	if err != nil {
		return models.Book{}, err
	}

	book.Author = bookInput.Author
	book.Title = bookInput.Title
	book.Publisher = bookInput.Publisher
	book.FiscalYear = bookInput.FiscalYear
	book.Isbn = bookInput.Isbn
	book.Qty = bookInput.Qty
	book.Rack = bookInput.Rack

	//method save untuk menyimpan data ke database
	err = database.DB.Save(&book).Error

	if err != nil {
		return models.Book{}, err
	}

	return book, nil
}

func (br *BookRepositoryImpl) Delete(id string) error {
	book, err := br.GetByID(id)

	if err != nil {
		return err
	}

	err = database.DB.Delete(&book).Error

	if err != nil {
		return err
	}

	return nil
}

func (br *BookRepositoryImpl) Restore(id string) (models.Book, error) {
	var trashedBook models.Book

	//method unscoped untuk mengakses semua data yang ada di database, baik dengan deleted_at ada maupun null
	err := database.DB.Unscoped().First(&trashedBook, "id = ?", id).Error

	if err != nil {
		return models.Book{}, err
	}

	trashedBook.DeletedAt = gorm.DeletedAt{}

	err = database.DB.Unscoped().Save(&trashedBook).Error

	if err != nil {
		return models.Book{}, err
	}

	return trashedBook, nil

}

// mengapus data secara permanen
func (br *BookRepositoryImpl) ForceDelete(id string) error {
	book, err := br.GetByID(id)

	if err != nil {
		return err
	}

	err = database.DB.Unscoped().Delete(&book).Error

	if err != nil {
		return err
	}

	return nil
}
