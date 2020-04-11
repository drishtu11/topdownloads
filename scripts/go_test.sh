#!/bin/sh
set -e # Stop script from running if there are any errors

cd httprequests
go test -v
cd ..

cd pqheap
go test -v
cd ..