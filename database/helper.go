package database

import "github.com/jinzhu/gorm"

// IsRecordNotFoundError returns true if error contains a RecordNotFound error.
func IsRecordNotFoundError(err error) bool {
	return gorm.IsRecordNotFoundError(err)
}
