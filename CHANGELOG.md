CHANGELOG
=========

v2.4.0 (02.09.2021)
-------------------

## 💔 Internal BC:

- 🔨 Pool, worker interfaces: payload now passed and returned by the pointer.

## 👀 New:

- ✏️ Long-awaited, reworked `Jobs` plugin with pluggable drivers. Now you can allocate/destroy pipelines in the runtime. Drivers included in the initial release: `RabbitMQ (0-9-1)`, `SQS v2`, `beanstalk`, `ephemeral` and local queue powered by the `boltdb`. [PR](https://github.com/spiral/roadrunner/pull/726)
-  ✏️Support for the IPv6 (`tcp|http(s)|empty [::]:port`, `tcp|http(s)|empty [::1]:port`, `tcp|http(s)|empty :// [0:0:0:0:0:0:0:1]:port`) for RPC, HTTP and other plugins. [RFC](https://datatracker.ietf.org/doc/html/rfc2732#section-2)
- ✏️ Support for the Docker images via GitHub packages.
- ✏️ Go 1.17 support for the all spiral packages.

## 🩹 Fixes:

- 🐛 Fix: add `debug` pool config key to the `.rr.yaml` configuration [reference](https://github.com/spiral/roadrunner-binary/issues/79).
- 🐛 Fix: fixed bug with goroutines waiting on the internal worker's container channel.
- 🐛 Fix: RR become unresponsive when new workers failed to re-allocate, [issue](https://github.com/spiral/roadrunner/issues/772).

## 📦 Packages:

- 📦 Update goridge to `v3.2.1`
- 📦 Update temporal to `v1.0.9`
- 📦 Update RR to `v2.4.0`
- 📦 Update endure to `v1.0.3`

## 📈 Summary:

- RR Milestone [2.4.0](https://github.com/spiral/roadrunner/milestone/29?closed=1)
- RR-Binary Milestone [2.4.0](https://github.com/spiral/roadrunner-binary/milestone/10?closed=1)

v2.3.2 (14.07.2021)
-------------------

## 🩹 Fixes:

- 🐛 Fix: Do not call the container's Stop method after the container stopped by an error.
- 🐛 Fix: Bug with ttl incorrectly handled by the worker [PR](https://github.com/spiral/roadrunner/pull/749)
- 🐛 Fix: Add `RR_BROADCAST_PATH` to the `websockets` plugin [PR](https://github.com/spiral/roadrunner/pull/749)

## 📈 Summary:

- RR Milestone [2.3.2](https://github.com/spiral/roadrunner/milestone/31?closed=1)

v2.3.1 (30.06.2021)
-------------------

- ✏️ Rework `broadcast` plugin. Add architecture diagrams to the `doc`
  folder. [PR](https://github.com/spiral/roadrunner/pull/732)
- ✏️ Add `Clear` method to the KV plugin RPC. [PR](https://github.com/spiral/roadrunner/pull/736)

## 🩹 Fixes:

- 🐛 Fix: Bug with channel deadlock when `exec_ttl` was used and TTL limit
  reached [PR](https://github.com/spiral/roadrunner/pull/738)
- 🐛 Fix: Bug with healthcheck endpoint when workers marked as invalid and stay is that state until next
  request [PR](https://github.com/spiral/roadrunner/pull/738)
- 🐛 Fix: Bugs with `boltdb` storage: [Boom](https://github.com/spiral/roadrunner/issues/717)
  , [Boom](https://github.com/spiral/roadrunner/issues/718), [Boom](https://github.com/spiral/roadrunner/issues/719)
- 🐛 Fix: Bug with incorrect Redis initialization and usage [Bug](https://github.com/spiral/roadrunner/issues/720)
- 🐛 Fix: Bug, Goridge duplicate error messages [Bug](https://github.com/spiral/goridge/issues/128)
- 🐛 Fix: Bug, incorrect request `origin` check [Bug](https://github.com/spiral/roadrunner/issues/727)

## 📦 Packages:

- 📦 Update goridge to `v3.1.4`
- 📦 Update temporal to `v1.0.8`

v2.3.0 (08.06.2021)
-------------------

## 👀 New:

- ✏️ Brand new `broadcast` plugin now has the name - `websockets` with broadcast capabilities. It can handle hundreds of
  thousands websocket connections very efficiently (~300k messages per second with 1k connected clients, in-memory bus
  on 2CPU cores and 1GB of RAM) [Issue](https://github.com/spiral/roadrunner/issues/513)
- ✏️ Protobuf binary messages for the `websockets` and `kv` RPC calls under the
  hood. [Issue](https://github.com/spiral/roadrunner/issues/711)
- ✏️ Json-schemas for the config file v1.0 (it also registered
  in [schemastore.org](https://github.com/SchemaStore/schemastore/pull/1614))
- ✏️ `latest` docker image tag supported now (but we strongly recommend using a versioned tag (like `0.2.3`) instead)
- ✏️ Add new option to the `http` config section: `internal_error_code` to override default (500) internal error
  code. [Issue](https://github.com/spiral/roadrunner/issues/659)
- ✏️ Expose HTTP plugin metrics (workers memory, requests count, requests duration)
  . [Issue](https://github.com/spiral/roadrunner/issues/489)
- ✏️ Scan `server.command` and find errors related to the wrong path to a `PHP` file, or `.ph`, `.sh`
  scripts. [Issue](https://github.com/spiral/roadrunner/issues/658)
- ✏️ Support file logger with log rotation [Wiki](https://en.wikipedia.org/wiki/Log_rotation)
  , [Issue](https://github.com/spiral/roadrunner/issues/545)

## 🩹 Fixes:

- 🐛 Fix: Bug with `informer.Workers` worked incorrectly: [Bug](https://github.com/spiral/roadrunner/issues/686)
- 🐛 Fix: Internal error messages will not be shown to the user (except HTTP status code). Error message will be in
  logs: [Bug](https://github.com/spiral/roadrunner/issues/659)
- 🐛 Fix: Error message will be properly shown in the log in case of `SoftJob`
  error:  [Bug](https://github.com/spiral/roadrunner/issues/691)
- 🐛 Fix: Wrong applied middlewares for the `fcgi` server leads to the
  NPE: [Bug](https://github.com/spiral/roadrunner/issues/701)

## 📦 Packages:

- 📦 Update goridge to `v3.1.0`

v2.2.1 (13.05.2021)
-------------------

## 🩹 Fixes:

- 🐛 Fix: revert static plugin. It stays as a separate plugin on the main route (`/`) and supports all the previously
  announced features.
- 🐛 Fix: remove `build` and other old targets from the Makefile.

---

v2.2.0 (11.05.2021)
-------------------

## 👀 New:

- ✏️ Reworked `static` plugin. Now, it does not affect the performance of the main route and persist on the separate
  file server (within the `http` plugin). Looong awaited feature: `Etag` (+ weak Etags) as well with the `If-Mach`
  , `If-None-Match`, `If-Range`, `Last-Modified`
  and `If-Modified-Since` tags supported. Static plugin has a bunch of new options such as: `allow`, `calculate_etag`
  , `weak` and `pattern`.
  ### Option `always` was deleted from the plugin.


- ✏️ Update `informer.List` implementation. Now it returns a list with the all available plugins in the runtime.

## 🩹 Fixes:

- 🐛 Fix: issue with wrong ordered middlewares (reverse). Now the order is correct.
- 🐛 Fix: issue when RR fails if a user sets `debug` mode with the `exec_ttl` supervisor option.
- 🐛 Fix: uniform log levels. Use everywhere the same levels (warn, error, debug, info, panic).

---

v2.1.1 (29.04.2021)
-------------------

## 🩹 Fixes:

- 🐛 Fix: issue with endure provided wrong logger interface implementation.

v2.1.0 (27.04.2021)
-------------------

## 👀 New:

- ✏️ Add support for `linux/arm64` platform for binaries in the RR releases.
- ✏️ New `service` plugin. Docs: [link](https://roadrunner.dev/docs/beep-beep-service)

## 🩹 Fixes:

- 🐛 Fix: logger didn't provide an anonymous log instance to a plugins w/o `Named` interface implemented.
- 🐛 Fix: http handler was without log listener after `rr reset`.

v2.0.4 (06.04.2021)
-------------------

## 👀 New:

- ✏️ Add support for `linux/arm64` platform for docker image (thanks @tarampampam).
- ✏️ Add dotenv file support (`.env` in working directory by default; file location can be changed using CLI
  flag `--dotenv` or `DOTENV_PATH` environment variable) (thanks @tarampampam).
- 📜 Add a new `raw` mode for the `logger` plugin to keep the stderr log message of the worker unmodified (logger
  severity level should be at least `INFO`).
- 🆕 Add Readiness probe check. The `status` plugin provides `/ready` endpoint which return the `204` HTTP code if there
  are no workers in the `Ready` state and `200 OK` status if there are at least 1 worker in the `Ready` state.
- 🆕 New option `unavailable_status_code` for the `status` plugin.

## 🩹 Fixes:

- 🐛 Fix: bug with the temporal worker which does not follow general graceful shutdown period.

## 📦 Updates:

- RR v2.0.4 - [Release](https://github.com/spiral/roadrunner/releases/tag/v2.0.4)
- RR-Temporal plugin v1.0.3 [Release](https://github.com/temporalio/roadrunner-temporal/releases/tag/v1.0.3)
- Endure v1.0.1 [Release](https://github.com/spiral/endure/releases/tag/v1.0.1)

v2.0.3 (29.03.2021)
-------------------

## 🩹 Fixes:

- 🐛 Fix: slow last response when reached `max_jobs` limit.

v2.0.2 (23.03.2021)
-------------------

- 🐛 Fix: Bug with required Root CA certificate for the SSL, now it's optional.
- 🆕 New: HTTP/FCGI/HTTPS internal logs instead of going to the raw stdout will be displayed in the RR logger at
  the `Info` log level.
- ⚡ New: Builds for the Mac with the M1 processor (arm64).
- 👷 Rework ServeHTTP handler logic. Use http.Error instead of writing code directly to the response writer. Other small
  improvements.

v2.0.1 (09.03.2021)
-------------------

- 🐛 Fix: incorrect PHP command validation
- 🐛 Fix: ldflags properly inject RR version
- ⬆️ Update: README, links to the go.pkg from v1 to v2
- 📦 Bump golang version in the Dockerfile and in the `go.mod` to 1.16
- 📦 Bump Endure container to v1.0.0.
- 📦 Bump Roadrunner-Temporal to v1.0.1 (release: ).

v2.0.0 (02.03.2021)
-------------------

- ✔️ Added shared server to create PHP worker pools instead of isolated worker pool in each individual plugin.
- 🧟 New plugin system with auto-recovery, easier plugin API.
- 📜 New `logger` plugin to configure logging for each plugin individually.
- 🔝 Up to 50% performance increase in HTTP workloads.
- ✔️ Added **[Temporal Workflow](https://temporal.io)** plugin to run distributed computations on scale.
- ✔️ Added `debug` flag to reload PHP worker ahead of request (emulates PHP-FPM behavior).
- ❌ Eliminated `limit` service, now each worker pool incluides `supervisor` configuration.
- 🆕 New resetter, informer plugins to perform hot reloads and observe loggers in a system.
- 💫 Exposed more HTTP plugin configuration options.
- 🆕 Headers, static and gzip services now located in HTTP config.
- 🆕 Ability to configure the middleware sequence.
- 💣 Faster Goridge protocol (eliminated 50% of syscalls).
- 💾 Added support for binary payloads for RPC (`msgpack`).
- 🆕 Server no longer stops when a PHP worker dies (attempts to restart).
- 💾 New RR binary server downloader.
- 💣 Echoing no longer breaks execution (yay!).
- 🆕 Migration to ZapLogger instead of Logrus.
- 💥 RR can no longer stuck when studding down with broken tasks in pipeline.
- 🧪 More tests, more static analysis.
- 💥 Created a new foundation for new KV, WebSocket, GRPC and Queue plugins.

v2.0.0-RC.3 (20.02.2021)
-------------------

- RR-Core update to v2.0.0-RC.3 version (release: [link](https://github.com/spiral/roadrunner/releases/tag/v2.0.0-RC.3))
- Temporal plugin update to v2.0.0-RC.2 version (
  release: [link](https://github.com/temporalio/roadrunner-temporal/releases/tag/v1.0.0-RC.2))

v2.0.0-RC.2 (11.02.2021)
-------------------

- RR-Core update to v2.0.0-RC.2 version
- Temporal plugin update to v2.0.0-RC.1 version
- Goridge update to v3.0.1 version
- Endure container update v1.0.0-RC.1 version
