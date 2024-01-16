package belajargolangembed

import (
	_ "embed"
	"fmt"
	"testing"
)

// Embed file ke String

//go:embed version.txt
var version string

//go:embed version.txt
var version2 string

func TestString(t *testing.T) {
	fmt.Println(version)
	fmt.Println(version2)
}
