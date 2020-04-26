# Hide or show a an SVG group (tag name `g`)
# Example call:
# cat in.svg | awk \
#     -v label_regex=mySvgGroupsLabel \
#     -v do_show=0 \
#     -f svg_hide_group.awk \
#     > "out.svg"
#
# Example SVG diff:
#     <g
#        id="group21352"
#        inkscape:label="mySvgGroupsLabel"
#-       style="display:inline">
#+       style="display:none">

BEGIN {
	in_group_tag=0
	id=-1
	inkscape_label=-1
	in_chosen_element=0
	modified_style=0
	if (length(label_regex) == 0) {
		print("Please set the objects inkscape:label (regex) to look for with '-v label_regex=\"my-label\"'") > "/dev/stderr"
		exit(52)
	}
	if (length(do_show) == 0) {
		print("Please set action to take (0=hide, 1=show) with '-v do_show=1'") > "/dev/stderr"
		exit(53)
	}
	if (do_show) {
		display_style_replace="inline"
	} else {
		display_style_replace="none"
	}
}

match($0, /[ \t]+<g( |$).*/) {
	if (!in_group_tag) {
		in_group_tag=1
	}
}

match($0, /[ \t]+id="[^"]*/) {
	if (in_group_tag) {
		id=substr($0, RSTART+1, RLENGTH-1)
		gsub(/[ \t]+id="/, "", id)
	}
}

match($0, /[ \t]+inkscape:label="[^"]*/) {
	if (in_group_tag) {
		inkscape_label=substr($0, RSTART+1, RLENGTH-1)
		gsub(/[ \t]+inkscape:label="/, "", inkscape_label)
		if (match(inkscape_label, label_regex) != 0) {
			in_chosen_element=1
		}
	}
}

match($0, /[ \t]+style="[^"]*/) {
	if (in_chosen_element) {
		style_old=substr($0, RSTART+1, RLENGTH-1)
		style_new=style_old
		if (!sub(/[ \t]display:[0-9a-zA-Z_-]+/, " display:" display_style_replace, style_new)) {
			if (!sub(/"display:[0-9a-zA-Z_-]+/, "\"display:" display_style_replace, style_new)) {
				style_new=style_new " display:" display_style_replace
			}
		}
		sub(style_old, style_new, $0)
		modified_style=1
	}
}

/>/ {
	if (in_group_tag) {
		in_group_tag=0
		if (in_chosen_element && !modified_style) {
			indent=$0
			sub(/[^ \t].*/, "", indent)
			sub(/>/, "\n" indent "style=\"display:" display_style_replace "\">", $0)
		}
	}
	in_chosen_element=0
}

{
	print($0)
}
