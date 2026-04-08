package pr197

import (
	"context"
	"fmt"
	"net/http"

	_ "embed"

	. "strings"

	h "net/http"
)

func UseImports() {
	fmt.Println(context.Background())
	_ = http.StatusOK
	_ = h.DefaultClient
	_ = Contains("hello", "ell")
}
