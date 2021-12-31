package window

// Sum the values within the window.
func Sum(iterator Iterator) float64 {
	var result = 0.0
	for iterator.Next() {
		bucket := iterator.Bucket()
		for _, p := range bucket.Points {
			result = result + p
		}
	}
	return result
}

// Avg the values within the window.
func Avg(iterator Iterator) float64 {
	var result = 0.0
	var count = 0.0
	for iterator.Next() {
		bucket := iterator.Bucket()
		for _, p := range bucket.Points {
			result = result + p
			count = count + 1
		}
	}
	return result / count
}

// Min the values within the window.
func Min(iterator Iterator) float64 {
	var result = 0.0
	var started = false
	for iterator.Next() {
		bucket := iterator.Bucket()
		for _, p := range bucket.Points {
			if !started {
				result = p
				started = true
				continue
			}
			if p < result {
				result = p
			}
		}
	}
	return result
}

// Max the values within the window.
func Max(iterator Iterator) float64 {
	var result = 0.0
	var started = false
	for iterator.Next() {
		bucket := iterator.Bucket()
		for _, p := range bucket.Points {
			if !started {
				result = p
				started = true
				continue
			}
			if p > result {
				result = p
			}
		}
	}
	return result
}
