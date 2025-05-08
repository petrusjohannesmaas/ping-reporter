package utils

import (
    "database/sql"
    "fmt"
    "log"
    "strconv"

    _ "github.com/lib/pq"
)

// Reading struct to store query results
type Reading struct {
    ID          int
    PPPusername string
    IP          string
    Identity    string
    MaxRtt      float64
}

// Extractor retrieves network latency data from the database
func Extractor() {
    // Define database connection string
    connStr := "postgres://postgres:<your-password>@localhost:5432/test?sslmode=disable"
    
    // Open the database connection
    db, err := sql.Open("postgres", connStr)
    if err != nil {
        log.Fatalf("Failed to open database: %v", err)
    }
    defer db.Close()

    // Verify database connectivity
    if err = db.Ping(); err != nil {
        log.Fatalf("Database unreachable: %v", err)
    }
    fmt.Println("Successfully connected to the database!")

    // Query to retrieve the highest avg_rtt_ms per record
    query := `
    WITH extracted AS (
        SELECT
            id,
            ppp_username,
            ip,
            elem->>'identity' AS identity,
            elem->>'avg_rtt_ms' AS max_rtt_ms,
            ROW_NUMBER() OVER (PARTITION BY id ORDER BY CAST(elem->>'avg_rtt_ms' AS DOUBLE PRECISION) DESC) AS rn
        FROM router_os_api.report,
        jsonb_array_elements(neighbors) AS elem
        WHERE elem->>'avg_rtt_ms' ~ '^[0-9]+(\.[0-9]+)?$' -- Ensure avg_rtt_ms is numeric
    )
    SELECT id, ppp_username, ip, identity, max_rtt_ms
    FROM extracted
    WHERE rn = 1; -- Select the highest avg_rtt_ms per record
    `

    // Execute the query
    rows, err := db.Query(query)
    if err != nil {
        log.Fatalf("Query execution failed: %v", err)
    }
    defer rows.Close()

    // Store query results
    var readings []Reading

    // Process each row
    for rows.Next() {
        var reading Reading
        var maxRttStr string

        // Scan row into Reading struct
        err = rows.Scan(&reading.ID, &reading.PPPusername, &reading.IP, &reading.Identity, &maxRttStr)
        if err != nil {
            log.Printf("Skipping row due to scan error: %v", err)
            continue
        }

        // Convert max_rtt_ms to float64
        reading.MaxRtt, err = strconv.ParseFloat(maxRttStr, 64)
        if err != nil {
            log.Printf("Skipping row due to parse error: %v", err)
            continue
        }

        // Append reading to the list
        readings = append(readings, reading)
    }

    // Verify successful iteration
    if err = rows.Err(); err != nil {
        log.Fatalf("Row iteration error: %v", err)
    }

    // Export results to an Excel file
    if err = WriteToExcel(readings, "output.xlsx"); err != nil {
        log.Fatalf("Failed to export data to Excel: %v", err)
    }

    fmt.Println("Data successfully exported to output.xlsx")
}