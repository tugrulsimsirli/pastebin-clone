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
	GetAllSnippetsOwn(userID uuid.UUID) (*[]models.SnippetResponseModel, error)
	GetAllSnippetsByUserID(userID uuid.UUID) (*[]models.SnippetResponseModel, error)
	GetSnippetByID(userID uuid.UUID, snippetID uuid.UUID) (*models.SnippetResponseModel, error)
	CreateSnippet(userID uuid.UUID, req models.CreateSnippetRequestModel) (*models.IdResponseModel, error)
	UpdateSnippet(userID uuid.UUID, snippetID uuid.UUID, req models.UpdateSnippetRequestModel) (*models.SnippetResponseModel, error)
	UpdateSnippetIsPublic(userID uuid.UUID, snippetID uuid.UUID, req models.BooleanRequestModel) (*models.SnippetResponseModel, error)
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

func (s *SnippetService) GetAllSnippetsOwn(userID uuid.UUID) (*[]models.SnippetResponseModel, error) {
	snippets, err := s.Repo.GetAllSnippetsOwn(userID)
	if err != nil {
		return nil, err
	}

	if snippets == nil {
		return nil, nil
	}

	response := &[]models.SnippetResponseModel{}
	err = mapper.Map(snippets, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (s *SnippetService) GetAllSnippetsByUserID(userID uuid.UUID) (*[]models.SnippetResponseModel, error) {
	snippets, err := s.Repo.GetAllSnippetsByUserID(userID)
	if err != nil {
		return nil, err
	}

	if snippets == nil {
		return nil, nil
	}

	response := &[]models.SnippetResponseModel{}
	err = mapper.Map(snippets, response)
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

	if snippet == nil {
		return nil, nil
	}

	response := &models.SnippetResponseModel{}
	err = mapper.Map(snippet, response)
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

	response := &models.IdResponseModel{ID: id}
	return response, nil
}

func (s *SnippetService) UpdateSnippet(userID uuid.UUID, snippetID uuid.UUID, req models.UpdateSnippetRequestModel) (*models.SnippetResponseModel, error) {
	snippet, err := s.Repo.GetSnippetByID(userID, snippetID)
	if err != nil || snippet == nil {
		return nil, err
	}

	updates := map[string]interface{}{}
	if req.Title != nil {
		updates["title"] = *req.Title
	}
	if req.Content != nil {
		updates["content"] = *req.Content
	}
	updates["modified_date"] = time.Now()

	if err := s.Repo.UpdateFields(snippetID, updates); err != nil {
		return nil, err
	}

	updatedSnippet, err := s.Repo.GetSnippetByID(userID, snippetID)
	if err != nil {
		return nil, err
	}

	response := &models.SnippetResponseModel{}
	err = mapper.Map(updatedSnippet, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (s *SnippetService) UpdateSnippetIsPublic(userID uuid.UUID, snippetID uuid.UUID, req models.BooleanRequestModel) (*models.SnippetResponseModel, error) {
	updates := map[string]interface{}{
		"is_public":     req.Bool,
		"modified_date": time.Now(),
	}

	if err := s.Repo.UpdateFields(snippetID, updates); err != nil {
		return nil, err
	}

	updatedSnippet, err := s.Repo.GetSnippetByID(userID, snippetID)
	if err != nil {
		return nil, err
	}

	response := &models.SnippetResponseModel{}
	err = mapper.Map(updatedSnippet, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (s *SnippetService) DeleteSnippet(userID uuid.UUID, snippetID uuid.UUID) error {
	return s.Repo.DeleteSnippet(snippetID)
}
