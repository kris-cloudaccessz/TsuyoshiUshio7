package model

import (
	"go-image/convert"
	"net/http"
	"strings"
)

type Goimg_req_t struct {
	Grayscale int
	Rotate    float64
	Width     uint
	Height    uint
	Quality   uint
	X         int
	Y         int
	P         int
	Download  int
	Format    string
}

func ParamHandler(req *Goimg_req_t, r *http.Request) {
	if len(r.FormValue("g")) == 0 {
		req.Grayscale = 0
	} else {
		req.Grayscale = convert.StringToInt(r.FormValue("g"))
	}

	req.Rotate = convert.StringToFloat64(r.FormValue("r"))
	req.Width = convert.StringToUint(r.FormValue("w"))
	req.Height = convert.StringToUint(r.FormValue("h"))

	q := convert.StringToUint(r.FormValue("q"))
	if q == 0 || q > 100 {
		req.Quality = 75
	} else {
		req.Quality = q
	}

	if len(r.FormValue("x")) == 0 {
		req.X = -1
	} else {
		req.X = convert.StringToInt(r.FormValue("x"))
	}

	if len(r.FormValue("y")) == 0 {
		req.Y = -1
	} else {
		req.Y = convert.StringToInt(r.FormValue("y"))
	}

	if len(r.FormValue("p")) == 0 {
		req.P = 1
	} else {
		req.P = convert.StringToInt(r.FormValue("p"))
	}

	if len(r.FormValue("d")) == 0 {
		req.Download = 0
	} else {
		req.Download = convert.StringToInt(r.FormValue("d"))
	}

	switch strings.ToLower(r.FormValue("f")) {
	case "jpeg", "jpg":
		req.Format = "jpeg"
	case "png":
		req.Format = "png"
	case "gif":
		req.Format = "gif"
	case "webp":
		req.Format = "webp"
	default:
		req.Format = "jpeg"
	}
}
