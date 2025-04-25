#!/bin/bash

echo "Manual Test Suite"
NAME=$1
PORT=$2

if [ $# -lt 2 ]; then
    echo "Provide the test name an port"
    exit 1
fi

function check_command_successful () {
    if [ $? -ne 0 ]; then
        echo $1 
        exit 1
    fi
}

function build_project () {
    go build
    check_command_successful "Build failure"
}

function test_create_project() {
    echo "Creating project $NAME"
    ./api-bouncer --create-project --project-name $NAME --port $PORT
    check_command_successful "Fail"
}

function test_create_request() {
    
}

# Setup
build_project
# Test creating a project
test_create_project
