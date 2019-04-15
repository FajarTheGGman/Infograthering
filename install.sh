echo "[!] Installing Package...."
apt-get install golang -y
go get -u github.com/mgutz/ansi
go run info.go
