#!/bin/sh
# Runs all the Escher tutorials.
# Requires the `escher` command available on the PATH.

which escher > /dev/null
if [ $? -ne 0 ]
then
	>&2 echo "Error: Could not find 'escher' in PATH"
	exit 1
fi

src_dir="$GOPATH/src/github.com/gocircuit/escher/src/"
tutorial_circuits="ShowIndex HelloWorld Break Debug Exec File TextMerge"

for circuit in $tutorial_circuits
do
	# run each tutorial for at most 2 seconds
	echo
	echo
	echo "################################################################################"
	echo "### Running Escher tutorial $circuit ..."
	echo "--------------------------------------------------------------------------------"
	timeout  --foreground --kill-after=2 --signal=SIGINT 3s \
		escher -src "$src_dir" "*tutorial.${circuit}Main"
	echo
	echo "################################################################################"
done

