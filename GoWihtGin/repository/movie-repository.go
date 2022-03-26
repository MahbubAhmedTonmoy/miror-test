package repository

import (
	"GoWithGin/entity"
	"database/sql"

	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

type MovieRepository interface {
	Save(movie entity.Movie)
	Update(movie entity.Movie)
	Delete(movie entity.Movie)
	FindAll() []entity.Movie
	FindById(Id int64) entity.Movie
	//CloseDB()
}

type database struct {
	connection *gorm.DB
}

func NewMovieRepository() MovieRepository {
	// github.com/denisenkom/go-mssqldb
	dsn := "sqlserver://localhost:49152?database=master&trusted+connection=yes"
	db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database")
	}
	db.AutoMigrate(&entity.Movie{}, &entity.Person{})
	return &database{
		connection: db,
	}
}

// func (db *database) CloseDB() {
// 	err := db.connection.Close()
// 	if err != nil {
// 		panic("Failed to close database")
// 	}
// }

func (db *database) Save(video entity.Movie) {
	db.connection.Create(&video)
}

func (db *database) Update(video entity.Movie) {
	db.connection.Save(&video)
}

func (db *database) Delete(video entity.Movie) {
	db.connection.Delete(&video)
}

func (db *database) FindAll() []entity.Movie {
	var videos []entity.Movie
	db.connection.Set("gorm:auto_preload", true).Find(&videos)
	return videos
}

func (db *database) FindById(Id int64) entity.Movie {
	var video entity.Movie
	db.connection.Where("id = @id", sql.Named("id", Id)).Find(&video)
	return video
}
