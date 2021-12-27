package xgo

import (
	"golang.org/x/sync/errgroup"
)

// ParallelWithError ...
func ParallelWithError(fns ...func() error) func() error {
	return func() error {
		eg := errgroup.Group{}
		for _, fn := range fns {
			eg.Go(fn)
		}

		return eg.Wait()
	}
}
