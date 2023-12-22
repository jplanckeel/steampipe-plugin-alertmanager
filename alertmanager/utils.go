package alertmanager

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"errors"
	"fmt"
	"net/http"
	"os"
	"path"

	clientruntime "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/prometheus/alertmanager/api/v2/client"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

const (
	defaultAmApiv2path = "/api/v2"
)

func connectAlertManager(ctx context.Context, d *plugin.QueryData) (*client.AlertmanagerAPI, error) {

	address := os.Getenv("ALERTMANAGER_ADDRESS")
	caCert := os.Getenv("ALERTMANAGER_CA_CERT")
	tlsCert := os.Getenv("ALERTMANAGER_TLS_CERT")
	tlsKey := os.Getenv("ALERTMANAGER_TLS_KEY")
	schemes := []string{"http"}
	alertmanager_path := "/"

	alertManagerConfig := GetConfig(d.Connection)
	if alertManagerConfig.Address != nil {
		address = *alertManagerConfig.Address
	}
	if address == "" {
		return nil, errors.New("'address' must be set in the connection configuration. Edit your connection configuration file and then restart Steampipe")
	}

	if alertManagerConfig.TlsCert != nil {
		tlsCert = *alertManagerConfig.TlsCert
	}

	if alertManagerConfig.CaCert != nil {
		caCert = *alertManagerConfig.CaCert
	}

	if alertManagerConfig.TlsKey != nil {
		tlsKey = *alertManagerConfig.TlsKey
	}

	var cert tls.Certificate
	var pool *x509.CertPool

	if caCert != "" {
		ca, err := os.ReadFile(caCert)
		if err != nil {
			return nil, fmt.Errorf("ca_cert error: %s", err.Error())
		}
		pool = x509.NewCertPool()
		pool.AppendCertsFromPEM(ca)
	}
	if tlsKey != "" && tlsCert != "" {
		var err error
		cert, err = tls.LoadX509KeyPair(tlsCert, tlsKey)
		if err != nil {
			return nil, fmt.Errorf("tls_key and tls_cert error: %s", err.Error())
		}
		schemes = []string{"https"}
	}

	httpClient := &http.Client{Transport: &http.Transport{
		TLSClientConfig: &tls.Config{
			RootCAs:      pool,
			Certificates: []tls.Certificate{cert},
		},
	}}

	cr := clientruntime.NewWithClient(address, path.Join(alertmanager_path, defaultAmApiv2path), schemes, httpClient)

	client := client.New(cr, strfmt.Default)

	return client, nil
}
