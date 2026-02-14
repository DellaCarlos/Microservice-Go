package use_cases

import (
	"microservice/internal/entities"
	"microservice/internal/repositories"
)

type createCategoryUseCase struct {
	repository repositories.ICategoryRepository
}

func NewCreateCategoryUseCase(repository repositories.ICategoryRepository) *createCategoryUseCase {
	return &createCategoryUseCase{repository}
}

func (u *createCategoryUseCase) Execute(name string) error {
	category, err := entities.NewCategory(name)

	if err != nil {
		return err
	}

	// TODO: Verificar se a categoria já existe antes de salvar

	// como estamos usando uma abstração de repositório,
	// não precisamos nos preocupar com a implementação específica
	// (in-memory, Postgres, etc.)
	err = u.repository.Save(category)
	if err != nil {
		return err
	}

	return nil
}
