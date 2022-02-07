build:
	go build -o app .
clean:
	rm -f app
	rm -f test_log.txt
	rm -rf test
test:
	go build -o app .
	sh test.sh
