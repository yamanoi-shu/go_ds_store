#!/bin/sh

# - test/
#   - test1/
#       - test.txt
#       - .DS_Store
#   - test2/
#       - test.txt
#       - .DS_Store/
#           - .DS_Store

function  generateFiles() {
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

./app ${PWD}/test > test_log.txt

rm -rf test

rm test_log.txt
