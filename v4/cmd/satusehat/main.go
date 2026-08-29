// Command-line tool for SATUSEHAT SDK version info.
package main

import (
	"fmt"
	"os"
)

var version = "v4.11.2"

func main() {
	fmt.Println("SATUSEHAT FHIR R4 SDK for Go")
	fmt.Println("Version:", version)
	fmt.Println("Docs: https://github.com/ivanwilliammd/satusehat-integration-go/tree/main/v4")
	os.Exit(0)
}
