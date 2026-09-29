# Price Calculator

A small Go command-line tool that reads a list of prices, applies several tax rates, and writes one JSON result file per rate. It was built to practise **Go interfaces, packages and error handling**. The pricing logic doesn't know where its data comes from or where it goes.

![Go](https://img.shields.io/badge/Go-1.22-00ADD8?logo=go&logoColor=white)
![License](https://img.shields.io/badge/license-MIT-green)

## How it works

```
prices.txt ──► FileManager.ReadLines ──► conversion.StringsToFloats
                                              │
                              TaxIncludePriceJob.Process (per tax rate)
                                              │
            result_<rate>.json ◄── FileManager.WriteLines (JSON)
```

For each tax rate (0%, 7%, 10% and 15%), a `TaxIncludePriceJob` loads the prices, calculates `price × (1 + rate)`, and writes the result.

## Design

The job depends on an interface, not a concrete type:

```go
type IOManager interface {
    ReadLines() ([]string, error)
    WriteLines(data interface{}) error
}
```

There are two implementations:

| Package | Reads from | Writes to |
|---|---|---|
| `file_manager` | a text file, one price per line | an indented JSON file |
| `cmd_manager` | interactive console input (type `exit` to finish) | stdout |

Swapping one for the other is a one-line change in `main.go`. No pricing code changes, which makes the job easy to test with a fake `IOManager`.

```
.
├── main.go          # runs one job per tax rate
├── prices/          # TaxIncludePriceJob: load → calculate → write
├── io_manager/      # IOManager interface
├── file_manager/    # file-based IOManager
├── cmd_manager/     # console-based IOManager
├── conversion/      # string → float64 parsing with error propagation
└── prices.txt       # sample input
```

## Usage

```bash
go run .
```

Given `prices.txt`:

```
9.99
10.49
15.89
12
```

`result_7.00.json` looks like this:

```json
{
  "tax_rate": 0.07,
  "input_prices": [9.99, 10.49, 15.89, 12],
  "tax_included_prices": {
    "9.99": "10.69",
    "10.49": "11.22",
    "12.00": "12.84",
    "15.89": "17.00"
  }
}
```

## Possible next steps

- [ ] Run the jobs concurrently with goroutines and collect errors over channels
- [ ] Table-driven unit tests with a fake `IOManager`
- [ ] Take tax rates and file paths as CLI flags
