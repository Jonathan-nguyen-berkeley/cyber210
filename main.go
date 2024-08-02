package main

import (
	"cyber210/final/chains"
	"fmt"
	"log"

	"go.uber.org/zap"
)

func main() {

	logger, err := zap.NewDevelopment()
	if err != nil {
		log.Panicf("Failed to create logger %s\n", err.Error())
	}
	log := logger.Sugar()

	for {
		fmt.Println("Which benchmark would you like to run?\n1) Straw\n2) Twig\n3) Brick\nPress Enter to exit")
		input := ""
		fmt.Scanln(&input)
		switch input {
		case "1":
			chains.RunStrawcoinBenchmark(log)

		case "2":
			chains.RunTwigcoinBenchmark(log)

		case "3":
			chains.RunBrickcoinBenchmark(log)

		default:
			return
		}
	}
}
