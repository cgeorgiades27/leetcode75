package asteroidcollision

func asteroidCollision(asteroids []int) []int {
	var winners []int
	for i := len(asteroids) - 1; i > 0; i-- {
		right := asteroids[i]
		left := asteroids[i-1]

		if !willCollide(left, right) {
			continue
		}

		if absv(left) == absv(right) {
			asteroids[i-1] = 0
			asteroids[i] = 0
		}

		if absv(left) > absv(right) {
			asteroids[i] = 0
		} else {
			asteroids[i-1] = 0
		}
	}

	for _, asteroid := range asteroids {
		if asteroid != 0 {
			winners = append(winners, asteroid)
		}
	}

	return winners
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
