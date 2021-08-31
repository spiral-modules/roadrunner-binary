package container

import (
	"github.com/spiral/roadrunner/v2/plugins/amqp"
	"github.com/spiral/roadrunner/v2/plugins/beanstalk"
	"github.com/spiral/roadrunner/v2/plugins/boltdb"
	"github.com/spiral/roadrunner/v2/plugins/broadcast"
	"github.com/spiral/roadrunner/v2/plugins/ephemeral"
	"github.com/spiral/roadrunner/v2/plugins/gzip"
	"github.com/spiral/roadrunner/v2/plugins/headers"
	httpPlugin "github.com/spiral/roadrunner/v2/plugins/http"
	"github.com/spiral/roadrunner/v2/plugins/informer"
	"github.com/spiral/roadrunner/v2/plugins/jobs"
	"github.com/spiral/roadrunner/v2/plugins/kv"
	"github.com/spiral/roadrunner/v2/plugins/logger"
	"github.com/spiral/roadrunner/v2/plugins/memcached"
	"github.com/spiral/roadrunner/v2/plugins/memory"
	"github.com/spiral/roadrunner/v2/plugins/metrics"
	"github.com/spiral/roadrunner/v2/plugins/redis"
	"github.com/spiral/roadrunner/v2/plugins/reload"
	"github.com/spiral/roadrunner/v2/plugins/resetter"
	rpcPlugin "github.com/spiral/roadrunner/v2/plugins/rpc"
	"github.com/spiral/roadrunner/v2/plugins/server"
	"github.com/spiral/roadrunner/v2/plugins/service"
	"github.com/spiral/roadrunner/v2/plugins/sqs"
	"github.com/spiral/roadrunner/v2/plugins/static"
	"github.com/spiral/roadrunner/v2/plugins/status"
	"github.com/spiral/roadrunner/v2/plugins/websockets"
	"github.com/temporalio/roadrunner-temporal/activity"
	temporalClient "github.com/temporalio/roadrunner-temporal/client"
	"github.com/temporalio/roadrunner-temporal/workflow"
)

// Plugins returns active plugins for the endure container. Feel free to add or remove any plugins.
func Plugins() []interface{} { //nolint:funlen
	return []interface{}{
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
		// service plugin
		&service.Plugin{},

		// ========= JOBS bundle
		&jobs.Plugin{},
		&amqp.Plugin{},
		&sqs.Plugin{},
		&beanstalk.Plugin{},
		&ephemeral.Plugin{},
		// =========

		// kv + ws plugin
		&memory.Plugin{},

		// KV + Jobs
		&boltdb.Plugin{},

		// broadcast via memory or redis
		// used in conjunction with Websockets, memory and redis plugins
		&broadcast.Plugin{},

		// ======== websockets broadcast bundle
		&websockets.Plugin{},
		&redis.Plugin{},
		// =========

		// ============== KV
		&kv.Plugin{},
		&memcached.Plugin{},
		// ==============

		// plugin to serve static files
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
	}
}
