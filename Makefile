PROJECT=github.com/iwittkau/ssh-select/cmd/sshs

all: install

install:
	go install ${PROJECT}

build-amd64:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=1 go build -o ./dist/sshs_amd64 ${PROJECT}

build-armv7:
	GOOS=linux GOARCH=arm GOARM=7 go build -o ./dist/sshs_armv7 ${PROJECT}

build-armv6:
	GOOS=linux GOARCH=arm GOARM=6 go build -o ./dist/sshs_armv6 ${PROJECT}

build-armv5:
	GOOS=linux GOARCH=arm GOARM=5 go build -o ./dist/sshs_armv5 ${PROJECT}