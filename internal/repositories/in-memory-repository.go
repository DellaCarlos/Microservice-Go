package repositories

// aqui pode ser adicionando a conexão com Postgres, MongoDB, etc.
// Por enquanto, vamos usar um repositório em memória para fins de demonstração.

import "microservice/internal/entities"

type inMemoryCategoryRepository struct {
	db []*entities.Category
}

func NewInMemoryCategoryRepository() *inMemoryCategoryRepository {
	return &inMemoryCategoryRepository{
		db: make([]*entities.Category, 0),
	}
}

func (r *inMemoryCategoryRepository) Save(category *entities.Category) error {
	// Adicionar ID antes de salvar (ID como responsabilidade do DB)
	r.db = append(r.db, category)

	return nil
}

func (r *inMemoryCategoryRepository) List() ([]*entities.Category, error) {
	return r.db, nil
}
