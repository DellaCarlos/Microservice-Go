package repositories

import "microservice/internal/entities"

type ICategoryRepository interface {
	Save(category *entities.Category) error
	List() ([]*entities.Category, error)
}
