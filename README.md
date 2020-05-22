# fabric-sdk-go-gm
国密改造版本的fabric go sdk
基于 fabric-sdk-go v1.0.0 



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
``
