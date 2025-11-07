## Minimal Viable Version — “Insomnia CLI: Simple Uptime Monitor”

### **Goal**

Monitor if a set of URLs (APIs, websites, internal services) are **up or down**, and log results to a local database.

---

### **Core Features (MVP Scope)**

1. **URL Monitoring**

   * Accept one or multiple URLs (via CLI args or config file).
   * Ping them using an HTTP GET request.
   * Record the response time, status code, and timestamp.

2. **Customizable Interval**

   * Use a flag like `--interval 30s` or `--interval 5m` to control how frequently checks happen.

3. **Database Logging (SQLite or PostgreSQL)**

   * Store each check in a table with fields like:

     ```
     id | url | status_code | response_time_ms | checked_at
     ```
   * Makes it easy to query uptime history.

4. **Simple CLI Output**

   * Print real-time status in the terminal (UP / DOWN).
   * Optionally, show average uptime for a URL.

---

### 🛠 Example CLI Usage

```bash
insomnia monitor --url https://example.com --interval 30s
```

**Output:**

```
[12:01:04] ✅ https://example.com (200) - 112ms
[12:01:34] ✅ https://example.com (200) - 120ms
[12:02:04] ❌ https://example.com (timeout)
```

---

### Optional Small Add-ons (for depth, not scope bloat)

* `--from-db` flag → show uptime stats from past runs.
* `--threshold` flag → alert if downtime > X minutes.
* Integrate `go-sqlite3` (lightweight) or `pgx` (if PostgreSQL).
* Config file (`insomnia.yaml`) for multiple URLs.

---

### **Suggested Folder Structure**

```
insomnia/
├── cmd/
│   └── root.go          # CLI entry point
├── internal/
│   ├── monitor/         # HTTP checking logic
│   ├── db/              # DB setup & insert/query functions
│   └── utils/           # Time, logging helpers
├── go.mod
└── main.go
```
