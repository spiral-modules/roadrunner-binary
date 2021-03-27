package container

import endure "github.com/spiral/endure/pkg/container"

func NewContainer(cfg Config) (*endure.Endure, error) {
	endureOptions := []endure.Options{
		endure.SetLogLevel(cfg.LogLevel),
		endure.RetryOnFail(cfg.RetryOnFail),
		endure.GracefulShutdownTimeout(cfg.GracePeriod),
	}

	if cfg.PrintGraph {
		endureOptions = append(endureOptions, endure.Visualize(endure.StdOut, ""))
	}

	return endure.NewContainer(nil, endureOptions...)
}
