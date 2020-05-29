package store

import (
	"github.com/Ghun2/fast-gorestapi/model"
	"github.com/jinzhu/gorm"
)

type UserStore struct {
	db *gorm.DB
}

func NewUserStore(db *gorm.DB) *UserStore {
	return &UserStore{
		 db,
	}
}

func (us *UserStore) Create(u *model.User) (err error) {
	return us.db.Create(u).Error
}

func (us *UserStore) GetByID(id uint) (*model.User, error) {
	var m model.User
	if err := us.db.First(&m, id).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		}
		return nil, err
	}
	return &m, nil
}

func (us *UserStore) GetByAuthID(aid string) (*model.User, error) {
	var m model.User
	if err := us.db.Where(&model.User{AuthID: aid}).First(&m).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		}
		return nil, err
	}
	return &m, nil
}

func (us *UserStore) Update(u *model.User) error {
	return us.db.Model(u).Update(u).Error
}

func (us *UserStore) Delete(u *model.User) error {
	return us.db.Delete(u).Error
}

func (us *UserStore) ListUsers() ([]model.User, error) {
	var users []model.User
	if err := us.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}