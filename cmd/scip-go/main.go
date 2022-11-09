package main

import (
	"fmt"

	"github.com/sourcegraph/scip-go/internal/index"
)

func main() {
	fmt.Println("scip-go")
	index.Parse()
}
