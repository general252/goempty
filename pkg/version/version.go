package version

import "fmt"

var (
	Version   string = "0.0.1"
	BuildTime string = "2006-01-02 15:04:05"
)

func ShowVersionInfo() {
	fmt.Printf("=================================\n")
	fmt.Printf("   version: %v\n", Version)
	fmt.Printf("build time: %v\n", BuildTime)
	fmt.Printf("=================================\n\n")
}
