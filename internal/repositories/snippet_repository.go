package repositories

import (
	"log"
	"pastebin-clone/internal/db"
	data_models "pastebin-clone/internal/db/data-models"
	"pastebin-clone/internal/mapper"
	"pastebin-clone/internal/repositories/dto"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type SnippetRepositoryInterface interface {
	GetAllSnippetsOwn(userID uuid.UUID) (*[]dto.SnippetDto, error)
	GetAllSnippetsByUserID(userID uuid.UUID) (*[]dto.SnippetDto, error)
	GetSnippetByID(userID uuid.UUID, snippetID uuid.UUID) (*dto.SnippetDto, error)
	CreateSnippet(snippet *data_models.Snippet) error
	UpdateFields(snippetID uuid.UUID, updates map[string]interface{}) error
	DeleteSnippet(snippetID uuid.UUID) error
}

type SnippetRepository struct{}

func NewSnippetRepository() SnippetRepositoryInterface {
	return &SnippetRepository{}
}

func (r *SnippetRepository) GetAllSnippetsOwn(userID uuid.UUID) (*[]dto.SnippetDto, error) {
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

func (r *SnippetRepository) GetAllSnippetsByUserID(userID uuid.UUID) (*[]dto.SnippetDto, error) {
	var snippets []data_models.Snippet
	err := db.DB.Where("user_id = ? and is_public = true", userID).Find(&snippets).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		log.Println(err)
		return nil, err
	}

	if len(snippets) == 0 {
		return nil, nil
	}

	response := &[]dto.SnippetDto{}
	err = mapper.Map(snippets, response)
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
	return db.DB.Create(snippet).Error
}

func (r *SnippetRepository) UpdateFields(snippetID uuid.UUID, updates map[string]interface{}) error {
	return db.DB.Model(&data_models.Snippet{}).Where("id = ?", snippetID).Updates(updates).Error
}

func (r *SnippetRepository) DeleteSnippet(snippetID uuid.UUID) error {
	return db.DB.Delete(&data_models.Snippet{}, snippetID).Error
}
