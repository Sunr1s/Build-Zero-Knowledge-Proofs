package main

import (
	"bufio"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math/big"
	"os"
	"strings"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Buffer(make([]byte, 1024*1024), 1024*1024)
	var g, p, q, priv, nonce, y, r *big.Int
	var ok1, ok2, ok3, ok4, ok5 bool
	var rBytes, yBytes []byte
	for sc.Scan() {
		parts := strings.Fields(sc.Text())
		if len(parts) == 0 {
			continue
		}
		if len(parts) == 10 && parts[0] == "G" && parts[2] == "P" &&
			parts[4] == "Q" && parts[6] == "X" && parts[8] == "K" {
			g, ok1 = new(big.Int).SetString(parts[1], 10)
			p, ok2 = new(big.Int).SetString(parts[3], 10)
			q, ok3 = new(big.Int).SetString(parts[5], 10)
			priv, ok4 = new(big.Int).SetString(parts[7], 10)
			nonce, ok5 = new(big.Int).SetString(parts[9], 10)

			if !ok1 || !ok2 || !ok3 || !ok4 || !ok5 {
				panic("invalid number in header")
			}

			y = new(big.Int).Exp(g, priv, p)
			r = new(big.Int).Exp(g, nonce, p)
			rBytes = r.Bytes()
			yBytes = y.Bytes()
		}

		if parts[0] == "SIGN" {
			if len(parts) != 2 || q == nil || priv == nil || nonce == nil || r == nil {
				continue
			}
			messageBytes, err := hex.DecodeString(parts[1])
			if err != nil {
				continue
			}
			data := append([]byte{}, rBytes...)
			data = append(data, messageBytes...)
			data = append(data, yBytes...)
			digest := sha256.Sum256(data)
			c := new(big.Int).SetBytes(digest[:])
			c.Mod(c, q)
			product := new(big.Int).Mul(c, priv)
			z := new(big.Int).Add(nonce, product)
			z.Mod(z, q)
			fmt.Println(r.Text(16), z.Text(16))
		}

	}
}
