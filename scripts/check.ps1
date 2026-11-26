$ErrorActionPreference = 'Stop'
Push-Location (Join-Path $PSScriptRoot '..')
try {
    $unformatted = & gofmt -l cmd internal
    if ($LASTEXITCODE -ne 0) { throw "gofmt failed (exit $LASTEXITCODE)." }
    if ($unformatted) { throw "Run gofmt on: $($unformatted -join ', ')" }

    & go vet ./...
    if ($LASTEXITCODE -ne 0) { throw "go vet failed (exit $LASTEXITCODE)." }

    & go test ./...
    if ($LASTEXITCODE -ne 0) { throw "go test failed (exit $LASTEXITCODE)." }

    & go mod verify
    if ($LASTEXITCODE -ne 0) { throw "go mod verify failed (exit $LASTEXITCODE)." }

    & go build ./...
    if ($LASTEXITCODE -ne 0) { throw "go build failed (exit $LASTEXITCODE)." }
} finally {
    Pop-Location
}
