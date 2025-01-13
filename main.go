package main

import (
	"enigma/config"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
)

type Enigma struct {
	SizeOfAlphabet int
	Rotors         [3]*config.RotorConfig
	Reflector      config.ReflectorConfig
	Key            string
}

func readEnigmaSettingsFromEnv() (Enigma, string, error) {
	var enigma Enigma

	message := os.Getenv("ENIGMA_MESSAGE")
	if message == "" {
		return enigma, "", fmt.Errorf("переменная окружения ENIGMA_MESSAGE не найдена")
	}

	rotorInput := os.Getenv("ENIGMA_KEY")
	if rotorInput == "" {
		return enigma, message, fmt.Errorf("переменная окружения ENIGMA_ROTORS не найдена")
	}
	rotorNums := strings.Split(rotorInput, " ")
	if len(rotorNums) != 3 {
		return enigma, message, fmt.Errorf("неверное количество номеров роторов в ENIGMA_ROTORS")
	}
	for i, numStr := range rotorNums {
		rotorNumAndDeltas := strings.Split(numStr, "-")
		num, _ := strconv.Atoi(rotorNumAndDeltas[0])
		delta, _ := strconv.Atoi(rotorNumAndDeltas[1])
		enigma.Rotors[i], _ = config.NewRotorConfig(fmt.Sprintf("rotors/rotor_%d.txt", num), delta)
	}

	reflectorInput := os.Getenv("ENIGMA_REFLECTOR")
	if reflectorInput == "" {
		return enigma, message, fmt.Errorf("переменная окружения ENIGMA_REFLECTOR не найдена")
	}
	reflectorNum, err := strconv.Atoi(reflectorInput)
	if err != nil {
		return enigma, message, fmt.Errorf("неверный формат номера рефлектора в ENIGMA_REFLECTOR: %s", reflectorInput)
	}

	reflectorConfigPath := fmt.Sprintf("reflectors/reflector_%d.txt", reflectorNum)
	enigma.Reflector, err = config.NewReflectorConfig(reflectorConfigPath)
	if err != nil {
		return enigma, message, fmt.Errorf("ошибка во время создания рефлектора: %s %w", reflectorConfigPath, err)
	}
	enigma.SizeOfAlphabet = len(enigma.Reflector)

	keyInput := os.Getenv("ENIGMA_KEY")
	if keyInput == "" {
		return enigma, message, fmt.Errorf("переменная окружения ENIGMA_KEY не найдена")
	}
	enigma.Key = keyInput

	return enigma, message, nil
}

func printEnigmaSettings(enigma Enigma) {

	fmt.Println("Настройки Энигмы:")

	for i, cfg := range enigma.Rotors {
		fmt.Printf("Ротор %d: ", i)
		fmt.Printf("Wiring1: %s, Wiring2: %s, Delta: %d\n", string(cfg.Wiring1), string(cfg.Wiring2), cfg.Delta)
	}

	fmt.Printf("Рефлектор: ")

	for key, value := range enigma.Reflector {
		fmt.Printf("%c->%c  ", key, value)
	}
	fmt.Printf("\nКлюч: %s\n", enigma.Key)
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Ошибка загрузки .env файла: %v", err)
	}

	enigma, message, err := readEnigmaSettingsFromEnv()

	if err != nil {
		fmt.Println("Ошибка чтения настроек:", err)
		return
	}

	printEnigmaSettings(enigma)

	encryptedMessage := ""

	for i := 0; i < len(message); i++ {
		pos1Rotor := config.FindIndex(enigma.Rotors[0].Wiring1, message[i])
		letter1Rotor := enigma.Rotors[0].Wiring2[(pos1Rotor+enigma.Rotors[0].Delta)%enigma.SizeOfAlphabet]

		pos2Rotor := config.FindIndex(enigma.Rotors[1].Wiring1, letter1Rotor)
		letter2Rotor := enigma.Rotors[1].Wiring2[(pos2Rotor+enigma.Rotors[1].Delta)%enigma.SizeOfAlphabet]

		pos3Rotor := config.FindIndex(enigma.Rotors[2].Wiring1, letter2Rotor)
		letter3Rotor := enigma.Rotors[2].Wiring2[(pos3Rotor+enigma.Rotors[2].Delta)%enigma.SizeOfAlphabet]

		posReflector := enigma.Reflector[letter3Rotor]

		pos4Rotor := config.FindIndex(enigma.Rotors[2].Wiring2, posReflector)
		letter4Rotor := enigma.Rotors[2].Wiring1[(pos4Rotor-enigma.Rotors[2].Delta+enigma.SizeOfAlphabet)%enigma.SizeOfAlphabet]

		pos5Rotor := config.FindIndex(enigma.Rotors[1].Wiring2, letter4Rotor)
		letter5Rotor := enigma.Rotors[1].Wiring1[(pos5Rotor-enigma.Rotors[1].Delta+enigma.SizeOfAlphabet)%enigma.SizeOfAlphabet]

		pos6Rotor := config.FindIndex(enigma.Rotors[0].Wiring2, letter5Rotor)
		letter6Rotor := enigma.Rotors[0].Wiring1[(pos6Rotor-enigma.Rotors[0].Delta+enigma.SizeOfAlphabet)%enigma.SizeOfAlphabet]

		enigma.Rotors[0].Delta = (enigma.Rotors[0].Delta + 1) % enigma.SizeOfAlphabet
		if enigma.Rotors[0].Delta == 0 {
			enigma.Rotors[1].Delta = (enigma.Rotors[1].Delta + 1) % enigma.SizeOfAlphabet
			if enigma.Rotors[1].Delta == 0 {
				enigma.Rotors[2].Delta = (enigma.Rotors[2].Delta + 1) % enigma.SizeOfAlphabet
			}
		}
		encryptedMessage += string(letter6Rotor)

	}
	fmt.Printf("encrypted message: %s \n", encryptedMessage)
	fmt.Printf("original message: %s", message)
}
