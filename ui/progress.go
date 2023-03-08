package ui

import "fmt"

type Progress struct {
	x   int
	y   int
	w   int
	max int
	now int
}

func NewProgress(x, y, w int) Progress {
	return Progress{x, y, w, 0, 0}
}

func (p *Progress) SetMax(max int) {
	p.max = max
}
func (p *Progress) SetNow(now int) {
	p.now = now
}
func (p *Progress) Print() {
	proc := float64(p.now) / float64(p.max)
	for i := 0; i < p.w; i++ {
		wproc := float64(i) / float64(p.w)
		if wproc < proc {
			fmt.Printf("%v█", setXY(p.x+i, p.y))
		} else {
			fmt.Printf("%v_", setXY(p.x+i, p.y))
		}
	}
}
func (p *Progress) MaxMin() (int, int) {
	return p.max, p.now
}
func (p *Progress) Resize(w int) {
	p.w = w
}
func (p *Progress) Move(x, y int) {
	p.x, p.y = x, y
}
