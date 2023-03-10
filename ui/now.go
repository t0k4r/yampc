package ui

import "fmt"

type Now struct {
	x int
	y int
	w int
}

func NewNow(x, y, w int) Now {
	return Now{x, y, w}
}

func (n *Now) Print(isPaused bool, idx, qLen int, title string, pos, dur int) {
	txt := ""
	fmt.Printf("%v%v", setXY(n.x, n.y), padding(n.w))
	if !isPaused {
		txt = fmt.Sprintf("%v▶", setXY(n.x, n.y))
	} else {
		txt = fmt.Sprintf("%v⏸", setXY(n.x, n.y))
	}
	if qLen != 0 {
		idx++
	}
	txt = fmt.Sprintf("%v %v/%v ", txt, idx, qLen)
	txt = fmt.Sprintf("%v%v", txt, cut([]rune(italic(bold(title))), n.w-len(txt)))
	txt = fmt.Sprintf("%v%v", txt, padding(n.w-len(txt)-11))
	txt += fmt.Sprintf("%v%v/%v", setXY(n.x+n.w-11, n.y), Time(pos), Time(dur))
	fmt.Print(txt)
	fmt.Print(hideCurs())
}
func (n *Now) Move(x, y int) {
	n.x, n.y = x, y
}
func (n *Now) Resize(w int) {
	n.w = w
}
