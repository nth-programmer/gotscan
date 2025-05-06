package main

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"os"
	"path/filepath"
)

func main() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}

type model struct {
	repos    []string
	cursor   int
	selected map[int]struct{}
}

func initialModel() model {
	r := GetGitRepos()
	return model{
		repos:    r,
		selected: make(map[int]struct{}),
	}
}

func GetGitRepos() []string {
	flags := os.Args

	cwd, err := os.Getwd()

	if err != nil {
		fmt.Println(err)
	}

	if len(flags) > 1 {
		cwd = flags[1]
	}

	m, err := filepath.Glob(fmt.Sprintf("%v/**/.git", cwd))

	if err != nil {
		fmt.Println(err)
	}

	return m
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit

		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			if m.cursor < len(m.repos) {
				m.cursor++
			}
		}

	}

	return m, nil
}

func (m model) View() string {
	s := ""
	if len(m.repos) == 0 {
		s += "No git repositories found."
		s += "\nPress q to quit.\n"
		return s
	}

	s += "Here are the Git Repositories in the current directory\n\n"

	for i, repo := range m.repos {

		cursor := " " // no cursor
		if m.cursor == i {
			cursor = ">" // cursor!
		}

		// Is this choice selected?
		/*
			 checked := " " // not selected
			if _, ok := m.selected[i]; ok {
				checked = "x" // selected!
			}
		*/

		// Render the row
		s += fmt.Sprintf("%s %s\n", cursor, repo)
	}

	s += "\nPress q to quit.\n"

	// Send the UI for rendering
	return s
}
