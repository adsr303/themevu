# Copilot instructions for themevu

## Purpose

This file gives repository-specific instructions to help Copilot and future assistant sessions understand, build, test, and navigate the codebase quickly.

## Build, test, and lint commands

- Build the project (all packages):
  - `go build ./...`
- Build and run a single executable for manual testing:
  - `go build -o themevu ./ && ./themevu -h`
- Run all tests:
  - `go test ./...`
- Run package tests only:
  - `go test ./colors`
  - `go test ./themes`
- Run a single test by name (package and `-run` regex):
  - `go test ./colors -run TestPermutateRGB`
  - `go test ./themes -run TestParseGogh`
  - Add `-v` for verbose output: `go test ./colors -run TestPermutateRGB -v`
- Linting: No repository-level linter config detected. Recommended quick checks:
  - `go vet ./...`
  - Add golangci-lint (optional) and run: `golangci-lint run`

## High-level architecture

- `main.go`
  - CLI entrypoint. Flags:
    - `-fg`: show colored text on default background (prints color codes instead of blocks)
    - `-permutate`: generate RGB permutations from a 6-digit hex and display each
    - `-theme <file>`: parse and render a theme file (Windows Terminal JSON or Gogh YAML)
  - Behavior: reads stdin lines when -theme is not provided, otherwise parses the theme file and renders a swatch using the simulation package.

- packages/roles
  - `colors/`
    - PermutateRGB validates 6-digit hex colors (accepts leading `#`) and returns permutations of the three RGB pairs.
    - Tests: `colors/permutate_test.go` covers valid and invalid inputs.
  - `themes/`
    - Parses theme formats (Gogh YAML, Windows Terminal JSON). Theme structs use `theme:"color_##"` tags to mark numbered colors.
    - `NumberedColors` uses reflection to extract exactly 16 numbered colors from a theme struct; callers expect a 16-element slice.
    - Tests use `go:embed` for YAML fixtures (see `Nanosecond.yml`).
  - `simulation/`
    - Renders theme titles and color tables in terminal using `charm.land/lipgloss`.
    - `PrintAsTable` expects a 16-color slice (0..15) and arranges them into regular/bright columns.

## Key conventions and gotchas

- Theme struct tags
  - Fields that represent numbered colors must carry a struct tag in the format `theme:"color_##"` where `##` is 01..16. `NumberedColors` relies on this to map fields to indexes.

- Hex color formats
  - `PermutateRGB` requires exactly six hexadecimal digits (with or without a leading `#`). 3-digit shorthand (e.g., `#abc`) is rejected by tests and code.

- Tests and fixtures
  - Use `go:embed` for theme fixtures if adding tests like `themes/gogh_test.go`. Tests reference embedded files from the package directory.

- Rendering expectations
  - `simulation.PrintAsTable` and `PrintTitle` assume inputs contain foreground/background/cursor colors and a 16-element color slice. Keep that contract when modifying themes or adding parsers.

- Package layout
  - Keep parser logic in `themes/`, color utilities in `colors/`, and terminal rendering in `simulation/` for separation of concerns.

## Files to check when modifying behavior

- `main.go` — CLI parsing and high-level program flow
- `themes/themes.go`, `themes/gogh.go`, `themes/winterm.go` — parsing and color extraction
- `colors/permutate.go` — color validation and permutation logic
- `simulation/simulation.go` — output layout and lipgloss usage
