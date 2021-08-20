package main

import (
	"fmt"
	"math"

	"github.com/chubuntuarc/go_crash_course/03_packages/stutil"
)

func main() {
	fmt.Println(math.Floor(2.7))
	fmt.Println(math.Ceil(2.7))
	fmt.Println(stutil.Reverse("olleh"))
}
