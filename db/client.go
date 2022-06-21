package db

import (
	"go-casbin/models"
	"log"
	"os"

	// "github.com/casbin/casbin/v2"
	// gormadapter "github.com/casbin/gorm-adapter/v3"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Instance *gorm.DB
var err error

func Connect() (*gorm.DB, error){
	Instance, err = gorm.Open(postgres.Open(os.Getenv("DB_URL")), &gorm.Config{})

	// Increase the column size to 512.
	// type CasbinRule struct {
	// 	ID    uint   `gorm:"primaryKey;autoIncrement"`
	// 	Ptype string `gorm:"size:512;uniqueIndex:unique_index"`
	// 	V0    string `gorm:"size:512;uniqueIndex:unique_index"`
	// 	V1    string `gorm:"size:512;uniqueIndex:unique_index"`
	// 	V2    string `gorm:"size:512;uniqueIndex:unique_index"`
	// 	V3    string `gorm:"size:512;uniqueIndex:unique_index"`
	// 	V4    string `gorm:"size:512;uniqueIndex:unique_index"`
	// 	V5    string `gorm:"size:512;uniqueIndex:unique_index"`
	// }

	// a, _ := gormadapter.NewAdapter("postgres", os.Getenv("DB_URL"), true)
	// a, _ := gormadapter.NewAdapterByDBWithCustomTable(Instance, &CasbinRule{})
	// e, _ := casbin.NewEnforcer("../config/rbac_model.conf", a)
	// e, _ := casbin.NewEnforcer("../config/rbac_model.conf", "../config/rbac_policy.csv")

	// e.LoadPolicy()

	// e.Enforce("alice", "data1", "read")

	// e.SavePolicy()

	// log.Println("adapter:", a)
	// log.Println("enforcer:", e)

	if err != nil {
		log.Fatalln("::: :: DB connection error", err)
	}
	log.Println("::: :: Connected to DB")
	Migrate()
	
	
	// log.Println("::: :: 1")
	// e.LoadPolicy()
	
	// log.Println("::: :: 2")
	// e.Enforce("alice", "data1", "read")
	
	// log.Println("::: :: 3")
	// e.SavePolicy()

		return Instance, err

}

func Migrate() {
	Instance.AutoMigrate(&models.User{})
	log.Println("::: :: DB Migration Completed")
}