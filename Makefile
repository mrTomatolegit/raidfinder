NAME=raidfinder
VERSION=1.0.0

WINCOMPDIR=bin/${NAME}-${VERSION}_windows
LINCOMPDIR=bin/${NAME}-${VERSION}_linux
DARCOMPDIR=bin/${NAME}-${VERSION}_darwin

OSFLAG :=
ifeq ($(OS),Windows_NT)
	OSFLAG += windows
else
	OSFLAG += darlin
endif

build:
	go build -o ./bin/${NAME}.exe ./cmd/raidfinder/main.go

run:
	go run main.go

wincompile:
	rm -rf ${WINCOMPDIR} ${WINCOMPDIR}.zip
	mkdir ${WINCOMPDIR}
	GOOS=windows GOARCH=amd64 go build -o ${WINCOMPDIR}/raidfinder.exe ./cmd/raidfinder/main.go
	touch ${WINCOMPDIR}/raidlist.txt
	touch ${WINCOMPDIR}/noraidlist.txt
ifeq ($(OSFLAG),windows)
	7z a -r ${WINCOMPDIR}.zip ./${WINCOMPDIR}/*
else
	zip -r ${WINCOMPDIR}.zip ./${WINCOMPDIR}/*
endif

lincompile:
	rm -rf ${LINCOMPDIR} ${LINCOMPDIR}.zip
	mkdir ${LINCOMPDIR}
	GOOS=linux GOARCH=amd64 go build -o ${LINCOMPDIR}/raidfinder ./cmd/raidfinder/main.go
	touch ${LINCOMPDIR}/raidlist.txt
	touch ${LINCOMPDIR}/noraidlist.txt
ifeq ($(OSFLAG),windows)
	7z a -r ${LINCOMPDIR}.zip ./${LINCOMPDIR}/*
else
	zip -r ${LINCOMPDIR}.zip ./${LINCOMPDIR}/*
endif

darcompile:
	rm -rf ${DARCOMPDIR} ${DARCOMPDIR}.zip
	mkdir ${DARCOMPDIR}
	GOOS=darwin GOARCH=amd64 go build -o ${DARCOMPDIR}/raidfinder ./cmd/raidfinder/main.go
	touch ${DARCOMPDIR}/raidlist.txt
	touch ${DARCOMPDIR}/noraidlist.txt
ifeq ($(OSFLAG),windows)
	7z a -r ${DARCOMPDIR}.zip ./${DARCOMPDIR}/*
else
	zip -r ${DARCOMPDIR}.zip ./${DARCOMPDIR}/*
endif

compile:
	@echo "Compiling for every OS and Platform"
	make wincompile
	make lincompile
	make darcompile
	rm -rf ${WINCOMPDIR} ${LINCOMPDIR} ${DARCOMPDIR}

all: build
