# Air-Clip security

The Phase 0 starter has no listening endpoint. See the [PRD](PRD.md#sec--security-and-privacy-requirements-mandatory-release-gates) for SEC-01–12 and the [test matrix](TEST_PLAN.md) for verification status.

LAN peers, discovery records, input and request framing are untrusted. The target design requires an approved network and authenticated, identity-verified TLS. Windows pull requires a separate local enablement. A lost phone or copied Shortcut can expose the bearer and may permit live reads while pull is enabled. Revoke the credential from Windows and pair locally again. Read-disable, pause, rotation and trusted CA removal are operational requirements for the later Windows build.

OS clipboard history, cloud sync, iCloud Shortcuts sync, paging, and backups can retain data outside the app. A malicious process in the same desktop session can access the Windows clipboard. iOS may replace a newer phone copy during a pull. Windows may lose clipboard content if commit fails after emptying it. The acceptance plan tests controls and documents remaining risk without changing device privacy settings.
