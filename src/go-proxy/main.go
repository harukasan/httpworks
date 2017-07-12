package main

import (
	"crypto/tls"
	"crypto/x509"
	"flag"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func main() {
	var urlstr string
	var addr string
	var cert string
	var key string
	var caCertFile string

	flag.StringVar(&urlstr, "url", "", "target url")
	flag.StringVar(&addr, "addr", "", "server address")
	flag.StringVar(&cert, "cert", "", "certificate file")
	flag.StringVar(&caCertFile, "ca-cert", "", "ca certificate file")
	flag.StringVar(&key, "key", "", "certificate key file")
	flag.Parse()

	u, err := url.Parse(urlstr)
	if err != nil {
		log.Fatalf("failed to parse url: %v", urlstr)
	}

	if caCertFile != "" {
		caCert, err := ioutil.ReadFile(caCertFile)
		if err != nil {
			log.Fatal("failed to read ca certificate file: %v", err)
		}

		caCertPool := x509.NewCertPool()
		caCertPool.AppendCertsFromPEM(caCert)

		tlsConfig := &tls.Config{
			RootCAs: caCertPool,
		}
		tlsConfig.BuildNameToCertificate()

		http.DefaultTransport.(*http.Transport).TLSClientConfig = tlsConfig
	}

	proxy := httputil.NewSingleHostReverseProxy(u)
	proxy.Transport = http.DefaultTransport
	log.Fatalf("serve: %v", http.ListenAndServeTLS(addr, cert, key, proxy))
}
