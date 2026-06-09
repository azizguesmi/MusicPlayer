package ui

import (
	"errors"

	tea "github.com/charmbracelet/bubbletea"
)

func Run(args []string) error {
	if len(args) == 0 {
		return errors.New("invalid args")
	}

	switch args[0] {

	case "-s":
		m, err := createSongModel(nil)
		if err != nil {
			return err
		}
		_, err = tea.NewProgram(m).Run()
		return err

	case "-p":
		m, err := createPlayListModel()
		if err != nil {
			return err
		}
		_, err = tea.NewProgram(m).Run()
		return err

	case "-create":
		if len(args) < 2 {
			return errors.New("missing playlist name")
		}
		m, err := NewCreatePlayListModel(args[1])
		if err != nil {
			return err
		}
		_, err = tea.NewProgram(m).Run()
		return err

	case "-pl":
		if len(args) < 2 {
			return errors.New("missing playlist name")
		}
		m, err := NewPlayListPlay(args[1])
		if err != nil {
			return err
		}
		_, err = tea.NewProgram(m).Run()
		return err
	}

	return errors.New("unknown command")
}
