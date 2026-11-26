# Air-Clip architectural decisions

## ADR-001 · Keep the checked-in v1 product baseline

**2026-11-26 · Accepted for planning.** The checkout contains the full owner-supplied v1.0 planning PRD and the request designates it normative. Preserve its wording and IDs. No independently supplied second version was available to compare. The PRD records several prior proposals. Owner approval and measured evidence are tracked separately. The initial file SHA-256 is recorded in [the checkpoint](PROGRESS.md).

## ADR-002 · Prove loopback TLS before provisioning any device

**2026-11-26 · Accepted for Phase 0.** Use Go standard packages, a synthetic clipboard, and test-scoped certificates on `127.0.0.1`. Use an explicit root pool and strict hostname/IP verification. Phase 0 test material stays in process memory. No third-party modules are needed for this slice. The locally installed Go 1.26.4 is useful for an isolated prototype test. The patched Go version and vulnerability results are release prerequisites. [Go release history](https://go.dev/doc/devel/release)

## ADR-003 · Own the native clipboard commit boundary

**2026-11-26 · Design recorded, native proof pending.** `golang.design/x/clipboard` exposes ordinary `Read` and `Write` with no caller-owned transaction spanning sequence check and mutation, so the critical path requires a narrow Win32 adapter. Evaluate `golang.org/x/sys/windows` for the native bindings before pinning. Write with a valid owner window, stage movable allocations, check sequence under the same `OpenClipboard`, then empty and transfer handle ownership. Windows can defer sequence increment for delayed rendering. The implementation must assess a conservative way to prove the snapshot and commit safe, and treat unknown state as unavailable. A failure after emptying is terminal and may have changed the destination. [Clipboard library API](https://pkg.go.dev/golang.design/x/clipboard) · [Windows bindings](https://pkg.go.dev/golang.org/x/sys/windows) · [OpenClipboard](https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-openclipboard) · [Sequence semantics](https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getclipboardsequencenumber) · [SetClipboardData](https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-setclipboarddata)

## ADR-004 · Prove direct-address trust before mDNS

**2026-11-26 · Proposed, owner approval required for LAN and trust.** Use a dedicated Air-Clip CA, a short-lived leaf with `serverAuth`, and an exact approved IPv4 SAN for the first LAN experiment. Prefer discarding the CA private key after issuance for a disposable probe. Retaining a CA key for renewal later requires an owner-approved storage and lifecycle plan. The iPhone receives only the verified public CA. The LAN probe needs approved network identity, firewall scoping and a different, ephemeral bearer. Verify Shortcuts with that trust before evaluated mDNS implementation. Apple's root-trust procedure and certificate guidance are documented. Shortcuts behavior and visible fingerprint verification on the owner's iOS build are open experiments. [Apple manual root trust](https://support.apple.com/en-us/102390) · [Apple certificate requirements](https://support.apple.com/en-us/103769)
