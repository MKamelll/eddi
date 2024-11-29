package main

import (
	"log"

	"github.com/gdamore/tcell/v2"
)

type Editor struct {
	lines          string
	curr_row       int
	curr_col       int
	default_cursor string
	width          int
	height         int
}

func NewEditor() Editor {
	return Editor{
		lines:          "",
		curr_row:       0,
		curr_col:       0,
		default_cursor: "|",
		width:          0,
		height:         0,
	}
}

func main() {

	editor := NewEditor()
	screen, err := tcell.NewScreen()
	if err != nil {
		log.Fatal("Couldn't create a screen: " + err.Error())
	}

	defer screen.Fini()

	if err := screen.Init(); err != nil {
		log.Fatal("Couldn't initialize the screen " + err.Error())
	}

	style := tcell.StyleDefault.Background(tcell.ColorBlack).Foreground(tcell.ColorWhite)
	screen.SetStyle(style)

	screen.Clear()
	drawText(screen, editor.curr_col, editor.curr_row, style, editor.default_cursor)

	for {

		ev := screen.PollEvent()

		switch ev := ev.(type) {
		case *tcell.EventKey:
			{

				switch ev.Key() {
				case tcell.KeyCtrlQ:
					return
				case tcell.KeyEnter:
					{
						editor.curr_row += 1
						editor.curr_col = 0
						break
					}
				case tcell.KeyBackspace:
				case tcell.KeyBackspace2:
					{
						if editor.curr_col > 0 {
							editor.curr_col -= 1
						}

						if len(editor.lines) > 0 {

							editor.lines = editor.lines[0 : len(editor.lines)-1]
						}

						drawText(screen, 0, editor.curr_row, style, editor.lines)
						drawText(screen, editor.curr_col, editor.curr_row, style, editor.default_cursor)
						drawText(screen, editor.curr_col+1, editor.curr_row, style, " ")

						break
					}
				default:
					{

						editor.curr_col += 1
						editor.lines += string(ev.Rune())
						drawText(screen, 0, editor.curr_row, style, editor.lines)
						drawText(screen, editor.curr_col, editor.curr_row, style, editor.default_cursor)
					}
				}

				screen.Show()

			}
		case *tcell.EventResize:
			{
				screen.Sync()
			}
		}

	}

}

func drawText(screen tcell.Screen, x, y int, style tcell.Style, str string) {
	for i, c := range str {
		screen.SetContent(x+i, y, c, nil, style)
	}
}
