# Gemini CLI

Gemini CLI is a command-line tool built using Golang, the Gemini API, and the Cobra package.

## Features

```
go run main.go
Gemini CLI is a command line tool for Gemini API

Usage:
  gemini [flags]
  gemini [command]

Available Commands:
  help        Help about any command
  image       Know details about an image
  search      Ask a question and get a response
  update      Update the Gemini CLI
  version     know the installed version of gemini cli

Flags:
  -h, --help   help for gemini

Use "gemini [command] --help" for more information about a command.
```

## Prerequisites

- Go 1.20+ installed on your machine
- A valid Gemini API key

## Installation

1. **Clone the Repository**

    ```bash
    git clone https://github.com/naman1402/gemini-cli-go.git
    cd gemini-cli-go
    ```

2. **Initialize the Go Module**

    ```bash
    go mod init github.com/naman1402/gemini-cli-go
    ```

3. **Set up Environment Variables**

    Replace `gemini_api_key` in your environment variables with your actual Gemini API key:

    ```bash
    export GEMINI_API_KEY=your_actual_gemini_api_key
    ```

4. **Build the Project**

    ```bash
    go build -o gemini-cli-go
    ```

5. **Run the Program**

    ```bash
    ./gemini-cli-go
    ```

### Resources

- [Gemini API](github.com/google/generative-ai-go/genai)
- [Cobra](github.com/spf13/cobra@latest)
- [Golang](https://golang.org/)
