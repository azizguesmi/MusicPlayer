package ui

import (
	"MuisicPlayer/Model"
	"MuisicPlayer/persistent"
	"fmt"
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


