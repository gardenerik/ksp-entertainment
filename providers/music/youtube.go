package music

import (
	"bytes"
	"encoding/json"
	"fmt"
	neturl "net/url"
	"os/exec"
	"strings"
)

type Youtube struct {
	Binary string
}

func (y Youtube) Identifier(url string) (string, error) {
	if !strings.HasPrefix(url, "http") {
		url = "https://" + url
	}

	u, err := neturl.Parse(url)
	if err != nil {
		return "", err
	}

	fmt.Println(u.Hostname(), strings.HasSuffix(u.Hostname(), "youtu.be"))
	u.Host = strings.ToLower(u.Host)
	u.Path = strings.Trim(u.Path, "/")

	// youtu.be/dQw4w9WgXcQ
	if strings.HasSuffix(u.Hostname(), "youtu.be") {
		return u.Path, nil
	}

	// youtube.com ...
	if strings.HasSuffix(u.Hostname(), "youtube.com") {
		// youtube.com/watch?v=dQw4w9WgXcQ
		if u.Path == "watch" {
			return u.Query().Get("v"), nil
		}

		// youtube.com/v/dQw4w9WgXcQ
		// youtube.com/embed/dQw4w9WgXcQ
		// youtube.com/e/dQw4w9WgXcQ
		parts := strings.Split(u.Path, "/")
		if parts[0] == "v" || parts[0] == "embed" || parts[0] == "e" {
			return parts[1], nil
		}
	}

	return "", fmt.Errorf("could not retrieve YouTube ID from %s", url)
}

func (y Youtube) StreamURL(identifier string) (string, error) {
	return fmt.Sprintf("https://youtube.com/watch?v=%s", identifier), nil
}

type YoutubeJson struct {
	Id       string `json:"id"`
	Title    string `json:"title"`
	Uploader string `json:"uploader"`
}

func (y Youtube) Metadata(identifier string) (SongMetadata, error) {
	url, err := y.StreamURL(identifier)
	if err != nil {
		return SongMetadata{}, err
	}

	var out bytes.Buffer
	cmd := exec.Command(y.Binary, "-j", url)
	cmd.Stdout = &out
	err = cmd.Run()
	if err != nil {
		return SongMetadata{}, err
	}

	var ytJson YoutubeJson
	err = json.Unmarshal(out.Bytes(), &ytJson)
	if err != nil {
		return SongMetadata{}, err
	}

	name, artist := ParseTitle(ytJson.Title)
	if artist == "" {
		artist = ytJson.Uploader
	}

	return SongMetadata{
		Provider: YOUTUBE,
		Name:     name,
		Artist:   artist,
		Album:    AlbumMetadata{},
		Cover:    fmt.Sprintf("https://i.ytimg.com/vi_webp/%s/maxresdefault.webp", ytJson.Id),
	}, nil
}
