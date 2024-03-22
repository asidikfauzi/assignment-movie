package main

import (
	"assignment-movie/common/database"
	"assignment-movie/models"
	"fmt"
	"gorm.io/gorm"
	"log"
	"time"
)

func main() {
	db, err := database.InitDatabase()
	if err != nil {
		panic("Failed to connect to database: " + err.Error())
	}

	InitMoviesSeed(db)

	fmt.Println("SUCCESSFULLY ADD SEEDER")
}

func InitMoviesSeed(db *gorm.DB) {
	movies := []models.Movies{
		{
			Title: "Pengabdi Setan 2 Comunion",
			Description: "dalah sebuah film horor Indonesia tahun 2022 yang disutradarai dan ditulis oleh Joko Anwar " +
				"sebagai sekuel dari film tahun 2017, Pengabdi Setan",
			Rating:    7.0,
			Image:     "image1.png",
			CreatedAt: time.Now(),
		},
		{
			Title: "Agak Laen",
			Description: "Demi mengejar mimpi untuk mengubah nasib, empat sekawan penjaga rumah hantu di pasar malam, " +
				"mencari cara baru menakuti pengunjung agar selamat dari kebangkrutan. Sialnya, usaha Bene, Jegel, " +
				"Boris dan Oki malah memakan korban jiwa salah satu pengunjungnya. Karena panik, korban tersebut mereka " +
				"kubur di dalam rumah hantu.Di luar dugaan, arwah si korban malah gentayangan, membuat rumah hantunya " +
				"jadi seram dan ramai pengunjung. Ketika polisi mulai menyelidiki, mereka pun terpaksa melakukan " +
				"berbagai persekongkolan konyol untuk menutupi kejadian sebenarnya. Bagaimana nasib mereka selanjutnya?",
			Rating:    9.0,
			Image:     "image2.png",
			CreatedAt: time.Now(),
		},
		{
			Title: "Avenger End Game",
			Description: "Melanjutkan Avengers Infinity War, dimana kejadian setelah Thanos berhasil mendapatkan semua " +
				"infinity stones dan memusnahkan 50% semua mahluk hidup di alam semesta. Akankah para Avengers berhasil " +
				"mengalahkan Thanos?",
			Rating:    9.5,
			Image:     "image3.png",
			CreatedAt: time.Now(),
		},
	}

	for _, movie := range movies {
		var existingMovie models.Movies
		if err := db.Where("title = ?", movie.Title).First(&existingMovie).Error; err == nil {
			log.Printf("Movie '%s' already exists, skipping.", movie.Title)
			continue
		}

		if err := db.Create(&movie).Error; err != nil {
			log.Printf("Failed to create customer: %s", err)
		} else {
			log.Printf("Movie '%s' created successfully.", movie.Title)
		}
	}
}
