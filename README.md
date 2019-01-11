# totp-gen
Generate TOTP token

## Usage

1. Secret as parameter
```bash
totp-gen -secret <BASE32 secret>
```

2. Secret from stdin
```bash
# Using macOS keychain 
security find-generic-password -a ${USER} -s TOTP/github.com -w | totp-gen

# Using pass
pass TOTP/github.com | totp-gen
```
