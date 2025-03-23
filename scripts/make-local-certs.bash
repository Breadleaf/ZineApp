#! /usr/bin/env bash

PROJECT_ROOT="$(git rev-parse --show-toplevel)"
CERT_DIR="$PROJECT_ROOT/frontend/certs"

mkdir -p $CERT_DIR

openssl req -x509 -nodes -days 365 \
    -newkey rsa:2048 \
    -keyout $CERT_DIR/privkey.pem \
    -out $CERT_DIR/fullchain.pem \
    -subj "/C=US/ST=State/L=City/O=ZineApp Dev/CN=zine-app.local"

echo "Local self-signed certs generated in $CERT_DIR"