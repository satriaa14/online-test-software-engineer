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
			{val: 1, want: "X\n"},
			{val: 2, want: " X\nX X\n X"},
			{val: 0, want: " \n"},
			{val: -10, want: " \n"},
		}
		for _, input := range testPatternLogic {
			got := OnlineTestSoftwareEngineer(input.val)
			if got != input.want {
				t.Fatalf("\nPattern does not match!\nIf value = %d , \nGot : \n\n%s\n\n\nWant : \n\n%s\n", input.val, got, input.want)
			} else {
				log.Printf("\nPattern match!\nIf value = %d, \nGot : \n\n%s\n\n\nWant :  \n\n%s\n", input.val, got, input.want)
			}
		}
	})
}
