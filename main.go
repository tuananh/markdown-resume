package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
)

func main() {
	// Define CLI flags
	input := flag.String("input", "", "Path to the HTML file to convert (required)")
	output := flag.String("output", "output.pdf", "Path to save the generated PDF")
	timeout := flag.Int("timeout", 30, "Timeout in seconds for the conversion process")
	flag.Parse()

	// Check for required input
	if *input == "" {
		fmt.Println("Error: The input HTML file is required.")
		flag.Usage()
		os.Exit(1)
	}

	// Ensure the input file exists
	if _, err := os.Stat(*input); os.IsNotExist(err) {
		fmt.Printf("Error: File %s does not exist.\n", *input)
		os.Exit(1)
	}

	// Run the conversion
	if err := convertHTMLToPDF(*input, *output, time.Duration(*timeout)*time.Second); err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("PDF successfully saved to %s\n", *output)
}

// convertHTMLToPDF converts an HTML file to a PDF
func convertHTMLToPDF(input string, output string, timeout time.Duration) error {
	// Read the HTML file
	html, err := os.ReadFile(input)
	if err != nil {
		return fmt.Errorf("failed to read HTML file: %v", err)
	}

	// Create a context with a timeout
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	// Create a new chromedp context
	ctx, cancelCtx := chromedp.NewContext(ctx)
	defer cancelCtx()

	// Define tasks to generate the PDF
	var pdfBuffer []byte
	tasks := chromedp.Tasks{
		chromedp.Navigate("data:text/html," + string(html)), // Load the HTML as a data URL
		chromedp.ActionFunc(func(ctx context.Context) error {
			var err error
			// This is where PrintToPDF is used
			pdfBuffer, _, err = page.PrintToPDF().WithPaperWidth(8.27).WithPaperHeight(11.69).Do(ctx)
			return err
		}),
	}

	// Run the tasks
	if err := chromedp.Run(ctx, tasks); err != nil {
		return fmt.Errorf("failed to generate PDF: %v", err)
	}

	// Save the PDF to the output file
	if err := os.WriteFile(output, pdfBuffer, 0644); err != nil {
		return fmt.Errorf("failed to save PDF: %v", err)
	}

	return nil
}
