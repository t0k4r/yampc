package main

import (
	"log"
	"time"
	"yampc/client"
	"yampc/ui"

	"github.com/gdamore/tcell/v2"
)

type QueueScreen struct {
	c *client.Client
	l ui.List[*uiSong]
	p ui.Progress
	n ui.Now
	s tcell.Screen
}

func NewQueueScreen(s tcell.Screen, c *client.Client) QueueScreen {
	w, h := s.Size()
	return QueueScreen{
		c, ui.NewList[*uiSong](1, 1, w, h-2), ui.NewProgress(1, h-1, w), ui.NewNow(1, h, w), s,
	}
}
func (q *QueueScreen) Tick() {
	q.l.SetItems(maps(q.c.Queue().Songs, func(sng client.Song) *uiSong {
		return &uiSong{sng}
	}))
	q.p.SetNow(int(q.c.Position().Seconds()))
	q.p.SetMax(int(q.c.Duration().Seconds()))
}

func (q *QueueScreen) Update(ev tcell.Event) {
	q.Tick()
	switch ev := ev.(type) {
	case *tcell.EventKey:
		switch ev.Key() {
		case tcell.KeyCtrlC:
			q.s.Clear()
			q.s.Fini()
			log.Fatal("CTRL+C")
		case tcell.KeyLeft:
			q.c.Seek(false, time.Second*10)
		case tcell.KeyRight:
			q.c.Seek(true, time.Second*10)
		case tcell.KeyRune:
			switch ev.Rune() {
			case 'P', 'p':
				q.c.Play()
			case '<':
				q.c.Prev()
			case '>':
				q.c.Next()
			case ' ':
				q.c.SetPause(!q.c.IsPaused())
			}
		case tcell.KeyEnter:
			idx := q.l.Index()
			if len(q.c.Queue().Songs) > idx {
				q.c.Index(uint(idx))
			}
		case tcell.KeyDelete:
			idx := q.l.Index()
			if len(q.c.Queue().Songs) > idx {
				q.c.Delete(uint(idx))
			}
		default:
			q.l.Update(ev)
		}
	case *tcell.EventResize:
		q.Resize()
	default:
	}
}
func (q *QueueScreen) Print() {
	q.l.Print()
	q.p.Print()
	title := ""
	pos := 0
	dur := 0
	if len(q.c.Queue().Songs) > int(q.c.Queue().Index) {
		sng := q.c.Queue().Songs[q.c.Queue().Index]
		pos = int(q.c.Position().Milliseconds())
		dur = int(sng.Ms)
		title = sng.Title
	}
	q.n.Print(q.c.IsPaused(), int(q.c.Queue().Index), len(q.c.Queue().Songs), title, pos, dur)
}
func (q *QueueScreen) Resize() {
	w, h := q.s.Size()
	q.l.Resize(w, h-2)
	q.p.Resize(w)
	q.n.Resize(w)
	q.p.Move(1, h-1)
	q.n.Move(1, h)
}
