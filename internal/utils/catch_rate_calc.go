package utils

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func CatchRateCalc(baseXP int) bool {
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))

	const maxXP = 400.0
	// Make basexp as ratio by /maxXP to compare with random number generated
	catchRate := 1 - float64(baseXP)/maxXP

	catchRate = math.Max(catchRate, 0.1) // Set Min catch rate at 10% math.Max return larger value between catchRate, and 0.1

	randomRoll := rng.Float64() // Generate random num between 0 and 1 (float64)
	fmt.Printf("Catch Rate: %.2f, Random Roll: %.2f\n", catchRate, randomRoll)
	return randomRoll < catchRate // If random num is less than catch rate, return true
}

// Probability explanation:
// With 0.4 catch rate, catch rate is 40%, the randomRoll need to be lower than 0.4. as per visual
// [0.0 ................ 0.4 | 0.4 .................. 1.0]
//   ^                   ^     ^                      ^
// roll = 0.22      <-- Success   Fail -->       roll = 0.61
//
//     The range from 0.0 to 0.4 represents success (40% of the total range).
//
//     The range from 0.4 to 1.0 represents failure (60% of the total range).
