// Package ascii_donut
// Created by Teocci.
// Author: teocci@yandex.com on 2021-09-01
// Based on [Donut math: how donut.c works](https://www.a1k0n.net/2011/07/20/donut-math.html)
package main

import (
	"fmt"
	"github.com/teocci/go-csv-samples/src/iolive"
	"math"
	"time"
)

const (
	luminance    = ".,-~:;=!*#$@"
	thetaSpacing = 0.07
	phiSpacing   = 0.02

	screenWidth  = 70
	screenHeight = 30

	R1 = 1
	R2 = 2
	K2 = 5
)

// K1 based on screen size: the maximum x-distance occurs
// roughly at the edge of the torus, which is at x=R1+R2, z=0.  we
// want that to be displaced 3/8ths of the width of the screen, which
// is 3/4th of the way from the center to the side of the screen.
// screenWidth*3/8 = K1*(R1+R2)/(K2+0)
// screenWidth*K2*3/(8*(R1+R2)) = K1
const K1 = screenWidth * K2 * 3 / (18 * (R1 + R2))

var (
	output = make([][]rune, screenWidth)
	zBuff  = make([][]float64, screenWidth)

	a, b float64

	writer *iolive.Writer
)

func main() {
	writer = iolive.New()
	// start listening for updates and render
	writer.Start()

	for i := float64(0); i <= 100; i++ {
		a, b = 0+0.7*i, 0+0.3*i
		renderFrame(a, b)
		time.Sleep(time.Millisecond * 80)
	}


	writer.Stop() // flush and stop rendering
}

func renderFrame(a float64, b float64) {
	// precompute sines and cosines of a and b
	cosA, sinA := math.Cos(a), math.Sin(a)
	cosB, sinB := math.Cos(b), math.Sin(b)

	for i := range output {
		output[i] = make([]rune, screenHeight)
	}
	for i := range output {
		for _ = range output[i] {
			output[i] = append(output[i], ' ')
		}
	}

	for i := range zBuff {
		zBuff[i] = make([]float64, screenHeight)
	}

	//_, _ = fmt.Fprintf(writer, "\x1b[2J")
	// theta goes around the cross-sectional circle of a torus
	for theta := float64(0); theta < 2*math.Pi; theta += thetaSpacing {
		// precompute sines and cosines of theta
		costheta := math.Cos(float64(theta))
		sintheta := math.Sin(float64(theta))

		// phi goes around the center of revolution of a torus
		for phi := float64(0); phi < 2*math.Pi; phi += phiSpacing {
			// precompute sines and cosines of phi
			cosphi := math.Cos(phi)
			sinphi := math.Sin(phi)

			// the x,y coordinate of the circle, before revolving (factored
			// out of the above equations)
			circlex := R2 + R1*costheta
			circley := R1 * sintheta

			// final 3D (x,y,z) coordinate after rotations, directly from
			// our math above
			x := circlex*(cosB*cosphi+sinA*sinB*sinphi) - circley*cosA*sinB
			y := circlex*(sinB*cosphi-sinA*cosB*sinphi) + circley*cosA*cosB
			z := K2 + cosA*circlex*sinphi + circley*sinA
			ooz := 1 / z // "one over z"

			// x and y projection.  note that y is negated here, because y
			// goes up in 3D space but down on 2D displays.
			xp := (int)(screenWidth/2 + K1*ooz*x)
			yp := (int)(screenHeight/2 - K1*ooz*y)
			if xp < 0 {
				xp = 0
			}
			if xp > screenWidth-1 {
				xp = screenWidth - 1
			}
			if yp < 0 {
				yp = 0
			}
			if yp > screenHeight-1 {
				yp = screenHeight - 1
			}

			// calculate luminance.  ugly, but correct.
			l := cosphi*costheta*sinB - cosA*costheta*sinphi - sinA*sintheta + cosB*(cosA*sintheta-costheta*sinA*sinphi)
			// l ranges from -sqrt(2) to +sqrt(2).  If it's < 0, the surface
			// is pointing away from us, so we won't bother trying to plot it.
			if l > 0 {
				// test against the z-buffer.  larger 1/z means the pixel is
				// closer to the viewer than what's already plotted.
				//fmt.Printf("(%d, %d)\n", xp, yp)
				if ooz > zBuff[xp][yp] {
					zBuff[xp][yp] = ooz
					luminanceIndex := int(l * 8)
					// luminanceIndex is now in the range 0..11 (8*sqrt(2) = 11.3)
					// now we lookup the character corresponding to the
					// luminance and plot it in our output:

					//_, _ = fmt.Fprintf(writer, "(%d, %d) | %d\n", xp, yp, luminanceIndex)
					output[xp][yp] = rune(luminance[luminanceIndex])
				}
			}
		}
	}
	//_, _ = fmt.Fprintf(writer, "(luminance len: %d\n", len(luminance))
	// now, dump output[] to the screen.
	// bring cursor to "home" location, in just about any currently-used
	// terminal emulation mode
	//_, _ = fmt.Fprintf(writer, "\x1b[H")
	for j := 0; j < screenHeight; j++ {
		for i := 0; i < screenWidth; i++ {
			_, _ = fmt.Fprintf(writer, "%c", output[i][j])
		}
		_, _ = fmt.Fprintf(writer, "%c", '\n')
	}
}
