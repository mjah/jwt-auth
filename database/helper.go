package database

import "github.com/jinzhu/gorm"

// IsRecordNotFoundError ...
func IsRecordNotFoundError(err error) bool {
	return gorm.IsRecordNotFoundError(err)
}
