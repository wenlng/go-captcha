/**
 * @Author Awen
 * @Date 2024/06/01
 * @Email wengaolng@gmail.com
 **/

package random

import (
	rand2 "crypto/rand"
	"math"
	"math/big"
	"math/rand"
	"time"
)

var rnd = newRnd()

// newRnd .
func newRnd() *rand.Rand {
	var seed = time.Now().UnixNano()
	var src = rand.NewSource(seed)
	return rand.New(src)
}

// Rand63n generates a 64-bit random number
func Rand63n(ri int64) {
	rnd.Int63n(ri)
}

// Rand31n generates a 32-bit random number
func Rand31n(ri int32) {
	rnd.Int31n(ri)
}

// Perm generates a random permutation
func Perm(n int) []int {
	return rand.Perm(n)
}

// RandInt generates a safe random number in the interval [n, m]
func RandInt(min, max int) int {
	if min > max {
		return max
	}

	if min < 0 {
		f64Min := math.Abs(float64(min))
		i64Min := int(f64Min)
		result, _ := rand2.Int(rand2.Reader, big.NewInt(int64(max+1+i64Min)))
		return int(result.Int64() - int64(i64Min))
	}

	result, _ := rand2.Int(rand2.Reader, big.NewInt(int64(max-min+1)))
	return int(int64(min) + result.Int64())
}
