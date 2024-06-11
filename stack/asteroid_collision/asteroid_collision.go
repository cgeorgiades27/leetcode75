package asteroidcollision

// asteroidCollision returns the asteroids that will remain after all collisions
func asteroidCollision(asteroids []int) []int {
	stackPtr := len(asteroids) - 1
	for i := len(asteroids) - 1; i > 0; i-- {
		// pop right
		right := asteroids[i]
		stackPtr--

		// pop left
		left := asteroids[i-1]
		stackPtr--

		if !willCollide(left, right) {
			stackPtr += 2
			continue
		}

		labsv := absv(left)
		rabsv := absv(right)

		if labsv == rabsv {
			continue
		}

		stackPtr++
		if labsv > rabsv {
			asteroids[stackPtr] = left
		} else {
			asteroids[stackPtr] = right
		}
	}

	if stackPtr < 0 {
		return []int{}
	}

	return asteroids[:stackPtr+1]
}

func willCollide(l, r int) bool {
	sum := l + r
	return sum < l || sum < r
}

func absv(n int) int {
	if n < 0 {
		return n * -1
	}
	return n
}
