package main

import "github.com/JoaoDanielRufino/gcloc/internal/app"

func main() {
	rootCmd := app.NewGClocCmd()

	rootCmd.Execute()
}
