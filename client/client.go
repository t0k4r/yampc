package client

import (
	"fmt"
	"log"
	"time"

	"github.com/imroc/req/v3"
)

type Song struct {
	SongID   uint   `json:"song_id"`
	AlbumID  uint   `json:"album_id"`
	ArtistID uint   `json:"artist_id"`
	Title    string `json:"title"`
	Album    string `json:"album"`
	Artist   string `json:"artist"`
	Ms       uint64 `json:"ms"`
}
type Album struct {
	AlbumID  uint   `json:"album_id"`
	Title    string `json:"title"`
	ArtistID uint   `json:"artis_idt"`
	Artist   string `json:"artist"`
	Songs    uint   `json:"songs"`
	Year     uint   `json:"year"`
}

type Now struct {
	Id    uint `json:"id"`
	Dur   uint `json:"dur"`
	Pos   uint `json:"pos"`
	Pause bool `json:"pause"`
}

type Queue struct {
	Index uint   `json:"index"`
	Songs []Song `json:"songs"`
}

type Client struct {
	addr  string
	conn  *req.Client
	now   *Now
	queue *Queue
}

type Query struct {
	Like string `json:"like"`
}

func New(addr string) *Client {
	c := Client{
		addr: addr,
		conn: req.NewClient(),
		now:  nil,
		queue: &Queue{
			Index: 0,
			Songs: []Song{},
		},
	}
	go c.update()
	return &c
}
func (c *Client) update() {
	for {
		time.Sleep(time.Microsecond * 500)
		c.getNow()
		c.getQueue()
	}
}
func (c *Client) getNow() {
	c.now = func() *Now {
		req, err := c.conn.R().Get(fmt.Sprintf("http://%v/ply/now", c.addr))
		if err != nil {
			log.Panic(err)
		}
		switch req.StatusCode {
		case 200:
			var now Now
			req.UnmarshalJson(&now)
			return &now
		default:
			return nil
		}
	}()
}
func (c *Client) getQueue() {
	c.queue = func() *Queue {
		req, err := c.conn.R().Get(fmt.Sprintf("http://%v/ply/queue", c.addr))
		if err != nil {
			log.Panic(err)
		}
		switch req.StatusCode {
		case 200:
			var now Queue
			req.UnmarshalJson(&now)
			return &now
		default:
			return nil
		}
	}()
}

func (c *Client) Play() {
	c.conn.R().Post(fmt.Sprintf("http://%v/ply/play", c.addr))
	c.getNow()
}
func (c *Client) Next() {
	c.conn.R().Post(fmt.Sprintf("http://%v/ply/next", c.addr))
	c.getNow()
}
func (c *Client) Prev() {
	c.conn.R().Post(fmt.Sprintf("http://%v/ply/prev", c.addr))
	c.getNow()
}
func (c *Client) Index(index uint) {
	c.conn.R().Post(fmt.Sprintf("http://%v/ply/index/%v", c.addr, index))
	c.getNow()
}
func (c *Client) Delete(index uint) {
	c.conn.R().Delete(fmt.Sprintf("http://%v/ply/index/%v", c.addr, index))
	c.getNow()
}
func (c *Client) PushSong(id uint) {
	c.conn.R().Post(fmt.Sprintf("http://%v/ply/queue/song/%v", c.addr, id))
	c.getQueue()
}
func (c *Client) PushAlbum(id uint) {
	c.conn.R().Post(fmt.Sprintf("http://%v/ply/queue/album/%v", c.addr, id))
	c.getQueue()
}

func (c *Client) Seek(forward bool, time time.Duration) {
	switch forward {
	case true:
		c.conn.R().Post(fmt.Sprintf("http://%v/ply/pos/seek/forw/%v", c.addr, time.Milliseconds()))
	case false:
		c.conn.R().Post(fmt.Sprintf("http://%v/ply/pos/seek/back/%v", c.addr, time.Milliseconds()))
	}
	c.getNow()
}
func (c *Client) Duration() time.Duration {
	c.getNow()
	if c.now != nil {
		return time.Duration(time.Millisecond * time.Duration(c.now.Dur))
	}
	return time.Second * 0
}
func (c *Client) Position() time.Duration {
	c.getNow()
	c.getNow()
	if c.now != nil {
		return time.Duration(time.Millisecond * time.Duration(c.now.Pos))
	}
	return time.Second * 0
}
func (c *Client) IsPaused() bool {
	c.getNow()
	if c.now != nil {
		return c.now.Pause
	} else {
		return false
	}
}
func (c *Client) SetPause(pause bool) {
	switch pause {
	case true:
		c.conn.R().Post(fmt.Sprintf("http://%v/ply/pause", c.addr))
	case false:
		c.conn.R().Post(fmt.Sprintf("http://%v/ply/unpause", c.addr))
	}
	c.getNow()
}
func (c *Client) Queue() Queue {
	return *c.queue
}
func (c *Client) QuerySongByID(id uint) *Song {
	req, err := c.conn.R().Get(fmt.Sprintf("http://%v/lib/song/%v", c.addr, id))
	if err != nil {
		log.Panic(err)
	}
	switch req.StatusCode {
	case 200:
		var song Song
		req.UnmarshalJson(&song)
		return &song
	default:
		return nil
	}

}
func (c *Client) QuerySongByTitle(like string) []Song {
	req, err := c.conn.R().SetBodyJsonMarshal(Query{like}).Post(fmt.Sprintf("http://%v/lib/song", c.addr))
	if err != nil {
		log.Panic(err)
	}
	var songs []Song
	err = req.UnmarshalJson(&songs)
	if err != nil {
		log.Fatal(err)
	}
	return songs
}
func (c *Client) QuerySongByAlbumID(id uint) []Song {
	req, err := c.conn.R().Get(fmt.Sprintf("http://%v/lib/song/album/%v", c.addr, id))
	if err != nil {
		log.Panic(err)
	}
	var songs []Song
	req.UnmarshalJson(&songs)
	return songs
}
func (c *Client) QueryAlbumByID(id uint) *Album {
	req, err := c.conn.R().Get(fmt.Sprintf("http://%v/lib/album/%v", c.addr, id))
	if err != nil {
		log.Panic(err)
	}
	switch req.StatusCode {
	case 200:
		var album Album
		req.UnmarshalJson(&album)
		return &album
	default:
		return nil
	}
}
func (c *Client) QueryAlbumByTitle(like string) []Album {
	req, err := c.conn.R().SetBodyJsonMarshal(Query{like}).Post(fmt.Sprintf("http://%v/lib/album", c.addr))
	if err != nil {
		log.Panic(err)
	}
	var albums []Album
	req.UnmarshalJson(&albums)
	return albums
}
