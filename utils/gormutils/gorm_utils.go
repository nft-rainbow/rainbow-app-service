package gormutils

import "gorm.io/gorm"

func IsRecordNotFoundError(err error) bool {
	return err == gorm.ErrRecordNotFound
}
