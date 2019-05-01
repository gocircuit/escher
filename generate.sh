#!/bin/bash
# *nix script to generate the HTML handbook
# from the sources in "github.com/gocircuit/escher/src/handbook".

commit="false"
if [ "$1" = "-c" ]
then
	commit="true"
fi

escher_repo="$GOPATH/src/github.com/gocircuit/escher"
src_dir="$escher_repo/src"

# Generate the static HTML pages
escher -src "$src_dir" "*handbook.main"

# Copy over the resources
cp -r "$src_dir/handbook/css" ./
cp -r "$src_dir/handbook/img" ./
cp -r "$src_dir/handbook/pdf" ./
rm -f css/font/.gitignore

# Create a git commit, if requested, and if there are local changes
local_changes=$(git status --porcelain)
if [ "$commit" = "true" -a "$local_changes" != "" ]
then
	echo "Committing ..."
	git add --all
	branch_name=$(cd "$escher_repo" ; git rev-parse --abbrev-ref HEAD)
	remote_and_branch_name=$(cd "$escher_repo" ; git for-each-ref --format='%(upstream:short)' $(git symbolic-ref -q HEAD))
	commit_description=$(cd "$escher_repo" ; git describe --tags --always)
	#commit_date=$(cd "$escher_repo" ; git log -1 --format="%at" | xargs -I{} date -d @{} +"%d. %B %Y %H:%M:%S")
	commit_date=$(cd "$escher_repo" ; git log -1 --format="%at" | xargs -I{} date -d @{} +"%d. %B %Y")
	#commit_time=$(cd "$escher_repo" ; git log -1 --format="%at" | xargs -I{} date -d @{} +"%H:%M:%S")
	git commit --quiet --message "latest as of $commit_date, generated from $remote_and_branch_name $commit_description" && \
		git push --quiet && \
		echo "Pushed!" || \
		echo "Failed!"
fi

