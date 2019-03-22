#!/bin/sh
# *nix script to generate the HTML handbook
# from the sources in "github.com/gocircuit/escher/src/handbook".

escher -src "$GOPATH/src/github.com/gocircuit/escher/src/" "*handbook.main"

