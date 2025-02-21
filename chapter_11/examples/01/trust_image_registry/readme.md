Golang sample_scaffold_working_structure

<pre>
trust_image_registry % go run main.go
Server running on http://10.0.0.97:8080

% mkdir ./images; touch images/image_00.bin
mkdir: ./images: File exists
</pre>

```
<b># Terminal 1</b>
% go run main.go
Server running on http://10.0.0.97:8080
Notifying issuer issuer@example.com: {
  "image_id": "images/image_00.bin",
  "issuer": "issuer@example.com",
  "signature": "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855",
  "status": "Revoked",
  "cve_url": "https://cve.mitre.org/cgi-bin/cvename.cgi?name=CVE-2025-1234"
}


<b># Terminal 2</b>
trust_image_registry % curl -X POST http://localhost:8080/sign -H "Content-Type: application/json" -d '{
  "image_id": "images/image_00.bin",
  "issuer": "issuer@example.com"
}'

{"image_id":"images/image_00.bin","issuer":"issuer@example.com","signature":"e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855","status":"Valid"}

trust_image_registry % curl -X POST http://localhost:8080/revoke -H "Content-Type: application/json" -d '{
  "image_id": "images/image_XX.bin",
  "cve_url": "https://cve.mitre.org/cgi-bin/cvename.cgi?name=CVE-2025-1234"
}'
Image not found

trust_image_registry % curl -X POST http://localhost:8080/revoke -H "Content-Type: application/json" -d '{
  "image_id": "images/image_00.bin", 
  "cve_url": "https://cve.mitre.org/cgi-bin/cvename.cgi?name=CVE-2025-1234"
}'
{"image_id":"images/image_00.bin","issuer":"issuer@example.com","signature":"e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855","status":"Revoked","cve_url":"https://cve.mitre.org/cgi-bin/cvename.cgi?name=CVE-2025-1234"}
```

<pre>

% ps
:
64067 ttys006    0:00.02 .../T/go-build3608737476/b001/exe/main <--- look for the ps number. Here: 64067
:
% kill -9 64067 <--- use ps number

</pre>
