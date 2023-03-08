package main

import (
	"fmt"
	"yampc/client"
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
		secs := u.Ms / 1000
		min := secs / 60
		sec := secs % 60
		if min < 10 {
			t = fmt.Sprintf("0%v", min)
		} else {
			t = fmt.Sprintf("%v", min)
		}
		if sec < 10 {
			t = fmt.Sprintf("%v:0%v", t, sec)
		} else {
			t = fmt.Sprintf("%v:%v", t, sec)
		}
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
