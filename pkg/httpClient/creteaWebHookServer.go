package httpClient

import (
	"flag"
	"github.com/muhammedkaya/tcpdump-webhook/pkg"
	"log"
	"net/http"
	"strconv"
)

func CreateWebHookServer() {
	flag.IntVar(&pkg.Parameters.Port, "port", 8443, "Webhook server port.")
	flag.StringVar(&pkg.Parameters.CertFile, "tlsCertFile", "/etc/webhook/certs/tls.crt", "File containing the x509 Certificate for HTTPS.")
	flag.StringVar(&pkg.Parameters.KeyFile, "tlsKeyFile", "/etc/webhook/certs/tls.key", "File containing the x509 private key to --tlsCertFile.")
	flag.Parse()

	http.HandleFunc("/mutate", HandleMutateEndpoint)
	log.Fatal(http.ListenAndServeTLS(":"+strconv.Itoa(pkg.Parameters.Port), pkg.Parameters.CertFile, pkg.Parameters.KeyFile, nil))
}
