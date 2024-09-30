package repositories

import (
	"pastebin-clone/internal/db"
	data_models "pastebin-clone/internal/db/data-models"
	"pastebin-clone/internal/mapper"
	"pastebin-clone/internal/repositories/dto"

	"github.com/google/uuid"
)

type SnippetRepositoryInterface interface {
	GetAllSnippetsByUser(userID uuid.UUID) (*[]dto.SnippetDto, error)
	GetSnippetByID(userID uuid.UUID, snippetID uuid.UUID) (*dto.SnippetDto, error)
	CreateSnippet(snippet *data_models.Snippet) error
	UpdateSnippet(snippet *data_models.Snippet) (*dto.SnippetDto, error)
	DeleteSnippet(snippetID uuid.UUID) error
}

type SnippetRepository struct{}

func NewSnippetRepository() SnippetRepositoryInterface {
	return &SnippetRepository{}
}

func (r *SnippetRepository) GetAllSnippetsByUser(userID uuid.UUID) (*[]dto.SnippetDto, error) {
	var snippets []data_models.Snippet
	if err := db.DB.Where("user_id = ?", userID).Find(&snippets).Error; err != nil {
		return nil, err
	}

	response := &[]dto.SnippetDto{}

	err := mapper.Map(snippets, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (r *SnippetRepository) GetSnippetByID(userID uuid.UUID, snippetID uuid.UUID) (*dto.SnippetDto, error) {
	var snippet data_models.Snippet
	if err := db.DB.Where("user_id = ? AND id = ?", userID, snippetID).First(&snippet).Error; err != nil {
		return nil, err
	}

	response := &dto.SnippetDto{}

	err := mapper.Map(snippet, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (r *SnippetRepository) CreateSnippet(snippet *data_models.Snippet) error {
	if err := db.DB.Create(snippet).Error; err != nil {
		return err
	}

	return nil
}

func (r *SnippetRepository) UpdateSnippet(snippet *data_models.Snippet) (*dto.SnippetDto, error) {
	if err := db.DB.Save(snippet).Error; err != nil {
		return nil, err
	}

	response := &dto.SnippetDto{}

	err := mapper.Map(snippet, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (r *SnippetRepository) DeleteSnippet(snippetID uuid.UUID) error {
	return db.DB.Delete(&data_models.Snippet{}, snippetID).Error
}
