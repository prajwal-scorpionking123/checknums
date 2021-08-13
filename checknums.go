package checknums

func CheckEven(number int) bool {
	if number%2 == 0 {
		return true
	} else {
		return false
	}
}

func CheckPalindrome(number int) bool {
	var remainder, temp int
	var reverse int = 0
	for {
		remainder = number % 10
		reverse = reverse*10 + remainder
		number /= 10

		if number == 0 {
			break // Break Statement used to exit from loop
		}
	}

	if temp == reverse {
		return true
	} else {
		return false
	}
}
