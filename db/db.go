package DB

import (
	"exons/utils"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db  *gorm.DB
	err error
)

//Open return a dastabase instance
func Open() *gorm.DB {
	db, err = gorm.Open("mysql", "root:root@/exons")
	utils.CheckErr(err, "No se puede abrir una conexión")

	return db
}

