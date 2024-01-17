package belajargolangembed

import (
	_ "embed"
	"fmt"
	"io/fs"
	"io/ioutil"
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

// Embed file ke []Byte

//go:embed logo.jpeg
var logo []byte

func TestByte(t *testing.T) {
	err := ioutil.WriteFile("logo_new.jpeg", logo, fs.ModePerm)
	if err != nil {
		panic(err)
	}
}
