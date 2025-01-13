package config

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type ReflectorConfig map[byte]byte

func NewReflectorConfig(filename string) (ReflectorConfig, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("ошибка открытия файла: %w", err)
	}
	defer file.Close()

	config := make(ReflectorConfig)
	scanner := bufio.NewScanner(file)
	lineNumber := 0

	for scanner.Scan() {
		lineNumber++
		line := strings.ToUpper(scanner.Text())
		pairs := strings.Split(line, " ")

		for _, pair := range pairs {
			pair = strings.TrimSpace(pair)
			if len(pair) != 2 {
				return nil, fmt.Errorf("неверный формат пары в строке %d: %s, должно быть 2 символа", lineNumber, pair)
			}

			if pair[0] < 'A' || pair[0] > 'Z' || pair[1] < 'A' || pair[1] > 'Z' {
				return nil, fmt.Errorf("недопустимый символ в паре в строке %d: %s", lineNumber, pair)
			}

			config[pair[0]] = pair[1]
			config[pair[1]] = pair[0]
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("ошибка чтения файла: %w", err)
	}

	return config, nil
}
