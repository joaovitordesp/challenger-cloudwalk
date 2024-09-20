package logparser

import (
	"bufio"
	"os"
	"regexp"
	"strings"
)

func ParseLogFile(filepath string) ([]Game, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var games []Game
	var currentGame *Game

	// Regex para capturar quem matou quem e com qual arma
	playerRegex := regexp.MustCompile(`Kill:\s\d+\s\d+\s\d+:\s(.+)\skilled\s(.+)\sby\s(.+)`)

	for scanner.Scan() {
		line := scanner.Text()

		// Inicia um novo jogo
		if strings.Contains(line, "InitGame") {
			currentGame = &Game{
				Players: make(map[string]Player),
			}
		}

		// Quando o jogo termina, adiciona à lista de jogos
		if strings.Contains(line, "ShutdownGame") {
			if currentGame != nil {
				games = append(games, *currentGame)
			}
		}

		// Processa as informações de "Kill"
		if strings.Contains(line, "Kill:") {
			currentGame.TotalKills++

			// Usa regex para extrair quem matou quem
			matches := playerRegex.FindStringSubmatch(line)
			if len(matches) == 4 {
				killer := matches[1]
				victim := matches[2]

				// Caso o <world> tenha matado alguém
				if killer == "<world>" {
					currentGame.Players[victim] = Player{
						Name:  victim,
						Kills: currentGame.Players[victim].Kills - 1,
					}
				} else {
					// Incrementa as kills do jogador assassino
					if _, exists := currentGame.Players[killer]; !exists {
						currentGame.Players[killer] = Player{Name: killer, Kills: 0}
					}
					currentGame.Players[killer] = Player{
						Name:  killer,
						Kills: currentGame.Players[killer].Kills + 1,
					}

					// Garante que a vítima seja registrada no mapa de jogadores
					if _, exists := currentGame.Players[victim]; !exists {
						currentGame.Players[victim] = Player{Name: victim, Kills: 0}
					}
				}
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return games, nil
}
