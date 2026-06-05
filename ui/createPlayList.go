package ui

import (
	model "MuisicPlayer/Model"
	"MuisicPlayer/persistent"

	tea "github.com/charmbracelet/bubbletea"
)

type CreatePlayListModel struct {
	name     string
	songs    []model.Song
	selected map[int]struct{}
	cursor   int
}

func NewCreatePlayListModel(name string) (*CreatePlayListModel, error) {
	songs, err := persistent.GetAllSongs()
	if err != nil {
		return nil, err
	}
	return &CreatePlayListModel{
		name:     name,
		songs:    songs,
		selected: make(map[int]struct{}),
	}, nil
}

func (m *CreatePlayListModel) Init() tea.Cmd {
	return nil
}

func (m *CreatePlayListModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return nil, tea.Quit
		case "enter":
			_, ok := m.selected[m.cursor]
			if ok {
				delete(m.selected, m.cursor)
			} else {
				m.selected[m.cursor] = struct{}{}
			}
		case "up":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down":
			if m.cursor < len(m.songs)-1 {
				m.cursor++
			}
		case "c":
			sel := make([]model.Song, 0, len(m.selected))
			for i := range m.selected {
				sel = append(sel, m.songs[i])
			}
			playlist := model.PlayList{Name: m.name, Songs: sel}
			err := persistent.AddPlayList(playlist)
			if err != nil {
				return nil, tea.Quit
			}
			return nil, tea.Quit
		}
	}
	return m, nil
}

func (m *CreatePlayListModel) View() string {
	s := "Create a Playlist named : " + m.name + "\n + Select songs to add to the playlist\n"

	for i, song := range m.songs {
		selected := " "
		if _, ok := m.selected[i]; ok {
			selected = "x"
		}
		cursor := " "
		if i == m.cursor {
			cursor = ">"
		}
		s += selected + cursor + " " + song.Title + " - " + song.Artist + "\n"
	}
	return s
}
