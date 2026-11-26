# Air-Clip progress

## 2026-11-26 22:24 WIB · Phase 0 discovery

**Status** — Prototype. Phase 0 is open. UJ-01, FR-01, SEC-01–04 and SEC-11 are active. No physical acceptance or performance measurements exist.

| Observation | Evidence |
| --- | --- |
| Repository | `C:/AIDAN/air-clip`, origin fetch and push `https://github.com/aidanismail/air-clip`, initial `main...origin/main` at `75cf629`. Tracked files `README.md`, `LICENSE`. Existing untracked `docs/PRD.md` was preserved. |
| PRD | Read all 417 lines of the sole local v1.0 copy, SHA-256 `7B1F1C5F19403F0406895466C4B2A5D5C9BE9C55ABA3B6FF15B4336D1B819FE8`. |
| Host | Windows build `26300.10075`, registry `26H2`, `CoreSingleLanguage`, reported product string `Windows 10 Home Single Language`. Registry product branding is unconfirmed. Host architecture x64. PowerShell Desktop 5.1.26100.10068. Device model and iOS build await owner confirmation. |
| Tools | Go 1.26.4 `windows/amd64`, `CGO_ENABLED=1`, Git 2.48.1.windows.1, RTK 0.45.0, GitHub CLI 2.78.0, GCC 6.3.0 target `mingw32`. `govulncheck` was not on PATH. The GCC target is 32-bit. |
| Permissions | Workspace file writes and read-only `.git` by default. `rtk proxy git switch -c phase/0-discovery-foundation` first failed, exit 128, unable to create `.git/refs/heads/phase`. The approved retry succeeded, exit 0. CIM OS query returned access denied, so host build fields came from the registry. |
| Baseline | `rtk proxy go build ./...` and `rtk proxy go test ./...` both exited 1 with `directory prefix . does not contain main module or its selected dependencies`. The Go command also reported denied creation of the roaming telemetry upload token. No source module or tests existed at baseline. |

**Changed files** — `AGENTS.md`, `docs/IMPLEMENTATION_PLAN.md`, `docs/DECISIONS.md`, `docs/PROGRESS.md`. No owner device configuration changes.

**Last successful command** — `rtk proxy git switch -c phase/0-discovery-foundation`, exit 0 on approved retry.

**Open decisions and owner gates** — Confirm iOS and Windows target builds, dedicated CA/fingerprint process, approved network and address, firewall, and separate permission for synthetic native clipboard use. Hosted CI and release publication decisions remain open.

**Next task** — Add a standard-library Go module, fake adapter, and TLS integration tests bound to IPv4 loopback. Run the plan's local checks and write owner-only probe instructions.

## 2026-11-26 22:33 WIB · Starter tree on main

**Status** — Starter tree verified. Phase 0 and all physical gates are open. The owner asked for the template tree on `main` before feature work. Local `main` and `phase/0-discovery-foundation` still point to `75cf629`. The scaffold is in the working tree on `main`, available for review before the first scaffold commit. Scope is structure, development tooling, and synthetic groundwork for FR-04–05, FR-08 and SEC-11. T09 and T18 have partial fake evidence only.

**Changed files** — `README.md`, `.gitignore`, `AGENTS.md`, `go.mod`, `cmd/air-clip/main.go`, `internal/clipboard/` and its `cliptest` adapter/test, tracked-folder placeholders under `internal/`, `scripts/check.ps1`, `testdata/README.md`, and the architecture, security, platform, planning, checkpoint and test docs under `docs/`. The owner-provided PRD hash matched discovery.

| Command or check | Result |
| --- | --- |
| `rtk proxy git switch main` | Exit 0, approved local metadata change. |
| `rtk proxy gofmt -w internal/clipboard` and `rtk proxy go test ./internal/clipboard/...` | Both exit 0. Fake snapshot ownership, cancellation and compare-and-write tests passed. |
| `rtk proxy powershell -NoProfile -File .\scripts\check.ps1` | Exit 0. Formatting, `go vet ./...`, `go test ./...`, `go mod verify` and native Windows amd64 `go build ./...` passed. The first successful run had roaming Go cache warnings. The rerun used `GOTOOLCHAIN=local`, `GOPROXY=off`, and temporary `APPDATA`, `GOMODCACHE` and `GOCACHE`. `go telemetry off` was scoped to that temporary `APPDATA` and exited 0. |
| `rtk proxy go test -count=25 ./internal/clipboard/cliptest` | Exit 0, synthetic tests repeated 25 times. |
| `rtk proxy go run ./cmd/air-clip` | Exit 0. Reports `Air-Clip is unconfigured. Clipboard sharing is unavailable.` |
| `rtk proxy go list -m all` | Exit 0 with `GOPROXY=off`, lists only Air-Clip. |
| `rtk proxy git diff --check` and working-tree content/path scans | Diff check exit 0, CRLF conversion notice for `README.md`. File-only `rg` scans of tracked, untracked and ignored files outside `.git` returned exit 1 with no matches for common credential patterns or runtime artifacts. Trailing-whitespace scan of scaffold files returned no matches. |
| `rtk proxy go test -race ./...` | Exit 1. Local `mingw32` compiler error `64-bit mode not compiled in`. Repeat with a validated Windows amd64 C compiler or private CI runner. |
| `rtk proxy govulncheck ./...` | Not run, executable unavailable on PATH. |
| `rtk proxy go version` with `GOTOOLCHAIN=go1.26.11` | Exit 1. Official Go module proxy was unreachable through the restricted shell network. The installed Go 1.26.4 is behind published security patches. Phase checks above used that installed version. |

**Last successful verification** — `rtk proxy go run ./cmd/air-clip`, exit 0. The build and synthetic tests did not require network access or native clipboard access.

**Open gates and owner actions** — Review the scaffold and authorize a local commit if desired. The optional offline Shortcuts steps in [iPhone feasibility](IOS_SHORTCUTS.md) can establish the actual iOS build and conversion labels. LAN identity, TLS trust, firewall access, native clipboard consent, real P0 tests, current-toolchain/security verification, and performance measurement remain open.

**Next task** — Once the initial scaffold is committed with owner authorization, advance `phase/0-discovery-foundation` to that baseline and add the verified loopback TLS health/typed-response tests. The separate owner-approved LAN probe follows after its trust and network prerequisites are known.
