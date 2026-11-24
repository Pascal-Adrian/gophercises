# Quiz Game

Solution to [Gophercises exercise 1](https://gophercises.com/). A timed command-line quiz
that reads `question,answer` pairs from a CSV file, asks them one by one, and reports how
many you got right when you finish or the time runs out.

Answers are compared case-insensitively, with surrounding whitespace trimmed.

## Usage

```
go run . [-filepath problems.csv] [-timeout 30] [-shuffle]
```

| Flag        | Default        | Description                                  |
| ----------- | -------------- | -------------------------------------------- |
| `-filepath` | `problems.csv` | CSV file with the problems                   |
| `-timeout`  | `30`           | Time limit for the whole quiz, in seconds    |
| `-shuffle`  | `false`        | Shuffle the problems before starting         |

## Testing

```
go test -v -race .
```

`readToChan` is covered by a table-driven test (`main_test.go`).

`getProblems` is deliberately kept as a single function with no test of its own. It is
a thin wrapper around `os.Open` and `csv.Reader.ReadAll` from the standard library: it
contains no logic of its own and only passes their results and errors back to the caller.
A test for it would mostly re-test the standard library, which is already tested upstream.
Splitting it further (e.g. opening the file separately from parsing) would add indirection
without adding anything to test. 
