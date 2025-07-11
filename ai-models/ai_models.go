package ai_models

import (
	"errors"
)

// GoDocsGenerator is an interface for generating Go documentation.
type GoDocsGenerator interface {
	// Generate produces documentation for the given path.
	Generate(path string) error
}

// AIModel represents an AI model with a name, API key, and a documentation generator.
type AIModel struct {
	Name      string
	ApiKey    string
	Generator GoDocsGenerator
}

// Option defines a function type for configuring an AIModel.
type Option func(*AIModel) *AIModel

// WithApiKey returns an Option that sets the API key of an AIModel.
func WithApiKey(key string) Option {
	return func(a *AIModel) *AIModel {
		a.ApiKey = key
		return a
	}
}

// WithName returns an Option that sets the name of an AIModel.
func WithName(name string) Option {
	return func(a *AIModel) *AIModel {
		a.Name = name
		return a
	}
}

// addGenerator assigns a suitable GoDocsGenerator to the AIModel based on its name.
// It returns an error if the model name is invalid.
func addGenerator(a *AIModel) (*AIModel, error) {
	if IsValidGPTChatModel(a.Name) {
		a.Generator = NewGPT(a.Name, a.ApiKey)

		return a, nil
	}
	return nil, errors.New("invalid model name")
}

// NewAIModel creates a new AIModel configured with the given options.
// It returns an error if the model name is invalid.
func NewAIModel(options ...Option) (*AIModel, error) {
	var a AIModel
	for _, option := range options {
		a = *option(&a)
	}

	return addGenerator(&a)
}
