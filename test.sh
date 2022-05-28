#!/bin/bash

# - test/
#   - test1/
#       - test.txt
#       - .DS_Store
#   - test2/
#       - test.txt
#       - .DS_Store/
#           - .DS_Store

function generateFiles() {
    mkdir test

    cd test

    touch .DS_Store

    touch test.txt

    mkdir test1

    cd test1

    touch .DS_Store

    cd ..

    mkdir test2

    cd test2

    touch test.txt
    mkdir .DS_Store

    cd .DS_Store

    touch .DS_Store

    cd ../../../

    return
}

export LOG_FILE="true"

generateFiles

./app ${PWD}/test

sed -e "s|!pwd!|$PWD|g" ./testdata/test_expect_log_template.txt > ./test_expect_log.txt

diff ./test_actual_log.txt ./test_expect_log.txt

if [ $? -eq 1 ]; then
    echo "Faild....."
    exit 1
elif [ $? -eq 1 ]; then
    echo "Success!"

fi

rm -rf test

rm test_actual_log.txt
rm test_expect_log.txt

