package class2

import (
	"bufio"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func HashExample(word string) []byte {
	h := sha256.New()
	h.Write([]byte(word))

	hashedWord := h.Sum(nil)

	return hashedWord
}

//HashString returns hashed string
func HashString(text string) (string, error) {
	hash := sha256.New()
	_, err := hash.Write([]byte(text))

	if err != nil {
		return "", err
	}

	hashBytes := hash.Sum(nil)
	return hex.EncodeToString(hashBytes), nil
}

//Funkcja cytająca plik

func ReadFile(fileName string) {

	data, err := ioutil.ReadFile(fileName)

	if err != nil {
		log.Fatalf("%v", err)
	}

	fmt.Print(string(data))

}

func OpenAndReadFile(fileName string) {

	file, err := os.Open(fileName)
	defer file.Close()

	if err != nil {
		log.Fatalf("%v", err)
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	//error checking
	if err := scanner.Err(); err != nil {
		log.Fatalf("%v", err)
	}
}
