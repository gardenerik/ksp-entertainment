package music

import "github.com/spf13/viper"

type Providers struct {
	Youtube Youtube
	Spotify Spotify
}

var GlobalProviders Providers

func (p *Providers) Prepare() {
	p.Youtube.Binary = viper.GetString("binaries.youtube_dl")

	p.Spotify.ClientId = viper.GetString("api.spotify.client_id")
	p.Spotify.ClientSecret = viper.GetString("api.spotify.client_secret")
}
