package server

import (
	"crypto/tls"
	"net/http"
	"time"
)

func New(mux *http.ServeMux, servAddr string) *http.Server {
	tls := &tls.Config{
		//选择密码套件
		PreferServerCipherSuites: true,
		// 椭圆曲线优先级
		CurvePreferences: []tls.CurveID{
			tls.CurveP256,
			tls.X25519,
		},
		MinVersion: tls.VersionTLS12,
		// CipherSuites 是受支持的密码套件
		CipherSuites: []uint16{
			tls.TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305,
			tls.TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305,
			tls.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256,
			tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,
		},
	}

	srv := &http.Server{
		Addr:         servAddr,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  20 * time.Second,
		Handler:      mux,
		TLSConfig:    tls,
	}
	return srv
}
