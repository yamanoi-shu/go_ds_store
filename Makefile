build:
	go build -o app .
clean:
	rm app
	rm  test_log.txt
	rm -rf test
test:
	go build -o app .
	sh test.sh
