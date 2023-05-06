package services

import (
	"perpustakaan/middlewares"
	"perpustakaan/models"
	"perpustakaan/repositories"
)

type AdminService struct {
	Repository repositories.AdminRepository
	jwtAuth    *middlewares.JWTConfig
}

func InitAdminService(jwtAuth *middlewares.JWTConfig) AdminService {
	return AdminService{
		Repository: &repositories.AdminRepositoryImpl{},
		jwtAuth:    jwtAuth,
	}
}

func (as *AdminService) Register(adminInput models.AdminInput) (models.Admin, error) {
	return as.Repository.Register(adminInput)
}
func (as *AdminService) Login(adminInput models.AdminInput) (string, error) {
	admin, err := as.Repository.GetByEmail(adminInput)

	//penyebab saat login success meskipun passwordnya salah
	// if err != nil {
	// 	return "", nil
	// }

	if err != nil {
		return "", err
	}

	token, err := as.jwtAuth.GenerateToken(int(admin.ID))

	if err != nil {
		return "", err
	}

	return token, nil
}
