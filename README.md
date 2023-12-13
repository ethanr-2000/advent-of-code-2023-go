# advent-of-code-go

This repo was forked from https://github.com/jacob-alt-del/advent-of-code-go, which itself is a trimmed down version of https://github.com/alexchao26/advent-of-code-go

# Runtimes
The table below shows the best 3 of 4 consecutive runtimes. Doesn't include `go build`.

| Day | Runtime | Day | Runtime | Day | Runtime | Day | Runtime | Day | Runtime |
|-----|---------|-----|---------|-----|---------|-----|---------|-----|---------|
| 1   | 72ms    | 6   | 71ms    | 11  | 120ms   | 16  |         | 21  |         |
| 2   | 77ms    | 7   | 120ms   | 12  | 2365ms  | 17  |         | 22  |         |
| 3   | 75ms    | 8   | 79ms    | 13  | 107ms   | 18  |         | 23  |         |
| 4   | 84ms    | 9   | 81m     | 14  |         | 19  |         | 24  |         |
| 5   | 696ms   | 10  | 122ms   | 15  |         | 20  |         | 25  |         |

# Usage

## Create template
* `make skeleton` - Template for the current day
* `make skeleton DAY=N` -  Template for Nth day
* `make skeleton DAY=N YEAR=M` - Template for Nth day and Mth year

## Testing
* `./scripts/run-tests.sh` - run all tests

#### In day folder
* `go test` - run tests
* `go run main.go` - run both parts
* `go run main.go -part <1 or 2>` - run part 1 or 2
