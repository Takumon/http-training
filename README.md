# http-training

## 手順

- ルート証明書作成
  - 256ビットの楕円曲線デジタル署名アルゴリズム(ECDSA)の認証局秘密鍵を作成

    ```bash
    openssl genpkey -algorithm ec -pkeyopt ec_paramgen_curve:prime256v1 -out ca.key
    ```

  - 証明書署名要求( CSR)を作成

    ```bash
    openssl req -new -sha256 -key ca.key -out ca.csr -config openssl.cnf
    ```

  - 証明書を自分の秘密鍵で署名して作成

    ```bash
    openssl x509 -in ca.csr -days 365 -req -signkey ca.key -sha256 -out ca.crt -extfile ./openssl.cnf -extensions CA
    ```

- 作成結果の確認
  - 秘密鍵の確認

    ```bash
    openssl ec -in ca.key -text
    ```

  - 証明書署名要求(CSR)の確認

    ```bash
    openssl req -in ca.csr -text
    ```

  - 証明書の確認

    ```bash
    openssl x509 -in ca.crt -text
    ```

- サーバー証明書作成
  - 256ビットの楕円曲線デジタル署名アルゴリズム(ECDSA)のサーバー秘密鍵を作成

    ```bash
    openssl genpkey -algorithm ec -pkeyopt ec_paramgen_curve:prime256v1 -out server.key
    ```

  - 証明書署名要求(CSR)を作成

    ```bash
    openssl req -new -nodes -sha256 -key server.key -out server.csr -config openssl.cnf
    ```

  - 証明書を自分の秘密鍵で署名して作成

    ```bash
    openssl x509 -req -days 365 -in server.csr -sha256 -out server.crt -CA ca.crt -CAkey ca.key -CAcreateserial -extfile ./openssl.cnf -extensions Server
    ```

- 作成結果の確認
  - ルート証明書の場合と同様なので略
