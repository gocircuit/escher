#!/bin/sh
# *nix script to generate the HTML handbook
# from the sources in "github.com/gocircuit/escher/src/handbook".

src_dir="$GOPATH/src/github.com/gocircuit/escher/src/"

escher -src "$src_dir" "*handbook.main"

cp -r "$src_dir/handbook/css" ./
cp -r "$src_dir/handbook/img" ./
cp -r "$src_dir/handbook/pdf" ./

rm -f css/font/.gitignore

