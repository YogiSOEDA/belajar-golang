package model

import "gorm.io/gorm"

type Car struct {
	Model
	Nama  string `json:"nama" gorm:"not null"`
	Tipe  string `json:"tipe" gorm:"not null"`
	Tahun string `json:"tahun" gorm:"not null"`
}

func (cr *Car) Create(db *gorm.DB) error {
	err := db.Model(Car{}).Create(&cr).Error
	if err != nil {
		return err
	}

	return nil
}

func (cr *Car) GetByID(db *gorm.DB) (Car, error) {
	res := Car{}

	err := db.Model(Car{}).Where("id = ?", cr.Model.ID).Take(&res).Error
	if err != nil {
		return Car{}, err
	}

	return res, nil
}

func (cr *Car) GetAll(db *gorm.DB) ([]Car, error) {
	res := []Car{}

	err := db.Model(Car{}).Find(&res).Error
	if err != nil {
		return []Car{}, err
	}

	return res, nil
}

func (cr *Car) UpdateOneByID(db *gorm.DB) error {
	err := db.Model(Car{}).
		Select("nama", "tipe", "tahun").
		Where("id = ?", cr.Model.ID).
		Updates(map[string]interface{}{
			"nama":  cr.Nama,
			"tipe":  cr.Tipe,
			"tahun": cr.Tahun,
		}).
		Error

	if err != nil {
		return err
	}

	return nil
}

func (cr *Car) DeleteByID(db *gorm.DB) error {
	err := db.Model(Car{}).Where("id = ?", cr.Model.ID).Delete(&cr).Error

	if err != nil {
		return err
	}

	return nil
}