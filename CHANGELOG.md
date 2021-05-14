CHANGELOG
=========

UNRELEASED
----------

## 👀 New:

- ✏️ Json-schemas for the config file v2.0 (it also registered in [schemastore.org](https://github.com/SchemaStore/schemastore/pull/1614))
- ✏️ `latest` docker image tag is supported now (but we strongly recommend to use versioned tag (like `1.2.3`) instead)

v2.2.1 (13.05.2021)
-------------------

## 🩹 Fixes:

- 🐛 Fix: revert static plugin. It stays as a separate plugin on the main route (`/`) and supports all the previously announced features.
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
