package alertmanager

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"errors"
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
	schemes := []string{"http"}
	alertmanager_path := "/"

	alertManagerConfig := GetConfig(d.Connection)
	if alertManagerConfig.Address != nil {
		address = *alertManagerConfig.Address
	}
	if address == "" {
		return nil, errors.New("'address' must be set in the connection configuration. Edit your connection configuration file and then restart Steampipe")
	}

	cr := clientruntime.New(address, path.Join(alertmanager_path, defaultAmApiv2path), schemes)

	if *alertManagerConfig.CertKey != "" && *alertManagerConfig.CertCA != "" && *alertManagerConfig.Cert != "" {
		cert, err := os.ReadFile(*alertManagerConfig.CertCA)
		if err != nil {
			return nil, errors.New("could not open certificate ca file")
		}
		caCertPool := x509.NewCertPool()
		caCertPool.AppendCertsFromPEM(cert)

		certificate, err := tls.LoadX509KeyPair(*alertManagerConfig.Cert, *alertManagerConfig.CertKey)
		if err != nil {
			return nil, errors.New("could not load certificate and certificate key file")
		}

		httpClient := &http.Client{Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				RootCAs:      caCertPool,
				Certificates: []tls.Certificate{certificate},
			},
		}}

		cr = clientruntime.NewWithClient(address, path.Join(alertmanager_path, defaultAmApiv2path), schemes, httpClient)
	}

	client := client.New(cr, strfmt.Default)

	return client, nil
}
