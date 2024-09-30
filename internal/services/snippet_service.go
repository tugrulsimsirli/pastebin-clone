// internal/services/snippet_service.go
package services

import (
	data_models "pastebin-clone/internal/db/data-models"
	"pastebin-clone/internal/http/models"
	"pastebin-clone/internal/repositories"
	"time"

	"github.com/google/uuid"
)

type SnippetServiceInterface interface {
	GetAllSnippetsByUser(userID uuid.UUID) ([]models.SnippetResponseModel, error)
	GetSnippetByID(userID uuid.UUID, snippetID uuid.UUID) (*models.SnippetResponseModel, error)
	CreateSnippet(userID uuid.UUID, req models.CreateSnippetRequestModel) (*models.SnippetResponseModel, error)
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
	for _, snippet := range snippets {
		response = append(response, models.SnippetResponseModel{
			ID:           snippet.ID,
			Title:        snippet.Title,
			Content:      snippet.Content,
			ViewCount:    snippet.ViewCount,
			CreatedDate:  snippet.CreatedDate,
			ModifiedDate: snippet.ModifiedDate,
			IsDeleted:    snippet.IsDeleted,
		})
	}

	return response, nil
}

func (s *SnippetService) GetSnippetByID(userID uuid.UUID, snippetID uuid.UUID) (*models.SnippetResponseModel, error) {
	snippet, err := s.Repo.GetSnippetByID(userID, snippetID)
	if err != nil {
		return nil, err
	}

	return &models.SnippetResponseModel{
		ID:           snippet.ID,
		Title:        snippet.Title,
		Content:      snippet.Content,
		ViewCount:    snippet.ViewCount,
		CreatedDate:  snippet.CreatedDate,
		ModifiedDate: snippet.ModifiedDate,
		IsDeleted:    snippet.IsDeleted,
	}, nil
}

func (s *SnippetService) CreateSnippet(userID uuid.UUID, req models.CreateSnippetRequestModel) (*models.SnippetResponseModel, error) {
	snippet := &data_models.Snippet{
		ID:           uuid.New(),
		UserID:       userID,
		Title:        req.Title,
		Content:      req.Content,
		ViewCount:    0,
		CreatedDate:  time.Now(),
		ModifiedDate: time.Now(),
		IsDeleted:    false,
	}

	createdSnippet, err := s.Repo.CreateSnippet(snippet)
	if err != nil {
		return nil, err
	}

	return &models.SnippetResponseModel{
		ID:           createdSnippet.ID,
		Title:        createdSnippet.Title,
		Content:      createdSnippet.Content,
		ViewCount:    createdSnippet.ViewCount,
		CreatedDate:  createdSnippet.CreatedDate,
		ModifiedDate: createdSnippet.ModifiedDate,
		IsDeleted:    createdSnippet.IsDeleted,
	}, nil
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

	return &models.SnippetResponseModel{
		ID:           updatedSnippet.ID,
		Title:        updatedSnippet.Title,
		Content:      updatedSnippet.Content,
		ViewCount:    updatedSnippet.ViewCount,
		CreatedDate:  updatedSnippet.CreatedDate,
		ModifiedDate: updatedSnippet.ModifiedDate,
		IsDeleted:    updatedSnippet.IsDeleted,
	}, nil
}

func (s *SnippetService) DeleteSnippet(userID uuid.UUID, snippetID uuid.UUID) error {
	// First, get the snippet to ensure it's owned by the user
	snippet, err := s.Repo.GetSnippetByID(userID, snippetID)
	if err != nil {
		return err
	}

	// Proceed to delete
	return s.Repo.DeleteSnippet(snippet.ID)
}
