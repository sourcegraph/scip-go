package initial

import (
  "slices"
)

func f(xs []int) int {
  for _, x := range slices.All(xs) {
    return x
  }
  return -1
}
