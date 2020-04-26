#!/usr/bin/env bash
# Builds the escher handbook.
# NOTE
# * Requires the `escher` command to be available on the PATH.
# * Requires the `inkscape` command to be available on the PATH.
# * Requires the AWK script `svg_hide_group.awk`,
#   which should already be in the same directory as this script.

# Exit immediately on each error and unset variable;
# see: https://vaneyckt.io/posts/safer_bash_scripts_with_set_euxo_pipefail/
#set -Eeuo pipefail
set -Eeu

script_dir=$(dirname "$(readlink -f "${BASH_SOURCE[0]}")")

svg_hide="$script_dir/svg_hide_group.awk"
#export ESCHER="$GOPATH/src/github.com/gocircuit/escher/src/"
export ESCHER="$script_dir/../src/"

cd "$ESCHER/handbook"

echo
echo "Removing previous build artifacts ..."
rm -f img/*-generated.*
rm -f img/*.png

echo
echo "Generating different views of a \"packed\" SVG (using AWK) ..."
svg_in="img/circuit.svg"

svg_out="img/circuit-parts-generated.svg"
echo -e "\t\"$svg_in\" --> \"$svg_out\""
cat "img/circuit.svg" | awk \
	-v label_regex=labels-instances -v do_show=0 -f "$svg_hide" \
	> "$svg_out"

svg_out="img/circuit-instances-generated.svg"
echo -e "\t\"$svg_in\" --> \"$svg_out\""
cat "img/circuit.svg" | awk \
	-v label_regex=labels-parts -v do_show=0 -f "$svg_hide" \
	> "$svg_out"

svg_out="img/circuit-raw-generated.svg"
echo -e "\t\"$svg_in\" --> \"$svg_out\""
cat "img/circuit.svg" | awk \
	-v label_regex='labels-.*' -v do_show=0 -f "$svg_hide" \
	> "$svg_out"

echo
echo "Generate widely-compatible versions of our SVG images (using inkscape) ..."
# This makes the generated versions be
# not just Inkscape compatible,
# but display correctly everywhere.
for svg_in in img/*-generated.svg
do
	svg_out=$(echo "$svg_in" | sed -e 's|-generated\.svg$|-plain-generated.svg|')
	echo -e "\t\"$svg_in\" --> \"$svg_out\""
	inkscape --without-gui "$svg_in" --export-text-to-path --export-plain-svg "$svg_out"
	rm "$svg_in"
done

echo
echo "Convert SVGs to PNGs (using inkscape) ..."
for svg in img/*.svg
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
