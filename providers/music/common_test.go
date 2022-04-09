package music

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseTitle(t *testing.T) {
	tests := []struct {
		title          string
		expectedName   string
		expectedArtist string
	}{
		{"Owl City - Fireflies (Official Music Video)", "Fireflies", "Owl City"},
		{"American Pie", "American Pie", ""},
		{"Train - Hey, Soul Sister (Official Video)", "Hey, Soul Sister", "Train"},
		{"Smash Mouth - I'm A Believer", "I'm A Believer", "Smash Mouth"},
		{"Rick Astley - Never Gonna Give You Up (Official Music Video)", "Never Gonna Give You Up", "Rick Astley"},
	}

	for _, test := range tests {
		name, artist := ParseTitle(test.title)
		assert.Equal(t, test.expectedName, name)
		assert.Equal(t, test.expectedArtist, artist)
	}
}
