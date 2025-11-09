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
    ./api-bouncer --create-project --project-name $NAME --project-port $PORT
    check_command_successful "Fail"
}

function test_create_request_get() {
    echo "Creating request GET /v1/1_test_get"
    ./api-bouncer --create-request --project-name $NAME --request-method GET --request-url /v1/1_test_get
    check_command_successful "Error to create a GET request"
}

function test_create_request_post() {
    echo "Creating POST request /v1/1_test_post"
    ./api-bouncer --create-request --project-name $NAME --request-method POST --request-url /v1/1_test_post
    check_command_successful "Error to create a POST request"
}

function test_create_get_response() {
    echo "Creating a response to /v1/1_test_get"
    ./api-bouncer --create-response --project-name $NAME --request-method GET --request-url /v1/1_test_get --response-status-code 200 --response-mime application/json --response-body "{\"test_get\":\"ok\"}" --response-identifier testing_1
    check_command_successful "Error inserting response to /v1/1_test_get"
}

function test_crete_post_response() {
    echo "Creating a response /v1/1_test_post"
    ./api-bouncer --create-response --project-name $NAME --request-method POST --request-url /v1/1_test_post --response-status-code 200 --response-mime application/json --response-body "{\"test_post\":\"ok\"}" --response-identifier testing_2
    check_command_successful
}

# Setup
build_project
# Test creating a project
test_create_project
test_create_request_get
test_create_request_post
test_create_get_response
test_crete_post_response

echo "Manual testing finished"