If you want your web server to use HTTPS instead of relying on external
applications, such as NGINX, you can easily do so if you already have a
valid certificate. If you don't, you can use OpenSSL to create one:
`> openssl genrsa -out server.key 2048`
`> openssl req -new -x509 -sha256 -key server.key -out server.crt -days 3650`
The first command generates a private key, while the second one creates a
public certificate that you need for the server. The second command will also
require a lot of additional information in order to create the certificate, from
country name to email address.
