package main

import (
    "encoding/csv"
    "encoding/json"
    "fmt"
    "os"
    "github.com/spf13/cobra"
)

func main() {
    var inputFile string
    var outputFile string

    // Define the root command
    var rootCmd = &cobra.Command{
        Use:   "csvtojl [input csv file] [output json lines file]",
        Short: "Convert CSV to JSON Lines",
        Args:  cobra.ExactArgs(2),
        Run: func(cmd *cobra.Command, args []string) {
            inputFile = args[0]
            outputFile = args[1]
            convertCSVToJSONL(inputFile, outputFile)
        },
    }

    // Execute the command
    if err := rootCmd.Execute(); err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
}

// Convert CSV to JSON Lines
func convertCSVToJSONL(inputFile string, outputFile string) {
    // Open the CSV file
    csvFile, err := os.Open(inputFile)
    if err != nil {
        fmt.Printf("Error opening CSV file: %v\n", err)
        return
    }
    defer csvFile.Close()

    // Create a CSV reader
    reader := csv.NewReader(csvFile)

    // Read all records
    records, err := reader.ReadAll()
    if err != nil {
        fmt.Printf("Error reading CSV file: %v\n", err)
        return
    }

    // Create or truncate the output file
    jsonFile, err := os.Create(outputFile)
    if err != nil {
        fmt.Printf("Error creating JSON lines file: %v\n", err)
        return
    }
    defer jsonFile.Close()

    // Convert each record to JSON and write to the output file
    for _, record := range records {
        jsonData, err := json.Marshal(record)
        if err != nil {
            fmt.Printf("Error converting record to JSON: %v\n", err)
            continue
        }
        jsonFile.Write(jsonData)
        jsonFile.Write([]byte("\n")) // Write a newline after each JSON object
    }

    fmt.Println("Conversion complete!")
}
