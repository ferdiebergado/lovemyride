package env

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

// LoadEnv loads environment variables from a .env file
func LoadEnv(envFile string) error {
	file, err := os.Open(envFile)

	if err != nil {
		return err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		// Skip empty lines and comments
		if len(line) == 0 || strings.HasPrefix(line, "#") {
			continue
		}

		// Split the line into key and value
		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			continue
		}
		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])

		// Set the environment variable
		if err := os.Setenv(key, value); err != nil {
			return err
		}
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}

func Must(v string) string {
	res, exists := os.LookupEnv(v)

	if !exists {
		fmt.Fprintf(os.Stderr, "%s not set!\n", v)
		os.Exit(1)
	}

	return res
}

func GetEnv(v string, def string) string {
	res, exists := os.LookupEnv(v)

	if !exists {
		log.Println(v + " does not exist, using default of " + def)
		return def
	}

	return res
}
