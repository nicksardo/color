# Simple ANSI Formatting

Console coloring library specifically for byte streams instead of per-string.

```go
p := color.Palette{Foreground:color.Red}
p.ApplyTo(os.Stdout)

fmt.Println("This is in red.")

p.Background = color.Blue
p.ApplyTo(os.Stdout)
fmt.Println("So will the next line, but with a blue background.")

p.Clear()
p.ApplyTo(os.Stdout)

fmt.Println("Back to normal")
```
