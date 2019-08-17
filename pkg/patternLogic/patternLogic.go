package patternlogic

import (
	"fmt"
	"strings"
)

var (
	// Initial for default returns pattern (returningPattern = "")
	// Initial cell Pattern (x, o)
	// Initial Space (noSpace and oneSpace)
	// Initial for new row (newRow)
	returningPattern, x, o, noSpace, oneSpace, newRow = "", "X", "O", "", " ", "\n"

	// Top Triangle Initial (i , j)
	// Bottom Triangle Initial (a, b)
	// Botton Triangele Iteration Initial (y)
	// Yes, there are integer
	i, j, a, b, y int
)

// function for OnlineTestSoftwareEngineer only accepting integer (rows)
// and return as string, which is return variable are "returningPattern"
func OnlineTestSoftwareEngineer(rows int) string {
	// Start of Top Triangle
	for i = 1; i <= rows; i++ {
		for j = i; j <= rows; j++ {
			if j == rows {
				returningPattern = fmt.Sprintf("%s%s", returningPattern, noSpace)
			} else {
				returningPattern = fmt.Sprintf("%s%s", returningPattern, oneSpace)
			}
		}
		topSwither := true
		topTemp := noSpace
		topSpaceTemp := noSpace
		for j = 1; j <= (2*i - 1); j++ {
			if i == rows {
				continue
			} else {
				if j <= i {
					if j%2 == 0 {
						topTemp = fmt.Sprintf("%s%s", topTemp, oneSpace)
						if j == 1 {
							topSpaceTemp = fmt.Sprintf("%s%s", topSpaceTemp, noSpace)
						} else {
							topSpaceTemp = fmt.Sprintf("%s%s", topSpaceTemp, oneSpace)
						}
					} else {
						if topSwither == true {
							topTemp = fmt.Sprintf("%s%s", topTemp, x)
							if j == 1 {
								topSpaceTemp = fmt.Sprintf("%s%s", topSpaceTemp, noSpace)
							} else {
								topSpaceTemp = fmt.Sprintf("%s%s", topSpaceTemp, oneSpace)
							}
							topSwither = false
						} else {
							topTemp = fmt.Sprintf("%s%s", topTemp, o)
							if j == 1 {
								topSpaceTemp = fmt.Sprintf("%s%s", topSpaceTemp, noSpace)
							} else {
								topSpaceTemp = fmt.Sprintf("%s%s", topSpaceTemp, oneSpace)
							}
							topSwither = true
						}
					}
				}
			}
		}
		reverseTop := func(s string) string {
			var reverseTemp string
			chars := strings.Split(s, "")
			length := len(chars)
			for i := length - 1; i >= 0; i-- {
				reverseTemp += chars[i]
			}
			return reverseTemp
		}
		returningPattern = fmt.Sprintf("%s%s", returningPattern, topSpaceTemp)
		returningPattern = fmt.Sprintf("%s%s", returningPattern, reverseTop(topTemp))
		if i != rows {
			returningPattern = fmt.Sprintf("%s%s", returningPattern, newRow)
		}
	}
	// End of Top Triangle

	// Start of Middle Pattern
	midSwitcher := true
	if rows%2 == 0 {
		temp := ""
		for m := 1; m < (rows); m++ {
			if m%2 == 0 {
				temp = fmt.Sprintf("%s%s", temp, oneSpace)
			} else {
				if midSwitcher == true {
					temp = fmt.Sprintf("%s%s", temp, x)
					midSwitcher = false
				} else {
					temp = fmt.Sprintf("%s%s", temp, o)
					midSwitcher = true
				}
			}
		}
		midReverse := func(s string) string {
			var reverseTemp string
			chars := strings.Split(s, "")
			length := len(chars)
			for i := length - 1; i >= 0; i-- {
				reverseTemp += chars[i]
			}
			return reverseTemp
		}
		returningPattern = fmt.Sprintf("%s%s", returningPattern, temp)
		returningPattern = fmt.Sprintf("%s%s", returningPattern, oneSpace)
		returningPattern = fmt.Sprintf("%s%s", returningPattern, midReverse(temp))
	} else {
		for m := 1; m < (rows); m++ {
			if m%2 == 0 {
				returningPattern = fmt.Sprintf("%s%s", returningPattern, oneSpace)
			} else {
				if midSwitcher == true {
					returningPattern = fmt.Sprintf("%s%s", returningPattern, x)
					midSwitcher = false
				} else {
					returningPattern = fmt.Sprintf("%s%s", returningPattern, o)
					midSwitcher = true
				}
			}
		}
		for n := 1; n < (rows + 1); n++ {
			if n%2 == 0 {
				returningPattern = fmt.Sprintf("%s%s", returningPattern, oneSpace)
			} else {
				if midSwitcher == true {
					returningPattern = fmt.Sprintf("%s%s", returningPattern, x)
					midSwitcher = false
				} else {
					returningPattern = fmt.Sprintf("%s%s", returningPattern, o)
					midSwitcher = true
				}
			}
		}
	}
	returningPattern = fmt.Sprintf("%s%s", returningPattern, newRow)
	// End of middle pattern

	// Start of Bottom down side triangle pattern
	dummy := j
	bottomSwitcher := true
	bottomTemp := noSpace
	for a = 1; a <= rows-1; a++ {
		for b = 0; b <= a; b++ {
			if b == 0 {
				continue
			} else {
				returningPattern = fmt.Sprintf("%s%s", returningPattern, oneSpace)
			}
		}
		dummy = dummy - 1
		for b = a; b < (dummy - 1); b++ {
			y++
			if b > rows {
				bottomTemp = fmt.Sprintf("%s%s", bottomTemp, oneSpace)
			} else {
				if y%2 == 0 {
					if bottomSwitcher == true {
						returningPattern = fmt.Sprintf("%s%s", returningPattern, x)
						bottomTemp = fmt.Sprintf("%s%s", bottomTemp, x)
						bottomSwitcher = false
					} else {
						returningPattern = fmt.Sprintf("%s%s", returningPattern, o)
						bottomTemp = fmt.Sprintf("%s%s", bottomTemp, o)
						bottomSwitcher = true
					}
				} else {
					if y == 1 {
						continue
					}
					returningPattern = fmt.Sprintf("%s%s", returningPattern, oneSpace)
				}
			}
		}
		y = 0
		if a == rows-1 {
			returningPattern = fmt.Sprintf("%s%s", returningPattern, x)
			continue
		}
		returningPattern = fmt.Sprintf("%s%s", returningPattern, newRow)
	}
	// End of Bottom down side triangle pattern

	// Returning variable that already merging with previous pattern
	return returningPattern
}
