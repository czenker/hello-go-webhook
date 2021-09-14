This is an **insecure** self-signed certificate for `localhost`. It is used to locally run
the webhook and for testing.

It should never find its way to a server.

## Recreate new certificate

````bash
openssl req -x509 -sha256 -nodes -newkey rsa:2048 -days 3650 -subj "/CN=localhost" -addext "subjectAltName = DNS:localhost" -keyout tls.key -out tls.crt
````

Reference: https://stackoverflow.com/a/65711669

`@TODO`: I will very like automate it, but if you update the certificate you have to update the
`caBundle` in [`test/manifests/webhook.yaml`](../manifests/webhook.yaml), too.

