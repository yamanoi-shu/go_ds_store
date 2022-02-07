build:
	go build -o app .
clean:
	rm -f app
	rm -rf test
test:
	rm -r test_actual_log.txt
	rm -rf test
	go build -o app .
	sh test.sh
