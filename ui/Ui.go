package ui

import (
	"MuisicPlayer/Model"
	"MuisicPlayer/persistent"
	"bufio"
	"fmt"
	"os"
	"strings"
)

type View interface {
	ViewSong()
	ViewAllSong()
}
type Ui struct{

}
func (u *Ui) ViewSong(s model.Song) {
	fmt.Printf("Title : %s | Artist | %s | %s\n", s.Title, s.Artist, s.Lenght)
}

func (u *Ui) ViewAllSongs() {
	t, err := persistent.GetAllSongs()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("      Title      |      Artist      |     Time      ")
	for _,s := range t {
		u.ViewSong(s)
	}
}

func (u *Ui) ViewPlayList(p model.PlayList) {
	fmt.Printf("Title : %s\n", p.Name)

}

func (u *Ui) ViewAllPlayList() {
	t, err := persistent.GetAllPlayLists()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("------PlayLists------")
	fmt.Println("      Title      ")

	for _,p := range t {
		u.ViewPlayList(p)
	}
}

func (u *Ui) ViewSongsBySearch(search string) {
	t, err := persistent.GetAllSongs()
	if err != nil {
		fmt.Println(err)
	}
	for _,s := range t {
		if strings.Contains(s.Title, search) {
			u.ViewSong(s)
		}
	}
} 

func (u *Ui) ViewPlayListBySeach(search string) {
	t, err := persistent.GetAllPlayLists()
	if err != nil {
		fmt.Println(err)
	}
	for _,s := range t {
		if strings.Contains(s.Name, search) {
			u.ViewPlayList(s)
		}
	}
} 

func (u *Ui)Run() {
	fmt.Println("type [p] to see playLists")
	fmt.Println("type [s] to see songs")

	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()

	choice := sc.Text()
	for choice != "p" && choice != "s" {
		fmt.Println("wrong choice")
		sc.Scan()
		choice = sc.Text()
	}

	switch choice {
		case "p" :
			u.ViewAllPlayList()
		case "s" :
			u.ViewAllSongs()
	}
}


