package main

import (
	"bufio"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"os"
	"strings"
)

// TODO (interactive-non-interactive): implement per the lesson description.

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Buffer(make([]byte, 1024*1024), 1024*1024)

	for sc.Scan() {

		parts := strings.Fields(sc.Text())
		if len(parts) != 3 || parts[0] != "PROOF" {
			continue
		}

		commitmentBytes, err := hex.DecodeString(parts[1])
		if err != nil {
			continue
		}

		statementBytes, err := hex.DecodeString(parts[2])
		if err != nil {
			continue
		}

		concatBytes := append(commitmentBytes, statementBytes...)
		hash := sha256.Sum256(concatBytes)

		result := hex.EncodeToString(hash[:16])
		fmt.Println(result)
	}
}
