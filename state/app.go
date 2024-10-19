package state

import (
	"github.com/charmbracelet/bubbletea"
	"github.com/vsm0/adt/stack"
	"gorm.io/gorm"
)

type Stack struct {
	*stack.Stack[tea.Model]
}

type Db struct {
	*gorm.DB
}

type App struct {
	stack *Stack
	db *Db
}

func NewApp(db *gorm.DB) *App {
	return &App{
		stack: &Stack{stack.New[tea.Model]()},
		db: &Db{db},
	}
}

func (a *App) Run() (tea.Model, error) {
	p := tea.NewProgram(a, tea.WithAltScreen())
	return p.Run()
}

func (a *App) Init() tea.Cmd {
	s := NewLogin(a.stack, a.db)
	a.stack.Push(s)
	return s.Init()
}

func (a *App) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	s, err := a.stack.Peek()
	if err != nil {
		panic(err)
	}

	return s.Update(msg)
}

func (a *App) View() string {
	s, err := a.stack.Peek()
	if err != nil {
		panic(err)
	}

	return s.View()
}
