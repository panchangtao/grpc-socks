package lib

import (
	"crypto/tls"
	"crypto/x509"
	"log"

	"google.golang.org/grpc/credentials"
)

var certPEMBlock = []byte(`-----BEGIN CERTIFICATE-----
MIIEMDCCAxigAwIBAgIJAO8uVrXywwnLMA0GCSqGSIb3DQEBBQUAMG0xCzAJBgNV
BAYTAkNIMQ0wCwYDVQQIEwQxMDI0MQ0wCwYDVQQHEwQxMDI0MQ0wCwYDVQQKEwQx
MDI0MQ0wCwYDVQQLEwQxMDI0MQ0wCwYDVQQDEwQxMDI0MRMwEQYJKoZIhvcNAQkB
FgQxMDI0MB4XDTE2MDMyNjAwMTkwMloXDTE5MDMyNjAwMTkwMlowbTELMAkGA1UE
BhMCQ0gxDTALBgNVBAgTBDEwMjQxDTALBgNVBAcTBDEwMjQxDTALBgNVBAoTBDEw
MjQxDTALBgNVBAsTBDEwMjQxDTALBgNVBAMTBDEwMjQxEzARBgkqhkiG9w0BCQEW
BDEwMjQwggEiMA0GCSqGSIb3DQEBAQUAA4IBDwAwggEKAoIBAQCsHh0cmk/GGgsA
YzKnexeCMgciVR5kFAYzXGWlirfBvRi1hmVJ14guWslKpoM40kuWx77tKhoilcQA
ACfsrmRrXJYZ6z5Y6oawXxjpMEkDXZdje09VPTiTUQaFTjcb7qq9l0AjdBonpMb3
4In9DwtyWEeQCJYo0gnxZcYOVwdjO8yskM80dgSjfrBMeIzV4bDDxajGQq+ce/gS
9t3TfColdQhXGFBY/KbOHPTBzxCAt2KN2VyiTFWdw2jhe1k/NRgKjAoMQWmQR9lq
NeKQ8MGhtGT1drsVHVPueT+CW1lmb4ec3ga3v/wiRxXDJRuimiAs2hFJUgN6fkBk
FGYSIe11AgMBAAGjgdIwgc8wHQYDVR0OBBYEFF9Wev8a9eJzLDMvT1wzzWVKvtlp
MIGfBgNVHSMEgZcwgZSAFF9Wev8a9eJzLDMvT1wzzWVKvtlpoXGkbzBtMQswCQYD
VQQGEwJDSDENMAsGA1UECBMEMTAyNDENMAsGA1UEBxMEMTAyNDENMAsGA1UEChME
MTAyNDENMAsGA1UECxMEMTAyNDENMAsGA1UEAxMEMTAyNDETMBEGCSqGSIb3DQEJ
ARYEMTAyNIIJAO8uVrXywwnLMAwGA1UdEwQFMAMBAf8wDQYJKoZIhvcNAQEFBQAD
ggEBAEWXMh+5nAIdaSwsDe/EkaP8qNr4oDj/8DRVRq7J/7b5ts1b1RI8VmdE/AMi
x2S7Dsh51JvG3u1Oss+PU44uLgGKP6RtBpqeo8uvlCgpWY2Qve5m2r/AR3M75AjB
Rpvdtj7r4BD7uDTYseDazqZpC4A7tKqF/PomkWseA+QQHUjedkEk30e/7nHdZZoV
s/i2kgkU7Rna1wqesihX8SRfLCDvyckuPFgimz0ry+TIpuFgm1orytjcjYsMv6ax
UG+yA/zpmkIfvSrv6Gc6u+2hfBZREMcbyaqRxFlTOxRvVeKt14+pdJ90tc5Jq0aJ
aiJQH2u5zeca9eFkLZzcTUZmgvw=
-----END CERTIFICATE-----
`)

var keyPEMBlock = []byte(`-----BEGIN RSA PRIVATE KEY-----
MIIEogIBAAKCAQEArB4dHJpPxhoLAGMyp3sXgjIHIlUeZBQGM1xlpYq3wb0YtYZl
SdeILlrJSqaDONJLlse+7SoaIpXEAAAn7K5ka1yWGes+WOqGsF8Y6TBJA12XY3tP
VT04k1EGhU43G+6qvZdAI3QaJ6TG9+CJ/Q8LclhHkAiWKNIJ8WXGDlcHYzvMrJDP
NHYEo36wTHiM1eGww8WoxkKvnHv4Evbd03wqJXUIVxhQWPymzhz0wc8QgLdijdlc
okxVncNo4XtZPzUYCowKDEFpkEfZajXikPDBobRk9Xa7FR1T7nk/gltZZm+HnN4G
t7/8IkcVwyUbopogLNoRSVIDen5AZBRmEiHtdQIDAQABAoIBAEMSfR+/VqUJUw40
mfHAOxoGatdLOkufrgbn08id9Rvvl6htlh0fe7css9J+bxZ+hOxeLJ35UTP3Duk9
JlHKZ+Gas/env6Ugx5oFhQyiP2GrYCppTDIYaGqoIZVjNICTEkBNp7XpMaQPR2Yj
P0K5USmfE0wivHlt2GgU1AiUi5F0gYSYpGoahbgG030Mv+GmKtX69/KwE3mgrNQ/
54Czd6r+nHaCE9g0DfHwkIJrHHaWNa7DMoU6ws0glAFWAz8znYCZxPZfMp2oSEDn
BjG//aGrT/jCvD+jrTcJhc/0sVhrDXPmqz49OXVTLjgRmwrEkjPrWnpRA5VKu4a7
GjBFyWECgYEA2lvtSw+MA4O5ed9WE7LyldQojXrBU/gdBKs/OpyMv1bJTYUMkaPn
BZ2fVTJV327bA8yMNrM/Reiqx373NGcw8h57WFbLjwMV72tvy7sOmg8MYTCiNarU
fsigr3cdAoAwonSkUUPJHsJrlLca5pN0KXajCDg80BnMvsYnf+b2V3kCgYEAycmT
RwLVYIQS7DDO/CUHTErIa6jte+bjTTW7G7nG/LhqPThS/b6ohzez6wztFspIldYs
G0tO25q4NZa9jg1noM2CesvsGwlhxNrYbAK9CtZ6fWXgeDMvF71Hh1F+8LwfD4It
tXGiU4pfUp+5OSKZN/jtqqeyYE/MywwXc4gfOt0CgYBIqFQCKO8u8DLUUbNDpMTB
hDHmOdWAikullRHZ/+N5e3hKOh5fi8lAfh1ZbQFT8oAf+H0jamuAaJYDAcViA4Au
4GOsllzvflhbLUWq5dhK/Pzijhs7fldsxHdrS1g0z9DfDa7rd4HBoXHIr1DdLm11
qos/He9mU19kj2zvSzvnCQKBgCMVxGDNclJUxIGCvwqCWbF/Mzfc6GXpsE3lcMIS
XDHm0roQSAXMl7rjCYpt9e9HfrVmxsZ8Ipr2XN8cdZr0Y7dG5E/7kvLkf7ZdotGs
7DetMSEKjKv5ok+LXpt9pQewfeoRZWct+d5yqb5Q/UCc7m0YACLzA4XRejc3xAAX
g+6VAoGAI5n5EAi10YAm1P+E7sWtipdxe+FTXyketmWoty6pVw6wWrAgT8/HpNyG
Z93rnVEtVNJ9kqNPqrn4Mn4320cJxGzVNrOKrS8NAlccbKqkuZ3fMr/oZ6/8o5aU
zXPJfyC/bgm20s7B029Ojmwy3ReoTY2oL2hCeRRiUXp82Az+wCQ=
-----END RSA PRIVATE KEY-----
`)

func ServerTLS() credentials.TransportCredentials {
	cert, err := tls.X509KeyPair(certPEMBlock, keyPEMBlock)
	if err != nil {
		log.Fatal(err)
	}

	return credentials.NewTLS(&tls.Config{Certificates: []tls.Certificate{cert}})
}

func ClientTLS() credentials.TransportCredentials {
	cp := x509.NewCertPool()
	if !cp.AppendCertsFromPEM(certPEMBlock) {
		log.Fatal("Credentials: Failed to append certificates")
	}

	return credentials.NewClientTLSFromCert(cp, "1024")
}
