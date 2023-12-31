package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Host struct {
	ID   uint   `gorm:"primaryKey;autoIncrementIncrement:true"`
	Name string `gorm:"uniqueIndex"`
	IP   string `gorm:"unique"`
}

func Connect(file string) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(file), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database")
	}

	return db
}

func GetHosts(db *gorm.DB) []Host {
	var hosts []Host

	db.Find(&hosts)

	return hosts
}

func AddHost(db *gorm.DB, name string, ip string) {
	db.Create(&Host{Name: name, IP: ip})
}

func DeleteHost(db *gorm.DB, name string) []Host {
	var hosts []Host

	db.Clauses(clause.Returning{}).Where("name = ?", name).Delete(&hosts)

	return hosts
}
