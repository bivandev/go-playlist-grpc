package main

import (
	"database/sql"
	"errors"
	"sync"
	"time"

	_ "github.com/lib/pq"
)

type Song struct {
	Name     string
	Duration int64
}

type Node struct {
	Song *Song
	Prev *Node
	Next *Node
}

type Playlist struct {
	db      *sql.DB
	head    *Node
	tail    *Node
	cur     *Node
	playing bool
	pause   bool
	mutex   sync.Mutex
}

func NewPlaylist(db *sql.DB) (*Playlist, error) {
	if db == nil {
		return nil, errors.New("database connection is nil")
	}

	p := &Playlist{db: db}

	err := p.loadPlaylist()
	if err != nil {
		return nil, err
	}

	return p, nil
}

func (p *Playlist) loadPlaylist() error {
	rows, err := p.db.Query("SELECT name, duration FROM songs WHERE playlist_id = 1 ORDER BY id")
	if err != nil {
		return err
	}
	defer rows.Close()

	var songs []*Song
	for rows.Next() {
		var name string
		var duration int64
		err = rows.Scan(&name, &duration)
		if err != nil {
			return err
		}
		songs = append(songs, &Song{Name: name, Duration: duration})
	}

	for _, song := range songs {
		p.AddSong(song)
	}

	currentSong := ""
	err = p.db.QueryRow("SELECT current_song FROM playlist WHERE id = 1").Scan(&currentSong)
	if err != nil && err != sql.ErrNoRows {
		return err
	}

	if currentSong != "" {
		cur := p.head
		for cur != nil {
			if cur.Song.Name == currentSong {
				p.cur = cur
				break
			}
			cur = cur.Next
		}
	}

	return nil
}

func (p *Playlist) AddSong(s *Song) error {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	newNode := &Node{Song: s}

	if p.head == nil {
		p.head = newNode
		p.tail = newNode

		_, err := p.db.Exec("INSERT INTO playlist (id, current_song, playing, pause) VALUES ($1, $2, $3, $4)", 1, s.Name, false, false)
		if err != nil {
			return err
		}
	} else {
		newNode.Prev = p.tail
		p.tail.Next = newNode
		p.tail = newNode

		_, err := p.db.Exec("INSERT INTO songs (playlist_id, name, duration) VALUES ($1, $2, $3)", 1, s.Name, s.Duration)
		if err != nil {
			return err
		}
	}

	return nil
}

func (p *Playlist) Play() error {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	if p.cur == nil {
		p.cur = p.head
	}

	if p.cur == nil {
		return errors.New("playlist is empty")
	}

	if p.playing {
		if p.pause {
			p.pause = false
		} else {
			return errors.New("playlist is already playing")
		}
	}

	p.playing = true

	go func() {
		for p.cur != nil && p.playing {
			if !p.pause {
				time.Sleep(time.Duration(p.cur.Song.Duration))
				p.Next()
			} else {
				time.Sleep(time.Second)
			}
		}
	}()

	return nil
}

func (p *Playlist) Pause() {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	p.pause = true
}

func (p *Playlist) DelSong(name string) error {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	node := p.head
	for node != nil && p.pause {
		if node.Song.Name == name {
			if node.Prev == nil {
				p.head = node.Next
			} else {
				node.Prev.Next = node.Next
			}

			if node.Next == nil {
				p.tail = node.Prev
			} else {
				node.Next.Prev = node.Prev
			}

			if node == p.cur {
				p.Next()
			}

			_, err := p.db.Exec("DELETE FROM songs WHERE name = $1", name)
			if err != nil {
				return err
			}

			return nil
		}

		node = node.Next
	}

	return errors.New("song not found")
}

func (p *Playlist) Next() error {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	if p.cur == nil {
		return errors.New("playlist is empty")
	}

	p.cur = p.cur.Next

	if p.cur == nil {
		p.playing = false
		return errors.New("end of playlist")
	}

	// Retrieve the current song from the database
	var currentSong string
	err := p.db.QueryRow("SELECT current_song FROM playlist WHERE id = 1").Scan(&currentSong)
	if err != nil {
		return err
	}

	// Update the current song in the database
	_, err = p.db.Exec("UPDATE playlist SET current_song = $1 WHERE id = 1", p.cur.Song.Name)
	if err != nil {
		return err
	}

	if currentSong != p.cur.Song.Name {
		p.playing = true
		p.pause = false
	}

	return nil
}

func (p *Playlist) Prev() error {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	if p.cur == nil {
		return errors.New("playlist is empty")
	}

	if p.cur.Prev == nil {
		return errors.New("beginning of playlist")
	}

	p.cur = p.cur.Prev

	// Retrieve the current song from the database
	var currentSong string
	err := p.db.QueryRow("SELECT current_song FROM playlist WHERE id = 1").Scan(&currentSong)
	if err != nil {
		return err
	}

	// Update the current song in the database
	_, err = p.db.Exec("UPDATE playlist SET current_song = $1 WHERE id = 1", p.cur.Song.Name)
	if err != nil {
		return err
	}

	if currentSong != p.cur.Song.Name {
		p.playing = true
		p.pause = false
	}

	return nil
}
