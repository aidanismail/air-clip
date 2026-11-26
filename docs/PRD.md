# Air-Clip — Product Requirements Document

Start Date: 26 November 2026 (planning baseline)

Review Status: Proposed for implementation; core concept approved, detailed requirements pending owner review

Document Status: v1.0 planning specification | Private personal project

People: Project owner / sole user; Codex coding agent (implementation assistant)

Repository: https://github.com/aidanismail/air-clip

## Background

### What is Air-Clip?
Air-Clip is a local-first, on-demand clipboard bridge between an iPhone 13 and an ASUS VivoBook running Windows. It transfers a single plain-text item or raster image in either direction over a reachable local Wi-Fi network. It replaces sending messages to oneself, emailing screenshots, uploading images or manually retyping text. It is not a continuous clipboard mirroring service, file transfer tool, or public product.

The Windows side is a native Go desktop tray application running in the signed-in user's interactive session. The phone side uses two manually configured Apple Shortcuts, optionally assigned to double and triple Back Tap. The laptop hosts the HTTPS API; the phone always initiates network requests. A user copies, invokes a shortcut, and pastes. No cloud server, account, internet access, database, native iOS application, browser extension, Mac or App Store distribution is required for the intended architecture. The system must be validated on the actual devices before daily use.

### Problem statement and user persona
The sole target user alternates between mobile and desktop applications and wants to transfer snippets, links, code, screenshots and photos without switching to a third-party messaging app. The user values simplicity, privacy, low overhead, predictable behavior, and maintainability of a GitHub-hosted project. The target is a personal installation, not acquisition, monetization or multi-user scale.

### Impact
- Convenience: complete each transfer in copy → one intentional phone action → paste, without opening a desktop UI.
- Privacy: avoid third-party upload, tracking, clipboard history and unattended background clipboard reads. Transport encryption and authenticated local access are mandatory for release.
- Reliability: avoid silent clipboard corruption, duplicate writes, stale content and accidental replacement of newer Windows clipboard items during network transfers.
- Learning: demonstrate Go networking, Win32 integration, protocol design, threat modeling, automated testing and agent-assisted software engineering.

### Assumptions and boundaries
- iPhone 13 with an installed iOS version supporting Shortcuts and Back Tap; actual iOS build and permission behavior must be recorded in Phase 0.
- ASUS VivoBook running Windows 10 or 11, actual edition/build and CPU architecture to be verified. Primary build Windows amd64 if supported by the laptop; other architectures out of scope until needed.
- Both devices can reach each other on an explicitly approved private LAN; guest networks, captive portals, client isolation, corporate restrictions and VPN routes may prevent connection.
- Windows session is signed in, unlocked, app running, and network allowed. Sleep, logoff, shutdown and locked screen are intentionally unavailable for clipboard transfer.
- The owner can authorize one-time local firewall configuration and iOS certificate installation. These are manual owner actions, not silent Codex actions.
- Scope, proposed limits, and choices below are design decisions, not measured performance or already-tested functionality. No repository modifications are implied by this document.

## Success Metrics

### Main Metrics
| ID | Metric | Proposed target / measurement |
|---|---|---|
| SM-01 | Four core transfer paths | 100% of 10 consecutive trials per direction and content type on target devices; actual pasting in target apps. |
| SM-02 | Correctness | Exact Unicode text round trip except documented Windows newline representation; decoded image dimensions and pixel content match normalized PNG, subject to documented clipboard/app conversions. |
| SM-03 | Task friction | After setup: copy, invoke one shortcut, paste; no desktop interaction except optional approval and failure recovery. |
| SM-04 | Security | All mandatory threat-model tests pass; no plaintext LAN endpoint, unauthorized read/write, clipboard content logs, or untrusted certificate bypass. |
| SM-05 | Safe writes | Zero duplicate commits for a repeated ticket within its lifetime; zero newer Windows clipboard overwrites in tested conflict scenarios. |

### Secondary Metrics
| ID | Metric | Proposed target / measurement |
|---|---|---|
| SM-06 | End-to-end text latency | p95 ≤ 2 seconds for ≤ 64 KiB text on stable home Wi-Fi, excluding user gestures/prompts; record 20 samples each direction. |
| SM-07 | End-to-end image latency | p95 ≤ 5 seconds for ≤ 2 MiB PNG on stable LAN, same exclusions; record actual device/network conditions. |
| SM-08 | Idle footprint | Target ≤ 40 MiB private working set and < 1% average CPU over five idle minutes; no clipboard polling; measure via Task Manager/PowerShell and report deviations. |
| SM-09 | Recovery | After resume/reconnect, within 15 seconds of an eligible network becoming available or a documented manual Resume action, no stale queued writes. |
| SM-10 | Maintainability | Clean build, pinned dependencies, automated tests, setup/uninstall docs, checkpoint files and signed-off manual matrix. |

Metrics are local test measurements only; no analytics SDK or transmission of telemetry. Record timings without logging clipboard bodies, content hashes, credentials or personal URLs. Report actual measurements versus targets, never invent results. Security gates cannot be traded for performance targets.

## Summary Proposal

Air-Clip v1 uses a small Windows tray process, two iOS Shortcuts, a versioned HTTPS interface, a random device credential, a deliberately trusted LAN, and direct read/write of the native Windows clipboard on explicit requests. Use no clipboard monitor, persistence, background transfer, or broadcast. Start with an isolated loopback test using synthetic fixtures; graduate to authenticated HTTPS and real-device smoke tests before enabling daily operation. Implement mDNS only after direct-address HTTPS succeeds, with a documented fallback that retains certificate identity checks.

| Aspect | Universal Clipboard | Manual self-message | Air-Clip proposal |
|---|---|---|---|
| Platform | Apple ecosystem | Any compatible messaging service | iPhone 13 + Windows laptop |
| Interaction | OS-integrated | Open app, send, open on other device | Copy → shortcut → paste |
| Transport | Apple-managed | Often internet/service dependent | Local HTTPS; internet unnecessary |
| Privacy | Platform-managed | Service policies and message history | No app-side history or telemetry |
| Deployment | Built into supported Apple devices | Existing account/apps | Local EXE + two Shortcuts |
| Strategy | Automatic sync | General messaging | Intentional push/pull, one item at a time |

Conceptual alternatives, not a current feature audit or commercial competitor analysis.

### Product scope and prioritization
| Priority | Included / excluded | Requirements |
|---|---|---|
| P0 High | Core | Bidirectional UTF-8 plain text, single PNG raster; strict input checks; encrypted HTTPS; pairing/revocation; safe clipboard commit; user-driven Shortcuts; session/network controls; bounded resource use. |
| P0 High | Deployment | Tray status, pause/resume, exit, initial setup, single instance, startup opt-in, direct address and mDNS or documented secure fallback, troubleshooting and uninstall. |
| P1 Medium | Polish | Back Tap guidance, shortcut home-screen/widget placement, copy safe setup address, local metadata status, optional per-transfer PC approval or restricted read mode. |
| P2 Later | Optional | More platforms, Bluetooth BLE, native iOS client, additional image formats, improved pairing UX; new PRD required. |
| Out of scope | Deliberately excluded | Continuous sync/watchers, clipboard history, multi-device sharing, files/directories, audio/video, rich text, drag and drop, cloud/remote access, UPnP, background iOS polling, automatic keyboard injection, analytics, app store, commercial distribution. |

## UI/UX Design

Figma link: Not created; no Figma dependency for v1. Use these textual wireframes and acceptance criteria as the interaction source of truth.

### UX principles
Keep Windows unobtrusive. Never display clipboard text or image thumbnails in notifications or tray menus. Show only actual server readiness, not a fabricated persistent iPhone connection. Confirm success only after the destination clipboard operation has succeeded. Preserve normal copy and paste behavior. Prefer action-oriented errors and avoid jargon where possible. Permissions and prompts are OS-controlled; do not promise that Back Tap works from every app or locked state.

### Windows tray interface (status: exploration)
| Control / display | Behavior |
|---|---|
| Status | Ready / Paused / Unpaired / Locked / Network unavailable / Error; no persistent “connected” claim. |
| Network | Approved adapter, active hostname/port, masked or non-secret setup details; copy address locally. |
| Pair / reset | Locally open pairing UI or setup CLI; display credential once after owner confirmation; rotate invalidates old credential and tickets. |
| Pause / resume | Pause disables remote read/write and invalidates outstanding tickets; resume only when eligible network/session and credential exist. |
| Launch on startup | Explicit toggle; use current-user sign-in startup entry, not SYSTEM service; report actual state. |
| Last operation | Direction/type/result/time only in memory; no content, size optional, clear on exit. |
| Help / exit | Open local setup instructions; graceful shutdown closes listener, withdraws discovery, invalidates tickets. |

### iPhone interface (status: exploration)
Two shortcuts: “Air-Clip — Send to PC” and “Air-Clip — Get from PC”. First run uses an explicitly configured HTTPS base URL and token; authenticate every API action. Show an optional confirmation if desired, OS paste/local network permissions, and success or actionable failure notification. Use a type selector only if automatic detection cannot be verified without coercing arbitrary files or text. Back Tap double = Send, triple = Get, optional; home screen or Shortcuts widget is fallback. Keep token out of shortcut names, notifications, URLs and shared exports. Do not publish a credential-bearing shortcut; warn that Shortcuts configuration may sync/back up via Apple account settings.

### User Journey — Phase 0 / Setup and validation

#### UJ-01 — Install and establish trust | High | FR-01, SEC-01–04
Journey: Download/build EXE → run → choose approved network → configure firewall → configure certificate trust → pair → configure two shortcuts → test health.

Pre-condition: Owner has access to Windows, iPhone and approved Wi-Fi; no previous pairing required.

Story: As the owner, I want a one-time setup that authenticates both ends so I can safely reuse the bridge.

Acceptance criteria:
- Startup is blocked for remote requests until configuration, certificate and credentials are valid; no insecure fallback.
- Owner explicitly approves network and firewall scope, trust anchor and pairing secret installation; steps explain removal and recovery.
- iPhone verifies certificate identity and authenticates to health endpoint; unknown token or invalid cert fails, no clipboard action.
- Troubleshooting identifies wrong host, isolated Wi-Fi, port conflict, denied permission, incorrect time and trust failure without exposing token.

#### UJ-02 — Reconnect after reboot | High | FR-02, FR-09
Journey: Sign in → optional autostart → verify network → tray Ready → invoke shortcut.

Pre-condition: Valid existing pairing and approved network; startup enabled by owner.

Story: As the owner, I want the utility ready without terminal or IDE.

Acceptance criteria: One instance runs in interactive session; startup toggle accurately reflects state; app can resume after sleep only if network/session eligible; no queued transfer resumes; non-approved networks remain paused; no Windows admin rights needed at runtime.

### User Journey — Phase 1 / Text

#### UJ-03 — iPhone to Windows text | High | FR-03, FR-05
Journey: Copy text → Send shortcut → select/validate text → request push ticket → upload → success → Ctrl+V.

Pre-condition: Windows Ready, valid token, phone clipboard contains plain text.

Story: As the owner, I want to paste text copied on my phone directly into desktop applications.

Acceptance criteria: Valid UTF-8 text is pasted into Notepad/VS Code; emoji, non-Latin text, multiline code and whitespace preserved; reject invalid/empty/NUL-containing input without modifying destination; a changed Windows sequence returns conflict; repeated same ticket does not reapply the write; success only after commit.

#### UJ-04 — Windows to iPhone text | High | FR-04, FR-06
Journey: Ctrl+C → Get shortcut → authenticated pull → decode type → Copy to Clipboard → Paste.

Pre-condition: Current Windows clipboard has supported text and read mode allowed.

Story: As the owner, I want to paste my current Windows clipboard on the phone.

Acceptance criteria: Read live clipboard only on request; response validated and decoded before iOS clipboard write; paste into Notes; empty/unsupported or locked clipboard yields actionable error; no auto-retry that might fetch newer content after failure; no server writes during pull.

### User Journey — Phase 2 / Images

#### UJ-05 — iPhone to Windows image | High | FR-03, FR-05, FR-07
Journey: Copy photo/screenshot → Send → normalize to PNG → request ticket → upload → paste in Paint/Word.

Pre-condition: Single supported image; valid size; ready bridge.

Story: As the owner, I want to paste an image rather than download an image file.

Acceptance criteria: Single raster received as a pasteable image; correct dimensions/orientation and colors for reference fixtures; document alpha handling; enforce byte and pixel ceilings; unsupported/multiple images rejected; malformed or interrupted upload never reaches clipboard write.

#### UJ-06 — Windows to iPhone image | High | FR-04, FR-06, FR-07
Journey: Copy screenshot/image → Get → read native image → normalize PNG → decode → Copy to Clipboard → paste in Notes or supported messaging app.

Pre-condition: Windows clipboard exposes supported raster format; file paths/HTML alone do not qualify.

Story: As the owner, I want a copied screenshot available as an image on iOS.

Acceptance criteria: Snipping Tool and Paint fixtures work; PNG decoded into an image object rather than base64 text or a file reference; errors leave existing iOS clipboard unchanged where possible; app-specific paste limitations documented; no Photos save required.

### User Journey — Phase 3 / Safety and recovery

#### UJ-07 — Conflict, interruption or invalid input | High | FR-05, FR-08
Journey: Start transfer → network drops, busy clipboard or new copy occurs → clear failure/conflict → owner deliberately retries from beginning.

Acceptance criteria: Timeout, payload rejection, auth failure and pre-commit conflict cause no write; response lost after commit reports uncertain outcome and tells user to check destination; no automatic replay with a fresh ticket; transient clipboard lock has bounded retry; partial OS write failure explicitly reported, never falsely guaranteed atomic rollback.

#### UJ-08 — Pause, revoke, exit and uninstall | High | FR-09, FR-10
Journey: Tray Pause/Reset/Exit → remote access disabled → optionally resume/re-pair/uninstall.

Acceptance criteria: Pause and session lock prevent new clipboard operations; revocation invalidates prior tokens and tickets; in-flight operations recheck eligibility immediately before commit; uninstall removes only Air-Clip-owned startup/firewall/config artifacts with consent, explains removing iOS trust profile and shortcuts, and does not clear unrelated clipboard contents.

## Technical Design

### Architecture and stack
| Layer | Decision |
|---|---|
| Language | Go, supported stable release selected and recorded in go.mod after environment check; Windows native executable. |
| HTTP/TLS | Go net/http, crypto/tls, crypto/x509, crypto/rand, crypto/subtle; TLS 1.2 minimum, prefer 1.3. |
| Clipboard | Interface with native Windows adapter and fake adapter. Evaluate golang.design/x/clipboard for ordinary format operations; use x/sys/windows and Win32 where required for compare-and-write, format handling and privacy markers. Pin validated version. |
| Images | Standard image/png, image/color; bounded decode/encode. Native DIB conversion only through audited/tested adapter. |
| Discovery | Evaluate maintained Go mDNS responder that explicitly claims hostname and updates A records, not merely DNS-SD service advertisement. IPv4-only v1. |
| UI | Lightweight maintained Windows tray package or narrow native implementation; no Electron/webview unless explicitly approved. |
| Persistence | Per-user config and credential verifier/certificate material only; no database or payload storage. |
| Tooling | Go modules, Git, PowerShell, gofmt, go vet, go test, govulncheck, GitHub Actions optional/private; no Docker/WSL runtime. |

Data flow: iPhone Shortcuts → HTTPS over private LAN → authentication + network/session gates → API validation → bounded clipboard executor → Win32 clipboard. Reverse pull traverses same path. Tray controls gate the server; mDNS only locates the endpoint. No outbound connection from Windows to iPhone. No clipboard state replication.

Go is chosen for native distribution and straightforward networking/concurrency. Rust remains viable but adds engineering complexity; Python remains viable for prototypes but requires runtime packaging. These are tradeoffs, not performance measurements.

### FR — Functional requirements
| ID | Priority | Specification |
|---|---|---|
| FR-01 | High | Local first-run setup, explicit network approval, TLS bootstrap, token generation, shortcut configuration instructions and authenticated health check. |
| FR-02 | High | Interactive user process, single instance via per-user/session guard, tray icon, startup opt-in; no admin or session-0 service at runtime. |
| FR-03 | High | Phone push with two-step ticket and typed payload; validate and commit at most once per ticket. |
| FR-04 | High | Phone pull (only when owner enables read access) returns one supported item from a consistent current Windows clipboard snapshot; never updates Windows. |
| FR-05 | High | Check native clipboard sequence under the same OpenClipboard acquisition used for commit; reject changed sequence; bounded retries; no watchers. |
| FR-06 | High | Strict supported-format selection, no automatic opening/execution of links/files/scripts; response and shortcut validate before writing iOS. |
| FR-07 | High | UTF-8 plain text and PNG, normalize single image, enforce limits on both paths; preserve compatible pixel/text semantics. |
| FR-08 | High | Structured errors, conservative retry, duplicate-ticket outcome, resource bounds, cancellation and recovery. |
| FR-09 | High | Pause/lock/network-change invalidation and graceful shutdown; no stale queued transfer. |
| FR-10 | High | Rotate/revoke credentials locally, remove app-owned artifacts, documented backup/upgrade/reset. |
| FR-11 | High | Stable mDNS name and secure direct-address fallback; network change re-evaluated before listening. |
| FR-12 | Medium | Local metadata-only status, optional Back Tap shortcuts, optional per-pull local confirmation (if not implemented, document residual risk and provide disabled read mode). |

### API contract v1 — proposed normative baseline
Base URL: https://air-clip-vivobook-<random-suffix>.local:8765 (example placeholder; actual unique hostname generated at setup). Prefix /v1; IPv4 LAN listener bound only to approved address; no wildcard listener or public interface. Every endpoint requires exactly one Authorization: Bearer <base64url random token> header. Require expected Host including port (case-insensitive DNS name comparison, normalize safely), configured HTTPS host; reject unexpected Origin if present, no browser CORS support, reject cross-site preflight. Do not trust X-Forwarded-For or proxy headers. Require application/json for POST; reject unknown/duplicate fields, extra JSON values, unexpected query parameters, invalid methods and transfer encodings not explicitly supported. No redirects. All responses Cache-Control: no-store, X-Content-Type-Options: nosniff. Do not put secret/ticket in URL or log headers/bodies.

| Method and path | Input | Successful response / semantics |
|---|---|---|
| GET /v1/health | No body | 200 JSON: ok, protocol:1, ready:true. No device identity or content. If disabled, connection unavailable or 503. |
| POST /v1/push-ticket | {} | 201 JSON: ok:true, ticket:<random>, expires_in_seconds:60. Stores current sequence, credential epoch and session/network epoch, no payload. |
| POST /v1/clipboard | ticket, type, data_b64 | 200 JSON: ok:true, result:applied (or already_applied for same committed ticket); no content echoed. |
| GET /v1/clipboard | No body | 200 JSON: ok:true, type:text or image, data_b64. Live snapshot, no writes. |

Example push body (synthetic): {"ticket":"EXAMPLE_NOT_VALID","type":"text","data_b64":"SGVsbG8="}. All data_b64 uses standard padded RFC 4648 base64 without whitespace; not encryption. JSON envelope chosen for predictable Shortcuts parsing and errors at modest size overhead; validate actual iOS Base64 Encode/Decode actions, text conversion and image type conversion in Phase 0. If evidence shows envelope unreliable, propose a versioned raw MIME alternative and update PRD/tests before adopting it. No silent protocol drift.

### Protocol errors and retry policy
| HTTP | Error code | Client behavior |
|---|---|---|
| 400 | invalid_request / invalid_payload | Reject; fix input, no write. |
| 401 | unauthorized | Re-pair or check credential; do not retry automatically. |
| 403 | forbidden / read_disabled | Check Windows network/approval settings. |
| 404 / 405 | not_found / method_not_allowed | Check shortcut version; no endpoint enumeration details. |
| 409 | clipboard_changed / ticket_conflict | Existing destination preserved before commit; user initiates fresh transfer. |
| 410 | ticket_expired | Start new transfer deliberately, not automatic retry. |
| 413 | payload_too_large | Use smaller text/image; no write. |
| 415 / 422 | unsupported_media / no_supported_item | Copy supported text or one image; no coercion of files. |
| 429 | rate_limited | Show retry later, no retry loop. |
| 503 | unavailable / clipboard_busy / paused | Check tray, session, network; bounded local retry already attempted. |
| 500 | clipboard_failed | Report OS failure, check destination before retry. |

Error envelope: {"ok":false,"error":{"code":"clipboard_changed","message":"Windows clipboard changed. Copy again and retry."}}. Error text static, sanitized, no raw exception, secret, data or arbitrary request content. Shortcuts may throw on non-2xx; exact status handling must be tested. If detailed error cannot be surfaced, show a generic failure rather than continuing to Copy to Clipboard. Network timeout after sending can mean committed but acknowledgement lost; tell owner outcome uncertain. Repeating the same ticket safely is permitted only while ticket is known and same payload is retained; default shortcuts do not auto-retry. GET pulls never change Windows, but an explicit new pull may retrieve different content.

### Transfer limits and resource policy (proposed defaults)
| Parameter | Default |
|---|---|
| Text decoded size | 256 KiB UTF-8; reject empty and embedded NUL; preserve whitespace, no Unicode normalization. |
| PNG decoded wire bytes | 12 MiB compressed PNG; maximum width and height 8192 each, ≤ 20,000,000 pixels; checked arithmetic and decode bounds. |
| JSON body | 17 MiB max including base64 overhead, headers ≤ 16 KiB; strict per-type decoded limit. |
| HTTP timeouts | Header 3 s, read/write 30 s, idle 15 s; adjust only after testing large images; no unlimited requests. |
| Concurrency | Two in-flight API requests max per process, one clipboard executor; bounded connection admission and TLS handshake handling; reject overload. |
| Rate limits | Global + per-peer capped buckets, baseline 30 authenticated requests/min burst 6; unauthenticated attempts 10/min per peer plus bounded global limiter. |
| Push tickets | 60 s validity, max 16 active/retained records total, one pending per credential; new ticket invalidates prior pending ticket. |
| Clipboard busy retry | Up to 5 attempts, 25/50/100/200 ms backoff within request deadline; no spin; recheck state. |
| Shutdown | Cancel pending operations, bounded grace period ≤ 5 s; do not claim interruption after irreversible OS write begins. |

Avoid compressed HTTP bodies; accept identity encoding only. Cap actual body reads regardless of Content-Length or chunked framing; enforce outgoing response limits too. Validate PNG signature and DecodeConfig before full decode, reject overflow/oversized dimensions/truncated images; decode and re-encode canonical PNG (strip ancillary metadata when feasible) with bounded memory; consider rejecting trailing bytes. No temporary files, image history, disk spooling, data dumps or crash-body logging. Memory cleanup best effort only; Go GC, OS paging and clipboard owners prevent guaranteed zeroization.

### Windows clipboard integration and safe commit
- Clipboard adapter contract: ReadSnapshot(ctx) → typed bytes or typed error; Sequence(ctx) → native sequence; CompareAndWrite(ctx, expectedSequence, item) → applied/conflict/failure; all through a serialized executor. Test fake adapter separately; native adapter must be tested on Windows.
- Text: support CF_UNICODETEXT, convert UTF-16 safely with terminating NUL and checked lengths, reject unpaired surrogate/embedded NUL as appropriate. Document CRLF/LF normalization policy, compare normalized text without trimming. Do not interpret HTML/RTF or file paths as files. V1 deterministic rule: native PNG/DIB image takes precedence over plain text if both valid, but reject CF_HDROP-only and unsupported formats; document tradeoff for copied web content.
- Image: inspect registered PNG, CF_DIBV5/CF_DIB and supported CF_BITMAP conversions; never assume package output equals native support. Validate DIB headers, dimensions, stride, masks, compression, palette and allocation before reading; handle top-down/bottom-up rows, BGRA/RGBA, alpha, 24/32-bit layouts; reject unsupported formats rather than guessing. Publish native PNG plus compatible DIB/DIBV5 where verified; test transparency, screenshots and common app paste. Normalize phone JPEG/HEIC via Shortcuts conversion only if it provides image input; wire PNG only, no arbitrary image file ingestion.
- Win32 ownership: use valid HWND in interactive session for OpenClipboard/EmptyClipboard; pin thread if required by chosen message loop; close on every path; copy data before releasing native handles; allocate and prepare all required formats before EmptyClipboard; use movable global memory, transfer ownership only on successful SetClipboardData, free on failure; no delayed rendering in v1. Never call EmptyClipboard before authentication, full validation and final sequence check while clipboard held. A library API that reads sequence and then separately acquires clipboard for Write cannot satisfy atomic compare-and-write; implement minimal native transaction if needed. Win32 empty+set is not fully transactional; a native failure after emptying may lose prior content or leave partial formats. Return error, best effort recover only if safe without overwriting newer data; do not promise rollback. Document this residual risk.
- Clipboard privacy: where verified, write Windows registered exclusion formats for history/cloud (ExcludeClipboardContentFromMonitorProcessing, CanIncludeInClipboardHistory, CanUploadToCloudClipboard) correctly; verify on target OS; warn that Air-Clip cannot control other apps/OS history/backups and advise disabling Windows clipboard history/cloud sync if sensitive. iOS clipboard may also be available to OS sync mechanisms. Do not silently alter OS settings.
- Pull: read under one native open and copy to owned memory; handle delayed rendering/sequence changes safely with timeout or fail; encode outside native lock. Recheck session/network/auth epoch before response. If OS calls cannot be cancelled, isolate/bound execution and fail closed without spawning unlimited goroutines.

### Sync loops, duplicates and ordering
Air-Clip has no automatic clipboard watcher, no continuous polling and no echo propagation. Push writes once; pull reads only; an explicit new user action is not a loop. Do not implement a watcher merely because Gemini suggested hash suppression. If future automatic sync is separately approved, require origin device IDs, monotonically tracked event IDs, sequence-aware suppression, bounded seen-ID TTL, and no echo on remote writes; hashes alone are insufficient after format conversion or intentional recopy.

Ticket state machine: ISSUED → PROCESSING → APPLIED / FAILED / CONFLICT / EXPIRED. Each ticket cryptographically random (at least 128 bits), bound to credential/session/network epoch, expected clipboard sequence, and expiry; record a keyed or SHA-256 digest of canonical type+decoded bytes in volatile memory solely for duplicate detection, never log/persist. Compare exact same payload on repeat; different body with same ticket returns conflict; simultaneous duplicate returns busy/conflict until terminal result. Validate and stage before commit; under serialized state lock and native clipboard acquisition, recheck epoch, expiry and sequence immediately before mutation. Store outcome before returning success; retain only bounded digest/outcome until ticket expiry. On partial native failure mark terminal failed and never automatically reapply. Process restart loses ticket state: reject old tickets, no claims of durable exactly-once. Revocation/pause/network/lock invalidate tickets including pending. No system-wide ordering guarantee for concurrent phone-side copy while response is in flight; Shortcuts cannot offer native atomic compare-and-write on iOS. Show user warning/optional confirm before replacing iPhone clipboard; never pretend iOS destination conflict is solved. A deliberate fresh transfer of identical data is allowed.

### SEC — Security and privacy requirements (mandatory release gates)
Threat model: an untrusted LAN peer can scan, spoof mDNS, intercept or replay traffic, send malformed requests, flood connections, or attempt to read/inject clipboard content; a lost phone or copied shortcut can expose its credential; an unrelated Windows process can race or change the clipboard. A compromised OS, administrator or malicious process in the same user session can bypass application-level protections; Air-Clip does not claim protection against those actors. An authorized phone is trusted to initiate clipboard operations, not to prove which human tapped it.

| ID | Required control and verification |
|---|---|
| SEC-01 | TLS only on LAN. Minimum TLS 1.2, prefer 1.3; use Go defaults for secure cipher selection. Never disable verification, accept any certificate, downgrade, or provide a plaintext redirect. Test wrong host, untrusted CA, expiry and invalid chain. |
| SEC-02 | Generate a unique local CA and server leaf at setup using crypto/rand and standard crypto/x509. CA with critical basic constraints and appropriate key usage; leaf with serverAuth and exact DNS SAN, plus explicit IP SAN only for approved direct-address fallback. Use modern keys/signatures, conservative validity (leaf ≤ 90 days, CA ≤ 365 days proposed); verify iOS current trust requirements, dates and SAN. Keep CA private key offline in protected owner-controlled storage or discard after issuance and document re-pairing; never export private keys to phone. |
| SEC-03 | Transfer only CA public certificate via owner-controlled channel (e.g. USB/offline transfer or local file share independently verified); compare SHA-256 fingerprint on trusted devices before installing. iOS Settings profile installation and Certificate Trust Settings full trust are manual actions; explain broad system trust, never instruct accepting mismatched fingerprints. Document renewal before expiry and removing CA profile; no unsafe fallback if trust fails. |
| SEC-04 | Pair locally only. Generate ≥ 256-bit random bearer credential, base64url without padding, display once via local setup; persist SHA-256 verifier rather than plaintext, constant-time compare fixed-length digest; no public pairing API/default credentials. Store verifier and private keys under per-user LocalAppData with restrictive ACL, optionally DPAPI user scope for key protection; never put token in argv, shell history, URL or Git. iPhone shortcut contains token; it is not hardware-bound or MFA. |
| SEC-05 | Authenticate every endpoint before reading body, sequence, or clipboard. Validate exact allowed Host, no CORS, no reverse proxy trust, deny unexpected Origin, only exact routes/methods, strict headers and JSON, response no-store. No HTML rendering, auto-paste keystrokes, shell execution, URL fetching, file paths or deserialization into executable objects. |
| SEC-06 | Bind only to approved IPv4 interface/address. Windows Private profile and explicit owner-approved network identity, not profile alone; check actual peer subnet and local listener address, ignore proxy headers; reject other interfaces, VPN/virtual adapters and unknown/public networks by default. Firewall inbound TCP scoped to executable, port, Private profile and approved local subnet/address; if using mDNS, separately restrict UDP 5353 to approved interface/subnet as OS permits. Never disable firewall, open router ports, enable UPnP, or request unnecessary admin rights. |
| SEC-07 | On pause, lock, suspend, adapter change, logoff or credential rotation, stop accepting requests, cancel queued work, increment epoch and invalidate tickets; recheck just before clipboard access/commit and before sending pull. Network identification can be spoofed; TLS + auth remain mandatory. Session event detection must be tested; unknown state fails closed. |
| SEC-08 | Validate type, byte/pixel ceilings, base64, UTF-8, PNG/DIB metadata and complete body; cap concurrency, total connections, timeouts, rate limiter state and memory. Invalid input must not reach clipboard. Avoid decompression bombs and integer overflow. |
| SEC-09 | No clipboard history, content logs, disk payloads, telemetry, crash uploads or analytics. Safe diagnostic codes only; no token/ticket/headers/content/hash/IP unless owner explicitly needs sanitized network diagnostics. Avoid logging raw errors or request URLs. OS clipboard history/cloud sync, iCloud Shortcuts sync, paging and backups are outside app control; document limitations. |
| SEC-10 | Local credential revocation and rotation invalidate sessions/tickets immediately; removal of lost-phone shortcuts is insufficient by itself. CA/key compromise requires replacement, removing old trust, rotation and firewall review. No remote admin endpoint. |
| SEC-11 | Restrict dependencies and CI: pin go.mod/go.sum, review maintainers/licenses/transitive code, run govulncheck and secret scan, no arbitrary scripts or untrusted CI artifacts, no release with known applicable critical/high vulnerabilities without owner-reviewed remediation. Do not run test fixtures on real clipboard without explicit consent. |
| SEC-12 | Pull access is explicitly enabled by owner; independent read-disable toggle and optional per-request approval for stronger confidentiality. A stolen token can read live Windows content while enabled and Ready. Warn prominently; pause during sensitive work. No password detection claim, no automatic content filtering that silently alters text. |

### Network discovery, TLS lifecycle and setup
Use one generated stable name such as air-clip-vivobook-a1b2.local (illustrative). Advertise hostname A record on approved adapter, probe for conflicts, refresh on DHCP/address changes, withdraw on pause/network loss. Do not advertise tokens, content or owner identity. DNS-SD service advertisement alone does not guarantee hostname resolution in Shortcuts. On name conflict fail safely and guide reconfiguration; do not silently change TLS hostname. Direct-IP fallback requires stable router DHCP reservation and certificate SAN for that exact IP; if IP changes, update address and reissue cert before use. Do not rely on mDNS to authenticate. Initial validation can use loopback with fake clipboard; real LAN test must use HTTPS with verified trust. Validate captive portal and client isolation behavior on the actual iPhone and LAN. If mDNS is blocked, secure direct-IP fallback is acceptable with documented loss of portability.

Certificate renewal: notify locally before expiry (proposed 14 days), document leaf reissue using protected CA or regenerate trust anchor with owner involvement if key discarded. No remote distribution of CA key. No silent rotation of trust roots or bypass on time skew. TLS certificates and trust setup may be the largest manual usability cost; Phase 0 must prove feasibility in Shortcuts before further investment.

### Proposed repository and component boundaries
```
air-clip/
  AGENTS.md                 # persistent Codex rules and release gates
  README.md                 # overview, build, install, troubleshooting
  go.mod / go.sum
  cmd/air-clip/             # app entrypoint, setup commands
  internal/api/             # versioned handlers, request/response models
  internal/auth/            # verifier, tickets, rate limits
  internal/clipboard/       # interface, fake, Windows native adapter
  internal/codec/           # UTF-8, PNG, native raster validation
  internal/network/         # approved interfaces, listener, mDNS
  internal/security/        # TLS bootstrap, protected configuration
  internal/lifecycle/       # session events, epoch, pause and shutdown
  internal/tray/            # controls and status; no content rendering
  scripts/                 # build, explicit setup/uninstall (PowerShell)
  docs/PRD.md               # canonical specification
  docs/ARCHITECTURE.md      # decisions and diagrams
  docs/SECURITY.md          # threat model and trust/revocation
  docs/IOS_SHORTCUTS.md     # action-by-action manual setup
  docs/WINDOWS_SETUP.md     # firewall, trust, startup, uninstall
  docs/TEST_PLAN.md         # automated + owner-only manual checks
  docs/IMPLEMENTATION_PLAN.md
  docs/PROGRESS.md          # resumable task checkpoints and evidence
  docs/DECISIONS.md         # dated ADRs and spec deviations
  testdata/                 # synthetic text and PNG fixtures only
  .gitignore
  .github/workflows/        # optional, private repository CI
```
Existing repository contents must be inspected by Codex first; reconcile layout without deleting existing work. Runtime files go to %LOCALAPPDATA%\\Air-Clip, not the repo. Document config version, atomic config updates, strict validation, owner-only ACL, backup/restore without exposing credentials, path canonicalization and no following untrusted reparse points during privileged operations. Store only approved interfaces/identity, hostname, port, flags, verifier, certificate paths and schema version. Do not store token, clipboard content, history or pending tickets. Ignore runtime configuration, keys, logs, builds and generated secret-bearing shortcuts; never ignore go.sum. Do not embed private keys or live configuration in EXE.

### iOS shortcut build specification (manual owner validation)
Send: Get Clipboard once → reject unsupported/empty/multiple input (explicit text/image selection if necessary) → text branch encode exact text to UTF-8/base64; image branch convert one image to PNG then base64 without line breaks → POST /v1/push-ticket with bearer token and empty JSON object → validate ok/ticket/expiry → POST /v1/clipboard with ticket/type/data_b64 → inspect success result, notify success only on applied/already_applied; otherwise stop and show safe failure or uncertain outcome. Ticket is obtained after local normalization but before upload; the destination check protects changes from issuance onward, not earlier. No automatic resubmission with new ticket. Source is a snapshot; a newer phone clipboard does not change this payload.

Get: GET /v1/clipboard with bearer → validate response schema, success, type and size (as practical in Shortcuts) → base64 decode → text branch convert to text, image branch explicitly get image object from decoded data (not a file URL or base64 string) → optional confirmation before replacing clipboard → Copy to Clipboard (local-only option if available and verified) → success notification. If request action throws or returns malformed/unknown result, stop before Copy to Clipboard. Never clear clipboard on error. Avoid clipboard previews, Quick Look of secrets, automatic external links or share sheet exporting token. Document exact action names/parameters observed on owner's iOS version; do not invent importable .shortcut files or claim testing on iOS simulator. Optional Double/Triple Back Tap under Settings > Accessibility > Touch > Back Tap; alternate widget/home screen/manual Shortcuts launch.

### Non-functional requirements and quality gates
- Performance: use Success Metrics targets; measure release-mode binary on the actual VivoBook, Wi-Fi and iPhone. No artificial benchmarks or blanket promises of native transfer speed.
- Accessibility: tray labels/status not color-only; clear error and confirmation messages; keyboard-accessible setup controls; no Back Tap requirement.
- Reliability: deterministic state machine (Unconfigured, Paused, Ready, Locked, NetworkUnavailable, Error, Exiting), serialized clipboard executor, startup retry only for eligible transient failures with bounded backoff; fail closed on corrupted config, unknown session/network, missing certificate or port conflict. No fallback listener on arbitrary port.
- Maintainability: idiomatic Go, small interfaces, build tags for Windows, injectable clock/random/clipboard for testing, explicit cancellation and error classes. Prefer standard library, no unnecessary UI framework or microservices.
- Distribution: native Windows release build (target architecture verified), version metadata, SHA-256 checksum, build instructions, documented unsigned EXE/SmartScreen behavior; never disable protections or bypass warnings. Optional signing is outside v1 if unavailable. Install per user in stable directory; firewall path must match actual EXE; updates verify checksum, stop old process, replace only app files and preserve compatible config, recheck firewall. Uninstall does not alter unrelated registry/firewall/trust artifacts. No automatic updates or remote download execution.

### Test plan and acceptance evidence
All automated fixtures synthetic. Native clipboard tests require a dedicated opt-in test mode and owner consent; hosted CI generally cannot prove interactive clipboard or Shortcuts. Mark every test Passed/Failed/Not run, attach safe command output and evidence. A Windows cross-build is compilation evidence only. Test using actual iOS/Windows versions and representative apps, record limitations without copying private content.

| ID | Scenario | Expected / evidence |
|---|---|---|
| T01 | Text in both directions, emoji, Indonesian/non-Latin, multiline, whitespace, long payload | Exact supported round trip; documented newline handling. |
| T02 | Text invalid UTF-8, NUL, empty, limit+1 | Rejected; destination unchanged before commit. |
| T03 | Image both directions, screenshot, photo, alpha, odd stride, top/bottom-down | Correct native paste, dimensions and documented normalized pixel behavior. |
| T04 | PNG malformed, truncated, oversized dimensions, overflow, huge base64, fake MIME, trailing JSON | Rejected before native allocation/commit; bounded memory. |
| T05 | Unsupported file/HTML/RTF/multiple image; mixed representations | Deterministic selection or safe rejection, no file access. |
| T06 | Missing/wrong/old/duplicate auth header; wrong Host/Origin; unknown method/query | No clipboard access; safe error, no CORS or redirects. |
| T07 | Untrusted/wrong-host/expired cert and spoofed mDNS | iOS rejects; no HTTP fallback; no secret exposed. |
| T08 | Valid ticket repeated; changed payload same ticket; simultaneous duplicate | At most one commit; terminal result or conflict. |
| T09 | Copy newer item after ticket; locked clipboard; delayed rendering; failed SetClipboardData | Conflict preserves newer item; bounded retry; partial write risk reported honestly. |
| T10 | Network interrupted before commit and after commit before response | No precommit write; uncertain outcome after commit; no silent retry. |
| T11 | Pause/revoke/lock/sleep/logoff/network change during upload/pull | No new access; pending work invalidated; recheck at boundary; no queued resume. |
| T12 | Guest Wi-Fi/AP isolation, public profile, unknown adapter, VPN, DHCP change, mDNS collision | Denied or safe reconnect/fallback; correct firewall scope; no unintended exposure. |
| T13 | Slow headers/body, concurrent flood, invalid auth flood, oversized chunked body | Bounded connections, limiter state, memory, time; healthy recovery. |
| T14 | Restart with expired tickets, corrupted config, port occupied, duplicate EXE | Fail closed, no duplicated listener, explicit error. |
| T15 | Fresh install, startup toggle, reboot, upgrade, uninstall, CA removal, credential rotation | Owner-controlled and reversible; no unrelated modifications. |
| T16 | iOS permission denied, Back Tap disabled, malformed/non-2xx response, phone new copy during pull | No overwrite on detected error; disclose iOS race; alternative shortcut launch. |
| T17 | Privacy inspection, Git secret scan, log/response grep using synthetic sentinel, no disk payload | No sensitive content in app-owned outputs or tracked files. |
| T18 | Repeat transfer, idle period, rapid manual alternating pushes/pulls | No autonomous reads/writes or echoes; intentional identical items allowed. |
| T19 | Windows clipboard history/cloud exclusion and iOS local-only if available | Verify supported markers/options; document residual OS behavior. |
| T20 | SM-01 through SM-10 measurement | Record target vs observed; no release claim before mandatory gates pass. |

Automated: unit tests for strict parsers, tickets/epochs with fake clock, fake clipboard call counts, auth, rate/concurrency limits, malformed input, PNG limits; httptest TLS integration and end-to-end fake adapter; fuzz JSON/base64/PNG/DIB metadata parsers using bounded fixtures; race detector on supported host, go vet, gofmt, go test ./..., govulncheck ./..., Windows build. Native tests validate Win32 ownership and compare-and-write, clipboard busy, formats and privacy markers. For blocked tools or OS-specific tests, record exact blocker and command for owner. Never equate a mock to physical validation.

## Risks & Concerns

| Risk | Impact | Mitigation / plan | Residual / gate |
|---|---|---|---|
| iOS clipboard/API or trust behavior differs | Core workflow fails | Phase 0 Shortcuts HTTPS, typed text/image roundtrip on actual iPhone; no native app fallback without approval | Critical; stop before claiming feasibility |
| TLS CA installed too broadly or key leaked | Other network trust compromised | Unique CA, verify fingerprint, protect/discard CA key, short leaf validity, removal instructions | User accepts trust scope; no insecure fallback |
| Lost phone or exported shortcut | Live clipboard read/injection | Strong token, local revoke, read-disable and pause; no token in Git | Bearer credential is not hardware-bound |
| Windows clipboard format/ownership errors | Corruption or partial write | Native adapter, validate and allocate before empty, test common formats | Win32 write not transactional after empty |
| Race with new clipboard contents | Unexpected overwrite | Ticket + native sequence under lock; no stale retry; optional confirmation on iPhone | iOS cannot atomically compare-and-set; user aware |
| Network identity/spoofed discovery | Wrong endpoint or exposure | Approved adapter, firewall, certificate name verification, auth | LAN availability varies |
| Malformed images/slow requests | CPU/RAM exhaustion | Bounded parsing, rate/concurrency/time caps, fuzz tests | Validate native format handling |
| Clipboard history outside app | Sensitive data retained by OS | Exclusion flags where verified, owner guidance | Not a global privacy guarantee |
| Agent scope creep / interrupted session | Incomplete or unsafe project | PRD + AGENTS + progress checkpoints, owner gates, small changes and test evidence | Never assert untested completion |
| Unsigned EXE, OS policy or security software | Installation blocked | Source build, checksum, explain signing/policy | No bypass; owner decides |

## Timeline

Sequence is milestone-based, not a deadline or promised duration. P0 security gates cannot be deferred to a post-release cleanup. Code can advance with mocks while manual device gates are pending, but no daily-use release can be declared until they pass.

| Phase | Deliverables | Exit gate |
|---|---|---|
| 0 — Discovery & feasibility | Inspect repository/tools/OS; verify library APIs; isolated fake adapter; HTTPS certificate/Shortcuts probe; text/image type probe; decisions and plan | Record manual outcomes or blocked items; owner approves trust approach; no real LAN plaintext. |
| 1 — Secure text foundation | Strict API, TLS/auth, ticket state machine, fake integration tests, native Win32 text adapter | Unit/integration tests pass; native compare-and-write proof; text two directions owner-tested. |
| 2 — Images | PNG normalization, safe DIB handling, limits, native image tests, iOS image shortcut | Both image directions pasted as actual images, invalid inputs safe. |
| 3 — Operational hardening | Network allowlist, mDNS or secure fallback, session events, tray, pause/revoke, startup, installer/uninstall | Threat tests, sleep/lock/reconnect, privacy and firewall checks pass. |
| 4 — Release candidate | Build EXE, checksum, docs, manual matrix, performance measurements, review | Owner sign-off; all P0 tests pass; no unresolved critical/high security gaps. |

### Definition of done / owner sign-off
- All four text/image paths work on the actual iPhone 13 and VivoBook; no third-party service involved. Tests demonstrate on-demand semantics and no automatic sync loop.
- No insecure LAN mode, unauthorized clipboard access, unsafe retry, unbounded parser or clipboard logging; known residual risks documented and understood.
- Owner can install, configure, pair, use, pause, revoke, renew certificates, troubleshoot, update and uninstall using provided docs; no terminal needed for ordinary daily use.
- Source compiles, automated test results and security/dependency checks are recorded, manual acceptance matrix signed off, and versioned EXE/checksum generated locally. Repository remains private as owner intends; no release publishing or pushing without explicit authorization.
- If any required gate remains untested/failed, status is prototype or release candidate, not complete. Record remaining work and the exact next command/action.

## Supporting Information

| Supporting Docs | Link / location |
|---|---|
| Original PRD template | https://docs.google.com/document/d/1VFT9wY2RHT8Grp95YvAs7WKJqWtXrawWdXInBdfx8mo/edit |
| Existing repository (owner-provided) | https://github.com/aidanismail/air-clip |
| Codex implementation prompt | CODEX_PROMPT.md (companion deliverable) |
| Apple Shortcuts API requests | https://support.apple.com/guide/shortcuts/request-your-first-api-apd58d46713f/ios |
| Apple Shortcuts input types | https://support.apple.com/guide/shortcuts/understanding-input-types-apd7644168e1/ios |
| Apple certificate trust | https://support.apple.com/en-us/102390 |
| Win32 clipboard operations | https://learn.microsoft.com/en-us/windows/win32/dataxchg/clipboard-operations |
| Win32 clipboard sequence | https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getclipboardsequencenumber |
| Win32 OpenClipboard | https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-openclipboard |
| Go clipboard library | https://pkg.go.dev/golang.design/x/clipboard |
| Go HTTP | https://pkg.go.dev/net/http |
| Go TLS | https://pkg.go.dev/crypto/tls |
| Codex persistent execution plans | https://developers.openai.com/cookbook/articles/codex_exec_plans |

Reference links support API research, not proof of device compatibility. Confirm latest platform documentation and dependency versions during implementation. This PRD follows the supplied template's section order and user-story format; expands Technical Design and replaces commerce metrics with personal reliability/security goals. Exact repository contents were not inspected in this planning session; Codex must reconcile them. Proposed limits, protocol, trust model and risk controls are design baselines for owner review; any change that weakens security or alters core scope requires explicit approval.
