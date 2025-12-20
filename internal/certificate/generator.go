package certificate

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/fogleman/gg"
)

var (
	outputDir = "output"
)

func Generate(name string) (string, error) {
	// Read environment variables at runtime
	templatePath := os.Getenv("CERT_TEMPLATE")
	fontPath := os.Getenv("CERT_FONT_PATH")

	// Ensure output directory exists
	if err := os.MkdirAll(outputDir, 0755); err != nil {
		return "", err
	}

	// Load base certificate
	img, err := gg.LoadImage(templatePath)
	if err != nil {
		return "", err
	}

	// Create drawing context
	dc := gg.NewContextForImage(img)

	// Load font (increased size for better visibility)
	if err := dc.LoadFontFace(fontPath, 85); err != nil {
		return "", err
	}

	// Set text color to solid black with full opacity
	dc.SetRGBA(0, 0, 0, 1)

	// Draw name centered on the blank line
	// X: centered horizontally with slight offset to the right
	// Y: positioned at approximately 51% from top to align with blank line
	dc.DrawStringAnchored(
		name,
		float64(dc.Width())/2+200,
		float64(dc.Height())*0.51,
		0.5,
		0.5,
	)

	// Clean filename
	safeName := strings.ReplaceAll(name, " ", "_")
	outPath := filepath.Join(outputDir, fmt.Sprintf("%s.png", safeName))

	// Save certificate
	if err := dc.SavePNG(outPath); err != nil {
		return "", err
	}

	return outPath, nil
}

// GenerateWithTemplate creates a certificate with a custom template
func GenerateWithTemplate(name, templatePath string) (string, error) {
	// Read font path from environment
	fontPath := os.Getenv("CERT_FONT_PATH")

	// Ensure output directory exists
	if err := os.MkdirAll(outputDir, 0755); err != nil {
		return "", err
	}

	// Load custom certificate template
	img, err := gg.LoadImage(templatePath)
	if err != nil {
		return "", err
	}

	// Create drawing context
	dc := gg.NewContextForImage(img)

	// Load font (increased size for better visibility)
	if err := dc.LoadFontFace(fontPath, 85); err != nil {
		return "", err
	}

	// Set text color to solid black with full opacity
	dc.SetRGBA(0, 0, 0, 1)

	// Draw name centered on the blank line
	// X: centered horizontally with slight offset to the right
	// Y: positioned at approximately 51% from top to align with blank line
	dc.DrawStringAnchored(
		name,
		float64(dc.Width())/2+200,
		float64(dc.Height())*0.51,
		0.5,
		0.5,
	)

	// Clean filename
	safeName := strings.ReplaceAll(name, " ", "_")
	outPath := filepath.Join(outputDir, fmt.Sprintf("%s.png", safeName))

	// Save certificate
	if err := dc.SavePNG(outPath); err != nil {
		return "", err
	}

	return outPath, nil
}
