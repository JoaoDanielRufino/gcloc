package main

import "github.com/JoaoDanielRufino/gcloc/internal/app"

// This is a comment
func main() {
	rootCmd := app.NewGClocCmd()

	// rootCmd execute
	rootCmd.Execute()
}

// End
