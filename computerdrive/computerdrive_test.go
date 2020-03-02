package computerdrive_test

import (
	"testing"

	"github.com/vdntruong/chl/computerdrive"
)

func TestSolution(t *testing.T) {
	testCases := []struct {
		desc, input, output string
	}{
		{
			desc: "Test case 01 of challenge",
			input: `my.song.mp3 11b 
			greatSong.flac 1000b 
			not3.txt 5b video.mp4 200b 
			game.exe 100b 
			mov!e.mkv 10000b`,
			output: `music 1011b
			images 0b
			movies 10200b
			other 105b`,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			if rs := computerdrive.Solution(tC.input); rs != tC.output {
				t.Errorf("got rs=%s, wants rs=%s", rs, tC.output)
			}
		})
	}
}
