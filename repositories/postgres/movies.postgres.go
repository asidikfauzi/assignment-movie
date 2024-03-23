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
		Where("(title ILIKE ? OR description ILIKE ?)", "%"+search+"%", "%"+search+"%").
		Limit(limit).
		Offset(offset).
		Order(orderBy).
		Find(&movies).Error; err != nil {
		return nil, 0, err
	}

	if err := m.DB.Model(&models.Movies{}).
		Where("deleted_at IS NULL").
		Where("(title ILIKE ? OR description ILIKE ?)", "%"+search+"%", "%"+search+"%").
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

func (m *Movies) GetByTitle(title string) (models.GetMovies, error) {
	var movie models.GetMovies

	if err := m.DB.Model(&models.Movies{}).
		Where("deleted_at IS NULL").
		Where("title = ?", title).
		First(&movie).Error; err != nil {
		return movie, err
	}

	return movie, nil
}

func (m *Movies) Create(movie models.Movies) error {
	if err := m.DB.Create(&movie).Error; err != nil {
		return err
	}

	return nil
}

func (m *Movies) Update(movie models.Movies) error {
	if err := m.DB.Where("id = ?", movie.ID).Updates(&movie).Error; err != nil {
		return err
	}

	return nil
}

func (m *Movies) CheckIfExists(movie models.Movies) (bool, error) {
	var count int64
	err := m.DB.Model(&models.Movies{}).
		Where("deleted_at IS NULL").
		Where("title = ?", movie.Title).
		Not("id = ?", movie.ID).
		Count(&count).Error

	if err != nil {
		return false, err
	}

	return count > 0, nil
}

func (m *Movies) Delete(movie models.Movies) error {
	if err := m.DB.Model(&models.Movies{}).Where("id = ?", movie.ID).
		UpdateColumn("deleted_at", movie.DeletedAt).Error; err != nil {
		return err
	}

	return nil
}
