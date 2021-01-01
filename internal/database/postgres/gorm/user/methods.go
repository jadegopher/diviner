package user

import (
	"telegram-pug/model"
)

func (u *user) CreateTableIfNotExists() error {
	if err := u.db.AutoMigrate(model.User{}); err != nil {
		return err
	}
	return nil
}

func (u *user) Select(id uint64) (*model.User, error) {
	ret := &model.User{}
	err := u.db.First(ret, id).Error
	if err != nil {
		return nil, err
	}
	return ret, nil
}

func (u *user) Insert(user model.User) (*model.User, error) {
	err := u.db.Create(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *user) Update(user model.User) (*model.User, error) {
	if err := u.db.Model(&user).Updates(user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *user) Delete(id uint64) error {
	if err := u.db.Delete(&model.User{}, id).Error; err != nil {
		return err
	}
	return nil
}
