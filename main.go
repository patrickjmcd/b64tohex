package main

import (
	"encoding/base64"
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "b64tohex",
	Short: "AWS IoT Core Lambdas",
	Long:  `Complete documentation is available at http://github.com/meshifyiot/aws-iot-core-lambdas`,
	Run: func(cmd *cobra.Command, args []string) {
		payload := args[0]
		log.Printf("Converting Payload: %s", payload)
		b64Decoded, err := base64.StdEncoding.DecodeString(payload)
		if err != nil {
			log.Fatalf("failed to base64 decode payload, %s", err)
		}
		log.Println("Raw Payload: ", b64Decoded)
		log.Printf("Hex Payload: %x", b64Decoded)
	},
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
