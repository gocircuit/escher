#!/usr/bin/env bash
# Builds the Escher handbook.
# NOTE
# * Requires the `escher` command available on the PATH.
# * Requires the `inkscape` command available on the PATH.
# * Requires the AWK script `svg_hide_group.awk`,
#   which should already be in the same directory as this script.

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
#src_dir="$GOPATH/src/github.com/hoijui/escher/src/"
# This way of defning src_dir ensures that we can use relative paths,
# while the script may still be called from anywhere,
# as long as the sources are to be found
# under the same relative path within the escher repo.
escher_src_dir="$repo_root/src"
svg_hide="$script_dir/svg_hide_group.awk"
src_dir="$escher_src_dir/handbook"
export ESCHER="$escher_src_dir"

rel_out_dir="${1:-}"
if [ "$rel_out_dir" = "" ]
then
	>&2 echo "Please supply a directory path to build the handbook in."
	exit 1
fi
out_dir=$(mkdir -p "$rel_out_dir"; cd "$rel_out_dir"; pwd)
if [ "$out_dir" = "$repo_root" ]
then
	>&2 echo "Please supply an output directory different then the escher repo root."
	exit 1
fi
if [ "$out_dir" = "$src_dir" ]
then
	>&2 echo "Please supply an output directory different then the hanbook source directory."
	exit 1
fi

echo "Building the handbook in '$out_dir'; press Ctrl+C to abort"
echo "Waiting 3 seconds ..."
sleep 3

echo
echo "Removing previous build artifacts ..."
rm -Rf "$out_dir/img"
rm -Rf "$out_dir/css"
rm -Rf "$out_dir/pdf"
rm -f "$out_dir/"*.html

echo
echo "Copying assets from sources to output directory ..."
echo -e "\tcss ..."
cp -r "$src_dir/css" "$out_dir/"
echo -e "\timg ..."
cp -r "$src_dir/img" "$out_dir/"
echo -e "\tpdf ..."
cp -r "$src_dir/pdf" "$out_dir/"
rm -f "$out_dir/css/font/.gitignore"

echo
echo "Generating different views of a \"packed\" SVG (using AWK) ..."
svg_in="$src_dir/img/circuit.svg"

svg_out="$out_dir/img/circuit-parts-generated.svg"
echo -e "\t\"$svg_in\" --> \"$svg_out\""
cat "$svg_in" | awk \
	-v label_regex=labels-instances -v do_show=0 -f "$svg_hide" \
	> "$svg_out"

svg_out="$out_dir/img/circuit-instances-generated.svg"
echo -e "\t\"$svg_in\" --> \"$svg_out\""
cat "$svg_in" | awk \
	-v label_regex=labels-parts -v do_show=0 -f "$svg_hide" \
	> "$svg_out"

svg_out="$out_dir/img/circuit-raw-generated.svg"
echo -e "\t\"$svg_in\" --> \"$svg_out\""
cat "$svg_in" | awk \
	-v label_regex='labels-.*' -v do_show=0 -f "$svg_hide" \
	> "$svg_out"

echo
echo "Generate widely-compatible versions of our SVG images (using inkscape) ..."
# This makes the generated versions be
# not just Inkscape compatible,
# but display correctly everywhere.
for svg_in in "$out_dir/img/"*-generated.svg
do
	svg_out=$(echo "$svg_in" | sed -e 's|-generated\.svg$|-plain-generated.svg|')
	echo -e "\t\"$svg_in\" --> \"$svg_out\""
	inkscape --without-gui "$svg_in" --export-text-to-path --export-plain-svg "$svg_out"
	rm "$svg_in"
done

echo
echo "Convert SVGs to PNGs (using inkscape) ..."
for svg in "$out_dir/img/"*.svg
do
	png=$(echo "$svg" | sed -e 's|\(-plain\)\?\(-generated\)\?\.svg$|.png|')
	echo -e "\t\"$svg\" --> \"$png\""
	inkscape --without-gui "$svg" --export-png "$png" > /dev/null
done

echo
echo "Building the handbook (using escher) ..."
escher "*handbook.main"

echo
echo "done."
