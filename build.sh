CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -gcflags=-trimpath=$GOPATH -asmflags=-trimpath=$GOPATH -ldflags "-w -s" -o scf_bootstrap
 zip -r scf_bootstrap.zip scf_bootstrap tools/
 rm scf_bootstrap