CHANGELOG
=========

UNRELEASED
----------

- Add support for `linux/arm64` platform for docker image
- Add dotenv file support

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
