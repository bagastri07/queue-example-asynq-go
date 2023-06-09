#!/usr/bin/env bash

# mockgen will generate mock interfaces
# for your changes from the master branch into the mock directory
# Example:
#   /pkg/order.go
#
#   Your mock will be generated to
#
#   /mocks/mock_pkg/order_mock.go
#
# Mock usage reference: https://github.com/golang/mock

# Get the current branch name
curBranch=$(git rev-parse --abbrev-ref HEAD)

# Install mockgen (if it hasn't been installed yet)
mockgen_version=$($(go env GOPATH)/bin/mockgen -version)
if [[ ${mockgen_version} == "" ]]; then
    go get -u github.com/golang/mock/...
    go install github.com/golang/mock/mockgen
fi

rm -r mocks

# Get changed files
files=$(find .)

for file in $files; do
    # Exclude certain directories and test files
    if [[ ${file} != "mocks"* && ${file} != "docker"* && ${file} != *"_test"* && ${file} == *".go" ]]; then
        # Generate the destination path for the mock file
        dest="internal/model/mock_model/$(basename ${file} | sed 's/\./_mock./')"

        # Remove the existing mock file if it doesn't exist in the source directory
        if [[ ! -f ${file} && -f ${dest} ]]; then
            rm ${dest}
            git add ${dest}
            continue
        fi

        # Generate the mock interface if the file contains an interface declaration
        if [[ -f ${file} && $(cat ${file} | grep -i ".* interface {" | wc -l) -ne 0 ]]; then
            $(go env GOPATH)/bin/mockgen -source=${file} -destination=${dest}
            git add ${dest}
            echo -e "${dest} is generated"
        fi
    fi
done

# Modify the "dest" to "mock_{name}" folder in the same directory as the generated interface
