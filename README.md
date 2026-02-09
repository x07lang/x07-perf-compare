# x07-perf-compare

Performance comparison benchmarks: **X07 vs C vs Rust vs Go**

This repo contains benchmark programs to compare performance and memory usage
between X07, C, Rust, and Go implementations of identical algorithms.

Community:

- Discord: https://discord.gg/59xuEuPN47
- Email: support@x07lang.org

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

## Performance Comparison (macOS, X07 v0.1.9)

Single-machine snapshot measured on February 9, 2026 (Apple M4 Max, macOS 15.2 / Darwin 25.2.0) using a local release build of `x07-host-runner` from `x07 0.1.9` (`cargo build -p x07-host-runner --release`) with default settings (100KB input, 5 iterations, 2 warmup).

Raw results:

- `snapshots/2026-02-09_macos_x07-0.1.9_host.json`
- `snapshots/2026-02-09_macos_x07-0.1.9_direct.json`

### Host Runner Mode (100KB input)

| Benchmark | X07 | C | Rust | Go | C vs X07 | Rust vs X07 | Go vs X07 |
|-----------|-----|---|------|----|----------|-------------|-----------|
| sum_bytes | 10.08ms | 3.10ms | 2.14ms | 2.64ms | 3.25x | 4.72x | 3.82x |
| word_count | 10.05ms | 3.86ms | 2.41ms | 3.34ms | 2.60x | 4.18x | 3.01x |
| rle_encode | 14.07ms | 3.97ms | 2.71ms | 2.55ms | 3.54x | 5.20x | 5.51x |
| byte_freq | 9.79ms | 3.14ms | 2.27ms | 2.79ms | 3.12x | 4.32x | 3.51x |
| fibonacci | 9.97ms | 1.80ms | 2.00ms | 2.26ms | 5.53x | 4.99x | 4.42x |

### Direct Binary Mode (100KB input)

| Benchmark | X07 | C | Rust | Go | C vs X07 | Rust vs X07 | Go vs X07 |
|-----------|-----|---|------|----|----------|-------------|-----------|
| sum_bytes | 2.20ms | 2.98ms | 2.13ms | 2.54ms | 0.74x | 1.03x | 0.87x |
| word_count | 2.48ms | 3.22ms | 2.24ms | 2.63ms | 0.77x | 1.11x | 0.94x |
| rle_encode | 2.34ms | 3.19ms | 2.36ms | 2.72ms | 0.73x | 0.99x | 0.86x |
| byte_freq | 3.31ms | 3.09ms | 2.20ms | 2.50ms | 1.07x | 1.50x | 1.32x |
| fibonacci | 2.04ms | 1.82ms | 1.85ms | 2.27ms | 1.12x | 1.10x | 0.90x |

### Compile Times (warm cache)

| Benchmark | X07 | C | Rust | Go |
|-----------|-----|---|------|----|
| sum_bytes | **27.5ms** | 50.2ms | 98.6ms | 1443.2ms |
| word_count | **29.1ms** | 42.1ms | 96.4ms | 1453.6ms |
| rle_encode | **28.3ms** | 42.4ms | 100.3ms | 1352.4ms |
| byte_freq | **31.3ms** | 44.1ms | 104.3ms | 1424.9ms |
| fibonacci | **26.5ms** | 39.6ms | 94.1ms | 1395.4ms |

On this run, X07 compiles ~1.3-1.8x faster than C and ~3.1-3.7x faster than Rust.

### Build Size & Memory

| Metric | X07 | C | Rust | Go |
|--------|-----|---|------|----|
| Binary Size | ~34.0 KiB | ~32.8-33.0 KiB | ~432-449 KiB | ~1337 KiB |
| Peak RSS | ~1.3-1.6 MiB | ~1.3-1.5 MiB | ~1.5-1.7 MiB | ~3.6-4.4 MiB |
