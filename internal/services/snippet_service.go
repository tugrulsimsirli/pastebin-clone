package services

import (
	data_models "pastebin-clone/internal/db/data-models"
	"pastebin-clone/internal/http/models"
	"pastebin-clone/internal/mapper"
	"pastebin-clone/internal/repositories"
	"time"

	"github.com/google/uuid"
)

type SnippetServiceInterface interface {
	GetAllSnippetsByUser(userID uuid.UUID) ([]models.SnippetResponseModel, error)
	GetSnippetByID(userID uuid.UUID, snippetID uuid.UUID) (*models.SnippetResponseModel, error)
	CreateSnippet(userID uuid.UUID, req models.CreateSnippetRequestModel) (*models.IdResponseModel, error)
	UpdateSnippet(userID uuid.UUID, snippetID uuid.UUID, req models.UpdateSnippetRequestModel) (*models.SnippetResponseModel, error)
	DeleteSnippet(userID uuid.UUID, snippetID uuid.UUID) error
}

type SnippetService struct {
	Repo repositories.SnippetRepositoryInterface
}

func NewSnippetService(repo repositories.SnippetRepositoryInterface) SnippetServiceInterface {
	return &SnippetService{
		Repo: repo,
	}
}

func (s *SnippetService) GetAllSnippetsByUser(userID uuid.UUID) ([]models.SnippetResponseModel, error) {
	snippets, err := s.Repo.GetAllSnippetsByUser(userID)
	if err != nil {
		return nil, err
	}

	var response []models.SnippetResponseModel

	err = mapper.Map(snippets, &response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (s *SnippetService) GetSnippetByID(userID uuid.UUID, snippetID uuid.UUID) (*models.SnippetResponseModel, error) {
	snippet, err := s.Repo.GetSnippetByID(userID, snippetID)
	if err != nil {
		return nil, err
	}

	var response *models.SnippetResponseModel

	err = mapper.Map(snippet, &response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (s *SnippetService) CreateSnippet(userID uuid.UUID, req models.CreateSnippetRequestModel) (*models.IdResponseModel, error) {
	id := uuid.New()

	snippet := &data_models.Snippet{
		ID:           id,
		UserID:       userID,
		Title:        req.Title,
		Content:      req.Content,
		ViewCount:    0,
		CreatedDate:  time.Now(),
		ModifiedDate: time.Now(),
		IsDeleted:    false,
	}

	err := s.Repo.CreateSnippet(snippet)
	if err != nil {
		return nil, err
	}

	response := &models.IdResponseModel{
		ID: id,
	}

	return response, nil
}

func (s *SnippetService) UpdateSnippet(userID uuid.UUID, snippetID uuid.UUID, req models.UpdateSnippetRequestModel) (*models.SnippetResponseModel, error) {
	snippet, err := s.Repo.GetSnippetByID(userID, snippetID)
	if err != nil {
		return nil, err
	}

	if req.Title != nil {
		snippet.Title = *req.Title
	}
	if req.Content != nil {
		snippet.Content = *req.Content
	}
	snippet.ModifiedDate = time.Now()

	updatedSnippet, err := s.Repo.UpdateSnippet(&data_models.Snippet{
		ID:           snippet.ID,
		UserID:       snippet.UserID,
		Title:        snippet.Title,
		Content:      snippet.Content,
		ViewCount:    snippet.ViewCount,
		CreatedDate:  snippet.CreatedDate,
		ModifiedDate: snippet.ModifiedDate,
		IsDeleted:    snippet.IsDeleted,
	})
	if err != nil {
		return nil, err
	}

	var response *models.SnippetResponseModel

	err = mapper.Map(updatedSnippet, &response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (s *SnippetService) DeleteSnippet(userID uuid.UUID, snippetID uuid.UUID) error {
	snippet, err := s.Repo.GetSnippetByID(userID, snippetID)
	if err != nil {
		return err
	}

	return s.Repo.DeleteSnippet(snippet.ID)
}
