package main

//
// Program for generating emergencies formatted for pressure seal forms for Arizona Treasure Hunt
// Written by: Andy Woodward - Committee 2019-2022
//
import (
	"flag"
	"fmt"
	"math/rand"
	"os"

	"github.com/phpdave11/gofpdf"
	"github.com/phpdave11/gofpdf/contrib/barcode"
	"github.com/phpdave11/gofpdf/contrib/gofpdi"
)

const (
	usage = `usage: %s

Program for generating emergencies formatted for pressure seal forms 
for Arizona Treasure Hunt
Written by: Andy Woodward - TH Committee 2019-2022
Email: awoodward@gmail.com

Options:
`
)

// Generate the PDF data for each car/team
func writeCar(car int, pages int, templateFile string, pdf *gofpdf.Fpdf,
	startingEmergency int, noSecurity bool, noBarcode bool) *gofpdf.Fpdf {
	count := startingEmergency - 1
	imp := gofpdi.NewImporter() // use this instead of gofpdi.ImportPage(). Otherwise, there is a memory leak.
	for i := 0; i < pages; i++ {
		// Address page
		tp := imp.ImportPage(pdf, templateFile, i+1, "/MediaBox")
		pdf.AddPage()
		imp.UseImportedTemplate(pdf, tp, 0, 0, 215.9, 279.4)
		if (i+1)%2 > 0 {
			count++
			if !noBarcode {
				barcodeData := fmt.Sprintf("%v-EM-%v", car, count)
				barcode.Barcode(
					pdf,
					barcode.RegisterCode128(pdf, barcodeData),
					15,
					202,
					35,
					8,
					false,
				)
			}
			pdf.MoveTo(15, 213)
			pdf.SetFont("Arial", "", 12)
			pdf.Cell(40, 10, fmt.Sprintf("Car - %v Emergency - %v", car, count))
		} else {
			// security page
			// Add some junk so people can't read the emergency with a flashlight
			if !noSecurity {
				pdf.SetLineWidth(.8)
				// Top Panel
				for i := 1; i < 13; i++ {
					for j := 1; j < 31; j++ {
						pdf.Circle(14+(6*float64(j)), 14+(6*float64(i)), 6, "D")

					}
				}
				// Bottom Panel
				for i := 1; i < 12; i++ {
					for j := 1; j < 31; j++ {
						pdf.Circle(14+(6*float64(j)), 200+(6*float64(i)), 6, "D")

					}
				}
				// Create a string of random characters
				noiseStr := ""
				for i := 0; i < 96; i++ {
					ascii := rand.Intn(90-48) + 48
					noiseStr = noiseStr + string(ascii)
				}
				pdf.SetFont("Arial", "", 12)
				pdf.MoveTo(15, 257)
				pdf.Cell(0, 0, noiseStr)
				pdf.MoveTo(15, 259.33)
				pdf.Cell(0, 0, noiseStr)
			}
		}
	}

	return pdf
}

func main() {
	carCount := 100
	flag.IntVar(&carCount, "carCount", 100, "Number of cars")
	pages := 52
	flag.IntVar(&pages, "pages", 52, "Number of pages in template")
	startingEmergency := 1
	flag.IntVar(&startingEmergency, "startingEmergency", 1, "First emergency in the sequence")
	singleFile := false
	flag.BoolVar(&singleFile, "singleFile", false, "Combine all into one file")
	noSecurity := false
	flag.BoolVar(&noSecurity, "noSecurity", false, "Disable printing of security hashing")
	noBarcode := false
	flag.BoolVar(&noBarcode, "noBarcode", false, "Disable printing of barcodes")
	templateFile := "template.pdf"
	flag.StringVar(&templateFile, "template", "template.pdf", "Emergency template file")
	saveFile := "emergencies.pdf"
	flag.StringVar(&saveFile, "file", "emergencies.pdf", "Destination file for single emergency file")
	flag.Usage = func() { // [4]
		fmt.Fprintf(flag.CommandLine.Output(), usage, os.Args[0])
		flag.PrintDefaults()
	}
	flag.Parse()
	var err error
	pdf := gofpdf.New("P", "mm", "Letter", "")
	for i := 1; i <= carCount; i++ {

		fmt.Printf("Car: %v\n", i)
		pdf = writeCar(i, pages, templateFile, pdf, startingEmergency, noSecurity, noBarcode)
		if !singleFile {
			filename := fmt.Sprintf("Car_%v.pdf", i)
			err = pdf.OutputFileAndClose(filename)
			_ = err
			pdf = gofpdf.New("P", "mm", "Letter", "")
		}

	}
	if singleFile {
		var sFile *os.File
		sFile, err = os.Create(saveFile)
		pdf.Output(sFile)
		sFile.Sync()
		sFile.Close()
	}
}
