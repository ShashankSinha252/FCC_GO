package main

import (
	"fmt"
)

const (
	isAdmin = 1 << iota
	isHeadquarters
	canSeeFinancials

	canSeeAfrica
	canSeeAsia
	canSeeEurope
	canSeeNorthAmerica
	canSeeSouthAmerica
)

func main() {
	var role byte = isAdmin | canSeeFinancials | canSeeEurope
	fmt.Printf("%b\n", role)
    fmt.Printf("Is admin : %v\n", isAdmin & role == isAdmin)
    fmt.Printf("Is HQ : %v\n", isHeadquarters & role == isHeadquarters)
}
