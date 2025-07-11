package ai_models

import (
	"errors"
)

type GoDocsGenerator interface {
	Generate(path string) error
}
type AIModel struct {
	Name      string
	ApiKey    string
	Generator GoDocsGenerator
}

type Option func(*AIModel) *AIModel

func WithApiKey(key string) Option {
	return func(a *AIModel) *AIModel {
		a.ApiKey = key
		return a
	}
}

func WithName(name string) Option {
	return func(a *AIModel) *AIModel {
		a.Name = name
		return a
	}
}

func addGenerator(a *AIModel) (*AIModel, error) {
	if IsValidGPTChatModel(a.Name) {
		a.Generator = NewGPT(a.Name, a.ApiKey)

		return a, nil
	}
	return nil, errors.New("invalid model name")
}

func NewAIModel(options ...Option) (*AIModel, error) {
	var a *AIModel
	for _, option := range options {
		a = option(a)
	}

	return addGenerator(a)
}
