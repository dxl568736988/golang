package workgroup

import "context"

type RunFunc func(<-chan struct{}) error

type Group struct {
	fns []RunFunc
}

func (g *Group) Add(fn RunFunc) {
	g.fns = append(g.fns, fn)
}

func (g *Group) Run() error {
	if len(g.fns) == 0 {
		return nil
	}

	stop := make(chan struct{})
	done := make(chan error, len(g.fns))

	for _, fn := range g.fns {
		go func(fn RunFunc) {
			done <- fn(stop)
		}(fn)
	}
	var err error
	for i := 0; i < cap(done); i++ {
		if err == nil {
			err = <-done
		} else {
			<-done
		}
		if i == 0 {
			close(stop)
		}
	}
	return err
}

// Context creates function for canceling execution using context.
func Context(ctx context.Context) RunFunc {
	return func(stop <-chan struct{}) error {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-stop:
			return nil
		}
	}
}