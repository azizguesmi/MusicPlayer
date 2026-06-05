package ui

import (
	model "MuisicPlayer/Model"
	"MuisicPlayer/persistent"
	"MuisicPlayer/player"
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

type songModel struct {
	songs    []model.Song
	cursor   int
	selected map[int]struct{}
}

func createSongModel(songs []model.Song) (*songModel, error) {
	if songs == nil {
		s, err := persistent.GetAllSongs()

		if err != nil {
			return nil, err
		}

		return &songModel{
			songs:    s,
			selected: make(map[int]struct{}),
		}, nil
	}

	return &songModel{
		songs:    songs,
		selected: make(map[int]struct{}),
	}, nil
}

func (m songModel) Init() tea.Cmd {
	return nil
}

func (m songModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:

		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			if m.cursor < len(m.songs)-1 {
				m.cursor++
			}
		case "enter":
			_, ok := m.selected[m.cursor]
			if ok {
				delete(m.selected, m.cursor)

			} else if len(m.selected) == 0 {
				m.selected[m.cursor] = struct{}{}
			}

		case "p":
			for i := range m.selected {
				curent_song := m.songs[i]

				p, err := player.NewPlaeyr(curent_song.Path)
				if err != nil {
					return m, nil
				}
				go func() {
					p.Play()
					p.Wait()
					p.Close()
				}()
			}
		}
	}
	return m, nil
}

func (m songModel) View() string {
	s := "All songs\n\n"
	for i, song := range m.songs {
		cursor := " "
		if m.cursor == i {
			cursor = ">"
		}
		checked := " "
		if _, ok := m.selected[i]; ok {
			checked = "X"
		}

		s += fmt.Sprintf("%s [%s] %s %s %s\n", cursor, checked, song.Title, song.Artist, song.Lenght)
	}

	s += "press q to quit \n"
	return s
}

func Run() {
	m, err := createSongModel(nil)
	if err != nil {
		fmt.Println("error %v", err)
	}
	p := tea.NewProgram(*m)
	if _, err := p.Run(); err != nil {
		fmt.Println("error %v", err)
		os.Exit(1)
	}
}
