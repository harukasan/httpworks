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
	"sync"
)

type pool struct {
	sync.Pool
}

func (p pool) Get() []byte {
	return p.Pool.Get().([]byte)
}

func (p pool) Put(b []byte) {
	p.Pool.Put(b)
}

func newPool() *pool {
	return &pool{
		Pool: sync.Pool{
			New: func() interface{} {
				return make([]byte, 4*1024)
			},
		},
	}
}

func main() {
	var urlstr string
	var addr string
	var cert string
	var key string
	var caCertFile string
	var usetls bool

	flag.StringVar(&urlstr, "url", "", "target url")
	flag.StringVar(&addr, "addr", "", "http server address")
	flag.BoolVar(&usetls, "tls", false, "use tls")
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
	proxy.BufferPool = newPool()
	proxy.Transport = http.DefaultTransport
	proxy.Transport.(*http.Transport).MaxIdleConnsPerHost = 100

	if usetls {
		log.Fatalf("serve: %v", http.ListenAndServeTLS(addr, cert, key, proxy))
	} else {
		log.Fatalf("serve: %v", http.ListenAndServe(addr, proxy))
	}
}
