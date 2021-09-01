package main

import (
	"fmt"
	"math"
)

const (
	luminance     = ".,-~:;=!*#$@"
	theta_spacing = 0.07
	phi_spacing   = 0.02

	screen_width  = 210
	screen_height = 90

	R1 = 1
	R2 = 2
	K2 = 5
)

// K1 based on screen size: the maximum x-distance occurs
// roughly at the edge of the torus, which is at x=R1+R2, z=0.  we
// want that to be displaced 3/8ths of the width of the screen, which
// is 3/4th of the way from the center to the side of the screen.
// screen_width*3/8 = K1*(R1+R2)/(K2+0)
// screen_width*K2*3/(8*(R1+R2)) = K1
const K1 = screen_width * K2 * 3 / (20 * (R1 + R2))

var (
	output = make([][]rune, screen_width)
	zBuff  = make([][]float64, screen_width)
)

func main() {
	renderFrame(200, 200)
}

func renderFrame(a float64, b float64) {
	// precompute sines and cosines of a and b
	cosA, sinA := math.Cos(a), math.Sin(a)
	cosB, sinB := math.Cos(b), math.Sin(b)

	for i := range output {
		output[i] = make([]rune, screen_height)
	}
	for i := range output {
		for _ = range output[i] {
			output[i] = append(output[i], ' ')
		}
	}

	for i := range zBuff {
		zBuff[i] = make([]float64, screen_height)
	}

	// theta goes around the cross-sectional circle of a torus
	for theta := float64(0); theta < 2*math.Pi; theta += theta_spacing {
		// precompute sines and cosines of theta
		costheta := math.Cos(float64(theta))
		sintheta := math.Sin(float64(theta))

		// phi goes around the center of revolution of a torus
		for phi := float64(0); phi < 2*math.Pi; phi += phi_spacing {
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
			xp := (int)(screen_width/2 + K1*ooz*x)
			yp := (int)(screen_height/2 - K1*ooz*y)
			if xp < 0 {
				xp = 0
			}
			if xp > screen_width-1 {
				xp = screen_width - 1
			}
			if yp < 0 {
				yp = 0
			}
			if yp > screen_height-1 {
				yp = screen_height - 1
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
					output[xp][yp] = rune(luminance[luminanceIndex])
				}
			}
		}
	}
	// now, dump output[] to the screen.
	// bring cursor to "home" location, in just about any currently-used
	// terminal emulation mode
	//fmt.Printf("\x1b[H")
	for j := 0; j < screen_height; j++ {
		for i := 0; i < screen_width; i++ {
			fmt.Printf("%c", output[i][j])
		}
		fmt.Printf("%c", '\n')
	}
}
