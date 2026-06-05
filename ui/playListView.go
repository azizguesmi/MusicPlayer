package ui

import (
	model "MuisicPlayer/Model"
	"MuisicPlayer/persistent"
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

type playListModle struct {
	PlayLists []model.PlayList
	cursor    int
	selected  map[int]struct{}
}

func createPlayListModel() (*playListModle, error) {
	ps, err := persistent.GetAllPlayLists()
	if err != nil {
		return nil, err
	}

	return &playListModle{
		PlayLists: ps,
		selected:  make(map[int]struct{}),
	}, nil
}

func (m playListModle) Init() tea.Cmd {
	return nil
}

func (m playListModle) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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
			if m.cursor < len(m.PlayLists)-1 {
				m.cursor++
			}
		case "enter":
			_, ok := m.selected[m.cursor]
			if ok {
				delete(m.selected, m.cursor)
			} else {
				m.selected[m.cursor] = struct{}{}
			}
		case "d":
			for i := range m.selected {
				persistent.DeletePlayList(m.PlayLists[i].Name)
			}
		}
	}
	return m, nil
}

func (m playListModle) View() string {

	s := "All Play lists\n\n"

	for i, playlist := range m.PlayLists {
		cursor := " "
		if m.cursor == i {
			cursor = ">"
		}
		checked := " "

		if _, ok := m.selected[i]; ok {
			checked = "X"
		}

		s += fmt.Sprintf("%s [%s] %s\n", cursor, checked, playlist.Name)
	}

	s += "press q to quit \n"

	return s
}
