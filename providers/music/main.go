package music

type Provider string

const (
	YOUTUBE Provider = "youtube"
)

type AlbumMetadata struct {
	Name   string
	Artist string
	Cover  string
}

type SongMetadata struct {
	Provider Provider
	Name     string
	Artist   string
	Album    AlbumMetadata
	Cover    string
}

type SongProvider interface {
	Identifier(url string) (string, error)
	StreamURL(identifier string) (string, error)
	Metadata(identifier string) (SongMetadata, error)
}
