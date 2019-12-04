package main

import (
	b "github.com/kukgini/my-go/behavioral"
)

func main() {
	worker := b.NewWorker(&b.PostMan{})
	worker.DailyRoutine()
} 