# Air-Clip acceptance and evidence

Use synthetic fixtures in automated tests. Physical results require the owner devices and separate permission for live clipboard fixtures, trust, firewall or device settings. Record command, exit code, device and app versions, pass/fail and sanitized evidence. See the [progress log](PROGRESS.md) for exact commands. Acceptance scenarios are open until all applicable parts pass.

| Test | Requirements | Evidence still required | Acceptance |
| --- | --- | --- | --- |
| T01 | FR-03–07, SM-01–03 | Exact text fixture tests and real Unicode text in both directions | Not run |
| T02 | FR-05–08, SEC-08 | Invalid UTF-8, NUL, empty and oversized text preserves destination | Not run |
| T03 | FR-03–07, SM-01–02 | PNG/DIB fixtures and actual screenshot/photo paste, both directions | Not run |
| T04 | FR-07–08, SEC-08 | Malformed/bounded PNG, DIB, JSON and Base64 test and fuzz evidence | Not run |
| T05 | FR-06–07, SEC-05 | Unsupported and mixed formats, deterministic selection | Not run |
| T06 | FR-01, SEC-04–05 | Strict HTTP rejection with zero clipboard calls | Not run |
| T07 | FR-01, FR-11, SEC-01–03, SEC-06 | TLS chain/identity tests and iPhone rejection behavior | Not run |
| T08 | FR-03, FR-05, FR-08, SM-05 | Expiry, concurrent replay, differing payload and terminal ticket results | Not run |
| T09 | FR-05, FR-08, SM-05 | Race-resistant native commit, delayed rendering, busy and partial failure | Not run |
| T10 | FR-08–09, SEC-07 | Controlled precommit and acknowledgement-loss interruption | Not run |
| T11 | FR-09, SEC-07, SEC-10 | Epoch cancellation plus lock/sleep/pause/revoke and reconnect | Not run |
| T12 | FR-09, FR-11, SEC-06–07 | Approved network, VPN, profile, DHCP, mDNS and firewall checks | Not run |
| T13 | FR-08, SEC-08 | Bounded slow, oversized and concurrent request handling | Not run |
| T14 | FR-02–03, FR-08–09, SEC-07 | Restart, invalid config, port conflict and single instance | Not run |
| T15 | FR-01–02, FR-09–10, SEC-02–04, SEC-06, SEC-10 | Owner installation, rollback, pairing, startup and uninstall | Not run |
| T16 | FR-04, FR-06, FR-08, FR-12, SEC-12 | Phone permissions, malformed response, race and manual launch | Not run |
| T17 | SEC-04–05, SEC-09, SEC-11 | Source/packaging scan and diagnostics sentinel tests | Not run |
| T18 | FR-03–05, SM-03, SM-05 | Adapter call counts and actual manual transfers/idle | Not run |
| T19 | FR-06, SEC-09, SEC-12 | Win32 exclusion and iOS local-only options, OS behavior | Not run |
| T20 | SM-01–10 | Actual device outcomes and timings/footprint/recovery | Not run |

Early fake-adapter checks exercise snapshot ownership, concurrent compare-and-write, deliberate identical writes, and cancellation. They inform T09 and T18 without closing those scenarios.
