package music

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

type Spotify struct {
	Token        string
	TokenExpiry  time.Time
	ClientId     string
	ClientSecret string
}

func (s *Spotify) Authenticate() error {
	client := http.Client{Timeout: 10 * time.Second}

	body := "grant_type=client_credentials"
	req, err := http.NewRequest("POST", "https://accounts.spotify.com/api/token", bytes.NewBufferString(body))
	if err != nil {
		return err
	}

	auth := fmt.Sprintf("%s:%s", s.ClientId, s.ClientSecret)
	auth = base64.StdEncoding.EncodeToString([]byte(auth))

	req.Header.Set("Authorization", fmt.Sprintf("Basic %s", auth))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Accept", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	var spotifyResponse struct {
		AccessToken string `json:"access_token"`
		ExpiresIn   int    `json:"expires_in"`
	}
	err = json.NewDecoder(resp.Body).Decode(&spotifyResponse)
	if err != nil {
		return err
	}

	s.Token = spotifyResponse.AccessToken
	s.TokenExpiry = time.Now().Add(time.Duration(spotifyResponse.ExpiresIn) * time.Second)

	return nil
}

func (s *Spotify) ImproveMetadata(meta *SongMetadata) (bool, error) {
	if s.Token == "" || s.TokenExpiry.Before(time.Now()) {
		err := s.Authenticate()
		if err != nil {
			return false, err
		}
		fmt.Printf("mame token %s\n", s.Token)
	}

	client := http.Client{Timeout: 10 * time.Second}
	query := url.Values{}
	query.Set("q", fmt.Sprintf("%s %s", meta.Name, meta.Artist))
	query.Set("type", "track")

	req, err := http.NewRequest("GET", "https://api.spotify.com/v1/search?"+query.Encode(), nil)
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", s.Token))
	req.Header.Set("Accept", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	var spotifyResponse struct {
		Tracks struct {
			Items []struct {
				Album struct {
					Artists []struct {
						Name string `json:"name"`
					} `json:"artists"`
					Images []struct {
						Height int    `json:"height"`
						Url    string `json:"url"`
						Width  int    `json:"width"`
					} `json:"images"`
					Name string `json:"name"`
				} `json:"album"`
				Artists []struct {
					Name string `json:"name"`
				} `json:"artists"`
				Name string `json:"name"`
			} `json:"items"`
		} `json:"tracks"`
	}
	err = json.NewDecoder(resp.Body).Decode(&spotifyResponse)
	if err != nil {
		return false, err
	}

	if len(spotifyResponse.Tracks.Items) < 1 {
		return false, nil
	}
	item := spotifyResponse.Tracks.Items[0]
	if len(item.Artists) > 0 {
		meta.Artist = item.Artists[0].Name
		meta.Name = item.Name

		meta.Album.Artist = item.Album.Artists[0].Name
		meta.Album.Name = item.Album.Name
		meta.Album.Cover = item.Album.Images[0].Url
		return true, nil
	}
	return false, nil
}
