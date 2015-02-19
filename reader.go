package main

import (
	"bufio"
	"os"
)

func readStdin() {
	scanner := bufio.NewScanner(os.Stdin)

	// Use a scanner to read from stdin and write this to stderr.
	for scanner.Scan() {
		if _, err := os.Stderr.WriteString(scanner.Text()); err != nil {
			Log.WithField("error", err.Error()).Errorf("Error writing scanned stdin text.")
		}
		if _, err := os.Stderr.WriteString("\n"); err != nil {
			Log.WithField("error", err.Error()).Errorf("Error writing newline.")
		}
	}

	if err := scanner.Err(); err != nil {
		Log.WithField("error", err.Error()).Errorf("Error scanning standard input.")
	}
}
