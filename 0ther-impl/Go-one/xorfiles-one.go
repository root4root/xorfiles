package main

import (
	"encoding/binary"
	"fmt"
	"io"
	"os"
)

// Just FYI:
// The kernel buffer, created by the pipe system call from the shell, is sized based on the page size for the system.
// The linux pipe buffers have changed to circular buffers (16 x 4KiB).
const bufsiz int = 524288 // Default 1024 * 512 bytes, unless STDIN is used

func main() {
	var firstFile, secondFile, outputFile *os.File

	/// Start processing CLI params ---->
	if (len(os.Args) < 2) || (len(os.Args) > 4) {
		fmt.Println("\nUsage:\n$ xorfiles {first-file | stdin} second-file [output-file]\n")
		os.Exit(22)
	}

	switch len(os.Args) {
	case 2:
		firstFile = os.Stdin
		secondFile = openRead(os.Args[1])
		outputFile = os.Stdout
	case 3:
		if _, err := os.Stat(os.Args[2]); err == nil {
			firstFile = openRead(os.Args[1])
			secondFile = openRead(os.Args[2])
			outputFile = os.Stdout
		} else {
			firstFile = os.Stdin
			secondFile = openRead(os.Args[1])
			outputFile = openWrite(os.Args[2])
		}
	default:
		firstFile = openRead(os.Args[1])
		secondFile = openRead(os.Args[2])
		outputFile = openWrite(os.Args[3])
	} /// <----- End processing CLI params

	defer func() {
		outputFile.Sync()

		outputFile.Close()
		firstFile.Close()
		secondFile.Close()
	}()

	firstStat, _ := firstFile.Stat()
	secondStat, _ := secondFile.Stat()

	if firstStat.Size() != secondStat.Size() {
		fmt.Fprintf(os.Stderr, "notice! Input files has different sizes\n")
	}

	firstBuffer := make([]byte, bufsiz)
	secondBuffer := make([]byte, bufsiz)
	outputBuffer := make([]byte, bufsiz)

	//Convert each 8 bytes from buffer to uint64 (decrease CPU operations by using x64 CPU registers)
	var operandA, operandB, xorResult uint64

	for {
		firstAmount, firstReadError := firstFile.Read(firstBuffer)
		secondAmount, secondReadError := secondFile.Read(secondBuffer[:firstAmount])
		minAmount := firstAmount

		if secondAmount < firstAmount {
			minAmount = secondAmount
		}

		if firstReadError == io.EOF || secondReadError == io.EOF {
			break
		}

		if minAmount&0b111 > 0 { //Same as dBatch.minSize % 8 > 0
			for i := 0; i < minAmount; i++ {
				outputBuffer[i] = firstBuffer[i] ^ secondBuffer[i]
			}

			writeData(outputFile, outputBuffer[:minAmount])
			continue
		}

		for i := 0; i < minAmount; i += 8 {
			operandA = binary.BigEndian.Uint64(firstBuffer[i:])
			operandB = binary.BigEndian.Uint64(secondBuffer[i:])
			xorResult = operandA ^ operandB
			binary.BigEndian.PutUint64(outputBuffer[i:], xorResult)
		}

		writeData(outputFile, outputBuffer[:minAmount])
	}

}

func writeData(file *os.File, data []byte) {
	_, err := file.Write(data)

	if err != nil {
		fmt.Fprintf(os.Stderr, "ERROR! %s\n", err.Error())
		os.Exit(5)
	}
}

func openRead(filePath string) *os.File {

	file, error := os.Open(filePath)

	if error != nil {
		fmt.Fprintf(os.Stderr, "ERROR! %s\n", error.Error())
		os.Exit(2)
	}

	return file
}

func openWrite(filePath string) *os.File {

	if _, err := os.Stat(filePath); err == nil {
		fmt.Fprintf(os.Stderr, "ERROR! File already exists (%s)\n", filePath)
		os.Exit(17)
	}

	file, error := os.Create(filePath)

	if error != nil {
		fmt.Fprintf(os.Stderr, "ERROR! %s", error.Error())
		os.Exit(1)
	}

	return file
}
