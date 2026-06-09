package player

import (
	model "MuisicPlayer/Model"
	"errors"
	"math/rand"
	"time"
)

type Queue struct {
	Trucks  []model.Song
	Current int
}

func NewSong(path, title, artist string) (*model.Song, error) {
	if path == "" || title == "" || artist == "" {
		return nil, errors.New("a Song Field is empty")
	}
	return &model.Song{Path: path, Title: title, Artist: artist}, nil
}

func NewQueue() *Queue {
	return &Queue{Trucks: []model.Song{}, Current: 0}
}

func NewQueueWithAr(s []model.Song) *Queue {
	return &Queue{Trucks: s, Current: 0}
}

func (q *Queue) EnQueue(s model.Song) {
	q.Trucks = append(q.Trucks, s)
}

func (q *Queue) DeQueue() model.Song {
	s := q.Trucks[len(q.Trucks)-1]
	q.Trucks = q.Trucks[:len(q.Trucks)-1]
	return s
}

func (q *Queue) Shuffle() {
	rand.Seed(time.Now().UnixNano())
	for i := len(q.Trucks) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		q.Trucks[i], q.Trucks[j] = q.Trucks[j], q.Trucks[i]
	}
}

func (q *Queue) IsEmpty() bool {
	return len(q.Trucks) == 0
}
