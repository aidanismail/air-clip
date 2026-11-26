# iPhone Shortcuts feasibility

The installed iOS version and Shortcuts behavior are pending owner verification. Apple documents manual TLS root trust under **Settings → General → About → Certificate Trust Settings**. An Air-Clip LAN probe first requires an approved endpoint and verification of the public CA SHA-256 fingerprint. Full trust is system-wide trust for that root. The owner manually enables it, then removes the test profile in **Settings → General → VPN & Device Management** when the probe ends. Keep the root private key on Windows or destroy it after issuing the probe leaf. [Apple root-trust documentation](https://support.apple.com/en-us/102390)

## Optional offline owner experiment

These temporary actions are candidates to verify on your iOS build. They use synthetic input and require no network, token, certificate or clipboard overwrite.

1. Record the iOS version and build from **Settings → General → About**.
2. Create a disposable manual shortcut. Add **Text** with your synthetic sample `Air-Clip ✓ 🌱`, then a second line. Add **Base64 Encode**, choose Encode and no line breaks if offered, then a second **Base64 Encode** set to Decode. Pass the decoded bytes to **Get Text from Input** if available and use **Show Result** to visually compare with the original.
3. Create a second temporary shortcut. Choose a test image you created specifically for the probe. Use **Select Photos → Convert Image** with PNG, then Base64 encode and decode as above. Try **Get Images from Input** on the decoded bytes and **Quick Look** only for this synthetic image. Record whether the result appears as an image. This checks type conversion. Real paste is a separate test.
4. Delete the two disposable shortcuts and the synthetic image when finished. Report build, exact action labels, observed output types and safe error categories. Keep screenshots and real clipboard content local.

The eventual manual Send and Get shortcuts follow the [PRD action specification](PRD.md#ios-shortcut-build-specification-manual-owner-validation). The Shortcuts API, HTTP error, clipboard, local-only copy, Back Tap, widget and permission paths will be confirmed on the actual phone before the installation guide is finalized. [Apple Shortcuts API limitations](https://support.apple.com/en-ie/guide/shortcuts/apd891a6c84e/ios) · [Apple Shortcuts input types](https://support.apple.com/en-gb/guide/shortcuts/apd7644168e1/ios)
