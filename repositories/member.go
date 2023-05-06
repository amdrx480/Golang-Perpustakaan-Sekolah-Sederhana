package repositories

//databaseMember lokalan tutorial
import (
	"perpustakaan/database"
	"perpustakaan/models"

	"gorm.io/gorm"
)

type MemberRepositoryImpl struct{}

func InitMemberRepository() MemberRepository {
	//penyebab merah pada dibawah ini adalah pada GetAll dan lainnya tidak sama dengan repository
	return &MemberRepositoryImpl{}
}

// method terdapat receiver
func (mr *MemberRepositoryImpl) GetAll() ([]models.Member, error) {
	var member []models.Member

	err := database.DB.Find(&member).Error

	if err != nil {
		return nil, err
	}

	return member, nil
}

func (mr *MemberRepositoryImpl) GetByID(id string) (models.Member, error) {

	var member models.Member

	err := database.DB.First(&member, "id = ?", id).Error

	if err != nil {
		return models.Member{}, err
	}

	return member, nil
}

func (mr *MemberRepositoryImpl) Create(memberInput models.MemberInput) (models.Member, error) {
	var createdMember models.Member = models.Member{
		Name:         memberInput.Name,
		Nis:          memberInput.Nis,
		Gender:       memberInput.Gender,
		Class:        memberInput.Class,
		PlaceOfBirth: memberInput.PlaceOfBirth,
		DateOfBirth:  memberInput.DateOfBirth,
		PhoneNumber:  memberInput.PhoneNumber,
	}

	result := database.DB.Create(&createdMember)

	if err := result.Error; err != nil {
		return models.Member{}, err
	}

	err := database.DB.Last(&createdMember).Error

	if err != nil {
		return models.Member{}, err
	}

	return createdMember, nil
}

func (mr *MemberRepositoryImpl) Update(memberInput models.MemberInput, id string) (models.Member, error) {
	member, err := mr.GetByID(id)

	if err != nil {
		return models.Member{}, err
	}

	member.Name = memberInput.Name
	member.Nis = memberInput.Nis
	member.Gender = memberInput.Gender
	member.Class = memberInput.Class
	member.PlaceOfBirth = memberInput.PlaceOfBirth
	member.DateOfBirth = memberInput.DateOfBirth
	member.PhoneNumber = memberInput.PhoneNumber

	err = database.DB.Save(&member).Error

	if err != nil {
		return models.Member{}, err
	}

	return member, nil
}

func (mr *MemberRepositoryImpl) Delete(id string) error {

	member, err := mr.GetByID(id)

	if err != nil {
		return err
	}

	err = database.DB.Delete(&member).Error

	if err != nil {
		return err
	}

	return nil
}

func (mr *MemberRepositoryImpl) Restore(id string) (models.Member, error) {
	var trashedMember models.Member

	//method unscoped untuk mengakses semua data yang ada di database, baik dengan deleted_at ada maupun null
	err := database.DB.Unscoped().First(&trashedMember, "id = ?", id).Error

	if err != nil {
		return models.Member{}, err
	}

	trashedMember.DeletedAt = gorm.DeletedAt{}

	err = database.DB.Unscoped().Save(&trashedMember).Error

	if err != nil {
		return models.Member{}, err
	}

	return trashedMember, nil

}

// mengapus data secara permanen
func (mr *MemberRepositoryImpl) ForceDelete(id string) error {
	member, err := mr.GetByID(id)

	if err != nil {
		return err
	}

	err = database.DB.Unscoped().Delete(&member).Error

	if err != nil {
		return err
	}

	return nil
}
