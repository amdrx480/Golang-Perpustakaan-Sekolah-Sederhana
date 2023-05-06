package repositories

import (
	"perpustakaan/database"
	"perpustakaan/models"

	"golang.org/x/crypto/bcrypt"
)

type AdminRepositoryImpl struct {
}

func InitAdminRepository() AdminRepository {
	return &AdminRepositoryImpl{}
}

func (ar *AdminRepositoryImpl) Register(adminInput models.AdminInput) (models.Admin, error) {
	//melakukan enkripsi pada password
	password, err := bcrypt.GenerateFromPassword([]byte(adminInput.Password), bcrypt.DefaultCost)

	if err != nil {
		return models.Admin{}, err
	}

	adminInput.Password = string(password)

	var admin models.Admin = models.Admin{
		Email:    adminInput.Email,
		Password: string(password),
	}

	//menyimpan data user ke database
	result := database.DB.Create(&admin)

	if err := result.Error; err != nil {
		return models.Admin{}, err
	}

	// mengembalikan data yang terakhir kali ditambahkan
	err = result.Last(&admin).Error
	if err != nil {
		return models.Admin{}, err
	}

	return admin, nil
}

func (ar *AdminRepositoryImpl) GetByEmail(adminInput models.AdminInput) (models.Admin, error) {
	var admin models.Admin

	//first untuk mendapatkan data yang pertama kali
	err := database.DB.First(&admin, "email = ?", adminInput.Email).Error

	if err != nil {
		return models.Admin{}, err
	}

	//melakukan pengecekan apakah email sama dengan password
	err = bcrypt.CompareHashAndPassword([]byte(admin.Password), []byte(adminInput.Password))

	if err != nil {
		return models.Admin{}, err
	}

	return admin, nil
}
