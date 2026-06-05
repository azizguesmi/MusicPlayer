package ui

import (
	"errors"

	tea "github.com/charmbracelet/bubbletea"
)

func Run(args []string) error {

	if len(args) == 0 {
		return errors.New("invalid args")
	}

	if args[0] == "-s" {
		model, err := createSongModel(nil)
		if err != nil {
			return err
		}
		_, err = tea.NewProgram(model).Run()
		if err != nil {
			return err
		}
	}

	if args[0] == "-p" {
		model, err := createPlayListModel()
		if err != nil {
			return err
		}
		_, err = tea.NewProgram(model).Run()
		if err != nil {
			return err
		}
	}

	if args[0] == "-play" && args[1] != "" {
		//player view
	}

	if args[0] == "-create" && args[1] != "" {
		model, err := NewCreatePlayListModel(args[1])
		if err != nil {
			return err
		}
		_, err = tea.NewProgram(model).Run()
		if err != nil {
			return err
		}
	}

	return nil
}
