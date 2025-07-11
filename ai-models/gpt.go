package ai_models

import (
	"context"
	projectParser "github.com/erfanmomeniii/autodocs/project-parser"
	"github.com/openai/openai-go"
	"github.com/openai/openai-go/option"
	"github.com/openai/openai-go/shared"
)

const GPTPrompt = `You are a Go documentation assistant. I will give you the content of a Go (.go) file. Your task is to **only add missing GoDoc comments** to all exported elements (functions, methods, structs, interfaces, constants, and variables). 

Your additions must follow the official [GoDoc documentation style](https://go.dev/doc/comment), which includes:

- Comments must start with the name of the exported element (e.g., "// MyFunction does ...")
- Comments must be full sentences and grammatically correct.
- Use a clear, concise, and informative style.
- Do **not** change any existing code.
- Do **not** add comments for unexported (lowercase) items unless already documented.
- Do **not** invent behavior — explain what the code **actually** does, based on the implementation.

Return the **full updated .go file** with comments inserted in the proper locations.

Here is the input Go file content:
`

var validChatModels = map[string]struct{}{
	shared.ChatModelGPT4_1:                           {},
	shared.ChatModelGPT4_1Mini:                       {},
	shared.ChatModelGPT4_1Nano:                       {},
	shared.ChatModelGPT4_1_2025_04_14:                {},
	shared.ChatModelGPT4_1Mini2025_04_14:             {},
	shared.ChatModelGPT4_1Nano2025_04_14:             {},
	shared.ChatModelO4Mini:                           {},
	shared.ChatModelO4Mini2025_04_16:                 {},
	shared.ChatModelO3:                               {},
	shared.ChatModelO3_2025_04_16:                    {},
	shared.ChatModelO3Mini:                           {},
	shared.ChatModelO3Mini2025_01_31:                 {},
	shared.ChatModelO1:                               {},
	shared.ChatModelO1_2024_12_17:                    {},
	shared.ChatModelO1Preview:                        {},
	shared.ChatModelO1Preview2024_09_12:              {},
	shared.ChatModelO1Mini:                           {},
	shared.ChatModelO1Mini2024_09_12:                 {},
	shared.ChatModelGPT4o:                            {},
	shared.ChatModelGPT4o2024_11_20:                  {},
	shared.ChatModelGPT4o2024_08_06:                  {},
	shared.ChatModelGPT4o2024_05_13:                  {},
	shared.ChatModelGPT4oAudioPreview:                {},
	shared.ChatModelGPT4oAudioPreview2024_10_01:      {},
	shared.ChatModelGPT4oAudioPreview2024_12_17:      {},
	shared.ChatModelGPT4oAudioPreview2025_06_03:      {},
	shared.ChatModelGPT4oMiniAudioPreview:            {},
	shared.ChatModelGPT4oMiniAudioPreview2024_12_17:  {},
	shared.ChatModelGPT4oSearchPreview:               {},
	shared.ChatModelGPT4oMiniSearchPreview:           {},
	shared.ChatModelGPT4oSearchPreview2025_03_11:     {},
	shared.ChatModelGPT4oMiniSearchPreview2025_03_11: {},
	shared.ChatModelChatgpt4oLatest:                  {},
	shared.ChatModelCodexMiniLatest:                  {},
	shared.ChatModelGPT4oMini:                        {},
	shared.ChatModelGPT4oMini2024_07_18:              {},
	shared.ChatModelGPT4Turbo:                        {},
	shared.ChatModelGPT4Turbo2024_04_09:              {},
	shared.ChatModelGPT4_0125Preview:                 {},
	shared.ChatModelGPT4TurboPreview:                 {},
	shared.ChatModelGPT4_1106Preview:                 {},
	shared.ChatModelGPT4VisionPreview:                {},
	shared.ChatModelGPT4:                             {},
	shared.ChatModelGPT4_0314:                        {},
	shared.ChatModelGPT4_0613:                        {},
	shared.ChatModelGPT4_32k:                         {},
	shared.ChatModelGPT4_32k0314:                     {},
	shared.ChatModelGPT4_32k0613:                     {},
	shared.ChatModelGPT3_5Turbo:                      {},
	shared.ChatModelGPT3_5Turbo16k:                   {},
	shared.ChatModelGPT3_5Turbo0301:                  {},
	shared.ChatModelGPT3_5Turbo0613:                  {},
	shared.ChatModelGPT3_5Turbo1106:                  {},
	shared.ChatModelGPT3_5Turbo0125:                  {},
	shared.ChatModelGPT3_5Turbo16k0613:               {},
}

func IsValidGPTChatModel(model string) bool {
	_, ok := validChatModels[model]
	return ok
}

type GPT struct {
	modelName string
	client    client
}

type client struct {
	c openai.Client
}

func (g *GPT) Generate(path string) error {
	p := projectParser.Parser{}

	files, err := p.AllFiles(path)
	if err != nil {
		return err
	}

	for _, file := range files {
		content, err := file.Read()
		if err != nil {
			return err
		}
		finalContent, err := g.addGoDocs(content)
		if err != nil {
			return err
		}
		err = file.Write(finalContent)
		if err != nil {
			return err
		}
	}
	return nil
}

func NewGPT(modelName string, apiKey string) *GPT {
	return &GPT{
		modelName: modelName,
		client: client{
			c: openai.NewClient(
				option.WithAPIKey(apiKey),
			),
		},
	}
}

func (g *GPT) addGoDocs(fileContent []byte) ([]byte, error) {
	result, err := g.client.c.Chat.Completions.New(context.Background(),
		openai.ChatCompletionNewParams{
			Messages: []openai.ChatCompletionMessageParamUnion{
				openai.UserMessage(
					GPTPrompt + string(fileContent)),
			},
			Model: g.modelName,
		})
	if err != nil {
		return nil, err
	}
	return []byte(result.Choices[0].Message.Content), nil
}
