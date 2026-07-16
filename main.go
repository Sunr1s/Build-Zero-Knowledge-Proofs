package main

import (
	"bufio"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"os"
	"strings"
)

func commitment(value, randomness string) string {
	hash := sha256.Sum256([]byte(value + randomness))
	return hex.EncodeToString(hash[:8])
}

func main() {
	// История коммитментов: идентификатор -> коммитмент.
	history := make(map[string]string)

	sc := bufio.NewScanner(os.Stdin)
	sc.Buffer(make([]byte, 1024*1024), 1024*1024)

	for sc.Scan() {
		parts := strings.Fields(sc.Text())
		if len(parts) == 0 {
			continue
		}

		switch parts[0] {
		case "COMMIT":
			if len(parts) != 4 {
				continue
			}
			id, value, randomness := parts[1], parts[2], parts[3]
			result := commitment(value, randomness)
			history[id] = result
			fmt.Println(result)

		case "VERIFY":
			if len(parts) != 4 {
				continue
			}
			id, value, randomness := parts[1], parts[2], parts[3]
			result := commitment(value, randomness)
			if saved, ok := history[id]; ok && saved == result {
				fmt.Println("OK")
			} else {
				fmt.Println("BAD")
			}
		}
	}
}
