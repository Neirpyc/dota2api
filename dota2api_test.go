package dota2api

import (
	"bytes"
	. "github.com/franela/goblin"
	"image"
	"image/color"
	"image/jpeg"
	"image/png"
	"net/http"
	"testing"
)

type mockClient struct {
	DoFunc func(req *http.Request) (*http.Response, error)
}

func (m *mockClient) Do(req *http.Request) (*http.Response, error) {
	return m.DoFunc(req)
}

func GetTestConfig() Config {
	return Config{
		Timeout:     1,
		SteamApiKey: "keyTEST",
	}
}

func getTestColor() color.RGBA {
	return color.RGBA{
		R: 42,
		G: 42 * 2,
		B: 42 * 3,
		A: 255,
	}
}

func getImageTest() image.Image {
	img := image.NewNRGBA(image.Rect(0, 0, 1, 1))
	img.Set(0, 0, getTestColor())
	return img
}

func getJpgTest() []byte {
	var b []byte
	buf := bytes.NewBuffer(b)
	_ = jpeg.Encode(buf, getImageTest(), &jpeg.Options{
		Quality: 100,
	})
	return buf.Bytes()
}

func getPngTest() []byte {
	var b []byte
	buf := bytes.NewBuffer(b)
	_ = png.Encode(buf, getImageTest())
	return buf.Bytes()
}

func validateTestImage(img image.Image) bool {
	r0, g0, b0, _ := img.At(0, 0).RGBA()
	r1, g1, b1, _ := getTestColor().RGBA()
	isOk := func(a, b uint32) bool {
		if a > b {
			return a-b < 255
		}
		return b-a < 255
	}
	return isOk(r0, r1) && isOk(g0, g1) && isOk(b0, b1)
}

func TestLoadConfig(t *testing.T) {
	g := Goblin(t)
	g.Describe("LoadConfigFromFile", func() {
		g.It("Should load without error", func() {
			_, err := LoadConfigFromFile("config.yaml")
			g.Assert(err).Equal(nil)
		})
	})
}
