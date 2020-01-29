package database

import "github.com/mjah/jwt-auth/logger"

// Migrate ...
func Migrate() {
	db, err := GetConnection()
	if err != nil {
		logger.Log().Fatal(err)
	}
	defer db.Close()

	db.AutoMigrate(&User{})
	db.AutoMigrate(&Role{})
	db.AutoMigrate(&TokenRevocation{})
	db.AutoMigrate(&EmailQueue{})

	adminRole := &Role{Role: "Admin"}
	db.FirstOrCreate(adminRole, adminRole)

	memberRole := &Role{Role: "Member"}
	db.FirstOrCreate(memberRole, memberRole)

	guestRole := &Role{Role: "Guest"}
	db.FirstOrCreate(guestRole, guestRole)
}
