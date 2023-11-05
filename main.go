package main

import (
	"fmt"
	"github.com/tazdij/die/editor"
)

func main() {
    fmt.Println("Hello from `die` (Duvall Industries Editor).")
    fmt.Println("This is a work in progress, and is not yet functional.")

    fmt.Println("Here is the about for the Tab:")
    fmt.Println(editor.GetTabAbout())

    fmt.Println("Here is the about for the Buffer:")
    fmt.Println(editor.GetBufferAbout())

    fmt.Println("\n\nDone.")
}
