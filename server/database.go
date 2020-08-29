package main

import (
   
    "os"
    "log"

    "github.com/jinzhu/gorm"
)

var db *gorm.DB

func init() {
    var err error
    var datasource string

    // DBMS := "mysql"
	// USER := "b6269a7b123b6c"
	// PASS := "c2129dd2"
	// PROTOCOL := "tcp(localhost:3306)"
    // DBNAME := "heroku_4cadff22e8acc17"

    // CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?charset=utf8&parseTime=True&loc=Local"

    
    if os.Getenv("DATABASE_URL") != "" {
        // Heroku用
datasource = "b6269a7b123b6c:c2129dd2@tcp(us-cdbr-east-02.cleardb.com:3306)/heroku_4cadff22e8acc17?parseTime=true&charset=utf8&loc=Local"
    } else {
        // ローカル用
        datasource = "b6269a7b123b6c:c2129dd2@/localhost?parseTime=true&charset=utf8&loc=Local"
    }
    db, err = gorm.Open("mysql",datasource )
    if err != nil {
      
        log.Fatalf("Got error when connect database, the error is '%v'", err)
    }

    db.AutoMigrate(&Product{})
}