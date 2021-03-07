package cli

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/spf13/cobra"
	"github.com/spiral/errors"
	"go.uber.org/multierr"
)

func init() {
	root.AddCommand(&cobra.Command{
		Use:   "serve",
		Short: "Start RoadRunner server",
		RunE:  handler,
	})
}

func handler(_ *cobra.Command, _ []string) error {
	const op = errors.Op("handle_serve_command")
	/*
		We need to have path to the config at the RegisterTarget stage
		But after cobra.Execute, because cobra fills up cli variables on this stage
	*/

	err := Container.Init()
	if err != nil {
		return errors.E(op, err)
	}

	errCh, err := Container.Serve()
	if err != nil {
		return errors.E(op, err)
	}

	// https://golang.org/pkg/os/signal/#Notify
	// should be of buffer size at least 1
	shutdownHandler := make(chan os.Signal, 2)
	signal.Notify(shutdownHandler, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)

	stop := make(chan struct{}, 2)

	go func() {
		// first catch - stop the container
		<-shutdownHandler
		// send signal to stop execution
		stop <- struct{}{}

		// after first hit we are waiting for the second
		// second catch - exit from the process
		<-shutdownHandler
		println("\n")
		println("exit forced")
		os.Exit(1)
	}()

	for {
		select {
		case e := <-errCh:
			err = multierr.Append(err, e.Error)
			log.Printf("error occurred: %v, plugin: %s", e.Error.Error(), e.VertexID)
			if !RetryOnFail {
				er := Container.Stop()
				if er != nil {
					err = multierr.Append(err, er)
					return errors.E(op, err)
				}
				return errors.E(op, err)
			}

		case <-stop:
			log.Printf("stop signal received, grace timeout is: %d seconds", uint64(GracePeriod.Seconds()))
			// stop the container after first signal
			err = Container.Stop()
			if err != nil {
				log.Printf("error occured during the stopping container: %v \n", err)
			}
			return nil
		}
	}
}
