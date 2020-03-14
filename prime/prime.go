package prime

import (
	"fmt"
	"os"
	"strconv"
)

func isPrime(n int) bool {
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

const primeResultPrint = "prime !"
const notPrimeResultPrint = "not prime :'("

func resultToPrint(isPrime bool) string {
	if isPrime {
		return primeResultPrint
	}
	return notPrimeResultPrint
}

func main() {
	number, _ := strconv.Atoi(os.Args[1])
	fmt.Println(number)
	fmt.Println(resultToPrint(isPrime(number)))
}
