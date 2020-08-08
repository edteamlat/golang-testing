#!/bin/sh

for d in $(go list ./...); do
	echo "Testeando el paquete $d"
	go test -v $d
done
