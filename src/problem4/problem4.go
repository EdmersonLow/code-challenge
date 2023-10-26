// Time O(n) Space O(1)
func sum_to_n_a(n int) int {
	// your code here
	if n < 0 {
        return 0
    }
    sum := 0
    for i := 1; i <= n; i++ {
        sum += i
    }
    return sum
}
// Time O(1) Space O(1)
func sum_to_n_b(n int) int {
	// your code here
	if n < 0 {
        return 0
    }
    return n * (n + 1) / 2
}
// Time O(n) Space O(n)
func sum_to_n_c(n int) int {
	// your code here
	if n <= 0 {
        return 0
    }
    return n + sum_to_n_c(n-1)
}