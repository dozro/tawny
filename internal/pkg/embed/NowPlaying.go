package embed

import (
	"bytes"
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	"image/png"

	"github.com/disintegration/imaging"
	"github.com/fogleman/gg"
	log "github.com/sirupsen/logrus"
	"golang.org/x/image/tiff"
)

const pathToFont = "assets/fonts/RubikMarker/RubikMarkerHatch-Regular.ttf"

func EmbedNowPlaying(trackTitle string, trackArtist, trackAlbum, trackAlbumCover, username string, now bool, outputformat string) (*bytes.Buffer, error) {
	// Albumcover herunterladen
	albumCoverResp, err := doHttpGetRequest(trackAlbumCover)
	if err != nil {
		return nil, fmt.Errorf("failed to download album cover: %w", err)
	}
	defer albumCoverResp.Body.Close()

	albumCover, _, err := image.Decode(albumCoverResp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to decode album cover: %w", err)
	}

	// Hintergrundbild laden
	bgImage, err := gg.LoadImage("assets/images/NowPlaying_Background.png")
	if err != nil {
		return nil, fmt.Errorf("failed to load background image: %w", err)
	}

	//artistSymbole, err := gg.LoadImage("assets/images/singer.png")
	//if err != nil {
	//	return nil, fmt.Errorf("failed to load artist symbole image: %w", err)
	//}

	imgWidth := bgImage.Bounds().Dx()
	imgHeight := bgImage.Bounds().Dy()

	dc := gg.NewContext(imgWidth, imgHeight)
	dc.DrawImage(bgImage, 0, 0)

	if err := dc.LoadFontFace(pathToFont, 50); err != nil {
		return nil, fmt.Errorf("failed to load font: %w", err)
	}

	dc.SetColor(color.White)

	const coverSize = 350
	albumCover = resizeImage(albumCover, coverSize, coverSize)
	dc.DrawImage(albumCover, 50, imgHeight/2-coverSize/2)

	textX := float64(50 + coverSize + 200)
	textWidth := float64(imgWidth) - textX - 60
	textY := float64(imgHeight/2 - 60)

	dc.SetColor(color.White)
	drawTextWithGlow(dc, truncateString(trackTitle, 40), textX-50, textY-100, 450, 50, color.White)

	dc.LoadFontFace(pathToFont, 28)
	dc.SetColor(color.RGBA{220, 220, 220, 255})
	dc.DrawStringWrapped(truncateString(trackAlbum, 35), textX, textY+60, 0, 0, textWidth, 1.3, gg.AlignLeft)

	dc.LoadFontFace(pathToFont, 24)
	dc.SetColor(color.RGBA{180, 180, 180, 255})
	dc.DrawStringWrapped(truncateString(trackArtist, 40), textX, textY+110, 0, 0, textWidth, 1.3, gg.AlignLeft)

	//artistSymbole = resizeImage(artistSymbole, 30, 30)
	//dc.DrawImage(artistSymbole, int(textX-30), int(textY+105))

	dc.LoadFontFace(pathToFont, 18)
	dc.SetColor(color.RGBA{220, 161, 161, 233})
	dc.DrawStringWrapped("Generated with Tawny based on last.fm data", 10, 390, 0, 0, textWidth+100, 1.3, gg.AlignLeft)
	if now {
		dc.DrawStringWrapped(fmt.Sprintf("%s is currently listening", username), 75, 15, 0, 0, textWidth+100, 1.3, gg.AlignLeft)
	} else {
		dc.DrawStringWrapped(fmt.Sprintf("%s was recently listening", username), 75, 15, 0, 0, textWidth+100, 1.3, gg.AlignLeft)
	}

	outputImage := dc.Image()
	buf := new(bytes.Buffer)
	switch outputformat {
	case "image/jpeg":
		{
			log.Debug("generating jpeg image")
			if err := jpeg.Encode(buf, outputImage, nil); err != nil {
				return nil, fmt.Errorf("failed to encode JPEG: %w", err)
			}
			break
		}
	case "image/tiff":
		{
			log.Debug("generating tiff image")
			if err := tiff.Encode(buf, outputImage, nil); err != nil {
				return nil, fmt.Errorf("failed to encode webp image: %w", err)
			}
		}
	default:
		{
			log.Debug("generating png image")
			if err := png.Encode(buf, outputImage); err != nil {
				return nil, fmt.Errorf("failed to encode PNG: %w", err)
			}
		}
	}

	return buf, nil
}

// Hilfsfunktion fürs Resizing
func resizeImage(img image.Image, width, height int) image.Image {
	return imaging.Resize(img, width, height, imaging.Lanczos)
}

// drawTextWithGlow zeichnet Text mit Glow / Neon-Effekt
func drawTextWithGlow(dc *gg.Context, text string, x, y float64, maxWidth float64, fontSize float64, glowColor color.Color) {
	r, g, b, _ := glowColor.RGBA()
	baseColor := color.RGBA{uint8(r >> 8), uint8(g >> 8), uint8(b >> 8), 255}

	for i := 10; i >= 2; i -= 2 {
		alpha := uint8(10 + (10-i)*3) // leicht wachsender alpha
		dc.SetColor(color.RGBA{baseColor.R, baseColor.G, baseColor.B, alpha})
		dc.DrawStringWrapped(text, x, y, 0, 0, maxWidth, 1.3, gg.AlignLeft)
	}

	dc.SetColor(color.White)
	dc.DrawStringWrapped(text, x, y, 0, 0, maxWidth, 1.3, gg.AlignLeft)
}
