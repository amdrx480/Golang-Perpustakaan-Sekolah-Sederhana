package services

import (
	"perpustakaan/models"
	"perpustakaan/repositories"
)

type MemberService struct {
	repository repositories.MemberRepository
}

func InitMemberService() MemberService {
	return MemberService{
		repository: &repositories.MemberRepositoryImpl{},
	}
}

func (ms *MemberService) GetAll() ([]models.Member, error) {
	return ms.repository.GetAll()
}
func (ms *MemberService) GetByID(id string) (models.Member, error) {
	return ms.repository.GetByID(id)
}
func (ms *MemberService) Create(memberInput models.MemberInput) (models.Member, error) {
	return ms.repository.Create(memberInput)
}
func (ms *MemberService) Update(memberInput models.MemberInput, id string) (models.Member, error) {
	return ms.repository.Update(memberInput, id)
}
func (ms *MemberService) Delete(id string) error {
	return ms.repository.Delete(id)
}
func (ms *MemberService) Restore(id string) (models.Member, error) {
	return ms.repository.Restore(id)
}
func (ms *MemberService) ForceDelete(id string) error {
	return ms.repository.ForceDelete(id)
}
