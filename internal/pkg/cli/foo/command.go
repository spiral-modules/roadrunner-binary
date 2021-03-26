package foo

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	endure "github.com/spiral/endure/pkg/container"
	"github.com/spiral/roadrunner/v2/plugins/config"
)

func NewCommand(cfgPlugin *config.Viper, container *endure.Endure) *cobra.Command {
	return &cobra.Command{
		Use:     "foo",
		Aliases: []string{"f"},
		Short:   "Foo command",
		RunE: func(*cobra.Command, []string) (err error) {
			_, err = fmt.Fprintf(os.Stdout, "\n%v\n%v\n", cfgPlugin, container)

			return
		},
	}
}
