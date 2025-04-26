# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [v0.2.1] - 2025-04-26
### Changed
- `NewTimeout` now returns a cancel function and an error channel

## [v0.2.0] - 2025-04-26
### Added
- `ReduceSeq`

### Changed
- `DateTimeFromStringPanic` -> `MustDateTimeFromString`
- `ResultOrPanic` -> `Must`
- `NilOrPanic` -> `NoErr`

### Removed
- `EnvConfig` (replaced by https://github.com/DAtek/env)
- `Logger` (there are very good quality loggers out there, eg. https://github.com/phuslu/log)
- `NewTimeoutMs`


## [v0.1.3] - 2023-04-17
### Changed
- Use time.Timer in `NewTimeout()`

## [v0.1.2] - 2023-04-11
### Fixed
- Make reduce function public

## [v0.1.1] - 2023-04-11 [YANKED]
### Added
- `func reduce[T any](arr []T, f func(accumulator, currentValue T) T) T`

## [v0.1.0] - 2023-04-07
