# x07-perf-compare

## Agent Entrypoint

Start here: https://x07lang.org/docs/getting-started/agent-quickstart

Performance comparison benchmarks: **X07 vs C vs Rust vs Go**

This repo contains the benchmark sources, runner, and published snapshots used to measure how X07 behaves against familiar systems languages on the same workloads.

Community:

- Discord: https://discord.gg/59xuEuPN47
- Email: support@x07lang.org

## What Is In This Repo

This repo includes:

- matching benchmark implementations in `x07/`, `c/`, `rust/`, `rust_cargo/`, and `go/`
- project-based X07 benchmarks in `projects/`
- the benchmark driver in `run_benchmarks.py`
- reproducible result snapshots under `snapshots/`

## Vision

The vision of `x07-perf-compare` is to give end users a simple, inspectable answer to a practical question:

"If X07 is designed for agents, does that force me to give up runtime speed or reasonable binary size?"

This repo exists so the answer is based on runnable code and published measurements, not marketing claims.

## How It Fits The X07 Ecosystem

`x07-perf-compare` is an evidence repo for the wider X07 story.

- [`x07`](https://github.com/x07lang/x07) is the language and toolchain being measured
- [`x07lang.org`](https://x07lang.org) and the core docs use the results here to explain the performance story
- other repos in the ecosystem focus on packaging, MCP, web UI, registry, and platform workflows
- this repo isolates one question: how fast the generated X07 programs run and how expensive they are to build

That makes this repo a supporting part of the whole, not a product surface on its own.

## Practical Usage

Use this repo when you want to:

- evaluate X07 against C, Rust, and Go on comparable workloads
- measure the impact of a toolchain change before a release
- generate fresh benchmark snapshots for docs, release notes, or internal review
- inspect exactly how the published numbers were produced

## Install And Use It Standalone

You can run this repo by itself with a released X07 toolchain plus the usual native compilers:

- install or download X07 from [`x07`](https://github.com/x07lang/x07)
- install Python, a C compiler, Rust, and Go
- point `run_benchmarks.py` at the X07 toolchain directory or `x07-host-runner`

If you want the shortest path, use a released toolchain bundle from GitHub Releases and keep this repo separate from the main X07 workspace.

## Use It As Part Of The X07 Ecosystem

In the full ecosystem, this repo is typically used after a compiler, runtime, or packaging change:

1. build or install the candidate X07 toolchain
2. run the benchmarks here
3. compare the fresh numbers with previous snapshots
4. feed the results back into release notes, docs, or go/no-go decisions

That keeps the performance story consistent across the language repo, the docs site, and release materials.

## Benchmarks

| Benchmark | Description |
|-----------|-------------|
| `sum_bytes` | Sum all bytes in input, output as u32_le |
| `word_count` | Count whitespace-separated words |
| `rle_encode` | Run-length encoding of byte stream |
| `byte_freq` | Byte frequency histogram (byte -> count mapping) |
| `fibonacci` | Iterative Fibonacci computation |

## Running Benchmarks

```bash
# Download the standalone X07 toolchain from GitHub Releases:
#   https://github.com/x07lang/x07/releases
# Extract it somewhere and either add it to PATH or point the runner at it:
#   tar -xzf x07-vX.Y.Z-macOS.tar.gz -C /path/to/x07-toolchain-dir

# Run all benchmarks with default settings (100KB input, 5 iterations)
python3 run_benchmarks.py --x07-toolchain /path/to/x07-toolchain-dir

# Run with custom input size (in KB)
python3 run_benchmarks.py --x07-toolchain /path/to/x07-toolchain-dir --size 1000

# Run specific benchmarks
python3 run_benchmarks.py --x07-toolchain /path/to/x07-toolchain-dir --benchmarks sum_bytes word_count

# More iterations for better statistics
python3 run_benchmarks.py --x07-toolchain /path/to/x07-toolchain-dir --iterations 10 --warmup 3

# Output as JSON
python3 run_benchmarks.py --x07-toolchain /path/to/x07-toolchain-dir --json > results.json

# Direct binary execution (no host runner overhead)
python3 run_benchmarks.py --x07-toolchain /path/to/x07-toolchain-dir --direct

# Size-focused native builds (passes through to x07-host-runner --cc-profile)
python3 run_benchmarks.py --x07-toolchain /path/to/x07-toolchain-dir --x07-cc-profile size

# Run against a different X07 toolchain (useful for before/after comparisons)
python3 run_benchmarks.py --x07-toolchain /path/to/other-x07-toolchain-dir
```

## Options

- `--x07-host-runner PATH`: Path to `x07-host-runner` (env: `X07_HOST_RUNNER`; default: search `PATH`)
- `--x07-toolchain DIR`: Extracted toolchain directory containing `x07-host-runner` (env: `X07_TOOLCHAIN`)
- `--size N`: Input size in KB (default: 100)
- `--iterations N`: Number of timed runs (default: 5)
- `--warmup N`: Warmup iterations before timing (default: 2)
- `--benchmarks`: Specific benchmarks to run
- `--json`: Output results as JSON
- `--direct`: Run X07 binaries directly without host runner overhead
- `--x07-cc-profile {default,size}`: Select the C toolchain profile for X07 builds

## Directory Structure

```
.
├── README.md
├── run_benchmarks.py  # Benchmark runner script
├── x07/               # X07 programs (.x07.json)
├── projects/          # Project-based X07 benchmarks (packages)
├── c/                 # C implementations
├── go/                # Go implementations (single-file)
├── rust/              # Rust implementations (single-file)
└── rust_cargo/        # Rust implementations with Cargo deps
```

## Requirements

- Python 3.10+
- Rust toolchain (rustc)
- C compiler (cc/gcc/clang)
- Go toolchain (`go`)
- X07 toolchain download (standalone binaries from https://github.com/x07lang/x07/releases)

## Metrics Collected

- **Mean time**: Average execution time across iterations
- **Min time**: Minimum (best) execution time
- **StdDev**: Standard deviation of execution times
- **Compile time**: One-time compilation overhead
- **Build size**: Final executable size (bytes on disk)
- **Peak RSS**: Peak resident set size from a single run (via `/usr/bin/time` if available)
- **Speedup**: Relative performance vs X07 (higher = faster)

## Notes

- **Host runner mode** (default): X07 times include host runner overhead (process spawn, JSON I/O, artifact loading). This represents the full execution environment overhead.
- **Direct binary mode** (`--direct`): Runs compiled X07 binaries directly, measuring only native code execution time. This gives a fairer comparison of generated code quality.
- **Peak RSS**: Measured by executing the compiled binaries directly once per language (so the metric is comparable across languages, independent of X07 host runner overhead).
- C and Rust programs are compiled with `-O3` / `opt-level=3`
- Go programs are compiled with `CGO_ENABLED=0` and `-ldflags "-s -w"`
- All programs use stdin/stdout for I/O
- Output correctness is verified (outputs must match across languages)

## Regex Benchmarks

This repo also contains `regex_*` programs, but they depend on the external `ext-regex` package, which is not published to the registry yet. They are not included in the default benchmark set.

## Performance Comparison (macOS, X07 v0.1.89)

Single-machine snapshot measured on March 17, 2026 (Apple M4 Max, macOS 26.3 / Darwin 25D125) using a local release build of `x07-host-runner` from `x07 v0.1.89` (`cargo build -p x07-host-runner --release`) with default settings (100KB input, 5 iterations, 2 warmup).

Compared with the previous published `v0.1.9` snapshot, this warmed `v0.1.89` run improved mean X07 runtime by about 10.7% in host-runner mode and 7.4% in direct mode, kept binary size effectively flat, and increased peak RSS by about 2.9%.

Raw results:

- `snapshots/2026-03-17_macos_x07-0.1.89_host.json`
- `snapshots/2026-03-17_macos_x07-0.1.89_direct.json`
- historical reference: `snapshots/2026-02-09_macos_x07-0.1.9_host.json`
- historical reference: `snapshots/2026-02-09_macos_x07-0.1.9_direct.json`

### Host Runner Mode (100KB input)

| Benchmark | X07 | C | Rust | Go | C vs X07 | Rust vs X07 | Go vs X07 |
|-----------|-----|---|------|----|----------|-------------|-----------|
| sum_bytes | 9.58ms | 2.93ms | 1.96ms | 2.44ms | 3.27x | 4.89x | 3.92x |
| word_count | 9.29ms | 2.96ms | 2.05ms | 2.55ms | 3.13x | 4.53x | 3.64x |
| rle_encode | 9.39ms | 2.80ms | 2.10ms | 2.45ms | 3.35x | 4.46x | 3.84x |
| byte_freq | 9.58ms | 2.86ms | 2.12ms | 2.37ms | 3.35x | 4.52x | 4.04x |
| fibonacci | 9.41ms | 1.53ms | 1.75ms | 2.01ms | 6.15x | 5.37x | 4.68x |

### Direct Binary Mode (100KB input)

| Benchmark | X07 | C | Rust | Go | C vs X07 | Rust vs X07 | Go vs X07 |
|-----------|-----|---|------|----|----------|-------------|-----------|
| sum_bytes | 2.06ms | 2.90ms | 2.15ms | 2.53ms | 0.71x | 0.96x | 0.81x |
| word_count | 2.43ms | 3.01ms | 2.12ms | 2.58ms | 0.81x | 1.14x | 0.94x |
| rle_encode | 2.22ms | 2.94ms | 2.10ms | 2.66ms | 0.76x | 1.06x | 0.84x |
| byte_freq | 3.21ms | 2.97ms | 2.12ms | 2.58ms | 1.08x | 1.51x | 1.24x |
| fibonacci | 1.62ms | 1.50ms | 1.81ms | 2.15ms | 1.07x | 0.89x | 0.75x |

### Compile Times (warm cache)

The compile figures below come from the warmed host-runner snapshot so they do not include the first-run cold-cache skew.

| Benchmark | X07 | C | Rust | Go |
|-----------|-----|---|------|----|
| sum_bytes | **20.0ms** | 43.6ms | 87.5ms | 1246.8ms |
| word_count | **21.1ms** | 41.1ms | 85.4ms | 1249.1ms |
| rle_encode | **21.6ms** | 41.4ms | 89.5ms | 1228.7ms |
| byte_freq | **23.4ms** | 41.5ms | 94.7ms | 1254.1ms |
| fibonacci | **22.2ms** | 38.7ms | 84.2ms | 1257.8ms |

On this run, X07 compiles about 1.7-2.2x faster than C and about 3.8-4.1x faster than Rust.

### Build Size & Memory

| Metric | X07 | C | Rust | Go |
|--------|-----|---|------|----|
| Binary Size | ~34.0-34.1 KiB | ~32.8-33.0 KiB | ~432-449 KiB | ~1337 KiB |
| Peak RSS | ~1.3-1.6 MiB | ~1.3-1.6 MiB | ~1.5-1.7 MiB | ~3.7-4.4 MiB |
