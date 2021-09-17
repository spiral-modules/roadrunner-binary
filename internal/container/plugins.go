package container

import (
	"github.com/spiral/roadrunner-plugins/v2/amqp"
	"github.com/spiral/roadrunner-plugins/v2/beanstalk"
	"github.com/spiral/roadrunner-plugins/v2/boltdb"
	"github.com/spiral/roadrunner-plugins/v2/broadcast"
	"github.com/spiral/roadrunner-plugins/v2/gzip"
	"github.com/spiral/roadrunner-plugins/v2/headers"
	httpPlugin "github.com/spiral/roadrunner-plugins/v2/http"
	"github.com/spiral/roadrunner-plugins/v2/informer"
	"github.com/spiral/roadrunner-plugins/v2/jobs"
	"github.com/spiral/roadrunner-plugins/v2/kv"
	"github.com/spiral/roadrunner-plugins/v2/logger"
	"github.com/spiral/roadrunner-plugins/v2/memcached"
	"github.com/spiral/roadrunner-plugins/v2/memory"
	"github.com/spiral/roadrunner-plugins/v2/metrics"
	"github.com/spiral/roadrunner-plugins/v2/redis"
	"github.com/spiral/roadrunner-plugins/v2/reload"
	"github.com/spiral/roadrunner-plugins/v2/resetter"
	rpcPlugin "github.com/spiral/roadrunner-plugins/v2/rpc"
	"github.com/spiral/roadrunner-plugins/v2/server"
	"github.com/spiral/roadrunner-plugins/v2/service"
	"github.com/spiral/roadrunner-plugins/v2/sqs"
	"github.com/spiral/roadrunner-plugins/v2/static"
	"github.com/spiral/roadrunner-plugins/v2/status"
	"github.com/spiral/roadrunner-plugins/v2/websockets"
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
		// =========

		// kv + ws + jobs plugin
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
