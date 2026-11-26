# Air-Clip architecture

The current executable reports `Unconfigured`. The in-memory adapter models sequence checks and snapshots for synthetic testing. The packages below identify the planned implementation boundaries.

| Path | Responsibility |
| --- | --- |
| `cmd/air-clip` | Windows app and local setup entry points |
| `internal/api` | V1 route, admission and typed response handling |
| `internal/auth` | Bearer verifier, volatile tickets and bounded limiters |
| `internal/clipboard` | Clipboard contract and native Windows adapter |
| `internal/clipboard/cliptest` | In-memory clipboard with call counts |
| `internal/codec` | Text and PNG/DIB validation and conversion |
| `internal/lifecycle` | Session/network eligibility, epochs and cancellation |
| `internal/network` | IPv4 network allowlist, binding and discovery |
| `internal/security` | Certificate provisioning and protected local state |
| `internal/tray` | Windows lifecycle controls and metadata-only status |

The iPhone initiates every request. The target path is Shortcuts → HTTPS → authentication and eligibility → typed validation → serialized Windows clipboard executor. The executor owns the native acquisition, final sequence check and mutation. On pull, copy a live snapshot before releasing the handle, encode outside the acquisition, and recheck eligibility before returning it. See the [PRD](PRD.md) for exact contracts and [ADRs](DECISIONS.md) for platform constraints.
