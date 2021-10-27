package parsers

import (
	"bytes"
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
