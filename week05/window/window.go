package window

type Bucket struct {
	Points []float64
	Count  int64
	next   *Bucket
}

// 向Points追加val
func (b *Bucket) Append(val float64) {
	b.Points = append(b.Points, val)
	b.Count++
}

// 累加val到Points
func (b *Bucket) Add(offset int, val float64) {
	b.Points[offset] += val
	b.Count++
}

// 重置Bucket
func (b *Bucket) Reset() {
	b.Points = b.Points[:0]
	b.Count = 0
}

// return the next bucket
func (b *Bucket) Next() *Bucket {
	return b.next
}

type Window struct {
	buckets []Bucket
	size    int
}

type Options struct {
	Size int
}

// NewWindow creates a new Window based on WindowOpts.
func NewWindow(opts Options) *Window {
	buckets := make([]Bucket, opts.Size)
	for offset := range buckets {
		buckets[offset].Points = make([]float64, 0)
		nextOffset := offset + 1
		if nextOffset == opts.Size {
			nextOffset = 0
		}
		buckets[offset].next = &buckets[nextOffset]
	}
	return &Window{buckets: buckets, size: opts.Size}
}

func (w *Window) ResetBuckets(offset int, count int) {
	for i := 0; i < count; i++ {
		w.ResetBucket(offset + i)
	}
}

// 重置桶
func (w *Window) ResetBucket(offset int) {
	w.buckets[offset%w.size].Reset()
}

// 重置窗口
func (w *Window) ResetWindow() {
	for offset := range w.buckets {
		w.ResetBucket(offset)
	}
}

// 向桶中追加val
func (w *Window) Append(offset int, val float64) {
	w.buckets[offset%w.size].Append(val)
}

// 向桶中累加val
func (w *Window) Add(offset int, val float64) {
	offset %= w.size
	if w.buckets[offset].Count == 0 {
		w.buckets[offset].Append(val)
		return
	}
	w.buckets[offset].Add(0, val)
}

// 返回指定索引的桶
func (w *Window) Bucket(offset int) Bucket {
	return w.buckets[offset%w.size]
}

// 返回窗口大小
func (w *Window) Size() int {
	return w.size
}

// 返回一个窗口迭代器
func (w *Window) Iterator(offset int, count int) Iterator {
	return Iterator{
		count: count,
		cur:   &w.buckets[offset%w.size],
	}
}
