package main

import (
	"log"
	"strconv"

	tea "github.com/charmbracelet/bubbletea"
)

type Editor struct {
	lines              []string
	curr_row           int
	curr_col           int
	default_cursor     string
	curr_window_width  int
	curr_window_height int
}

func NewEditor() Editor {
	return Editor{
		lines:              []string{},
		curr_row:           0,
		curr_col:           0,
		default_cursor:     "|",
		curr_window_width:  0,
		curr_window_height: 0,
	}
}

func (e Editor) Init() tea.Cmd {
	return tea.ClearScreen
}

func (e Editor) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		{
			switch msg.String() {
			case "ctrl+q":
				return e, tea.Quit
			case "backspace":
				{
					if len(e.lines) > 0 {
						e.lines = e.lines[0 : len(e.lines)-1]
					}

					break
				}
			case "enter":
				{
					e.curr_row += 1
					e.lines = append(e.lines, "\n")
					e.lines = append(e.lines, strconv.Itoa(e.curr_row+1)+"~ ")
					break
				}
			default:
				{
					e.lines = append(e.lines, msg.String())
					e.curr_col += 1
					break
				}
			}
		}
	}

	if len(e.lines) < 1 {
		e.lines = append(e.lines, "1~ ")
	}

	return e, nil
}

func (e Editor) View() string {
	s := ""
	for _, str := range e.lines {
		s += str
	}

	s += e.default_cursor

	return s
}

func main() {
	p := tea.NewProgram(NewEditor())
	if _, err := p.Run(); err != nil {
		log.Fatal("EddiError: " + err.Error())
	}
}
