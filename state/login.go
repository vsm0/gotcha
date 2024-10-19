package state

import (
	"github.com/charmbracelet/bubbletea"
)

type Login struct {
	stack *Stack
	db *Db
}

func NewLogin(s *Stack, db *Db) *Login {
	return &Login{
		stack: s,
		db: db,
	}
}

func (s *Login) Init() tea.Cmd {
	return tea.SetWindowTitle("Account Page")
}

func (s *Login) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return s, nil
}

func (s *Login) View() string {
	return ""
}
