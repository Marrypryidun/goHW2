package main
import "fmt"

func main() {
	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(s)

	s = s[:0]
	fmt.Println(s)

	s = s[:4]
	fmt.Println(s)

	// чому так не можна?
	/*s = s[:0]
	fmt.Println(s)*/

	s = s[:cap(s)]
	fmt.Println(s)

	s = s[2:6]
	fmt.Println(s)

	s = s[:cap(s)]
	fmt.Println(s)
}