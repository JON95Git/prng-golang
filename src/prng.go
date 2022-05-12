package prng

import (
	"fmt"
	"log"
	"os"
)

func readUrandom(sizeToRead int) []byte {
	file, err := os.Open("/dev/urandom")
	if err != nil {
		log.Fatal("Failed to open urandom file")
	}
	defer file.Close()

	buffer := make([]byte, sizeToRead)
	_, err = file.Read(buffer)
	if err != nil {
		log.Fatal("Failed to read urandom file")
	}

	return buffer
}

// CreatePRNGFile creates a file of random data using a pseudo random number generator
func CreatePRNGFile(fileName string, fileSizeBytes int) {
	file, err := os.Create(fileName)
	if err != nil {
		log.Fatal("Failed to create PRNG file")
	}
	defer file.Close()

	bytesWritten, err := file.Write(readUrandom(fileSizeBytes))
	if err != nil {
		log.Fatal("Failed to write PRNG data")
	}
	fmt.Printf("Bytes written: %d\n", bytesWritten)
}
