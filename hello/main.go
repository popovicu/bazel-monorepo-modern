package main

import (
	"fmt"

	"github.com/popovicu/bazel-monorepo-modern/greeting"
)

func main() {
	fmt.Println(greeting.For("Bazel"))
}
