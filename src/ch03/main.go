package main

import (
	"net/http"
)

//type Server struct {
//	Addr           string
//	Handler        Handler
//	ReadTimeout    time.Duration
//	WriteTimeout   time.Duration
//	MaxHeaderBytes int
//	TLSConfig      *tls.Config
//	TLSNextProto   map[string]func(*Server, *tls.Conn, Handler)
//	ConnState      func(net.Conn, ConnState)
//	ErrorLog       *log.Logger
//}

/*
func main() {
	//最も簡単なサーバー立ち上げ↓
	//http.ListenAndServe("", nil)

	server := http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: nil,
	}
	server.ListenAndServe()
}
*/

/*
httpsによるサービス
cert.pemがSSL証明書、key.pemがサーバー用の秘密鍵
*/
func main() {
	server := http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: nil,
	}
	server.ListenAndServeTLS("cert.pem", "key.pem")
}
