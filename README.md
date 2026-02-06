# x07-perf-compare

Performance comparison benchmarks: **X07 vs C vs Rust**

This repo contains benchmark programs to compare performance and memory usage
between X07, C, and Rust implementations of identical algorithms.

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
├── rust/              # Rust implementations (single-file)
└── rust_cargo/        # Rust implementations with Cargo deps
```

## Requirements

- Python 3.10+
- Rust toolchain (rustc)
- C compiler (cc/gcc/clang)
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
- All programs use stdin/stdout for I/O
- Output correctness is verified (outputs must match across languages)

## Regex Benchmarks

This repo also contains `regex_*` programs, but they depend on the external `ext-regex` package, which is not published to the registry yet. They are not included in the default benchmark set.

## Performance Comparison (macOS, X07 v0.0.94)

Single-machine snapshot measured on February 6, 2026 using a local release build of `x07-host-runner` from `x07 v0.0.94` (`cargo build -p x07-host-runner --release`) with default settings (100KB input, 5 iterations, 2 warmup).

### Host Runner Mode (100KB input)

| Benchmark | X07 | C | Rust | C vs X07 | Rust vs X07 |
|-----------|---------|---|------|--------------|-----------------|
| sum_bytes | 10.76ms | 3.56ms | 2.61ms | 3.03x | 4.13x |
| word_count | 10.88ms | 4.44ms | 2.83ms | 2.45x | 3.85x |
| rle_encode | 10.99ms | 4.31ms | 2.92ms | 2.55x | 3.76x |
| byte_freq | 10.16ms | 4.13ms | 3.00ms | 2.46x | 3.39x |
| fibonacci | 10.99ms | 2.16ms | 2.63ms | 5.09x | 4.17x |

### Direct Binary Mode (100KB input)

| Benchmark | X07 | C | Rust | C vs X07 | Rust vs X07 |
|-----------|---------|---|------|--------------|-----------------|
| sum_bytes | 2.72ms | 3.61ms | 2.69ms | 0.75x | 1.01x |
| word_count | 2.75ms | 3.61ms | 2.59ms | 0.76x | 1.06x |
| rle_encode | 2.67ms | 3.51ms | 2.57ms | 0.76x | 1.04x |
| byte_freq | 3.65ms | 3.50ms | 2.69ms | 1.04x | 1.36x |
| fibonacci | 2.06ms | 1.81ms | 2.19ms | 1.14x | 0.94x |

### Compile Times

| Benchmark | X07 | C | Rust |
|-----------|---------|---|------|
| sum_bytes | **11.7ms** | 46.0ms | 95.7ms |
| word_count | **13.4ms** | 43.7ms | 92.5ms |
| rle_encode | **13.6ms** | 45.2ms | 99.1ms |
| byte_freq | **12.8ms** | 44.2ms | 99.7ms |
| fibonacci | **13.6ms** | 43.1ms | 95.4ms |

On this run, X07 compiles ~3.2-3.9x faster than C and ~6.9-8.2x faster than Rust.

### Build Size & Memory

| Metric | X07 | C | Rust |
|--------|---------|---|------|
| Binary Size | ~34.0 KiB | ~32.8-33.0 KiB | ~432-449 KiB |
| Peak RSS | ~1.3-1.6 MiB | ~1.3-1.5 MiB | ~1.5-1.7 MiB |
