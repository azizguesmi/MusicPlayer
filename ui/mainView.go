package ui

import tea "github.com/charmbracelet/bubbletea"

type Page int

const (
	homePage Page = iota
	songPage
	playlistPage
)

type MainModel struct {
	page              Page
	songPageModel     songModel
	playListPageModle playListModle
}

func crateMainModel() MainModel {
	return MainModel{page: homePage}
}

func (m MainModel) Init() tea.Cmd {
	return nil
}

func (m MainModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		case "s":
			m.page = songPage
			return m, nil
		case "p":
			m.page = playlistPage
			return m, nil
		}
	}
	return m, nil
}

func (m MainModel) View() string {

}
