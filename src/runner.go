package src

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

type Runner struct {
	config *Config
}

func NewRunner() *Runner {
	return &Runner{
		config: GetDefaultConfig(),
	}
}

func (r *Runner) Start() {
	styles := DefaultStyles()

	// Step 1: Check if Makefile exists
	exists, err := CheckMakefileExists(r.config.MakefilePath)
	if err != nil {
		fmt.Println(styles.Text("⚠️  Failed to check Makefile: "+err.Error(), styles.ErrorColor))
		return
	}

	if !exists {
		fmt.Println(styles.Text("⚠️  Makefile not found in current directory", styles.ErrorColor))
		return
	}

	// Step 2: Parse Makefile
	targets, err := ParseMakefile(r.config.MakefilePath)
	if err != nil {
		fmt.Println(styles.Text("⚠️  Failed to parse Makefile: "+err.Error(), styles.ErrorColor))
		return
	}

	if len(targets) == 0 {
		fmt.Println(styles.Text("⚠️  No targets found in Makefile", styles.ErrorColor))
		return
	}

	// Step 3: Show targets view
	var selected string
	TargetsView(targets, r.config, &selected)

	if selected == ExitSignal || selected == "" {
		return
	}

	// Parse result: "target|name" or "settings|key"
	parts := strings.Split(selected, "|")
	if len(parts) != 2 {
		return
	}

	pageType := parts[0]
	itemName := parts[1]

	// Handle settings
	if pageType == "settings" {
		r.handleSettings(itemName)
		return
	}

	// Handle target execution
	if pageType == "target" {
		r.executeTarget(itemName)
	}
}

func (r *Runner) handleSettings(settingKey string) {
	styles := DefaultStyles()

	switch settingKey {
	case "makefile":
		fmt.Println()
		fmt.Println(styles.Text("Current Makefile path: "+r.config.MakefilePath, styles.ThistleColor))
		fmt.Println(styles.Text("(Settings feature coming soon)", styles.FooterColor))
		fmt.Println()
	default:
		return
	}
}

func (r *Runner) executeTarget(targetName string) {
	styles := DefaultStyles()

	fmt.Println()
	fmt.Println(styles.Text(fmt.Sprintf("⏳ Executing target: %s", targetName), styles.ThistleColor))
	fmt.Println()

	// Execute make command
	cmd := exec.Command("make", targetName)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	err := cmd.Run()
	if err != nil {
		fmt.Println()
		fmt.Println(styles.Text(fmt.Sprintf("⚠️  Error executing target: %v", err), styles.ErrorColor))
		fmt.Println()
	} else {
		fmt.Println()
		fmt.Println(styles.Text("✓ Target executed successfully!", styles.AquamarineColor))
		fmt.Println()
	}
}
