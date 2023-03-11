package main

import (
	"log"
	"yampc/client"
	"yampc/ui"

	"github.com/gdamore/tcell/v2"
)

type SongSearchScreen struct {
	c *client.Client
	i ui.Input
	l ui.List[*uiSong]
	s tcell.Screen
}

func NewSongSearchScreen(s tcell.Screen, c *client.Client) SongSearchScreen {
	w, h := s.Size()
	return SongSearchScreen{
		c: c,
		i: ui.NewIntput(1, 1, w),
		l: ui.NewList[*uiSong](1, 2, w, h-1),
		s: s,
	}
}
func (s *SongSearchScreen) Update(ev tcell.Event) {
	switch ev := ev.(type) {
	case *tcell.EventKey:
		switch ev.Key() {
		case tcell.KeyCtrlC:
			s.s.Clear()
			s.s.Fini()
			log.Fatal("CTRL+C")
		case tcell.KeyEnter:
			song := *s.l.Selected()
			if song != nil {
				s.c.PushSong(song.SongID)
				if !s.c.IsPaused() && len(s.c.Queue.Songs) == 1 {
					s.c.Play()
				}
			}
			s.i.Update(ev)
		case tcell.KeyUp, tcell.KeyDown:
			s.l.Update(ev)
			s.Print()
		default:
			s.i.Update(ev)
			s.l.SetItems(maps(s.c.QuerySongByTitle(s.i.Value()), func(alb client.Song) *uiSong {
				return &uiSong{alb}
			}))
			s.Print()
		}

	case *tcell.EventResize:
		s.Resize()
	}
}
func (s *SongSearchScreen) Print() {
	s.l.Print()
	s.i.Print()
}
func (s *SongSearchScreen) Resize() {
	w, h := s.s.Size()
	s.i.Resize(w)
	s.l.Resize(w, h-1)
}
