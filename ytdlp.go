package main

type YTVideo struct {
	Id    string `json:"id"`
	Url   string `json:"url"`
	Title string `json:"title"`
}

type YTPlaylist struct {
	Title   string    `json:"title"`
	Entries []YTVideo `json:"entries"`
}