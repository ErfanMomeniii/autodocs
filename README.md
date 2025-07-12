<div id="top"></div>

<p align="center">
  <a href="https://github.com/erfanmomeniii/autodocs">
    <img src="./assets/images/logo.jpg" alt="AutoDocs Logo" width="90%">
  </a>
</p>

<p align="center">
  <em>Focus on coding—let AI handle your Go documentation.</em>
</p>

<p align="center">
  <a href="https://pkg.go.dev/github.com/erfanmomeniii/autodocs" target="_blank">
    <img src="https://img.shields.io/badge/Go-1.24.4+-00ADD8?style=for-the-badge&logo=go" alt="Go Version" />
  </a>
  <img src="https://img.shields.io/badge/License-MIT-magenta?style=for-the-badge" alt="License" />
  <img src="https://img.shields.io/badge/Version-v1.0.0-red?style=for-the-badge" alt="Version" />
</p>

---

**<i>autodocs</i>** is a lightweight command-line tool that uses AI to automatically generate GoDoc-style comments for your Go code. It enhances code readability and maintainability by analyzing your Go source files and adding concise, meaningful documentation to your exported functions, structs, interfaces, and methods.

Designed to fit naturally into your development workflow, `autodocs` can be used as part of your CI process or as a local tool to keep your codebase well-documented with minimal effort.

---

## ❓ Why AutoDocs?

- 📚 **Improve code readability** with AI-generated GoDoc-style comments.
- ⚡ **Save time** by avoiding repetitive, manual documentation tasks.
- 🧠 **Model-aware**: Customize output using powerful models like `gpt-4o`.
- 🔧 **CI-friendly**: Use in local development or integrate into pipelines.


## 📦 Installation

```bash
go install github.com/erfanmomeniii/autodocs@latest
```

## 🚀 Usage

After installation, you can run autodocs on your Go project like this:

```bash
autodocs run --path ./your-project --model gpt-4o --apikey your-api-key
```
Or use environment variable for the API key:

```bash
export AUTODOCS_API_KEY=your-api-key
autodocs run --path ./your-project
```

This command will automatically generate GoDoc-style comments for all exported declarations in your project.

## ⚙️ Available Flags

| Flag         | Shorthand | Default             | Description                                                                 |
|--------------|-----------|---------------------|-----------------------------------------------------------------------------|
| `--path` &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;    | `-p`      | `./`                | Path to the Go project to document                                          |
| `--model`    | `-m`      | `gpt-4o`            | Model name (e.g. `gpt-4o`, `claude-3-opus`)                                 |
| `--apikey`   | `-k`      | `$AUTODOCS_API_KEY` | API key used for authentication. Can also be set via environment variable. |

## 🔌 Supported AI Providers
| Provider           | Status        | Notes                                              |
|--------------------|---------------|----------------------------------------------------|
| ✅ OpenAI           | Supported     | All models (e.g., `gpt-4o`, `gpt-3.5-turbo`, etc.) |
| 🕒 Anthropic        | Coming soon   | All Claude models                                  |
| 🕒 Google (Gemini)  | Coming soon   | All Gemini models                                  |
| 🕒 Mistral          | Coming soon   | All Mixtral/Mistral models                         |
| 🕒 Meta (LLaMA)     | Coming soon   | All LLaMA models                                   |
| 🕒 Cohere           | Coming soon   | All Command R models                               |
| 🕒 Local LLMs       | Planned       | e.g., Ollama, LM Studio                            |

Models are selected using the `--model` flag. AutoDocs will route to the correct provider based on the model name.


## Contributing

Pull requests are welcome. For changes, please open an issue first to discuss what you would like to change. Please make
sure to update tests as appropriate.