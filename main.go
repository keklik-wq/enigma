package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"enigma/config"

	"github.com/joho/godotenv"
)

type Enigma struct {
	Rotors	[3]*config.RotorConfig
	Reflector config.ReflectorConfig
	Key string
}

func main() {
	err := godotenv.Load()
	if err != nil {
			log.Fatalf("Ошибка загрузки .env файла: %v", err)
	}

	enigmaSettings, err := readEnigmaSettingsFromEnv()
	if err != nil {
			fmt.Println("Ошибка чтения настроек:", err)
			return
	}

	fmt.Println("Настройки Энигмы:")
	fmt.Printf("Роторы: %v\n", enigmaSettings.Rotors)
	fmt.Printf("Рефлектор: %d\n", enigmaSettings.Reflector)
	fmt.Printf("Ключ: %s\n", enigmaSettings.Key)
}

func readEnigmaSettingsFromEnv() (Enigma, error) {
	var enigma Enigma

	rotorInput := os.Getenv("ENIGMA_ROTORS")
	if rotorInput == "" {
			return enigma, fmt.Errorf("переменная окружения ENIGMA_ROTORS не найдена")
	}
	rotorNums := strings.Split(rotorInput, " ")
	if len(rotorNums) != 3 {
			return enigma, fmt.Errorf("неверное количество номеров роторов в ENIGMA_ROTORS")
	}
	for i, numStr := range rotorNums {
			num, err := strconv.Atoi(numStr)
			if err != nil {
					return enigma, fmt.Errorf("неверный формат номера ротора в ENIGMA_ROTORS: %s", numStr)
			}
			enigma.Rotors[i], err = config.NewRotorConfig(fmt.Sprintf("rotors/rotor_%d.txt", num))
	}

	reflectorInput := os.Getenv("ENIGMA_REFLECTOR")
	if reflectorInput == "" {
			return enigma, fmt.Errorf("переменная окружения ENIGMA_REFLECTOR не найдена")
	}
	reflectorNum, err := strconv.Atoi(reflectorInput)
	if err != nil {
			return enigma, fmt.Errorf("неверный формат номера рефлектора в ENIGMA_REFLECTOR: %s", reflectorInput)
	}

	enigma.Reflector, err = config.NewReflectorConfig(fmt.Sprintf("reflectors/reflector_%d.txt", reflectorNum))

keyInput := os.Getenv("ENIGMA_KEY")
if keyInput == "" {
	return enigma, fmt.Errorf("переменная окружения ENIGMA_KEY не найдена")
}
enigma.Key = keyInput

	return enigma, nil
}
