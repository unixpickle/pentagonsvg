package main

import (
	"bytes"
	"fmt"
	"math"
	"math/rand"
	"strconv"
)

const PentagonCount = 18
var FillColor = "white"

func main() {
	var svgData bytes.Buffer

	svgData.WriteString(`<?xml version="1.0" encoding="UTF-8" standalone="no"?>`)
	svgData.WriteString(`<svg xmlns="http://www.w3.org/2000/svg"
 xmlns:xlink="http://www.w3.org/1999/xlink" viewBox="0 0 1 1" preserveAspectRatio="xMidYMid slice"
 version="1.1">`)

	svgData.WriteString("<defs><polygon points=\"")
	for i := 0; i < 5; i++ {
		angle := math.Pi * 2 * float64(i) / 5
		if i != 0 {
			svgData.WriteString(" ")
		}
		svgData.WriteString(formatFloat(math.Cos(angle)))
		svgData.WriteString(",")
		svgData.WriteString(formatFloat(math.Sin(angle)))
	}
	svgData.WriteString(`" id="pentagon" fill="`)
	svgData.WriteString(FillColor)
	svgData.WriteString(`" /></defs>`)

	for i := 0; i < PentagonCount; i++ {
		radius := randomRadius()
		opacity := randomOpacity()
		x := rand.Float64()
		y := rand.Float64()
		rotation := rand.Float64() * 360

		svgData.WriteString(`<use xlink:href="#pentagon" fill-opacity="`)

		svgData.WriteString(formatFloat(opacity))
		svgData.WriteString(`"`)

		svgData.WriteString(` transform="`)
		
		svgData.WriteString("translate(")
		svgData.WriteString(formatFloat(x))
		svgData.WriteString(" ")
		svgData.WriteString(formatFloat(y))
		svgData.WriteString(")")
		
		svgData.WriteString(" scale(")
		svgData.WriteString(formatFloat(radius))
		svgData.WriteString(")")

		svgData.WriteString(" rotate(")
		svgData.WriteString(formatFloat(rotation))
		svgData.WriteString(")")

		svgData.WriteString("\" />")
	}

	svgData.WriteString(`</svg>`)

	fmt.Println(svgData.String())
}

func formatFloat(f float64) string {
	return strconv.FormatFloat(f, 'f', -1, 64)
}

func randomOpacity() float64 {
	return rand.Float64()*0.22 + 0.02
}

func randomRadius() float64 {
	return 0.05 + (math.Pow(rand.Float64(), 15)+1)*0.075
}
