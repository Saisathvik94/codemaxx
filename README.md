# CodeMaxx 

**CodeMaxx** is a cross-platform CLI tool that helps developers understand, and fix code using multiple large language models (LLMs).

## Installation

Choose your OS:

### **Windows (PowerShell)**

```powershell
iwr https://raw.githubusercontent.com/Saisathvik94/codemaxx/main/scripts/install.ps1 -UseBasicParsing | iex
```
> Run as **Administrator**

### **Linux/macOS**

```bash
curl -sSL https://raw.githubusercontent.com/Saisathvik94/codemaxx/main/scripts/install.sh | sudo bash
```

> After installation, `codemaxx` will be available on your PATH.

### Usage 

Run the built‑in help to see all available commands:
```
codemaxx --help
```

**Adding API Keys for LLMs**

Before using AI features, you need to configure your API keys for supported models:

```
# Open the keys manager
codemaxx keys

# Example workflow:
# 1. Select a model (OpenAI, Anthropic, Perplexity, Gemini, etc.)
# 2. Add your API key for selected model
# 3. Save the configuration
```

After setting your keys, you can use CodeMaxx to explain, fix, or analyze code with your preferred language model.

### Available Commands

Below are the main commands currently supported by CodeMaxx:

**1. fix**

- fix the errors in the code, and refactor the code.
```
codemaxx fix <path> 
codemaxx fix <path> --prompt "Extra Instructions"
```
**2. explain**

- explain the code 

```
codemaxx explain <path> 
codemaxx explain <path> --prompt "Explain this code indetail"
```

**3.keys**

- used to add API keys 

```
codemaxx keys
```

**4. models**

- used to select the provider or LLM

```
codemaxx models
```

## How it Works?

**1. Select the Model**

CodeMaxx supports multiple large language models (LLMs), including OpenAI, Anthropic, Gemini, and Perplexity. Users can choose which model to use for code explanations or fixes.

**2. Provide Your Code**

You can point CodeMaxx to a file or even pipe code directly through the CLI. The tool reads your code and prepares it for analysis.

**3. Analyze and Explain**

Using the selected LLM, CodeMaxx generates human-readable explanations of your code logic, helping you understand complex functions, algorithms, or structures.

**4. Detect Issues and Suggest Fixes**

CodeMaxx can automatically detect coding issues, bugs, or inefficiencies. It provides suggestions for fixes and, optionally, can apply them directly to your files.



**5. Cross-Platform CLI**

CodeMaxx works seamlessly on Windows, Linux, and macOS. The installation script downloads the correct binary for your platform automatically.

## Safety Notes

**Review AI Suggestions** – LLMs provide suggestions based on patterns and learned data. Always review changes before committing to production code.

**Private Code** – If your project contains sensitive data, avoid sending it to external LLMs without proper security measures. Be cautious with proprietary or confidential code.

**Version Control Recommended** – Use Git or another version control system to track changes and easily revert if needed.

## Future Improvements

- **Expanded LLM Support** – Include more models and allow users to plug in their own APIs.

- **Interactive Code Review** – Provide step-by-step explanations, inline suggestions, and reasoning for fixes directly in the terminal or IDE.

- **Project-Wide Refactoring** – Analyze and improve entire projects, not just single files, with intelligent dependency awareness.

- **Custom LLM Integrations** – Allow users to plug in their own language models or API keys for personalized AI behavior.

## Contributors

- Pull requests are welcome!

- Fork the repo

- Create a feature branch

- Commit your changes

- Open a PR

## MIT License

Free to use, modify, and distribute.

