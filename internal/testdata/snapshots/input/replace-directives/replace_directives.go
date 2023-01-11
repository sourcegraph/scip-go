package replacers

import (
	"fmt"

	"github.com/dghubble/gologin"
)

func Something() {
	fmt.Println(gologin.DefaultCookieConfig)
}
