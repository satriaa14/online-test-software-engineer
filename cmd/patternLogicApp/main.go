package main

import (
	"flag"
	"fmt"

	"github.com/satriaa14/online-test-software-engineer/pkg/patternlogic"
)

func main() {
	var defaultValue = flag.Int("val", 5, "type your age")
	flag.Parse()
	fmt.Print(patternlogic.OnlineTestSoftwareEngineer(*defaultValue))
}
