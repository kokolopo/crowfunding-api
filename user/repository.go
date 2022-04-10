package user

import "gorm.io/gorm"

type Repository interface {
	Save(user User) (User, error)
}

type repository struct {
	DB *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Save(user User) (User, error) {

	err := r.DB.Create(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}
