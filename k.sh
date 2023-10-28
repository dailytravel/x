#!/bin/bash

# Create the certs directory if it doesn't exist
mkdir -p certs/x509

# Create a configuration file for OpenSSL with the SAN extension
cat > openssl.cnf <<EOL
[req]
req_extensions = v3_req
distinguished_name = req_distinguished_name
[req_distinguished_name]
[v3_req]
subjectAltName = @alt_names
[alt_names]
DNS.1 = api.trip.express
EOL

# Generate a self-signed server certificate and private key with SANs
openssl req -x509 -newkey rsa:4096 -keyout certs/x509/server_key.pem -out certs/x509/server_cert.pem -days 365 -nodes -subj "/CN=api.trip.express" -config openssl.cnf

# Generate a self-signed CA certificate and private key
openssl req -x509 -newkey rsa:4096 -keyout certs/x509/ca_key.pem -out certs/x509/ca_cert.pem -days 365 -nodes -subj "/CN=CertificateAuthority"

# Secure the private key files by setting restrictive permissions
chmod 600 certs/x509/server_key.pem certs/x509/ca_key.pem

# Clean up the temporary configuration file
rm openssl.cnf

echo "TLS certificate and key files, and CA certificate files generated in 'certs/x509/' directory."
