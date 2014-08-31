/*
	Requirements:

		npm install dagre underscore

	Run with:

		node layout.js < circuit.json

	This program computes a drawing layout for a circuit topology specified in JSON format:

		{
			"Peer": { "A":{}, "B":{}, "C":{}},
			"Match": {
				"F": ["A", "B"],
				"G": ["C", "A"],
				"H": ["C", "B"]
			},
			"Valve": { "X":{}, "Y":{} }
		}
*/

var dagre = require("dagre"),
	_ = require("underscore");

// Read JSON graph definition from standard input.
process.stdin.setEncoding('utf8');
var source = "";
process.stdin.on('readable', function() {
	var chunk = process.stdin.read();
	if (chunk !== null) {
		source += chunk;
	}
});
process.stdin.on('end', function() {
	var data = JSON.parse(source);
	console.log(JSON.stringify(render(data)));
});

function renderWithDagre(data) {
	var g = new dagre.Digraph(); // Create a new directed graph
	_.each(data["Peer"], function(value, key) { // Add nodes
		g.addNode(key, {width: 100, height: 100});
	});
	_.each(data["Match"], function(value, key) { // Add nodes
		g.addEdge(key, value[0], value[1]);
	});
	var layout = dagre.layout() // Compute layout
				.debugLevel(0)
				.rankDir("LR")
				.run(g);
	var y = {
		Node: {},
		Edge: {}
	};
	layout.eachNode(function(u, v) { 
		y.Node[u] = {
			X: v.x,
			Y: v.y,
		}
	});
	layout.eachEdge(function(e, u, w, v) { 
		y.Edge[e] = {
			Points: v.points
		}
	});
	return y;
}
