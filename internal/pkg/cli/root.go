package cli

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
	endure "github.com/spiral/endure/pkg/container"
	"github.com/spiral/roadrunner-binary/v2/internal/pkg/cli/foo"
	"github.com/spiral/roadrunner-binary/v2/internal/pkg/container"
	dbg "github.com/spiral/roadrunner-binary/v2/internal/pkg/debug"
	"github.com/spiral/roadrunner-binary/v2/internal/pkg/meta"
	"github.com/spiral/roadrunner/v2/plugins/config"
)

// NewCommand creates root command.
func NewCommand(cmdName string) *cobra.Command {
	const (
		envDotenv = "DOTENV_PATH"
	)

	var (
		cfgFile  string   // path to the .rr.yaml
		workDir  string   // working directory
		dotenv   string   // path to the .env file
		debug    bool     // debug mode
		override []string // override config values
	)

	// next variables will be overwritten on pre run action (with configured objects)
	var (
		configPlugin    = &config.Viper{}
		endureContainer = &endure.Endure{}
	)

	cmd := &cobra.Command{
		Use:           cmdName,
		SilenceErrors: true,
		SilenceUsage:  true,
		Version:       fmt.Sprintf("%s (build time: %s, %s)", meta.Version(), meta.BuildTime(), runtime.Version()),
		PersistentPreRunE: func(*cobra.Command, []string) error {
			if cfgFile != "" {
				if absPath, err := filepath.Abs(cfgFile); err == nil {
					cfgFile = absPath // switch config path to the absolute

					// force working absPath related to config file
					if err = os.Chdir(filepath.Dir(absPath)); err != nil {
						return err
					}
				}
			}

			if workDir != "" {
				if err := os.Chdir(workDir); err != nil {
					return err
				}
			}

			if v, ok := os.LookupEnv(envDotenv); ok { // read path to the dotenv file from environment variable
				dotenv = v
			}

			if dotenv != "" {
				_ = godotenv.Load(dotenv) // error ignored because dotenv is optional feature
			}

			cfg := &config.Viper{Path: cfgFile, Prefix: "rr", Flags: override}

			containerCfg, err := container.NewConfig(cfgFile)
			if err != nil {
				return err
			}

			c, err := container.NewContainer(*containerCfg)
			if err != nil {
				return err
			}

			if err = c.Register(containerCfg); err != nil { // register config plugin
				return err
			}

			for i, plugins := 0, container.Plugins(); i < len(plugins); i++ { // register container plugins
				if err = c.Register(plugins[i]); err != nil {
					return err
				}
			}

			if debug {
				srv := dbg.NewServer()
				go func() { _ = srv.Start(":6061") }() // TODO implement graceful server stopping
			}

			// overwrite
			*configPlugin, *endureContainer = *cfg, *c //nolint:govet

			return nil
		},
	}

	cmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", ".rr.yaml", "config file")
	cmd.PersistentFlags().StringVarP(&workDir, "WorkDir", "w", "", "working directory") // TODO change to `workDir`?
	cmd.PersistentFlags().StringVarP(&dotenv, "dotenv", "", ".env", fmt.Sprintf("dotenv file [$%s]", envDotenv))
	cmd.PersistentFlags().BoolVarP(&debug, "debug", "d", false, "debug mode")

	cmd.PersistentFlags().StringArrayVarP(
		&override,
		"override",
		"o",
		nil,
		"override config value (dot.notation=value)",
	)

	cmd.AddCommand(foo.NewCommand(configPlugin, endureContainer))

	return cmd
}
