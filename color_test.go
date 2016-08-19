package color

import (
	"fmt"
	"os"
	"testing"
)

func TestColor(t *testing.T) {
	p := Palette{Foreground: Red}
	p.ApplyTo(os.Stdout)

	fmt.Println("This is in red.")

	p.Background = Blue
	p.ApplyTo(os.Stdout)
	fmt.Println("So will the next line, but with a blue background.")

	p.Clear()
	p.ApplyTo(os.Stdout)

	fmt.Println("Back to normal")
}
