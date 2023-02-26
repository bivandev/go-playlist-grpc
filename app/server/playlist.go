package main

import (
	"errors"
	"sync"
	"time"
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
	head    *Node
	tail    *Node
	cur     *Node
	playing bool
	pause   bool
	mutex   sync.Mutex
}

func NewPlaylist() *Playlist {
	return &Playlist{}
}

func (p *Playlist) AddSong(s *Song) {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	newNode := &Node{Song: s}

	if p.head == nil {
		p.head = newNode
		p.tail = newNode
	} else {
		newNode.Prev = p.tail
		p.tail.Next = newNode
		p.tail = newNode
	}
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

	return p.Play()
}

func (p *Playlist) Prev() error {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	if p.cur == nil {
		return errors.New("playlist is empty")
	}

	p.cur = p.cur.Prev

	if p.cur == nil {
		p.playing = false
		return errors.New("beginning of playlist")
	}

	return p.Play()
}
