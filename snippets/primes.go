package snippets

// IsPrime is an unoptimized function to verify if the number is prime or not
func IsPrime(number int) bool {
	if number < 2 {
		return false
	}
	for i := 2; i < number; i++ {
		if number%i == 0 {
			return false
		}
	}
	return true
}
