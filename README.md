# fabric-sdk-go-gm
base from fabric 1.4.4 and fabric-sdk-go v1.0.0 beta

This application only supports the GM version of the network.

usage method:

The configuration of bccsp needs to be modified for the application of client:

```
# BCCSP config for the client. Used by GO SDK.

BCCSP:
  Security:
    enabled: true
    # SW/GM SHA2/GMSM3
    default:
      provider: "GM"
    hashAlgorithm: "GMSM3"
    softVerify: true
    Level: 256
```

# Remaining Problem
- TLS is not supported in the current version