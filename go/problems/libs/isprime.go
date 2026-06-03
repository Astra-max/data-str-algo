package libs

//import "fmt"

func IsPrime(num int) bool {
	if num < 2 {
		return false
	}
	if num == 2 {
		return true
	}
	if num % 2 == 0 {
		return false
	}

	for i:= 3; i*i<=num; i+=2 {
		if num % i == 0 {
			return false
		}
	}
	return true
}

func Sieve(num int) []int {
	if num < 2 {
		return []int{}
	}

	isCompo := make([]bool, num+1)
	isCompo[0] = true
	isCompo[1] = true

	for i:=2; i*i<=num; i++ {
		if !isCompo[i] {
			for mult := i*i; mult <= num; mult += i {
				isCompo[mult] = true
			}
		}
	}

	prime := []int{}

	for i:=2; i<=num; i++ {
		if !isCompo[i] {
		prime = append(prime, i)

		}
	}
	return prime
}