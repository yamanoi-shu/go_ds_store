build:
	go build -o app .
clean:
	rm -f app
	rm -rf test
	rm test_actual_log.txt
	rm test_expect_log.txt
test:
	rm -f test_actual_log.txt
	rm -rf test
	go build -o app .
	bash test.sh
