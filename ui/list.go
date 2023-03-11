package ui

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
)

type ListItem interface {
	Col(int) string
	Cols() int
}

type List[T ListItem] struct {
	x     int
	y     int
	w     int
	h     int
	idx   int
	items []T
}

func NewList[T ListItem](x, y, w, h int) List[T] {
	return List[T]{x, y, w, h, 0, nil}
}
func (l *List[T]) SetItems(items []T) {
	l.items = items
	if len(l.items) <= l.idx {
		l.idx = len(l.items) - 1
	}
}
func (l *List[T]) Selected() *T {
	if len(l.items) != 0 {
		return &l.items[l.idx]
	}
	return nil
}
func (l *List[T]) Index() int {
	return l.idx
}
func (l *List[T]) Update(ev tcell.Event) {
	switch ev := ev.(type) {
	case *tcell.EventKey:
		switch ev.Key() {
		case tcell.KeyUp:
			if l.idx > 0 {
				l.idx--
			}
		case tcell.KeyDown:
			if l.idx < len(l.items)-1 {
				l.idx++
			}
		case tcell.KeyCtrlC:
			panic("CTRL+C")
		}
	}
}
func (l *List[T]) Print() {
	var disp []T
	start := 0
	if len(l.items) < l.h-1 {
		disp = l.items
	} else {
		if l.idx < l.h-1 {
			disp = l.items[:l.h-1]
		} else {
			start = l.idx - l.h + 2
			disp = append(disp, l.items[start:l.idx+1]...)
		}
	}
	if len(disp) != 0 {
		cols := disp[0].Cols()
		endCol := 0
		for c := 0; c < cols; c++ {
			endCol = c
		}
		endColW := len(disp[0].Col(endCol))
		collW := (l.w - endColW) / (cols - 1)
		endPad := l.w - (collW * (disp[0].Cols() - 1)) - endColW
		endIdx := disp[0].Cols() - 1
		lastI := 0
		for i, item := range disp {
			text := ""
			for j := 0; j < cols; j++ {
				coll := cut([]rune(item.Col(j)), collW)
				if j == endIdx {
					coll = fmt.Sprintf("%v%v", padding(endPad), coll)
				}
				if j == cols-1 {
					text += fmt.Sprintf("%v%v", setXY(l.x+collW*(j), l.y+i), coll)

				} else {
					text += fmt.Sprintf("%v%v%v", setXY(l.x+collW*(j), l.y+i), coll, padding(collW))

				}
			}
			if i+start == l.idx {
				text = italic(bold(text))
			}
			fmt.Print(text)
			lastI = i
		}
		lastI++
		for i := lastI; i < l.h; i++ {
			fmt.Printf("%v%v", setXY(l.x, l.y+i), padding(l.w))
		}

	} else {
		for i := 0; i < l.h; i++ {
			fmt.Printf("%v%v", setXY(l.x, l.y+i), padding(l.w))
		}
	}
	if len(l.items) != 0 {
		fmt.Printf(italic("%v%v%v%v/%v"), setXY(l.x, l.y+l.h-1), padding(l.w), setXY(l.x, l.y+l.h-1), l.idx+1, len(l.items))
	} else {
		fmt.Printf(italic("%v%v%v%v0/0"), setXY(l.x, l.y+l.h-1), padding(l.w), setXY(l.x, l.y+l.h-1), clearLine())
	}
}

func (l *List[T]) Resize(w, h int) {
	l.w, l.h = w, h
}
func (l *List[T]) Move(x, y int) {
	l.x, l.y = x, y
}
