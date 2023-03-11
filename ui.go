package main

import (
	"fmt"
	"yampc/client"
	"yampc/ui"
)

type uiSong struct {
	client.Song
}

func (u *uiSong) Col(id int) string {
	t := ""
	switch id {
	case 0:
		t = u.Title
	case 1:
		t = u.Artist
	case 2:
		t = u.Album
	case 3:
		t = ui.Time(int(u.Ms))
	}
	return t
}
func (u *uiSong) Cols() int {
	return 4
}

type uiAlbum struct {
	client.Album
}

func (u *uiAlbum) Col(id int) string {
	t := ""
	switch id {
	case 0:
		t = u.Title
	case 1:
		t = u.Artist
	case 2:
		t = fmt.Sprintf("%v", u.Year)
	}
	return t
}
func (u *uiAlbum) Cols() int {
	return 3
}
