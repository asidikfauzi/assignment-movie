package postgres

import (
	"assignment-movie/domain"
	"assignment-movie/models"
	"gorm.io/gorm"
)

type Movies struct {
	DB *gorm.DB
}

func NewMoviesPostgres(conn *gorm.DB) domain.MoviePostgres {
	return &Movies{
		DB: conn,
	}
}

func (m *Movies) GetAll(limit, offset int, orderBy, search string) ([]models.GetMovies, int64, error) {
	var (
		movies []models.GetMovies
		count  int64
	)
	if err := m.DB.Model(&models.Movies{}).
		Where("deleted_at IS NULL").
		Where("title ILIKE ?", "%"+search+"%").
		Or("description ILIKE ?", "%"+search+"%").
		Limit(limit).
		Offset(offset).
		Order(orderBy).
		Find(&movies).Error; err != nil {
		return nil, 0, err
	}

	if err := m.DB.Model(&models.Movies{}).
		Where("deleted_at IS NULL").
		Where("title ILIKE ?", "%"+search+"%").
		Or("description ILIKE ?", "%"+search+"%").
		Count(&count).Error; err != nil {
		return nil, 0, err
	}

	return movies, count, nil
}

func (m *Movies) GetByID(id int) (models.GetMovies, error) {
	var movie models.GetMovies

	if err := m.DB.Model(&models.Movies{}).
		Where("deleted_at IS NULL").
		Where("id = ?", id).
		First(&movie).Error; err != nil {
		return movie, err
	}

	return movie, nil
}
