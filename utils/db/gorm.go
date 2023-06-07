package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

func GormMysql() *gorm.DB {
	db, err := gorm.Open(mysql.Open("root:@tcp(docker.host.internal:3306)/api_integration?charset=utf8mb4&parseTime=True&loc=UTC"),
		&gorm.Config{})
	//dsn := fmt.Sprintf("root:@tcp(%s:3306)/crud?charset=utf8mb4&parseTime=True&loc=UTC", os.Getenv("DB_HOST"))
	//db, err := gorm.Open(mysql.Open(dsn),
	//	&gorm.Config{})
	if err != nil {
		log.Println("gorm.open", err)
	}
	return db

}
