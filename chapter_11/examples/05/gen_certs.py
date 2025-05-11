#!/usr/bin/env python3

import os
import sys
from cryptography import x509
from cryptography.x509.oid import NameOID
from cryptography.hazmat.primitives.asymmetric import rsa
from cryptography.hazmat.primitives import hashes, serialization
from datetime import datetime, timedelta, timezone

def create_self_signed_cert(index, output_dir):
    key = rsa.generate_private_key(public_exponent=65537, key_size=2048)
    subject = issuer = x509.Name([
        x509.NameAttribute(NameOID.COUNTRY_NAME, u"US"),
        x509.NameAttribute(NameOID.ORGANIZATION_NAME, u"ExampleOrg"),
        x509.NameAttribute(NameOID.COMMON_NAME, u"example.com"),
    ])
    cert = (
        x509.CertificateBuilder()
        .subject_name(subject)
        .issuer_name(issuer)
        .public_key(key.public_key())
        .serial_number(x509.random_serial_number())
        .not_valid_before(datetime.now(timezone.utc))
        .not_valid_after(datetime.now(timezone.utc) + timedelta(days=365))
        .sign(key, hashes.SHA256())
    )

    key_path = os.path.join(output_dir, f"key{index}.pem")
    cert_path = os.path.join(output_dir, f"cert{index}.pem")

    with open(key_path, "wb") as f:
        f.write(key.private_bytes(
            encoding=serialization.Encoding.PEM,
            format=serialization.PrivateFormat.TraditionalOpenSSL,
            encryption_algorithm=serialization.NoEncryption(),
        ))
    with open(cert_path, "wb") as f:
        f.write(cert.public_bytes(serialization.Encoding.PEM))

if __name__ == "__main__":
    output_dir = "gen_certs"
    os.makedirs(output_dir, exist_ok=True)
    N = int(sys.argv[1]) if len(sys.argv) > 1 else 5
    for i in range(N):
        create_self_signed_cert(i, output_dir)
    print(f"✅ Generated {N} certs in '{output_dir}/'")

"""
openssl x509 -in cert_gen/gen_certs/cert0.pem -text -noout

for cert in cert_gen/gen_certs/cert*.pem; do
    openssl x509 -in "$cert" -text -noout
done

find cert_gen/gen_certs/ -name "cert*.pem" -exec openssl x509 -in {} -text -noout \;

"""