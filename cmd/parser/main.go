package main

import (
	"fmt"

	logparser "github.com/test-cloudwalk/quake-log/pkg/logparser"
)

func main() {
	logFile := "logs/quake_log.txt"

	games, err := logparser.ParseLogFile(logFile)
	if err != nil {
		fmt.Printf("Error parsing log file: %v\n", err)
		return
	}

	for i, game := range games {
		fmt.Printf("Game_%d:\n", i+1)
		fmt.Printf("Total Kills: %d\n", game.TotalKills)
		fmt.Printf("Players: %v\n", game.PlayerNames())

		fmt.Println("Kills:")
		for player, kills := range game.Kills() {
			fmt.Printf("  %s: %d kills\n", player, kills)
		}

		fmt.Println("--------------------------------------")
	}
}
