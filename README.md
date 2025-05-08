# **Ping Reporter (RouterOS RTT Calculator)**

### **Overview**
`ping-reporter` is a Go-based utility that analyzes network latency using data from **RouterOS** devices. The tool connects to a PostgreSQL database, retrieves information about network neighbors, and identifies the worst-performing connections based on **highest round-trip times (RTTs)**. 

This allows network administrators to pinpoint connectivity issues and optimize network performance.

---

## **Project Structure**
The project is structured into the following key components:

```
ping_reporter/
│── main.go          # Entry point for the application
│── utils/
│   ├── extractor.go # Handles database operations and data extraction
│   ├── excel.go     # Handles exporting extracted data to an Excel file
│── go.mod           # Go module definition
│── README.md        # Project documentation
```

### **Workflow**
1. `main.go` calls `utils.Extractor()`, which:
   - **Connects** to the PostgreSQL database.
   - **Extracts** neighbors' RTT metrics from RouterOS reports.
   - **Identifies** the worst connection by selecting the **highest RTT per record**.
   - **Exports** the results to an Excel file for analysis.

---

## **Features**
✅ **Extracts Highest RTTs** – Pinpoints network connections with the worst performance.  
✅ **Filters Numeric RTT Values** – Ensures only valid metrics are considered.  
✅ **Exports Data to Excel** – Saves structured reports for easy review.  
✅ **PostgreSQL Integration** – Retrieves and processes RouterOS JSONB data efficiently.  
✅ **Scalable Design** – Modular components allow for future enhancements.  

### **Future Enhancements**
🔹 **Include MAC Address for Each Bad Connection**  
🔹 **Add `max_rtt` Alongside the Current RTT Metrics**  
🔹 **Implement SMTP Email Module for Report Delivery**  
🔹 **Optimize SQL Queries for Faster Processing**  

---

## **Installation**
### **Prerequisites**
Ensure you have:
- Go installed (`>= 1.18`)
- PostgreSQL database configured with RouterOS data

### **Setup**
1. Clone the repository:
   ```bash
   git clone https://github.com/petrusjohannesmaas/ping-reporter
   cd ping-reporter
   ```

2. Install dependencies:
   ```bash
   go mod tidy
   ```

3. Configure database connection in `extractor.go`:
   ```go
   connStr := "postgres://postgres:<your-password>@localhost:5432/test?sslmode=disable"
   ```

---

## **Usage**
Run the tool with:
```bash
go run main.go
```
After execution, the processed data will be saved to `output.xlsx`. Future versions will also support email-based report sharing.


---

### **Acknowledgments**
This tool relies on **RouterOS** data for network analysis, offering deep insights into connection performance.

---
