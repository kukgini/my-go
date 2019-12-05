package main

import (
	b "github.com/kukgini/my-go/patterns/behavioral"
)

func main() {
	worker := b.NewWorker(&b.PostMan{})
	worker.DailyRoutine()
} 