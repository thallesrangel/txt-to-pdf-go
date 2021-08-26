package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/jung-kurt/gofpdf"
)

func main() {
	file := "test.txt"

	content, err := ioutil.ReadFile(file)

	if err != nil {
		log.Fatalf("%s file not found", file)
	}

	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Helvetica", "", 16)
	pdf.MultiCell(190, 5, string(content), "0", "0", false)
	err = pdf.OutputFileAndClose("test.pdf")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("PDF Created")
}
