# Aare-Exporter

Aare-Exporter is a small metrics exporter written in Go and made for usage with Prometheus.
The River Aare is the river that flows through the capital of Switzerland. Bern.

It exposes temperature data provided via the [Aare API by Purpl3.net](https://api.purpl3.net/aare/doc/) as metrics for Prometheus.

## Usage

- Clone the repository and change into the project directory.
- Make sure you have the latest Go version installed on your system.

Enter in terminal:

```
go run main.go
```

The metrics are avaiable on port **3005**.
Make sure to add the exporter to Prometheus as a scraping target.

---

### TODO

- Configuration file

