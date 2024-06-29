
// GetRateFunc returns the rate/sec for the previous setup
// A bool of true will simply return the rate
func GetRateFunc(binsize int) func(bool) float64 {
	bins := make([]int, binsize)
	var sum int
	lastUpdate := time.Now().Round(time.Second)

	return func(get bool) float64 {
		current := int(time.Now().Round(time.Second).Second()) % binsize
		last := int(lastUpdate.Second()) % binsize
		// Zero out any bins which may have passed
		elapsed := time.Since(lastUpdate).Round(time.Second).Seconds()
		if int(elapsed) >= binsize { // Past binsize seconds? CLear everything out
			sum = 0
			clear(bins)
		} else {
      // Zero everything from one second past last position to current position
			for {
				if last == current {
					break
				}
				last = (last + 1) % binsize
				sum -= bins[last]
				bins[last] = 0
			}
		}
		if !get {  // Add!
			sum++
			bins[current]++
			lastUpdate = time.Now().Round(time.Second)
			return float64(sum) / float64(binsize)
		}
		return float64(sum) / float64(binsize)
	}
}
