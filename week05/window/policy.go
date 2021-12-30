package window

import (
	"sync"
	"time"
)

// 按照一定时间间隔统计的策略
type RollingPolicy struct {
	mu     sync.RWMutex
	size   int
	window *Window
	offset int

	bucketDuration time.Duration
	lastAppendTime time.Time
}

type RollingPolicyOpts struct {
	BucketDuration time.Duration
}

func NewRollingPolicy(window *Window, opts RollingPolicyOpts) *RollingPolicy {
	return &RollingPolicy{
		window: window,
		size:   window.Size(),
		offset: 0,

		bucketDuration: opts.BucketDuration,
		lastAppendTime: time.Now(),
	}
}

func (r *RollingPolicy) timespan() int {
	v := int(time.Since(r.lastAppendTime) / r.bucketDuration)
	if v > -1 {
		return v
	}
	return r.size
}

func (r *RollingPolicy) apply(f func(offset int, val float64), val float64) {
	r.mu.Lock()
	defer r.mu.Unlock()

	// 计算时间间隔 取整
	timespan := r.timespan()
	if timespan > 0 {
		// 从哪里开始 清除过期的 bucket
		start := (r.offset + 1) % r.size
		// 存入bucket的offset
		end := (r.offset + timespan) % r.size
		// 判断时间间隔过长超出窗口
		if timespan > r.size {
			timespan = r.size
		}
		// 清除过期的 bucket
		r.window.ResetBuckets(start, timespan)
		// 重新更新offset
		r.offset = end
		// 更新插入时间 取整
		r.lastAppendTime = r.lastAppendTime.Add(time.Duration(timespan * int(r.bucketDuration)))
	}
	f(r.offset, val)
}

func (r *RollingPolicy) Append(val float64) {
	r.apply(r.window.Append, val)
}

// Add adds the given value to the latest point within bucket.
func (r *RollingPolicy) Add(val float64) {
	r.apply(r.window.Add, val)
}

// 聚合函数
func (r *RollingPolicy) Reduce(f func(iterator Iterator) float64) (val float64) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	// 如果是先添加元素在调 Reduce方法，此时的timespan为0, 因为在apply方法中更新了lastAppendTime
	timespan := r.timespan()
	// 迭代次数
	if count := r.size - timespan; count > 0 {
		// 因为加了1所以初始情况是从 1、2、0 开始
		offset := r.offset + timespan + 1
		// 让timespan收敛于size
		if offset >= r.size {
			offset = offset - r.size
		}
		// window.Iterator返回Iterator结构体记录当前开始位置和迭代次数
		val = f(r.window.Iterator(offset, count))
	}
	// 进不去if条件返回0
	return val
}
