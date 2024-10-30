package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"os"

	"gopkg.org/proc"
)

var ErrInvalidInput = errors.New("invalid file provided")

func isELF(f *os.File) (bool, error) {
	if _, err := f.Seek(0, io.SeekStart); err != nil {
		return false, fmt.Errorf("failed to seek to start of file: %w", err)
	}

	header, err := io.ReadAll(io.LimitReader(f, 32))
	if err != nil {
		return false, fmt.Errorf("failed to read file header: %w", err)
	}

	hAsStr := string(header)

	return hAsStr[:4] == "\u007FELF", nil
}

func main() {
	path := "./hello-world/main"

	bin, err := os.Open(path)
	if err != nil {
		log.Fatalf("failed to open %s: %v", path, err)
	}

	elf, err := isELF(bin)
	if err != nil {
		log.Fatalf("couldn't determine whether bin was ELF: %v", err)
	}

	if !elf {
		log.Fatalf("%v: input file isn't an ELF", ErrInvalidInput)
	}

	self, err := proc.Self()
	if err != nil {
		log.Fatalf("failed to get information about the current process: %v", err)
	}

	fmt.Printf("%d\n", self.PID)

}
