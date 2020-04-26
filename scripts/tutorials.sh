#!/usr/bin/env bash
# Runs all the Escher tutorials.
# NOTE
# * Requires the `escher` command available on the PATH.

# Exit immediately on each error and unset variable;
# see: https://vaneyckt.io/posts/safer_bash_scripts_with_set_euxo_pipefail/
#set -Eeuo pipefail
set -Eeu

script_dir=$(dirname "$(readlink -f "${BASH_SOURCE[0]}")")
repo_root="$(cd $script_dir; cd ..; pwd)"
# NOTE We do not use this path,
#      even though it would make the script position independent,
#      because it would break (or worse: run the wrong code)
#      when working on a fork of the repository.
#src_dir="$GOPATH/src/github.com/gocircuit/escher/src/"
# This way of defning src_dir ensures that we can use relative paths,
# while the script may still be called from anywhere,
# as long as the sources are to be found
# under the same relative path within the escher repo.
src_dir="$repo_root/src"
tutorials_dir="$src_dir/tutorial"

which escher > /dev/null
if [ $? -ne 0 ]
then
	>&2 echo "Error: Could not find 'escher' in PATH"
	exit 1
fi

cd "$repo_root"

if [ "${1:-}" = "" ]
then
	find "$tutorials_dir" -regex '.*/[A-Z][^/]*.escher' > /dev/null
	if [ $? -ne 0 ]
	then
		>&2 echo "Error: No tutorials found in '$(pwd)/$tutorials_dir'."
		exit 2
	fi

	tutorial_circuits=$(find "$tutorials_dir" -regex '.*/[A-Z][^/]*.escher' \
		| xargs basename --multiple --suffix '.escher')
else
	tutorial_circuits="$1"
fi
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

