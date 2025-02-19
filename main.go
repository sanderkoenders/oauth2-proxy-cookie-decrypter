package main

import (
	"fmt"
	"strings"
	"encoding/base64"
	"bytes"
	"io"
	"encoding/json"
	"os"

	"github.com/oauth2-proxy/oauth2-proxy/v7/pkg/encryption"
	"github.com/pierrec/lz4/v4"
	"github.com/vmihailenco/msgpack/v5"
	"github.com/joho/godotenv"
)

type SessionState struct {
	AccessToken  string `msgpack:"at,omitempty"`
	IDToken      string `msgpack:"it,omitempty"`
}

func main() {
	var cookie = os.Args[1]

	secret, err := getCookieSecret()
	if err != nil {
		fmt.Printf("Could not load cookie secret: %v", err)
	}

	cipher, err := encryption.NewCFBCipher(encryption.SecretBytes(secret))

	if err != nil {
		fmt.Printf("error initialising cipher: %v", err)
	}

	encrypted, err := base64.URLEncoding.DecodeString(strings.Split(cookie, "|")[0])
	if err != nil {
		fmt.Println("error decoding string")
	}

	decrypted, err := cipher.Decrypt(encrypted)

	if err != nil {
		fmt.Printf("error decrypting the session state: %w", err)
	}

	packed := decrypted

	packed, err = lz4Decompress(decrypted)
	if err != nil {
		fmt.Println("Could not decompress")
	}
	
	var ss SessionState
	err = msgpack.Unmarshal(packed, &ss)

	if err != nil {
		fmt.Println("error unmarshalling data to session state: %w", err)
	}

	jsonData, err := json.MarshalIndent(ss, "", "\t")
	if err != nil {
		fmt.Println("Error converting struct to JSON: %v", err)
	}

	fmt.Println(string(jsonData))
}

func getCookieSecret() (string, error) {
	err := godotenv.Load()
	if err != nil {
		return "", fmt.Errorf("Error loading .env file: %v", err)
	}

	cookieSecret := os.Getenv("COOKIE_SECRET")
	if cookieSecret == "" {
		return "", fmt.Errorf("COOKIE_SECRET not set in .env file")
	}

	return cookieSecret, nil
}

// lz4Decompress decompresses with LZ4
func lz4Decompress(compressed []byte) ([]byte, error) {
	reader := bytes.NewReader(compressed)
	buf := new(bytes.Buffer)
	zr := lz4.NewReader(nil)
	zr.Reset(reader)
	_, err := io.Copy(buf, zr)
	if err != nil {
		return nil, fmt.Errorf("error copying lz4 stream to buffer: %w", err)
	}

	payload, err := io.ReadAll(buf)
	if err != nil {
		return nil, fmt.Errorf("error reading lz4 buffer: %w", err)
	}

	return payload, nil
}