package parsers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strings"
	"time"
)

type RTVSParser struct{}

type RTVSArchiveData struct {
	IP      string `json:"IP"`
	Request string `json:"request"`
	Clip    struct {
		DatetimeCreate string `json:"datetime_create"`
		Title          string `json:"title"`
		Description    string `json:"description"`
		Image          string `json:"image"`
		Sources        []struct {
			Src  string `json:"src"`
			Type string `json:"type"`
		} `json:"sources"`
		Mediaid  string `json:"mediaid"`
		Backend  string `json:"backend"`
		Previews string `json:"previews"`
		Length   string `json:"length"`
	} `json:"clip"`
}

var rtvsArchiveJsonRegex = regexp.MustCompile(`"(//www\.rtvs\.sk/json/archive.+)"`)

func (R RTVSParser) CanParse(url string) bool {
	return strings.Contains(url, "rtvs.sk/televizia/archiv/")
}

func (R RTVSParser) GetArchiveData(url string) (RTVSArchiveData, error) {
	embedUrl := strings.Replace(url, "televizia/archiv", "embed/archive", 1)
	client := http.Client{
		Timeout: time.Second * 5,
	}

	res, err := client.Get(embedUrl)
	if err != nil {
		return RTVSArchiveData{}, err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return RTVSArchiveData{}, err
	}

	bodyString := string(body)
	groups := rtvsArchiveJsonRegex.FindStringSubmatch(bodyString)
	if len(groups) != 2 {
		return RTVSArchiveData{}, fmt.Errorf("cannot find JSON url in RTVS embed")
	}

	res, err = client.Get("https:" + groups[1])
	if err != nil {
		return RTVSArchiveData{}, err
	}
	defer res.Body.Close()
	body, err = ioutil.ReadAll(res.Body)
	if err != nil {
		return RTVSArchiveData{}, err
	}

	var data RTVSArchiveData
	err = json.Unmarshal(body, &data)
	return data, err
}

func (R RTVSParser) GetStreamURL(url string) string {
	archive, err := R.GetArchiveData(url)
	if err != nil {
		log.Println("RTVS stream URL error: ", err)
		return ""
	}
	for _, source := range archive.Clip.Sources {
		if source.Type == "application/x-mpegurl" {
			return source.Src
		}
	}
	log.Printf("Could not find RTVS stream URL for %v\n", url)
	return ""
}

func (R RTVSParser) GetName(url string) string {
	archive, err := R.GetArchiveData(url)
	if err != nil {
		return fmt.Sprintf("(rtvs err: %v)", err)
	}
	return archive.Clip.Title
}
