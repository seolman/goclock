package main

import (
	"fmt"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	time time.Time
}

func initialModel() model {
	return model{
		time: time.Now(),
	}
}

func (m model) Init() tea.Cmd {
	return tea.Every(time.Second, func(t time.Time) tea.Msg {
		return t
	})
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "q" || msg.String() == "ctrl+c" {
			return m, tea.Quit
		}
	case time.Time:
		m.time = msg
		return m, m.Init()
	}
	return m, nil
}

func getASCIIDigit(d int) []string {
	digits := [][]string{
		{
			"███",
			"█ █",
			"█ █",
			"█ █",
			"███",
		},
		{
			" █ ",
			" █ ",
			" █ ",
			" █ ",
			" █ ",
		},
		{
			"███",
			"  █",
			"███",
			"█  ",
			"███",
		},
		{
			"███",
			"  █",
			"███",
			"  █",
			"███",
		},
		{
			"█ █",
			"█ █",
			"███",
			"  █",
			"  █",
		},
		{
			"███",
			"█  ",
			"███",
			"  █",
			"███",
		},
		{
			"███",
			"█  ",
			"███",
			"█ █",
			"███",
		},
		{
			"███",
			"  █",
			"  █",
			"  █",
			"  █",
		},
		{
			"███",
			"█ █",
			"███",
			"█ █",
			"███",
		},
		{
			"███",
			"█ █",
			"███",
			"  █",
			"███",
		},
	}
	return digits[d]
}

func (m model) View() string {
	timeStr := m.time.Format("15:04:05")
	var display []string

	// Initialize 5 empty lines for the ASCII art
	for i := 0; i < 5; i++ {
		display = append(display, "")
	}

	// Build each line of the display
	for _, char := range timeStr {
		if char == ':' {
			// Add colon separator
			for i := 0; i < 5; i++ {
				if i == 1 || i == 3 {
					display[i] += " █ "
				} else {
					display[i] += "   "
				}
			}
			continue
		}

		digit := int(char - '0')
		digitLines := getASCIIDigit(digit)
		for i := 0; i < 5; i++ {
			display[i] += digitLines[i] + " "
		}
	}

	// Join the lines into a single string
	result := "\n  ASCII Clock\n\n"
	for _, line := range display {
		result += "  " + line + "\n"
	}
	result += "\n  (press q to quit)\n"

	return result
}

func main() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Error running program: %v", err)
		return
	}
}
