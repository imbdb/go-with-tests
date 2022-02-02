package dependencyinjection

import (
	"fmt"
	"io"
)

func Greet(iowriter io.Writer, name string) {
	fmt.Fprintf(iowriter, "Hello, %s!", name)
}
