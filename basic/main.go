package main

import "fmt"

type car struct {
	brand      string
	model      string
	frontWheel wheel
	backWheel  wheel
}

type wheel struct {
	radius   int
	material string
}

type product struct {
	name        string
	description string
	brand       struct {
		name string
		info string
	}
}

func printVariableType(v interface{}) {
	switch t := v.(type) {
	case int:
		fmt.Printf("%T\n", t)
	case string:
		fmt.Printf("%T\n", t)
	default:
		fmt.Printf("We don't have this %T type\n", t)
	}
}

func main() {
	lambo := car{
		brand: "Lamborginig",
		model: "v1",
		frontWheel: wheel{
			radius:   120,
			material: "iron",
		},
		backWheel: wheel{},
	}

	unknown := car{}

	watch := product{
		name:        "Watch",
		description: "This is the watch",
		brand: struct {
			name string
			info string
		}{
			name: "Jumanji",
			info: "This is the premium watch company",
		},
	}

	fmt.Printf("My watch %#v\n", watch)
	fmt.Printf("My car %#v\n", lambo)
	fmt.Printf("Not define car %#v\n", unknown)

	printVariableType(10)
	printVariableType("ad")
	printVariableType(23.12)
	// Logger()
	RunPingPong()
	// RunProcess()
	runReadWrite()
}
