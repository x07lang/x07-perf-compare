# x07-perf-compare

Performance comparison benchmarks for X07 vs C, Rust, and Go.

This repo contains the benchmark sources, runner, and published snapshots used to measure X07 runtime behavior, compile time, binary size, and memory use on the same workloads as familiar systems languages.

**Start here:** [`run_benchmarks.py`](run_benchmarks.py) · [`snapshots/`](snapshots/) · [`x07lang/x07`](https://github.com/x07lang/x07)

## What This Repo Is For

Use `x07-perf-compare` when you want to:

- evaluate X07 against C, Rust, and Go on comparable workloads
- measure a toolchain change before release
- regenerate benchmark evidence for docs or release notes
- inspect exactly how the published numbers were produced

This is an evidence repo, not a product surface. Its job is to make the performance story inspectable.

## Benchmarks

- `sum_bytes`
- `word_count`
- `rle_encode`
- `byte_freq`
- `fibonacci`

## Quick Start

Prerequisites:

- Python 3
- Rust toolchain
- C compiler
- Go toolchain
- released X07 toolchain or `x07-host-runner`

Run the default suite:

```sh
python3 run_benchmarks.py --x07-toolchain /path/to/x07-toolchain-dir
```

Useful variants:

```sh
python3 run_benchmarks.py --x07-toolchain /path/to/x07-toolchain-dir --size 1000
python3 run_benchmarks.py --x07-toolchain /path/to/x07-toolchain-dir --benchmarks sum_bytes word_count
python3 run_benchmarks.py --x07-toolchain /path/to/x07-toolchain-dir --iterations 10 --warmup 3
python3 run_benchmarks.py --x07-toolchain /path/to/x07-toolchain-dir --json > results.json
python3 run_benchmarks.py --x07-toolchain /path/to/x07-toolchain-dir --direct
```

## Repo Layout

- `x07/`: benchmark programs written in X07
- `projects/`: project-style X07 benchmarks
- `c/`, `rust/`, `rust_cargo/`, `go/`: comparison implementations
- `snapshots/`: published result snapshots
- `run_benchmarks.py`: benchmark driver

## How It Fits The X07 Ecosystem

- [`x07`](https://github.com/x07lang/x07) is the toolchain being measured
- `x07-perf-compare` publishes the supporting evidence
- docs and release materials pull from the results here instead of making ungrounded performance claims
