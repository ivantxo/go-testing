package main

import (
	"fmt"
)

const (
	// bit shifting operations
	isAdmin = 1 << iota
	isHeadquarters
	canSeeFinancials

	canSeeAfrica
	canSeeAsia
	canSeeEurope
	canSeeSouthAmerica
	canSeeNorthAmerica
)

func main() {
	var roles byte = isAdmin | canSeeFinancials | canSeeEurope
	fmt.Printf("Roles: %b\n", roles)

	fmt.Printf("Is Admin? %v\n", isAdmin&roles == isAdmin)
	fmt.Printf("Is Head Quarters? %v\n", isHeadquarters&roles == isHeadquarters)
}
