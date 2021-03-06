package model

import (
    "GoFeedBot/config"
    "github.com/jinzhu/gorm"
)

var db *gorm.DB

func Init(config config.Config) {
    var err error
    db, err = gorm.Open("sqlite3", "data.db")
    if err != nil {
        panic(err)
    }

    db.AutoMigrate(&Feed{}, &Subscribe{}, &Item{})

    if config.Debug {
        db = db.Debug()
    }
}
