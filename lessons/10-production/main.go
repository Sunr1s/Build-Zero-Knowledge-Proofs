package main

import (
	"bufio"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"os"
	"strings"
)

// domainSeparatedHash returns SHA-256(protocolID + ":" + data)
// as a lowercase hexadecimal string.
func domainSeparatedHash(protocolID string, data []byte) string {
	message := []byte(protocolID)
	message = append(message, ':')
	message = append(message, data...)
	digest := sha256.Sum256(message)

	return hex.EncodeToString(digest[:])
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Buffer(make([]byte, 1024*1024), 1024*1024)

	for sc.Scan() {
		parts := strings.Fields(sc.Text())
		if len(parts) != 2 {
			continue
		}

		protocolID := parts[0]
		data, err := hex.DecodeString(parts[1])
		if err != nil {
			continue
		}

		fmt.Println(domainSeparatedHash(protocolID, data))
	}
}
