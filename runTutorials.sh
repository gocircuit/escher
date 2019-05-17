#!/bin/sh
# Runs all the Escher tutorials.
# NOTE Requires the `escher` command available on the PATH.
# NOTE Needs to be run from the escher repo root

which escher > /dev/null
if [ $? -ne 0 ]
then
	>&2 echo "Error: Could not find 'escher' in PATH"
	exit 1
fi

# NOTE We do not use this path,
#      even though it would make the script position independent,
#      because it would break (or worse: run the wrong code)
#      when working on a fork of the repository.
#src_dir="$GOPATH/src/github.com/gocircuit/escher/src/"
# This way of defning src_dir ensure that we can use relative paths,
# while the script may still be called from anywhere,
# as long as the tutorials are to be found
# in the correct relative path within the escher repo.
call_path="`dirname $0`"
repo_root="`cd $call_path; pwd`"
src_dir="$repo_root/src"
tutorials_dir="$src_dir/tutorial"

find "$tutorials_dir" -regex '.*/[A-Z][^/]*.escher' > /dev/null
if [ $? -ne 0 ]
then
	>&2 echo "Error: No tutorials found in '`pwd`$tutorials_dir'."
	exit 2
fi

tutorial_circuits=`find "$tutorials_dir" -regex '.*/[A-Z][^/]*.escher' \
	| xargs basename --multiple --suffix '.escher'`
export ESCHER="$src_dir"

for circuit in $tutorial_circuits
do
	echo
	echo
	echo "################################################################################"
	echo "### Running Escher tutorial $circuit ..."
	echo "--------------------------------------------------------------------------------"
	## run each tutorial for at most 2 seconds
	#timeout  --foreground --kill-after=2 --signal=SIGINT 3s \
		escher "*tutorial.${circuit}Main"
	echo
	echo "################################################################################"
done

