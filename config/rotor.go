package config

import (
        "bufio"
        "fmt"
        "os"
        "strings"
)

// RotorConfig представляет конфигурацию ротора.
type RotorConfig struct {
        Wiring1 [26]byte
        Wiring2 [26]byte
        Delta   int
}

// NewRotorConfig создает новую конфигурацию ротора из файла.
func NewRotorConfig(filename string) (*RotorConfig, error) {
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

        if len(line1) != 26 || len(line2) != 26 {
                return nil, fmt.Errorf("каждая строка должна содержать 26 букв, а не %d и %d", len(line1), len(line2))
        }

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


        config := &RotorConfig{}
        copy(config.Wiring1[:], []byte(line1))
        copy(config.Wiring2[:], []byte(line2))
        config.Delta = 0 // Начальное значение дельты, можно настроить по желанию

        return config, nil
}

// func main() {
//         config, err := NewRotorConfig("rotors/rotor_1.txt") // Замените на имя вашего файла
//         if err != nil {
//                 fmt.Println("Ошибка:", err)
//                 return
//         }

//         fmt.Println("Конфигурация ротора:")
//         fmt.Printf("Wiring1: %s\n", string(config.Wiring1[:]))
//         fmt.Printf("Wiring2: %s\n", string(config.Wiring2[:]))
//         fmt.Printf("Delta: %d\n", config.Delta)
// }