package ecs

import (
	"container/heap"
	"time"
)

// run on next frame
type Dispatcher interface {
	Dispatch(func())
	Dispatchn(uint64, func())
	Tick(tickCnt uint64)
}

func NewDispatcher() Dispatcher {
	d := &implDispatcher{}
	d.tasks = make(pqTasks, 0, 10)
	heap.Init(&d.tasks)
	return d
}

type implDispatcher struct {
	tasks pqTasks
}

type task struct {
	job  func()
	tick uint64
	time time.Time
}

type pqTasks []*task

func (pq pqTasks) Len() int { return len(pq) }
func (pq pqTasks) Less(i, j int) bool {
	if pq[i].tick == pq[j].tick {
		return pq[i].time.Before(pq[j].time)
	}
	return pq[i].tick < pq[j].tick
}
func (pq pqTasks) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}
func (pq *pqTasks) Push(x interface{}) {
	*pq = append(*pq, x.(*task))
}
func (pq *pqTasks) Pop() interface{} {
	old := *pq
	n := len(old)
	top := old[n-1]
	*pq = old[:n-1]
	return top
}
func (pq pqTasks) Top() *task {
	return pq[0]
}

func (d *implDispatcher) Dispatch(f func()) {
	heap.Push(&d.tasks, &task{job: f, time: time.Now()})
}

func (d *implDispatcher) Dispatchn(n uint64, f func()) {
	heap.Push(&d.tasks, &task{job: f, tick: n, time: time.Now()})
}

func (d *implDispatcher) Tick(tickCnt uint64) {
	for d.tasks.Len() > 0 {
		if d.tasks.Top().tick > tickCnt {
			break
		}
		heap.Pop(&d.tasks).(*task).job()
	}
}
