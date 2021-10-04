# Integration Tests

This directory contains integration tests that push the asa-go library outside the scope of the unit tests or the examples projects. 
These integration tests will make mutating calls against a production Apple Search Ads team, are invoked manually, and take a while to run. 
The unit tests by comparison are fast to run, do not involve the network, and run automatically on every commit.

To run the integration tests, provide something similar to this command line invocation.

```shell
env \
      ASA_INTEGRATION_OID="..."
	  ASA_INTEGRATION_KID="..." \
	  ASA_INTEGRATION_TID="..." \
	  ASA_INTEGRATION_CID="..." \
	  ASA_INTEGRATION_PRIVATE_KEY_PATH="..." \
      go test -v -tags=integration ./test/integration
```

Much like the examples, the integration tests require the presence of at least 5 of 6 different environment variables. 
If you want to know what they do, please consult the repo's general documentation on [authentication](../README.md#Authentication):

- `ASA_INTEGRATION_OID` - organization ID
- `ASA_INTEGRATION_KID` – key ID
- `ASA_INTEGRATION_TID` – team ID
- `ASA_INTEGRATION_CID` – client ID
- `ASA_INTEGRATION_PRIVATE_KEY` – path to a private key
- `ASA_INTEGRATION_PRIVATE_KEY_PATH` – path to a private key

Only one of either `ASA_INTEGRATION_PRIVATE_KEY` or `ASA_INTEGRATION_PRIVATE_KEY_PATH` is required; if both are provided, `ASA_INTEGRATION_PRIVATE_KEY` will take precedence. 
Since the Apple Search Ads API requires an authenticated session, you must have valid credentials to run these tests.