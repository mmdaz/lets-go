package concurrency_practice

import (
	"fmt"
	"io"
	"os"
	"time"
)

func StartEcho() {
	go echo(os.Stdin, os.Stdout)
	go echo(os.Stdin, os.Stdout)
	time.Sleep(30 * time.Second)
	fmt.Println("Timed out.")
	os.Exit(0)
}

func echo(in io.Reader, out io.Writer) {
	_, _ = io.Copy(out, in)
}
