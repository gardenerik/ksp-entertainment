package parsers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/spf13/viper"
	"log"
	"os/exec"
	"strings"
)

type YoutubeDLParser struct{}

func (y YoutubeDLParser) CanParse(url string) bool {
	return strings.Contains(url, "youtu")
}

func (y YoutubeDLParser) GetStreamURL(url string) string {
	return url
}

func (y YoutubeDLParser) GetName(url string) string {
	log.Printf("Invoking youtube-dl to get video name for %v\n", url)
	cmd := exec.Command(viper.GetString("binaries.youtube_dl"), "-e", url)
	var out bytes.Buffer
	cmd.Stdout = &out

	err := cmd.Run()
	if err != nil {
		return fmt.Sprintf("(yt-dl err: %v)", err)
	}

	return strings.Trim(out.String(), "\n ")
}

type YoutubeVideoMetadata struct {
	Id          string      `json:"id"`
	Title       string      `json:"title"`
	Duration    float64     `json:"duration"`
}

func (m YoutubeVideoMetadata) URL() string {
	return fmt.Sprintf("https://youtube.com/watch?v=%v", m.Id)
}

func (y YoutubeDLParser) GetPlaylistList(url string) []YoutubeVideoMetadata {
	cmd := exec.Command(viper.GetString("binaries.youtube_dl"), "-j", "--flat-playlist", url)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		fmt.Printf("GetPlaylistList yt-dl err: %v\n", err)
		return []YoutubeVideoMetadata{}
	}

	var items []YoutubeVideoMetadata
	for _, data := range strings.Split(out.String(), "\n") {
		if data == "" {
			continue
		}

		var item YoutubeVideoMetadata
		err := json.Unmarshal([]byte(data), &item)
		if err != nil {
			fmt.Printf("GetPlaylistList JSON err: %v\n", err)
			continue
		}
		items = append(items, item)
	}
	return items
}