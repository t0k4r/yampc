package main

import (
	"log"
	"yampc/client"
	"yampc/ui"

	"github.com/gdamore/tcell/v2"
)

type AlbumSearchScreen struct {
	c *client.Client
	i ui.Input
	l ui.List[*uiAlbum]
	s tcell.Screen
}

func NewAlbumSearchScreen(s tcell.Screen, c *client.Client) AlbumSearchScreen {
	w, h := s.Size()
	return AlbumSearchScreen{
		c: c,
		i: ui.NewIntput(1, 1, w),
		l: ui.NewList[*uiAlbum](1, 2, w, h-1),
		s: s,
	}
}
func (s *AlbumSearchScreen) Update(ev tcell.Event) {
	switch ev := ev.(type) {
	case *tcell.EventKey:
		switch ev.Key() {
		case tcell.KeyCtrlC:
			s.s.Clear()
			s.s.Fini()
			log.Fatal("CTRL+C")
		case tcell.KeyEnter:
			album := *s.l.Selected()
			if album != nil {
				s.c.PushAlbum(album.AlbumID)
				if !s.c.IsPaused() && len(s.c.Queue.Songs) == int(album.Songs) {
					s.c.Play()
				}
			}
		case tcell.KeyUp, tcell.KeyDown:
			s.l.Update(ev)
			s.Print()
		default:
			s.i.Update(ev)
			s.l.SetItems(maps(s.c.QueryAlbumByTitle(s.i.Value()), func(alb client.Album) *uiAlbum {
				return &uiAlbum{alb}
			}))
			s.Print()
		}
	case *tcell.EventResize:
		s.Resize()
	}
}
func (s *AlbumSearchScreen) Print() {
	s.l.Print()
	s.i.Print()
}
func (s *AlbumSearchScreen) Resize() {
	w, h := s.s.Size()
	s.i.Resize(w)
	s.l.Resize(w, h-1)
}
