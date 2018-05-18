package globals

// run on next frame
type Dispatcher interface {
	Dispatch(func())
	Tick()
}

type implDispatcher struct {
	tasks []func()
}

func NewDispatcher() Dispatcher {
	return &implDispatcher{tasks: make([]func(), 0, 10)}
}

func (d *implDispatcher) Dispatch(f func()) {
	d.tasks = append(d.tasks, f)
}

func (d *implDispatcher) Tick() {
	for _, f := range d.tasks {
		f()
	}
	d.tasks = d.tasks[:0]
}
