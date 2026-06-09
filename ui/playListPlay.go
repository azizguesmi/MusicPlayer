package ui

import (
	model "MuisicPlayer/Model"
	"MuisicPlayer/persistent"
	"MuisicPlayer/player"
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

type PlayListPlay struct {
	playlist    *model.PlayList
	currentSong *model.Song
	queue       *player.Queue
	player      *player.Player
}

/* ---------------- Messages ---------------- */

type changeMusicMsg struct {
	song model.Song
}

type musicEndingMsg struct{}

func sendChangeMusicMsg(song model.Song) tea.Cmd {
	return func() tea.Msg {
		return changeMusicMsg{song: song}
	}
}

func sendMusicEndingMsg() tea.Cmd {
	return func() tea.Msg {
		return musicEndingMsg{}
	}
}

/* ---------------- Constructor ---------------- */

func NewPlayListPlay(playlistname string) (*PlayListPlay, error) {

	playlist, err := persistent.GetPlayListByName(playlistname)
	if err != nil {
		return nil, err
	}

	queue := player.NewQueueWithAr(playlist.Songs)

	return &PlayListPlay{
		playlist: playlist,
		queue:    queue,
	}, nil
}

/* ---------------- Bubble Tea ---------------- */

func (p *PlayListPlay) Init() tea.Cmd {
	fmt.Println("init")

	if p.queue.IsEmpty() {
		return tea.Quit
	}

	song := p.queue.DeQueue()
	return sendChangeMusicMsg(song)
}

func (p *PlayListPlay) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	/* ---------- change song ---------- */
	case changeMusicMsg:
		p.currentSong = &msg.song

		pl, err := player.NewPlaeyr(msg.song.Path)
		if err != nil {
			fmt.Println("player error:", err)
			return p, nil
		}

		p.player = pl
		go p.player.LaunchSong()

		return p, nil

	/* ---------- key controls ---------- */
	case tea.KeyMsg:

		switch msg.String() {

		case "n":
			if p.queue.IsEmpty() {
				return p, tea.Quit
			}
			song := p.queue.DeQueue()
			return p, sendChangeMusicMsg(song)

		case "p", "space":
			if p.player == nil {
				return p, nil
			}
			if p.player.IsPaused() {
				p.player.Resume()
			} else {
				p.player.Pause()
			}
			return p, nil

		case "ctrl+c", "q":
			return p, tea.Quit
		}

	/* ---------- auto next song ---------- */
	case musicEndingMsg:
		if p.queue.IsEmpty() {
			return p, tea.Quit
		}
		song := p.queue.DeQueue()
		return p, sendChangeMusicMsg(song)
	}

	/* ---------- safety check ---------- */
	if p.player != nil && p.player.IsClosed() {
		return p, sendMusicEndingMsg()
	}

	return p, nil
}

/* ---------------- View ---------------- */

func (p *PlayListPlay) View() string {
	if p.currentSong == nil {
		return "Loading playlist...\n"
	}
	return fmt.Sprintf(
		"Playlist: %s\nNow playing: %s",
		p.playlist.Name,
		p.currentSong.Title,
	)
}
