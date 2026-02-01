package src

import (
"bufio"
"os"
"strings"
)

// ParseMakefile reads a Makefile and extracts all targets with their descriptions
func ParseMakefile(filepath string) ([]MakeTarget, error) {
file, err := os.Open(filepath)
if err != nil {
return nil, err
}
defer file.Close()

var targets []MakeTarget
var lastComment string
var lines []string
scanner := bufio.NewScanner(file)

// Read all lines first
for scanner.Scan() {
lines = append(lines, scanner.Text())
}

if err := scanner.Err(); err != nil {
return nil, err
}

// Parse targets
for i := 0; i < len(lines); i++ {
line := lines[i]
trimmedLine := strings.TrimSpace(line)

// Skip empty lines
if trimmedLine == "" {
lastComment = ""
continue
}

// Check if it's a comment (description for next target)
if strings.HasPrefix(trimmedLine, "#") {
// Extract comment text
comment := strings.TrimPrefix(trimmedLine, "#")
comment = strings.TrimSpace(comment)
if comment != "" {
lastComment = comment
}
continue
}

// Check if it's a target definition (contains : but not := or =)
if strings.Contains(line, ":") && !strings.Contains(line, ":=") && !strings.Contains(line, "=") {
// Extract target name (before the colon)
parts := strings.SplitN(line, ":", 2)
if len(parts) >= 1 {
targetName := strings.TrimSpace(parts[0])

// Skip if it starts with tab (it's a recipe, not a target)
if strings.HasPrefix(line, "\t") {
continue
}

// Skip special targets that start with . (like .PHONY)
if strings.HasPrefix(targetName, ".") {
continue
}

// Skip if contains spaces (might be a variable or pattern rule)
if strings.Contains(targetName, " ") {
continue
}

// Skip if empty
if targetName == "" {
continue
}

// Extract recipe lines (up to 3 lines that start with tab)
var recipe []string
for j := i + 1; j < len(lines) && len(recipe) < 3; j++ {
recipeLine := lines[j]
if strings.HasPrefix(recipeLine, "\t") {
// Remove leading tab and trim
cleanLine := strings.TrimPrefix(recipeLine, "\t")
cleanLine = strings.TrimSpace(cleanLine)
if cleanLine != "" {
recipe = append(recipe, cleanLine)
}
} else if strings.TrimSpace(recipeLine) != "" {
// Stop at first non-empty, non-tab line
break
}
}

target := MakeTarget{
Name:        targetName,
Description: lastComment,
Recipe:      recipe,
IsPhony:     false,
}

targets = append(targets, target)
}
lastComment = ""
} else {
// Reset comment if line is not a target or comment
lastComment = ""
}
}

return targets, nil
}

// CheckMakefileExists verifies if a Makefile exists in the current directory
func CheckMakefileExists(filepath string) (bool, error) {
_, err := os.Stat(filepath)
if err != nil {
if os.IsNotExist(err) {
return false, nil
}
return false, err
}
return true, nil
}
