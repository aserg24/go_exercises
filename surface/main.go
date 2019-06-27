package main

import (
	"fmt"
	"log"
	"math"
	"net/http"
)

const (
	width, height = 600, 320            // canvas size in pixels
	cells         = 100                 // number of grid cells
	xyrange       = 30.0                // axis ranges (-xyrange..+xyrange)
	xyscale       = width / 2 / xyrange // pixels per x or y unit
	zscale        = height * 0.4        // pixels per z unit
	angle         = math.Pi / 6         // angle of x, y axes (=30°)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)

func main() {
	handler := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "image/svg+xml")
		fmt.Fprintf(w, "<svg xmlns='http://www.w3.org/2000/svg' "+
			"style='stroke: grey; fill: white; stroke-width: 0.7' "+
			"width='%d' height='%d'>", width, height)

		for i := 0; i < cells; i++ {
			for j := 0; j < cells; j++ {
				ax, ay, az, ok1 := corner(i+1, j)
				bx, by, bz, ok2 := corner(i, j)
				cx, cy, cz, ok3 := corner(i, j+1)
				dx, dy, dz, ok4 := corner(i+1, j+1)

				z := (az + bz + cz + dz) / 4.0
				zmin := -0.2
				zmax := 0.5
				r := int(255. * (z - zmin) / (zmax - zmin))
				if r > 255 {
					r = 255
				}
				if r < 0 {
					r = 0
				}
				b := 255 - r

				if ok1 && ok2 && ok3 && ok4 {
					fmt.Fprintf(w, "<polygon points='%g,%g %g,%g %g,%g %g,%g' fill='rgb(%v, %v, %v)'/>\n",
						ax, ay, bx, by, cx, cy, dx, dy, r, 0, b)
				}
			}
		}
		fmt.Fprintln(w, "</svg>")
	}
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))

}

func corner(i, j int) (float64, float64, float64, bool) {
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// Compute surface height z.
	z, ok := f(x, y)
	if !ok {
		return 0, 0, 0, ok
	}

	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy, z, ok
}

func f(x, y float64) (float64, bool) {
	r := math.Hypot(x, y) // distance from (0,0)
	res := math.Sin(r) / r
	if math.IsNaN(res) || math.IsInf(res, 0) {
		return res, false
	}
	return res, true
}
