# Air-Clip repository instructions

@C:\Users\aidan\.codex\RTK.md

- Use `docs/PRD.md` as the normative baseline. Keep requirement IDs stable and record material protocol, scope, or security decisions in `docs/DECISIONS.md`. Get owner approval before changing scope, trust assumptions, security defaults, or the devices.
- At handoff or resume, read this file, the PRD, `docs/IMPLEMENTATION_PLAN.md`, `docs/PROGRESS.md`, and `docs/DECISIONS.md`. Reconcile Git status, the actual diff, and command results before editing or retrying a setup operation.
- Work in focused phase or topic branches. Preserve owner edits. Get express authorization for commits, pushes, publishing, uploads, or remote repository changes. Use `rtk proxy` to capture tool output when exact evidence matters.
- Use synthetic data, fake clipboard tests, and IPv4 loopback by default. Real clipboard access, trust installation, startup changes, LAN enablement, and privileged actions require owner approval. Keep app payloads, tokens, keys, and configuration out of Git, logs, and agent context.
- Use small Go packages with the standard library first. Review and pin external dependencies. Keep the Win32 ownership and compare-and-write boundary explicit. Use tests that exercise the normal entry point for each feature and distinguish automated, native, and physical evidence.
- Update the plan, test matrix, and checkpoint when a significant change or verification result lands. Include exact commands, exit codes, changed files, unresolved gates, and the next task. Owner sign-off and passing physical-device P0 gates are required for daily-use status.
