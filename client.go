//go:build ignore

package main

import (
  "crypto/tls"
  "crypto/x509"
  "log"
  "net/http"
  "net/http/httputil"
  "os"
)

func main() {
  cert, err := os.ReadFile("conf/ca.crt")
  if err != nil {
    panic(err)
  }
  certPool := x509.NewCertPool()
  certPool.AppendCertsFromPEM(cert)
  tlsConfig := &tls.Config{
    RootCAs: certPool,
  }
  tlsConfig.BuildNameToCertificate()

  // クライアントを生成
  client := &http.Client{
    Transport: &http.Transport{
      TLSClientConfig: tlsConfig,
    },
  }

  // 通信を行う
  resp, err := client.Get("https://localhost:18443")
  if err != nil {
    panic(err)
  }
  defer resp.Body.Close()
  dump, err := httputil.DumpResponse(resp, true)
  if err != nil {
    panic(err)
  }
  log.Println(string(dump))
}
