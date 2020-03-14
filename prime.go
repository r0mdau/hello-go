package premier

import (
	"fmt"
	"os"
	"strconv"
)

func premier(n int) bool {
	if n%2 == 0 {
		return n == 2
	}

	for d := 3; (d * d) <= n; d += 2 {
		if n%d == 0 {
			return false
		}
	}
	return true
}

func main() {
	number, _ := strconv.Atoi(os.Args[1])
	fmt.Println(number)

	if premier(number) {
		fmt.Println("premier !")
	} else {
		fmt.Println("pas un nombre premier")
	}

}
