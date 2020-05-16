package lib

import (
	"app/service/fabric-sdk-go-gm/internal/github.com/hyperledger/fabric/bccsp"
	"app/service/fabric-sdk-go-gm/internal/github.com/hyperledger/fabric/bccsp/gm"
	"app/service/fabric-sdk-go-gm/internal/github.com/tjfoc/gmsm/sm2"
	"crypto"
	"crypto/x509/pkix"
	"encoding/asn1"
	"fmt"
	"github.com/cloudflare/cfssl/csr"
	"github.com/cloudflare/cfssl/log"
	"net"
	"net/mail"
)

// cloudflare 证书请求 转成 国密证书请求
func generate(priv crypto.Signer, req *csr.CertificateRequest, key bccsp.Key) (csr []byte, err error) {
	log.Info("xx entry generate")
	sigAlgo := signerAlgo(priv)
	if sigAlgo == sm2.UnknownSignatureAlgorithm {
		return nil, fmt.Errorf("Private key is unavailable")
	}
	log.Info("xx begin create sm2.CertificateRequest")
	var tpl = sm2.CertificateRequest{
		Subject:            req.Name(),
		SignatureAlgorithm: sigAlgo,
	}
	for i := range req.Hosts {
		if ip := net.ParseIP(req.Hosts[i]); ip != nil {
			tpl.IPAddresses = append(tpl.IPAddresses, ip)
		} else if email, err := mail.ParseAddress(req.Hosts[i]); err == nil && email != nil {
			tpl.EmailAddresses = append(tpl.EmailAddresses, email.Address)
		} else {
			tpl.DNSNames = append(tpl.DNSNames, req.Hosts[i])
		}
	}

	if req.CA != nil {
		err = appendCAInfoToCSRSm2(req.CA, &tpl)
		if err != nil {
			err = fmt.Errorf("sm2 GenerationFailed")
			return
		}
	}
	if req.SerialNumber != "" {

	}
	csr, err = gm.CreateSm2CertificateRequestToMem(&tpl, key)
	log.Info("xx exit generate")
	return
}

func signerAlgo(priv crypto.Signer) sm2.SignatureAlgorithm {
	switch pub := priv.Public().(type) {
	case *sm2.PublicKey:
		switch pub.Curve {
		case sm2.P256Sm2():
			return sm2.SM2WithSM3
		default:
			return sm2.SM2WithSM3
		}
	default:
		return sm2.UnknownSignatureAlgorithm
	}
}

// appendCAInfoToCSRSm2 appends CAConfig BasicConstraint extension to a CSR
func appendCAInfoToCSRSm2(reqConf *csr.CAConfig, csreq *sm2.CertificateRequest) error {
	pathlen := reqConf.PathLength
	if pathlen == 0 && !reqConf.PathLenZero {
		pathlen = -1
	}
	val, err := asn1.Marshal(csr.BasicConstraints{true, pathlen})

	if err != nil {
		return err
	}

	csreq.ExtraExtensions = []pkix.Extension{
		{
			Id:       asn1.ObjectIdentifier{2, 5, 29, 19},
			Value:    val,
			Critical: true,
		},
	}

	return nil
}
