package main

const minInt32 = -2147483648
const maxInt32 = 2147483647

func reverse(x int) int {
	var out int
	for x != 0 {
		out = out*10 + x%10
		x /= 10
		if out < minInt32 || out > maxInt32 {
			return 0
		}
	}
	return out
}

func main() {
	result := reverse(123)
	println(result) // Output: 321

	result = reverse(-123)
	println(result) // Output: -321

	result = reverse(120)
	println(result) // Output: 21

	result = reverse(1534236469)
	println(result) // Output: 0 (out of bounds)
}
