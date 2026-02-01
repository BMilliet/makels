# makels

<div align="center">
  
**Interactive Makefile target selector for the terminal**

Fast, beautiful menu selection for your Makefile targets with fuzzy search

</div>

---

## 🎯 What is makels?

`makels` is a terminal UI tool that lists all available targets in your Makefile and lets you select and execute them interactively. No more typing `make <tab><tab>` or scrolling through your Makefile to find the right target!

## ✨ Features

- 🎨 **Beautiful terminal UI** with smooth scrolling and green-themed colors
- 🔍 **Fuzzy search** - quickly find targets by typing `/` and searching
- 📝 **Smart descriptions** - shows comment descriptions or code preview for each target
- ⚡ **Fast navigation** - keyboard-driven with vim-style keybindings
- 🎯 **Direct execution** - select a target and it runs immediately
- 🌈 **Syntax highlighting** - different colors for target names, descriptions, and code

## 📦 Installation

### Build from source

```bash
# Clone the repository
git clone <your-repo-url>
cd makels

# Install dependencies
make deps

# Build and install
make build
```

This will create the binary at `~/.makels/makels`.

### Add to PATH

To use `makels` globally, add this to your shell config (`~/.zshrc` or `~/.bashrc`):

```bash
export PATH="$HOME/.makels:$PATH"
```

Then reload your shell:

```bash
source ~/.zshrc  # or source ~/.bashrc
```

## 🚀 Usage

Simply run `makels` in any directory containing a Makefile:

```bash
makels
```

### Keyboard Controls

- **↑↓** or **j/k** - Navigate up/down through targets
- **/** - Enter search mode
- **Enter** - Execute selected target
- **q** or **Esc** - Quit
- **Esc** (in search) - Cancel search

### Search Mode

Press `/` to enter fuzzy search mode. Type any characters and the list will filter to matching targets:

```
🔍 Search: bui_
```

This will match targets like:
- `build`
- `build-prod`
- `rebuild`

## 📖 How it works

`makels` parses your Makefile and extracts:

1. **Target names** - any line ending with `:`
2. **Descriptions** - comments with `##` above the target
3. **Code preview** - first 3 lines of the target's recipe (if no description)

### Example Makefile

```makefile
.PHONY: help build run test

help: ## Show this help message
	@echo "Available targets:"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST)

build: ## Build the application
	@echo "Building..."
	go build -o app .

run:
	@echo "Running..."
	go run main.go

test: ## Run tests
	go test -v ./...
```

In `makels`, you'll see:
- `help` - "Show this help message"
- `build` - "Build the application"
- `run` - Shows code preview: `@echo "Running..."` (faded)
- `test` - "Run tests"

## 🎨 UI Preview

```
Makefile Targets

╭──────────────────────────────────────────────────────────────────────╮
│ build                                                                │
│ make                                                                 │
│ Build the application                                                │
╰──────────────────────────────────────────────────────────────────────╯

╭──────────────────────────────────────────────────────────────────────╮
│ run                                                                  │
│ make                                                                 │
│ @echo "Running..."                                                   │
│ go run main.go                                                       │
╰──────────────────────────────────────────────────────────────────────╯

  / search • ↑↓/jk navigate • enter select • q/esc quit
```

## 🛠️ Development

### Requirements

- Go 1.25.6 or later
- Make

### Available Make targets

```bash
make help      # Show available targets
make deps      # Install dependencies
make build     # Build the binary
make run       # Run without building
make fmt       # Format code
make test      # Run tests
make lint      # Run linters
make clean     # Remove build artifacts
```

### Project Structure

```
makels/
├── main.go              # Entry point
├── Makefile             # Build commands
└── src/
    ├── constants.go     # Constants (ExitSignal, etc)
    ├── makefile_parser.go  # Makefile parsing logic
    ├── models.go        # Data structures
    ├── runner.go        # Main execution flow
    ├── style.go         # UI styles and colors
    └── targets_view.go  # Interactive list view
```

## 🎨 Color Palette

The UI uses a green-themed color palette:

- **Light Green** (#A8E6A3) - Headings
- **Emerald Green** (#6BCF6B) - Selected items
- **Mint Green** (#90EE90) - Labels
- **Aqua Green** (#5FD3A6) - "make" command
- **Pale Green** (#D4F1D4) - Help text
- **Faded Gray** (#5A5A5A) - Code previews

## 🤝 Contributing

Contributions are welcome! Feel free to open issues or submit pull requests.

## 📄 License

MIT License - see LICENSE file for details

## 🙏 Acknowledgments

Built with:
- [Bubble Tea](https://github.com/charmbracelet/bubbletea) - Terminal UI framework
- [Lip Gloss](https://github.com/charmbracelet/lipgloss) - Style definitions

Inspired by similar projects like `postless` and `terminal-gameplay`.

---

<div align="center">
Made with 💚 for developers who love the terminal
</div>
