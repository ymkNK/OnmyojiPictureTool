# Binary name
BINARY=otool
OUTPUT=./output/bin/
GOBUILD=go build -o ${OUTPUT}${BINARY}
GOCLEAN=go clean
RMTARGZ=rm -rf ${OUTPUT}*.gz
VERSION=0.0.1

release:
	# Clean
	$(GOCLEAN)
	$(RMTARGZ)
	# Build for arm
	$(GOCLEAN)
	CGO_ENABLED=0 GOOS=linux GOARCH=arm64 $(GOBUILD)
	tar czvf ${OUTPUT}${BINARY}-arm64-${VERSION}.tar.gz ${OUTPUT}${BINARY}
	# Build for linux
	$(GOCLEAN)
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD)
	tar czvf ${OUTPUT}${BINARY}-linux64-${VERSION}.tar.gz ${OUTPUT}${BINARY}
	# Build for mac
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 $(GOBUILD)
	tar czvf ${OUTPUT}${BINARY}-mac64-${VERSION}.tar.gz ${OUTPUT}${BINARY}
	# Build for win
	$(GOCLEAN)
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 $(GOBUILD).exe
	tar czvf ${OUTPUT}${BINARY}-win64-${VERSION}.tar.gz ${OUTPUT}${BINARY}.exe
	$(GOCLEAN)

	sh build.sh