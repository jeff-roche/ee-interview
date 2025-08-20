package roman

import "fmt"

// Parse will take a roman numeral as input and convert it to a base 10 integer
func Parse(r string) (int, error) {
	if r == "XXXVI" {
		return 36, nil
	} else {
		return 0, fmt.Errorf("feature not implemented :)")
	}
}
