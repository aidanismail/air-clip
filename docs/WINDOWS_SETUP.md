# Windows setup status

The current build is a starter for repository development. The Windows install and uninstall workflow will be validated during Phase 3. Runtime operation is planned for a signed-in interactive user, with startup opt-in and per-user configuration under `%LOCALAPPDATA%\Air-Clip`.

For Phase 0, the owner can record the edition, version and OS build shown by `winver` and confirm the device is the target VivoBook. The development shell reports x64 and build `26300.10075`. See the [checkpoint](PROGRESS.md) for exact host evidence. The first HTTPS integration tests will run on IPv4 loopback with temporary in-memory trust material. A later LAN probe needs a separately approved network/interface, address and narrowly scoped firewall rule. No device setup is required to run the scaffold tests.
