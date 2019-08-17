package patternlogic

import (
	"log"
	"testing"
)

func TestOnlineTestSoftwareEngineer(t *testing.T) {
	t.Run("Test for pattern generator", func(t *testing.T) {
		var testPatternLogic = []struct {
			val  int
			want string
		}{
			{val: 1, want: "\nX"},
			{val: -10, want: ""},
		}
		for _, input := range testPatternLogic {
			got := OnlineTestSoftwareEngineer(input.val)
			if got != got {
				t.Fatalf("Pattern does not match, got : \n %s \n \n want : \n %s ", got, input.want)
			} else {
				log.Printf("Pattern match, got : \n %s \n want :  %s ", got, input.want)
			}
		}
	})
}
