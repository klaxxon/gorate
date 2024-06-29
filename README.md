# gorate
Simple golang closure to track and return a rate.

This is a per second rate tracker.  For an average rate over the past ten seconds:
<code>
x := stats.GetRateFunc(10)  // Save ten seconds of data

for {
  x(false)  // Count event
  fmt.Printf("Rate %f\n", x(true))  // Display current event rate
  time.Sleep(time.Millisecond * 500)
}
</code>

