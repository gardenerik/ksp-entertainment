package music

import "testing"
import "github.com/stretchr/testify/assert"

func TestYoutube_Identifier(t *testing.T) {
	urls := []string{
		"https://youtu.be/dQw4w9WgXcQ",
		"https://www.youtu.be/dQw4w9WgXcQ",
		"http://www.youtu.be/dQw4w9WgXcQ",
		"www.youtu.be/dQw4w9WgXcQ",
		"https://youtube.com/watch?v=dQw4w9WgXcQ",
		"https://www.youtube.com/watch?v=dQw4w9WgXcQ",
		"youtube.com/watch?v=dQw4w9WgXcQ",
		"https://youtube.com/embed/dQw4w9WgXcQ",
		"https://www.youtube.com/embed/dQw4w9WgXcQ",
		"youtube.com/embed/dQw4w9WgXcQ",
		"https://youtube.com/v/dQw4w9WgXcQ",
		"https://www.youtube.com/v/dQw4w9WgXcQ",
		"youtube.com/v/dQw4w9WgXcQ",
		"https://youtube.com/e/dQw4w9WgXcQ",
		"https://www.youtube.com/e/dQw4w9WgXcQ",
		"youtube.com/e/dQw4w9WgXcQ",
	}

	for _, url := range urls {
		id, err := Youtube{}.Identifier(url)
		assert.Nil(t, err)
		assert.Equal(t, "dQw4w9WgXcQ", id)
	}
}

func TestYoutube_Metadata(t *testing.T) {
	metadata, err := Youtube{Binary: "/usr/bin/youtube-dl"}.Metadata("dQw4w9WgXcQ")
	assert.Nil(t, err)
	assert.Equal(t, YOUTUBE, metadata.Provider)
	assert.Equal(t, "Never Gonna Give You Up", metadata.Name)
	assert.Equal(t, "Rick Astley", metadata.Artist)
	assert.Equal(t, "", metadata.Album.Name)
}
