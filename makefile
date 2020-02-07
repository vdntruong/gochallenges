.SILENT: clean

all: testb

testb:
	go test -v -count=1 ./businessman