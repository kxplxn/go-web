package _030103_structuringApps_

import "fmt"

type Car struct {
	Make  string
	Model string
	Year  int
}

type Person struct {
	FirstName string
	LastName  string
}

// Ignore below. This file should only contain the above.

func Demo() {
	fmt.Println("\n030103 Basics: Structuring Apps")
	fmt.Println("To demo this chapter, investigate the source code in chapter directory,")
	fmt.Println("run `go run 01-basics/03-structuringApps/cmd/webserver/main.go from project root,")
	fmt.Println("visit localhost:8080, and browse: `car/html`, `/car/plaintext`, `/person/html`, and `/person/plaintext`.")
}
