# Air-Clip

Air-Clip is a planned on-demand clipboard bridge between an iPhone and a Windows PC. The phone initiates each transfer through Apple Shortcuts over authenticated HTTPS on an approved local Wi-Fi network. V1 carries UTF-8 text or one PNG image.

This starter tree builds a Go executable that reports the unconfigured state. The in-memory fake clipboard is available for synthetic tests. Device use begins after the feasibility and setup gates in the [implementation plan](docs/IMPLEMENTATION_PLAN.md).

## Development

Use Windows PowerShell and a patched Go release on a supported Go 1.26 or newer line. The module depends on the Go standard library. Run from the repository root:

```powershell
.\scripts\check.ps1
go run ./cmd/air-clip
```

If local PowerShell policy prevents running the script, the equivalent commands are `gofmt -l cmd internal`, `go vet ./...`, `go test ./...`, `go mod verify`, and `go build ./...`. Actual tool versions and test outcomes are recorded in the [progress log](docs/PROGRESS.md).

## Project references

- [Product requirements](docs/PRD.md), [implementation plan](docs/IMPLEMENTATION_PLAN.md), and [project decisions](docs/DECISIONS.md)
- [Architecture](docs/ARCHITECTURE.md), [security](docs/SECURITY.md), and [test plan](docs/TEST_PLAN.md)
- [Windows setup status](docs/WINDOWS_SETUP.md) and [iPhone Shortcuts experiments](docs/IOS_SHORTCUTS.md)
