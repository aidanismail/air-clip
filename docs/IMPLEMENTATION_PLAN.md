# Air-Clip implementation plan

This plan tracks the [v1 PRD](PRD.md). PRD phases are integration gates. Repository-local work can continue on independent tests while a manual gate is open. Follow [the checkpoint](PROGRESS.md) for actual results.

## SDLC and engineering flow

1. Start each slice from an agreed baseline. Name the PRD IDs and tests, identify trust boundaries and failure behavior, and record any design question that needs owner approval.
2. Write or update focused regression tests alongside implementation. Exercise success, denied requests, cancellation, and irreversible operations. Keep production code small and leave platform assumptions in the test matrix until observed.
3. Review the diff as a Go maintainer, Windows engineer, security reviewer, QA engineer, and docs maintainer. Inspect ownership, error paths, resource bounds, dependency changes, secret exposure, and recovery. Reconcile docs and checkpoint before handing off.
4. Present each change as a separately reviewable slice with requirement IDs, actual evidence and limitations, manual instructions where applicable, and a rollback path. Close a phase after its evidence and owner gates pass.

## Branch and change policy

Start with a scaffold on `main`, as requested. Once the scaffold is committed with owner authorization, integrate that commit into the existing `phase/0-discovery-foundation` branch before implementing the probe. Create subsequent branches after their prerequisites are integrated. Use `phase/1-secure-text`, `phase/2-images`, `phase/3-windows-operations`, and `phase/4-release-candidate` as milestone branches. Use short topic branches off the current reviewed phase baseline when a slice merits separate review. If a manual gate is pending, keep independent synthetic work in its own topic and label the prerequisite.

The first slice contains the module, entry point, component folders, planning docs, local checks, and the fake clipboard contract with tests. Follow with small commits for the TLS probe, then protocol, native text, image pipeline, Windows lifecycle, packaging, and evidence. Commit only with explicit authorization. A push, PR, merge, ruleset change, tag or publication is a separately authorized action. When private PRs are authorized, use recorded test evidence and owner review before integration. Keep commits independently understandable and avoid rewriting shared history.

## CI, delivery, and supply-chain security

| Stage | Planned execution and evidence |
| --- | --- |
| Local feedback | `gofmt -l`, `go vet ./...`, `go test ./...`, `go mod verify`, scope and secret scan of tracked and new files. Run targeted tests during coding and again from the reviewed tree. |
| Private CI, if enabled by owner | Windows amd64 vet, unit and TLS integration tests, module verification and build. Run the race suite on a supported hosted runner with a validated C toolchain. Add bounded fuzz smoke tests and a longer owner-triggered run for parsers. Capture pass/fail per job. |
| Continuous security | Patch the supported Go line and verify the selected toolchain. Review module source, maintainers, license and transitive dependencies. Pin modules and reviewed Actions by full commit SHA. Run pinned `govulncheck`, scan source/diff and candidate packaging for secrets, review findings against actual platform and package reachability. |
| CI trust | Plan `pull_request` and trusted-branch checks with read-only `GITHUB_TOKEN`, no runtime secrets, no untrusted checkouts in privileged workflows, no dependence on untrusted artifacts, and bounded job time. Use hosted CI purely for synthetic work. Store only sanitized test evidence. |
| Continuous delivery | Build a local Windows amd64 candidate from reviewed, identifiable source with a patched toolchain. Record Go version, build command, source revision, EXE version metadata, SHA-256, and support status. Validate install and rollback on the owner devices. Dist artifacts remain local and ignored. Publish only by separate authorization after sign-off. |

[GitHub Actions secure use reference](https://docs.github.com/en/actions/reference/security/secure-use) · [Go security and release history](https://go.dev/doc/devel/release)

## Agent workflow and skills

Use the checked-in PRD, plan, checkpoint, ADRs, and test matrix as durable agent context. Go and Win32 work use repository tools and primary documentation. Keep documentation updates in this checkout. Use any artifact or repository skill only when the task requires its workflow. Keep owner evidence sanitized and use dedicated tests for security invariants. Review one small diff at a time, update the checkpoint, and on interruption resume from the last verified command. If a decision affects scope, trust, platform settings, or a destructive action, pause that dependency and advance independent tests.

## Phase 0 · Discovery and feasibility

- [x] Inspect the checkout, remotes, branch, environment and PRD. Establish the baseline, create this plan, and start ADRs and traceability.
- [x] Prepare and verify the initial scaffold with an unconfigured entry point, fake adapter and local checks.
- [ ] Exercise a verified IPv4 loopback HTTPS health/typed-response probe against the synthetic clipboard. Cover T06, T07, T17 and T18 at the Phase 0 boundary.
- [ ] Confirm owner hardware/OS builds and run local Shortcuts conversion experiments. Resolve approval for a dedicated local CA, restricted disposable LAN probe, firewall and native fixtures before running them.
- [ ] Validate Shortcuts trust, typed JSON/Base64 requests/responses, true image paste and expected failures on the real devices. Investigate native delayed rendering and clipboard ownership prerequisites in parallel. Covers UJ-01, FR-01, SEC-01–04, SEC-11 and groundwork for T01, T03, T07, T09, T16 and T19.

**Exit gate** — Record actual device evidence or explicit blockers, owner approves trust design, probe remains HTTPS-only. A blocked device check leaves Phase 0 open.

## Phase 1 · Secure text

- [ ] Create strict V1 route, Host, Origin, header, JSON and text validation, auth verifier, timeouts, bounded admission, and safe errors. Test with an approved local TLS identity and synthetic adapter.
- [ ] Implement volatile ticket lifecycle and deduplication, serialized clipboard access, eligibility epochs and cancellation. Cover uncertain acknowledgement and irreversible failure. Prove pull/read-disable and zero autonomous operations with fake call counts.
- [ ] Implement Win32 text snapshot/write under one acquisition, valid HWND and handle ownership. Exercise newer copies, delayed rendering, busy retries, partial failures, UTF-16 and privacy markers. Document newline policy and validate both Shortcuts text paths.

**Exit gate** — Automated tests pass and both native text paths, including compare-and-write conflicts, pass on the actual devices. Maps to FR-03–06, FR-08, SEC-04–05, SEC-07–09, SEC-12 and T01–02, T06, T08–11, T13, T16–19.

## Phase 2 · PNG and Win32 raster

- [ ] Implement bounded canonical PNG validation and fuzz JSON/Base64/PNG/DIB metadata. Validate structure, dimensions, stride, palette, masks, orientation, alpha, trailing input and memory budgets.
- [ ] Stage native image allocations before mutation, publish verified representations, and deterministically select supported read formats. Exercise synthetic fixtures, including malformed native metadata.
- [ ] Verify iPhone screenshot/photo to Paint or Word and Windows screenshot/Paint to Notes. Record image type, dimensions, pixel/alpha behavior and permission handling.

**Exit gate** — Both real image paths paste as raster, rejection preserves clipboard before commit, and bounds are verified. Maps to FR-03–08, SEC-08–09 and T03–05, T09, T16, T18–19.

## Phase 3 · Windows operations and hardening

- [ ] Implement strict approved interface, Private profile and network identity, real peer/local-address checks, and fail-closed lifecycle. Validate lock, suspend, logoff, resume, adapter change and restart.
- [ ] Implement per-user configuration, protected credential verifier and leaf key storage, rotation, local pairing, opt-in pull, startup and tray controls. Test single instance and state transitions.
- [ ] Prove direct-address TLS with a matching IP SAN and reserved lease. Evaluate host A record discovery with conflict handling, refresh/withdrawal and standalone UDP scoping. Document secure fallback if mDNS is blocked.
- [ ] Provide reversible setup, firewall, renewal, uninstall and metadata-only diagnostics. Verify least privilege and owner-only execution. Finish CI/security validation and review.

**Exit gate** — Threat tests, scoped real firewall, disconnect/recovery, privacy and lifecycle pass. Maps to FR-01–02, FR-08–12, SEC-01–12 and T06–07, T10–19.

## Phase 4 · Local release candidate

- [ ] Audit FR, SEC, UJ and SM traceability. Execute T01–T20, record all outcomes, and collect SM-01–10 actual measurements.
- [ ] Verify clean packaging, platform build, version metadata and checksum. Record unsigned EXE behavior and prove reversible setup, upgrade and uninstall.
- [ ] Obtain owner sign-off following all mandatory gates. Record the approved revision, integrity checksum and unresolved non-release findings.

**Exit gate** — All P0 and security gates pass on owner hardware, no unresolved critical/high security gaps, and owner approves release. Until then the status is prototype or release candidate.
