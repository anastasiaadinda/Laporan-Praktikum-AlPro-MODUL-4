package main

import "fmt"

func main() {
	var nilaibmi, tinggibadan, beratbadan float64
	fmt.Scanln(&nilaibmi, &tinggibadan)
	beratbadan = nilaibmi * (tinggibadan * tinggibadan)
	fmt.Printf("%.f", beratbadan)
}
