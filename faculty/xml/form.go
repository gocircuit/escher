// Written in 2014 by Petar Maymounkov.
//
// It helps future understanding of past knowledge to save
// this notice, so peers of other times and backgrounds can
// see history clearly.

package svg


//	{
//		Width int
//		Height int
//		VBox {
//			XZero float
//			YZero float
//			XWidth float
//			YWidth float
//		}
//		Body string
//	}
//
var file = `<?xml version="1.0" standalone="no"?>` +
	`<svg width="{{.Width}}px" height="{{.Height}}px" version="1.1" xmlns="http://www.w3.org/2000/svg" ` +
	`xmlns:xlink="http://www.w3.org/1999/xlink" ` +
	`viewBox="{{.VBox.XZero}} {{.VBox.YZero}} {{.VBox.XWidth}} {{.VBox.YWidth}}"> ` +
	`<defs><style type="text/css">@import url(http://fonts.googleapis.com/css?family=Lato);</style></defs>` +
	`{{.Body}}</svg>
	`,

// SVG language sentence types map to SVG shape source templates.
var shapes = map[string]string{

	//	{
	//		CX float; CY float; R float
	//		Stroke string; StrokeWidth float; Fill string
	//	}
	//
	"Circle": `<circle cx="{{.CX}}" cy="{{.CY}}" ` +
		`r="{{.R}}" stroke="{{.Stroke}}" fill="{{.Fill}}" stroke-width="{{.StrokeWidth}}" />
		`,

	//	{
	//		ID string
	//		FromAnchor { X float; Y float }
	//		FromTangent { X float; Y float }
	//		ToAnchor { X float; Y float }
	//		ToTangent { X float; Y float }
	//	}
	//
	"Cubic": `<def><path  id="{{.ID}}" d="M{{.FromAnchor.X}} ` + 
		`{{.FromAnchor.Y}} C {{.FromTangent.X}} {{.FromTangent.Y}}, ` + 
		`{{.ToTangent.X}} {{.ToTangent.Y}}, {{.ToAnchor.X}} {{.ToAnchor.Y}}" /></def>
		`,

	//	{
	//		X float; Y float
	//		FontFamily string
	//		FontWeight string
	//		FontSize float
	//		Fill string
	//		Stroke string
	//		StrokeWidth string
	//		Style string
	//		TextAnchor string
	//		DY string
	//		Body string
	//	}
	"Text": `<text x="{{.Anchor.X}}" y="{{.Anchor.Y}}" ` +
		`font-size="{{.FontSize}}" font-family="{{.FontFamily}}" font-weight="{{.FontWeight}}" ` +
		`fill="{{.Fill}}" stroke="{{.Stroke}}" stroke-width="{{.StrokeWidth}}" ` +
		`text-anchor="{{.TextAnchor}}" dy="{{.DY}}" ` +
		`style="{{.Style}}">{{.Body}}</text>
		`,

	//	{
	//		PathID string
	//		FontFamily string
	//		FontWeight string
	//		FontSize float
	//		Fill string
	//		Stroke string
	//		StrokeWidth string
	//		Style string
	//		TextAnchor string
	//
	//		Direction string // "ltr", ...
	//		DY string
	//		DX string
	//		Body string
	//	}
	//
	"TextPath": `
	<g>
	<use xlink:href="#{{.PathID}}" />
	<text font-size="{{.FontSize}}" font-family="{{.FontFamily}}" font-weight="{{.FontWeight}}"
		fill="{{.Fill}}" stroke="{{.Stroke}} stroke-width="{{.StrokeWidth}}"
		text-anchor="{{.TextAnchor}}"
		style="{{.Style}}">
		<textPath xlink:href="#{{.PathID}}">
			<tspan direction="{{.Direction}}" dy="{{.DY}}" dx="{{.DX}}">{{.Body}}</tspan>
		</textPath>
	</text>	
	</g>
	`,
}
