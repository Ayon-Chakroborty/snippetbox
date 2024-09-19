package mocks

import (
	"time"

	"snippetbox.ayonchakroborty.net/internal/models"
)

var mockSnippet = models.Snippet{
	ID: 1,
	Title: "mock1",
	Content: "mock1",
	Created: time.Now(),
	Expires: time.Now(),
}

type SnippetModel struct{}

func (m *SnippetModel) Insert(title string, content string, expires int) (int, error){
	return 2, nil
}

func (m *SnippetModel) Get(id int) (models.Snippet, error){
	switch id{
	case 1:
		return mockSnippet, nil
	default:
		return models.Snippet{}, models.ErrNoRecord 
	}
}

func (m *SnippetModel) GetLatest() ([]models.Snippet, error){
	return []models.Snippet{mockSnippet}, nil
}