package resources

import "time"

// MockedTime is variable to use in tests to mock time.Now when using CurrentTime
var MockedTime time.Time

// CurrentTime use this when time.Now is needed to be mocked
func CurrentTime() time.Time {
	if MockedTime.IsZero() {
		return time.Now()
	}

	return MockedTime
}
