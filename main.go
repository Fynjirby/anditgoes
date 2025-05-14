package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"
)

var (
	today    = time.Now().Format("02.01.2006")
	home, _  = os.UserHomeDir()
	filePath = filepath.Join(home, "Documents", "anditgoes", "notes.toml")
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("\n", "anditgoes...")
		return
	}

	do := os.Args[1]
	arg := strings.Join(os.Args[2:], " ")
	switch do {
	case "write":
		writeNote(arg)
	case "read":
		if len(os.Args) < 3 {
			readNote(today)
			return
		}
		readNote(arg)
	case "today":
		readNote(today)
	case "clear":
		if len(os.Args) < 3 {
			clearDate(today)
			return
		}
		clearDate(arg)
	case "help", "man":
		fmt.Println("\n", "commands: \n  write <text> - write a note of today with <text> \n  read <date> - read all notes for <date> \n  clear <date> - clear all notes for <date> \n date format 02.01.2006")
	default:
		fmt.Println("\n", "command not found")
	}
}

func writeNote(note string) {
	err := os.MkdirAll(filepath.Dir(filePath), 0755)
	if err != nil {
		fmt.Println("\n", "error creating directory:", err)
		return
	}

	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("\n", "error opening or creating file", err)
		return
	}
	defer file.Close()

	write := fmt.Sprintf("[%s]\n%s\n\n", today, note)

	if _, err := file.WriteString(write); err != nil {
		fmt.Println("\n", "error writing file", err)
		return
	}

	fmt.Println("\n", "note added")
}

func readNote(date string) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println("\n", "no notes found.")
		return
	}

	if date == "today" {
		date = today
	}

	lines := strings.Split(string(data), "\n")
	dateHeader := "[" + date + "]"
	inSection := false
	var notes []string

	for _, line := range lines {
		line = strings.TrimSpace(line)

		if strings.HasPrefix(line, "[") && strings.HasSuffix(line, "]") {
			if line == dateHeader {
				inSection = true
			} else {
				inSection = false
			}
			continue
		}

		if inSection && line != "" {
			notes = append(notes, line)
		}
	}

	if len(notes) == 0 {
		fmt.Println("\n", "nothing")
		return
	}

	for _, note := range notes {
		fmt.Println("\n", note)
	}
}

func clearDate(date string) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println("\n", "no notes found.")
		return
	}

	lines := strings.Split(string(data), "\n")
	dateHeader := "[" + date + "]"
	inSection := false
	var notes []string

	for _, line := range lines {
		trimmed := strings.TrimSpace(line)

		if strings.HasPrefix(trimmed, "[") && strings.HasSuffix(trimmed, "]") {
			if trimmed == dateHeader {
				inSection = true
			} else {
				inSection = false
				notes = append(notes, line)
			}
			continue
		}

		if !inSection {
			notes = append(notes, line)
		}
	}

	err = os.WriteFile(filePath, []byte(strings.Join(notes, "\n")), 0644)
	if err != nil {
		fmt.Println("\n", "error writing file")
		return
	}

	if date == today {
		date = "today"
	}
	fmt.Println("\n", "all notes for", date, "were deleted")
}
