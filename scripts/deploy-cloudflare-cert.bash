#! /usr/bin/env bash

ZONE="zine-app.com"

PROJECT_ROOT="$(git rev-parse --show-toplevel)"
CERT_DIR="$PROJECT_ROOT/frontend/certs"

mkdir -p $CERT_DIR

cloudflared cert create "$ZONE" \
    --host "$ZONE,*.$ZONE" \
    --type origin-pull \
    --output "$CERT_DIR/origin.pem" \
    --key "$CERT_DIR/origin-key.pem"

mv "$CERT_DIR/origin.pem" "$CERT_DIR/fullchain.pem"
mv "$CERT_DIR/origin-key.pem" "$CERT_DIR/privkey.pem"

echo "Cloudflare origin cert installed to $CERT_DIR"