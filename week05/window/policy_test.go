package window

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func GetRollingPolicy() *RollingPolicy {
	w := NewWindow(Options{Size: 3})
	return NewRollingPolicy(w, RollingPolicyOpts{BucketDuration: 100 * time.Millisecond})
}

func TestRollingPolicy_Add(t *testing.T) {
	// test func timespan return real span
	tests := []struct {
		timeSleep []int
		offset    []int
		points    []int
	}{
		{
			timeSleep: []int{150, 51},
			offset:    []int{1, 2},
			points:    []int{1, 1},
		},
		{
			timeSleep: []int{90, 250},
			offset:    []int{0, 0},
			points:    []int{1, 1},
		},
		{
			timeSleep: []int{150, 300, 600},
			offset:    []int{1, 1, 1},
			points:    []int{1, 1, 1},
		},
	}

	for _, test := range tests {
		t.Run("test policy add", func(t *testing.T) {
			var totalTs, lastOffset int
			timeSleep := test.timeSleep
			policy := GetRollingPolicy()
			for i, n := range timeSleep {
				totalTs += n
				time.Sleep(time.Duration(n) * time.Millisecond)
				policy.Add(float64(test.points[i]))
				offset, points := test.offset[i], test.points[i]
				assert.Equal(t, points, int(policy.window.buckets[offset].Points[0]),
					fmt.Sprintf("error, time since last append: %vms, last offset: %v", totalTs, lastOffset))
				lastOffset = offset
			}
		})
	}
}
