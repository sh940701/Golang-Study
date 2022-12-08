func getValue(value int) (int) {
	return value
}

func main() {
	x := getValue(10)
	y := 20

	switch x {
	case 10:
		fmt.Println("x == 10")
	case 20:
		fmt.Println("x == 20")
	}

	switch y {
	case 10:
		fmt.Println("y == 10")
	case 20:
		fmt.Println("y == 20")
	}
}