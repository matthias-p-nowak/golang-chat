

# create the right keys for the push notification
[confirm]
vapid:
    openssl ecparam -name prime256v1 -genkey -noout -out secrets/vapid.private.key
    openssl ec -in secrets/vapid.private.key -pubout -out secrets/vapid.public.key

[working-directory: 'server']
go-mods:
    #!/bin/env bash
    go get ./...
    go mod tidy