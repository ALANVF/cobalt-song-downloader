package main

import (
	"fmt"
	"strings"
	"bytes"
	"os"
	"os/exec"
	"encoding/json"
	"net/http"
	"time"
	"regexp"
)

func getSong(client *http.Client, url string) *CobaltResponse {
	request := CobaltRequest{
		Url: url,
		Quality: "max",
		Format: CobaltFormat_mp3,
		FilePattern: CobaltFilePattern_basic,
		IsAudioOnly: true,
	}

	requestString, err := json.Marshal(request)
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest(
		"POST",
		"https://co.wuk.sh/api/json",
		bytes.NewBuffer(requestString),
	)
	if err != nil {
		panic(err)
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	
	response := new(CobaltResponse)
	if err = json.NewDecoder(resp.Body).Decode(response); err != nil {
		panic(err)
	}

	return response
}

func createSongFile(title string, response *CobaltResponse) (filename string) {
	rxBrackets := regexp.MustCompile("\\s*\\[(.+?)\\]\\s*")

	resp, err := http.Get(response.Url)
	if err != nil {
		panic(err)
	}
	filename = title
	filename = strings.ReplaceAll(filename, "*", "")
	filename = strings.ReplaceAll(filename, ":", "")
	filename = strings.ReplaceAll(filename, `"`, "")
	filename = strings.ReplaceAll(filename, "|", "")
	filename = strings.ReplaceAll(filename, " ??", "")
	filename = strings.ReplaceAll(filename, "?? ", "")
	filename = strings.ReplaceAll(filename, "??", "")
	filename = strings.ReplaceAll(filename, "?", "")
	filename = strings.ReplaceAll(filename, "/", "")
	filename = strings.ReplaceAll(filename, "\\", "")
	filename = strings.ReplaceAll(filename, " (Original Song)", "")
	filename = strings.ReplaceAll(filename, " (Original Mix)", "")

	filename = rxBrackets.ReplaceAllString(filename, " $1")
	if !strings.Contains(filename, ".mp3") {
		filename += ".mp3"
	}

	fmt.Println("Downloading song: " + filename);
	
	file, err := os.Create("./" + filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	
	file.ReadFrom(resp.Body)
	file.Sync()
	
	
	fmt.Println("Song downloaded: " + filename);

	return
}

func main() {
	client := &http.Client{}

	fmt.Print("Enter playlist url: ")
	var url string
	_, err := fmt.Scanln(&url)
	if err != nil {
		panic(err)
	}

	// This assumes that yt-dlp is in your PATH
	cmd := exec.Command("yt-dlp.exe", url, "-s", "--flat-playlist", "--dump-single-json")

	var playlist YTPlaylist
	output, err := cmd.Output()
	if err != nil {
		// I have no clue why it returns an exit status of 1 sometimes
		fmt.Println(err)
		//panic(err)
	}
	json.Unmarshal(output, &playlist)

	playlist.Title = strings.ReplaceAll(playlist.Title, ":", "")
	playlist.Title = strings.ReplaceAll(playlist.Title, " / ", " ")
	playlist.Title = strings.ReplaceAll(playlist.Title, "/", " ")
	os.Mkdir("./" + playlist.Title, os.ModeDir)
	os.Chdir("./" + playlist.Title)

	fmt.Println("Downloading playlist: " + playlist.Title)

	for _, entry := range playlist.Entries {
		// Weird bug at the end of playlists sometimes
		if entry.Title == "" {
			fmt.Println(entry)
			break
		}

		song := getSong(client, entry.Url)

		if song.Status == CobaltStatus_error {
			panic("error: something went wrong downloading the song: " + entry.Title)
		} else if song.Status == CobaltStatus_rate_limit {
			fmt.Println("Rate limited. Pausing for 5 seconds")
			dur, _ := time.ParseDuration("5s")
			time.Sleep(dur)
			song = getSong(client, entry.Url)
		}

		_ = createSongFile(entry.Title, song)
	}
}