package cli

import (
	"fmt"
	"log"
	"net/http/pprof"
	"net/rpc"
	"os"
	"path/filepath"
	"time"

	"github.com/spiral/errors"

	goridgeRpc "github.com/spiral/goridge/v3/pkg/rpc"
	"github.com/spiral/roadrunner/v2/plugins/gzip"
	"github.com/spiral/roadrunner/v2/plugins/headers"
	httpPlugin "github.com/spiral/roadrunner/v2/plugins/http"
	"github.com/spiral/roadrunner/v2/plugins/informer"
	"github.com/spiral/roadrunner/v2/plugins/logger"
	"github.com/spiral/roadrunner/v2/plugins/metrics"
	"github.com/spiral/roadrunner/v2/plugins/reload"
	"github.com/spiral/roadrunner/v2/plugins/resetter"
	rpcPlugin "github.com/spiral/roadrunner/v2/plugins/rpc"
	"github.com/spiral/roadrunner/v2/plugins/server"
	"github.com/spiral/roadrunner/v2/plugins/static"
	"github.com/spiral/roadrunner/v2/plugins/status"
	"github.com/temporalio/roadrunner-temporal/activity"
	temporalClient "github.com/temporalio/roadrunner-temporal/client"
	"github.com/temporalio/roadrunner-temporal/workflow"

	"github.com/spiral/roadrunner/v2/plugins/config"

	"net/http"

	"github.com/spf13/cobra"
	endure "github.com/spiral/endure/pkg/container"
)

const EndureKey string = "endure"

// EndureConfig represents container configuration
type EndureConfig struct {
	GracePeriod time.Duration `mapstructure:"grace_period"`
	PrintGraph  bool          `mapstructure:"print_graph"`
	RetryOnFail bool          `mapstructure:"retry_on_fail"` //TODO check for races, disabled at the moment
	LogLevel    string        `mapstructure:"log_level"`
}

var (
	// WorkDir is working directory
	WorkDir string
	// CfgFile is path to the .rr.yaml
	CfgFile string
	// Debug mode
	Debug bool
	// Container is the pointer to the Endure container
	Container   *endure.Endure
	RetryOnFail bool
	cfg         *config.Viper
	override    []string
	root        = &cobra.Command{
		Use:           "rr",
		SilenceErrors: true,
		SilenceUsage:  true,
		Version:       fmt.Sprintf(Version+" Build time: %s", BuildTime),
	}
)

func Execute() {
	if err := root.Execute(); err != nil {
		// exit with error, fatal invoke os.Exit(1)
		log.Fatal(err)
	}
}

func init() {
	root.PersistentFlags().StringVarP(&CfgFile, "config", "c", ".rr.yaml", "config file (default is .rr.yaml)")
	root.PersistentFlags().StringVarP(&WorkDir, "WorkDir", "w", "", "work directory")
	root.PersistentFlags().BoolVarP(&Debug, "debug", "d", false, "debug mode")

	root.PersistentFlags().StringArrayVarP(
		&override,
		"override",
		"o",
		nil,
		"override config value (dot.notation=value)",
	)

	cobra.OnInitialize(func() {
		if CfgFile != "" {
			if absPath, err := filepath.Abs(CfgFile); err == nil {
				CfgFile = absPath

				// force working absPath related to config file
				if err := os.Chdir(filepath.Dir(absPath)); err != nil {
					panic(err)
				}
			}
		}

		if WorkDir != "" {
			if err := os.Chdir(WorkDir); err != nil {
				panic(err)
			}
		}

		cfg = &config.Viper{}
		cfg.Path = CfgFile
		cfg.Prefix = "rr"
		// override flags if exist
		cfg.Flags = override

		endureCfg := initEndureConfig()
		if endureCfg == nil {
			panic("endure config should not be nil")
		}

		var lvl endure.Level

		switch endureCfg.LogLevel {
		case "debug":
			lvl = endure.DebugLevel
		case "info":
			lvl = endure.InfoLevel
		case "warning":
			lvl = endure.WarnLevel
		case "error":
			lvl = endure.ErrorLevel
		case "panic":
			lvl = endure.PanicLevel
		case "fatal":
			lvl = endure.FatalLevel
		default:
			panic(fmt.Sprintf("unknown log level, provided: %s, availabe: debug, info, warning, error, panic, fatal\n", endureCfg.LogLevel))
		}

		var err error
		if endureCfg.PrintGraph {
			Container, err = endure.NewContainer(nil, endure.SetLogLevel(lvl), endure.RetryOnFail(endureCfg.RetryOnFail), endure.SetStopTimeOut(endureCfg.GracePeriod), endure.Visualize(endure.StdOut, ""))
		} else {
			Container, err = endure.NewContainer(nil, endure.SetLogLevel(lvl), endure.RetryOnFail(endureCfg.RetryOnFail), endure.SetStopTimeOut(endureCfg.GracePeriod))
		}

		if err != nil {
			log.Fatal(err)
		}

		err = Container.RegisterAll(
			// logger plugin
			&logger.ZapLogger{},
			// metrics plugin
			&metrics.Plugin{},
			// http server plugin
			&httpPlugin.Plugin{},
			// reload plugin
			&reload.Plugin{},
			// informer plugin (./rr workers, ./rr workers -i)
			&informer.Plugin{},
			// resetter plugin (./rr reset)
			&resetter.Plugin{},
			// rpc plugin (workers, reset)
			&rpcPlugin.Plugin{},
			// server plugin (NewWorker, NewWorkerPool)
			&server.Plugin{},

			// static
			&static.Plugin{},
			// headers
			&headers.Plugin{},
			// checker
			&status.Plugin{},
			// gzip
			&gzip.Plugin{},

			// temporal plugins
			&activity.Plugin{},
			&workflow.Plugin{},
			&temporalClient.Plugin{},

			// register config
			cfg,
		)
		if err != nil {
			log.Fatal(err)
		}

		// if debug mode is on - run debug server
		if Debug {
			go runDebugServer()
		}
	})
}

func initEndureConfig() *EndureConfig {
	c := &config.Viper{
		Path:   CfgFile,
		Prefix: "rr",
	}
	err := c.Init()
	if err != nil {
		panic(err)
	}

	// init default config
	if !c.Has(EndureKey) {
		return &EndureConfig{
			GracePeriod: time.Second * 30,
			PrintGraph:  false,
			RetryOnFail: false,
			LogLevel:    "debug",
		}
	}

	e := &EndureConfig{}
	err = c.UnmarshalKey(EndureKey, e)
	if err != nil {
		panic(err)
	}

	if e.LogLevel == "" {
		e.LogLevel = "error"
	}

	if e.GracePeriod == 0 {
		e.GracePeriod = time.Second * 30
	}

	RetryOnFail = e.RetryOnFail

	return e
}

// RPCClient is using to make a requests to the ./rr reset, ./rr workers
func RPCClient() (*rpc.Client, error) {
	rpcConfig := &rpcPlugin.Config{}

	c := &config.Viper{
		Path:   CfgFile,
		Prefix: "rr",
	}
	err := c.Init()
	if err != nil {
		return nil, err
	}

	if !c.Has(rpcPlugin.PluginName) {
		return nil, errors.E("rpc service disabled")
	}

	err = c.UnmarshalKey(rpcPlugin.PluginName, rpcConfig)
	if err != nil {
		return nil, err
	}
	rpcConfig.InitDefaults()

	conn, err := rpcConfig.Dialer()
	if err != nil {
		return nil, err
	}

	return rpc.NewClientWithCodec(goridgeRpc.NewClientCodec(conn)), nil
}

// debug server
func runDebugServer() {
	mux := http.NewServeMux()
	mux.HandleFunc("/debug/pprof/", pprof.Index)
	mux.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
	mux.HandleFunc("/debug/pprof/profile", pprof.Profile)
	mux.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	mux.HandleFunc("/debug/pprof/trace", pprof.Trace)
	srv := http.Server{
		Addr:    ":6061",
		Handler: mux,
	}

	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
