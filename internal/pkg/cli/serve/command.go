package serve

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/spiral/roadrunner-binary/v2/internal/pkg/container"

	"github.com/spf13/cobra"
	"github.com/spiral/errors"
	"github.com/spiral/roadrunner/v2/plugins/config"
	"go.uber.org/multierr"
)

func NewCommand(cfgPlugin *config.Viper) *cobra.Command {
	return &cobra.Command{
		Use:   "serve",
		Short: "Start RoadRunner server",
		RunE: func(*cobra.Command, []string) error {
			const op = errors.Op("handle_serve_command")

			// create endure container config
			containerCfg, err := container.NewConfig(cfgPlugin.Path)
			if err != nil {
				return errors.E(op, err)
			}

			// create endure container
			endureContainer, err := container.NewContainer(*containerCfg)
			if err != nil {
				return errors.E(op, err)
			}

			// register config plugin
			if err = endureContainer.Register(cfgPlugin); err != nil {
				return errors.E(op, err)
			}

			// register another container plugins
			for i, plugins := 0, container.Plugins(); i < len(plugins); i++ {
				if err = endureContainer.Register(plugins[i]); err != nil {
					return errors.E(op, err)
				}
			}

			// init container and all services
			if err = endureContainer.Init(); err != nil {
				return errors.E(op, err)
			}

			// start serving the graph
			errCh, err := endureContainer.Serve()
			if err != nil {
				return errors.E(op, err)
			}

			oss, stop := make(chan os.Signal, 2), make(chan struct{}, 1)
			signal.Notify(oss, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)

			go func() {
				// first catch - stop the container
				<-oss
				// send signal to stop execution
				stop <- struct{}{}

				// after first hit we are waiting for the second
				// second catch - exit from the process
				<-oss
				fmt.Println("exit forced")
				os.Exit(1)
			}()

			for {
				select {
				case e := <-errCh:
					fmt.Printf("error occurred: %v, plugin: %s\n", e.Error, e.VertexID)

					if !containerCfg.RetryOnFail {
						if er := endureContainer.Stop(); er != nil {
							return errors.E(op, multierr.Append(e.Error, er))
						}

						return errors.E(op, e.Error)
					}

				case <-stop: // stop the container after first signal
					fmt.Printf("stop signal received, grace timeout is: %d seconds\n", uint64(containerCfg.GracePeriod.Seconds()))

					if err = endureContainer.Stop(); err != nil {
						fmt.Printf("error occurred during the stopping container: %v\n", err)
					}

					return nil
				}
			}
		},
	}
}
