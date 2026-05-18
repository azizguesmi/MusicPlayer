package persistent

import (
	"MuisicPlayer/Model"
	"encoding/json"
	"errors"
	"os"
)
var save_file = "saveSong.json"
var save_playList = "savePlayList.json"


func GetAllSongs () ([]model.Song, error)  {

	file, err := os.ReadFile(save_file)
	if err != nil {
		return nil, errors.New("Error while reading the json file") 
	}
	var t []model.Song
	if err = json.Unmarshal(file, &t); err != nil {
		return nil, errors.New("error reading json")
	}	
	return t, nil
}

func SaveSong(s model.Song) error {
	t, err := GetAllSongs()
	if err != nil {
		return err
	}

	t = append(t,s)
	t_json, err := json.MarshalIndent(t, "","  ")

	if err != nil {
		return err
	}
	if err = os.WriteFile(save_file, t_json, 0664); err != nil {
		return errors.New("error saving song")
	}
	return nil
}

func DeleteSong(name string) error {
	t, err := GetAllSongs()

	if err != nil {
		return err
	}

	for i,s := range t {
		if s.Title == name {
			t = append(t[:i-1], t[i+1:]... )
		}
	}
	t_json, err := json.MarshalIndent(t,""," ") 

	if err != nil {
		return errors.New("error writing to json")
	}

	if err = os.WriteFile(save_file, t_json, 0644); err != nil {
		return  errors.New("error writing to file")
	}
	return nil
}


func GetSong(title string) (model.Song, error) {
	
	t, err := GetAllSongs()
	if err != nil {
		return model.Song{}, err
	}

	for _,s := range t {
		if s.Title == title {
			return s,nil
		}
	}
	return model.Song{},errors.New("song not found")
}

func GetAllPlayLists() ([]model.PlayList, error) {
	f,err := os.ReadFile(save_playList)

	if err != nil {
		return nil, errors.New("error opening playlist files")
	}

	var t []model.PlayList
	if err = json.Unmarshal(f, &t); err != nil {
		return nil, errors.New("json error")
	}

	return t, nil
}

func getPlayListByTitle(title string) (*model.PlayList, error)  {
	t,err := GetAllPlayLists()
	if err != nil {
		return nil, err
	}
	for _,p := range t {
		if p.Name == title {
			return  &p,nil
		}
	}
	return nil, errors.New("playlist not found")
}

func AddPlayList(playList model.PlayList) error {
	_,err := getPlayListByTitle(playList.Name)
	if err != nil {
		if err.Error() == "playlist not found" {
			
			t, err := GetAllPlayLists()
			t = append(t, playList)
			t_json, err := json.MarshalIndent(t,""," ")
			if  err != nil {
				return errors.New("error e coding to json")
			}
			if err = os.WriteFile(save_playList, t_json, 0644); err != nil {
				return errors.New("error while saving to file")
			}
			return nil
		}
		return err
	}
	return errors.New("Play list already exists")
}

