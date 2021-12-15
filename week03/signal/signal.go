package signal

import (
	"errors"
	"homework/week03/workgroup"
	"os"
	"os/signal"
)

// Signal creates function for canceling execution using os signal.
func Signal(sig ...os.Signal) workgroup.RunFunc {
	return func(stop <-chan struct{}) error {
		if len(sig) == 0 {
			sig = append(sig, os.Interrupt)
		}
		done := make(chan os.Signal, len(sig))
		defer close(done)

		signal.Notify(done, sig...)
		defer signal.Stop(done)

		select {
		case <-stop:
			return nil
		case <-done:
			return errors.New("exist by os signal")
		}
	}
}
