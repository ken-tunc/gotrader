package gotrader

import (
	"gorm.io/gorm"
)

func OpenDb(dialector gorm.Dialector, config *gorm.Config) (*gorm.DB, error) {
	db, err := gorm.Open(dialector, config)
	if err != nil {
		return nil, err
	}

	if err = db.AutoMigrate(&Candle{}); err != nil {
		return nil, err
	}

	return db, nil
}
