// generallyeric -> generic for short
package generallyeric

import "fmt"

func Print[T any](s []T) {
	for _, v := range s {
		fmt.Print(v)
	}
}
