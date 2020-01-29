package database

// Migrate ...
func Migrate() {
	db := GetConnection()
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
