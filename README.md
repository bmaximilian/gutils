# G-Utils
Some CLI Tools written in Go.
> Mainly for private purposes

## Jira CLI instructions
```bash
# Convert PKCS12 Cert to pem key and cert
openssl pkcs12 -in <username>.p12 -nokeys -out <username>.cer.pem
openssl pkcs12 -in <username>.p12 -nodes -nocerts -out <username>.key.pem

# Add path of <username>.cer.pem (as jira.cert) and <username>.key.pem (as jira.key) to config file

# Generate an auth token for jira
gutils jira token get -u <jira_username> -p <jira_password>

# Add token as jira.token to config file

# run your "gutils jira ..." commands
```
