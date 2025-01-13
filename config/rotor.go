package config

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type RotorConfig struct {
	Wiring1 []byte
	Wiring2 []byte
	Delta   int
}

func NewRotorConfig(filename string, delta int) (*RotorConfig, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("ошибка открытия файла: %w", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("ошибка чтения файла: %w", err)
	}

	if len(lines) != 2 {
		return nil, fmt.Errorf("файл должен содержать ровно 2 строки, а не %d", len(lines))
	}

	line1 := strings.ToUpper(lines[0])
	line2 := strings.ToUpper(lines[1])

	for _, char := range line1 {
		if char < 'A' || char > 'Z' {
			return nil, fmt.Errorf("первая строка содержит недопустимый символ: %c", char)
		}
	}

	for _, char := range line2 {
		if char < 'A' || char > 'Z' {
			return nil, fmt.Errorf("вторая строка содержит недопустимый символ: %c", char)
		}
	}

	if len(line1) != len(line2) {
		return nil, fmt.Errorf("строки должны быть одинаковой длины, но получены %d и %d", len(line1), len(line2))
	}

	config := &RotorConfig{
		Wiring1: []byte(line1),
		Wiring2: []byte(line2),
		Delta:   delta, // Начальное значение дельты, можно настроить по желанию
	}

	return config, nil
}

func FindIndex(arr []byte, target byte) int {
	// Итерируемся по массиву, чтобы найти индекс
	for i, v := range arr {
		if v == target {
			return i
		}
	}
	// Если элемент не найден, возвращаем -1
	return -1
}
