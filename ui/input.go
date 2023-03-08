package ui

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
)

type Input struct {
	x     int
	y     int
	w     int
	idx   int
	value []rune
}

func NewIntput(x, y, w int) Input {
	return Input{
		x:     x,
		y:     y,
		w:     w,
		idx:   0,
		value: []rune{},
	}
}
func (i *Input) Update(ev tcell.Event) {
	switch ev := ev.(type) {
	case *tcell.EventKey:
		switch ev.Key() {
		case tcell.KeyBackspace, tcell.KeyBackspace2:
			if i.idx != 0 {
				i.value = append(i.value[:i.idx-1], i.value[i.idx:]...)
				i.idx--
			}
		case tcell.KeyDelete:
			if i.idx < len(i.value) {
				i.value = append(i.value[:i.idx], i.value[i.idx+1:]...)
			}
		case tcell.KeyLeft:
			if i.idx != 0 {
				i.idx--
			}
		case tcell.KeyRight:
			if i.idx < len(i.value) {
				i.idx++
			}
		case tcell.KeyCtrlC:
			panic("CTRL+C")
		case tcell.KeyRune:
			i.value = append(append(i.value[:i.idx], ev.Rune()), i.value[i.idx:]...)
			i.idx++
		}
	}
}
func (i *Input) Print() {
	if len(i.value) < i.w-2 {
		fmt.Printf(italic("%v> %v%v%v%v"), setXY(i.x, i.y), string(i.value), padding(i.w-2-len(i.value)), setXY(i.x+i.idx+2, i.y), showCursr())
	} else if i.idx < i.w-2 {
		fmt.Printf(italic("%v> %v…%v%v"), setXY(i.x, i.y), string(i.value[:i.w-3]), setXY(i.x+i.idx+2, i.y), showCursr())
	} else if i.idx == len(i.value) {
		fmt.Printf(italic("%v> …%v%v%v"), setXY(i.x, i.y), string(i.value[i.idx-(i.w-3):]), setXY(i.x+i.w, i.y), showCursr())
	} else {
		w := i.w - 4
		text := string(i.value[i.idx-w : i.idx])
		fmt.Printf(italic("%v> …%v…%v%v"), setXY(i.x, i.y), text, setXY(i.x+i.w-1, i.y), showCursr())
	}
}
func (i *Input) Value() string {
	return string(i.value)
}

func (i *Input) Resize(w int) {
	i.w = w
}
func (i *Input) Move(x, y int) {
	i.x, i.y = x, y
}
