package window

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWindow_ResetWindow(t *testing.T) {
	opts := Options{Size: 3}
	window := NewWindow(opts)
	for i := 0; i < opts.Size; i++ {
		window.Append(i, 1)
	}
	window.ResetWindow()
	for i := 0; i < opts.Size; i++ {
		assert.Equal(t, len(window.Bucket(i).Points), 0)
	}
}

func TestWindow_Bucket(t *testing.T) {
	opts := Options{Size: 3}
	window := NewWindow(opts)
	for i := 0; i < opts.Size; i++ {
		window.Append(i, 1)
	}
	window.ResetBucket(0)
	assert.Equal(t, window.Bucket(0).Points, []float64{})
	assert.Equal(t, window.Bucket(1).Points[0], float64(1))
	assert.Equal(t, window.Bucket(2).Points[0], float64(1))
}

func TestWindow_ResetBuckets(t *testing.T) {
	opts := Options{Size: 3}
	window := NewWindow(opts)
	for i := 0; i < opts.Size; i++ {
		window.Append(i, 1)
	}
	window.ResetBuckets(0, opts.Size)
	for i := 0; i < opts.Size; i++ {
		assert.Equal(t, len(window.Bucket(i).Points), 0)
	}
}

func TestWindow_Append(t *testing.T) {
	opts := Options{Size: 3}
	window := NewWindow(opts)
	for i := 0; i < opts.Size; i++ {
		window.Append(i, 1)
	}
	for i := 0; i < opts.Size; i++ {
		assert.Equal(t, window.Bucket(i).Points[0], float64(1))
	}
}

func TestBucket_Add(t *testing.T) {
	opts := Options{Size: 3}
	window := NewWindow(opts)
	window.Append(0, 1)
	window.Add(0, 100)
	assert.Equal(t, window.Bucket(0).Points[0], float64(101))
}
