package main

import (
	"fmt"
	"os"

	"github.com/yankeguo/zrand"
)

var (
	special   = zrand.Immediate("_.-")
	operation = zrand.Combine(
		zrand.Random(zrand.Letters, 1),
		zrand.Shuffle(
			zrand.Combine(
				zrand.Random(zrand.Uppers, 3),
				zrand.Random(zrand.Lowers, 3),
				zrand.Random(zrand.Numerics, 3),
				zrand.Random(special, 1),
			),
		),
		zrand.Random(zrand.Combine(zrand.Letters, zrand.Numerics), 3),
		zrand.Shuffle(
			zrand.Combine(
				zrand.Random(zrand.Uppers, 3),
				zrand.Random(zrand.Lowers, 3),
				zrand.Random(zrand.Numerics, 3),
				zrand.Random(special, 1),
			),
		),
		zrand.Random(zrand.Letters, 1),
	)
)

func main() {
	fmt.Fprintln(os.Stdout, zrand.BuildString(operation))
}
