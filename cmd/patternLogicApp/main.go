package main

import (
	"flag"
	"fmt"

	patternlogic "github.com/satriaa14/online-test-software-engineer/pkg/patternLogic"
)

func main() {
	var defaultValue = flag.Int("val", 5, "type your age")
	flag.Parse()
	fmt.Print(patternlogic.OnlineTestSoftwareEngineer(*defaultValue))
}
