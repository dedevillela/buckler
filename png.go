package main

import (
	"net/http"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"os"
)

func makePngShield (w http.ResponseWriter, d Data) {
	w.Header().Add("content-type", "image/png");

	fi, _ := os.Open("edge.png");
	edge, _ := png.Decode(fi);
	mask := image.NewRGBA(image.Rect(0, 0, 100, 19));
    draw.Draw(mask, edge.Bounds(), edge, image.ZP, draw.Src);

	sr := image.Rect(2, 0, 3, 19);
	for i := 3; i <= 97; i++ {
		dp := image.Point{i, 0};
		r := sr.Sub(sr.Min).Add(dp);
    	draw.Draw(mask, r, edge, sr.Min, draw.Src);
	}

	sr = image.Rect(0, 0, 1, 19);
	dp := image.Point{99, 0};
	r := sr.Sub(sr.Min).Add(dp);
	draw.Draw(mask, r, edge, sr.Min, draw.Src);

	sr = image.Rect(1, 0, 2, 19);
	dp = image.Point{98, 0};
	r = sr.Sub(sr.Min).Add(dp);
	draw.Draw(mask, r, edge, sr.Min, draw.Src);

	img := image.NewRGBA(image.Rect(0, 0, 100, 19));
	blue := color.RGBA{0, 0, 255, 255};
    draw.Draw(img, img.Bounds(), &image.Uniform{blue}, image.ZP, draw.Src);
	red := color.RGBA{255, 0, 0, 255};
	rect := image.Rect(0, 0, 50, 19);
    draw.Draw(img, rect, &image.Uniform{red}, image.ZP, draw.Src);


	dst := image.NewRGBA(image.Rect(0, 0, 100, 19));
	draw.DrawMask(dst, dst.Bounds(), img, image.ZP, mask, image.ZP, draw.Over);
	png.Encode(w, dst);
}
