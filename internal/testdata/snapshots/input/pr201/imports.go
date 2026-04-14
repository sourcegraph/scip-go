package pr201

import (
	"context"
	"net/http"

	_ "embed"

	. "strings"

	fmt "fmt"
	h "net/http"
)

func UseImports() {
	fmt.Println(context.Background())
	_ = http.StatusOK
	_ = h.DefaultClient
	_ = Contains("hello", "ell")
}
