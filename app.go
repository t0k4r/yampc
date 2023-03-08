package main

import (
	"log"
	"time"
	"yampc/client"

	"github.com/gdamore/tcell/v2"
)

type ScreenType int

const (
	Album ScreenType = iota
	Song
	Queue
)

type App struct {
	ass       AlbumSearchScreen
	sss       SongSearchScreen
	qs        QueueScreen
	s         tcell.Screen
	tickev    chan (interface{})
	tcelev    chan (tcell.Event)
	nowScreen ScreenType
}

func NewApp() App {
	s, err := tcell.NewScreen()
	if err != nil {
		log.Fatalf("%+v", err)
	}
	if err := s.Init(); err != nil {
		log.Fatalf("%+v", err)
	}
	c := client.New("127.0.0.1:2137")
	return App{
		ass:       NewAlbumSearchScreen(s, c),
		sss:       NewSongSearchScreen(s, c),
		qs:        NewQueueScreen(s, c),
		s:         s,
		tickev:    make(chan interface{}),
		tcelev:    make(chan tcell.Event),
		nowScreen: Queue,
	}
}
func (a *App) tick() {
	for {
		time.Sleep(time.Millisecond * 500)
		a.tickev <- true
	}

}
func (a *App) event() {
	for {
		ev := a.s.PollEvent()
		a.tcelev <- ev
	}
}
func (a *App) run() {
	go a.tick()
	go a.event()
	for {
		select {
		case <-a.tickev:
			switch a.nowScreen {
			case Queue:
				a.qs.Tick()
				a.qs.Print()
			case Song:
				a.sss.Print()
			case Album:
				a.ass.Print()
			}
		case ev := <-a.tcelev:
			switch ev := ev.(type) {
			case *tcell.EventKey:
				switch ev.Key() {
				case tcell.KeyF1:
					a.nowScreen = Queue
				case tcell.KeyF2:
					a.nowScreen = Song
				case tcell.KeyF3:
					a.nowScreen = Album
				case tcell.KeyCtrlC:
					a.s.Clear()
					a.s.Fini()
					log.Fatal("\x1B?25h CTRL+C")
				}
			case *tcell.EventResize:
				a.ass.Resize()
				a.sss.Resize()
				a.qs.Resize()
			}
			switch a.nowScreen {
			case Queue:
				a.qs.Update(ev)
				a.qs.Print()
			case Song:
				a.sss.Update(ev)
				a.sss.Print()
			case Album:
				a.ass.Update(ev)
				a.ass.Print()
			}
		}

	}
}
