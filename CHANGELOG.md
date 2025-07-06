# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

## [1.1.1] - 2025-07-06

### Changed
- Simplified CI/CD infrastructure by consolidating workflows into a single main workflow
- Improved dependency verification with explicit checks for go.mod and go.sum
- Enhanced error handling for tests and security scanning
- Ensured consistent CI setup across all branches

### Removed
- Removed redundant workflow files (security-scan.yml, gosec-matrix.yml, golangci-lint.yml, go-test.yml)
- Removed dependabot.yml to reduce maintenance overhead

## [1.1.0] - 2025-07-06

### Added
- Added new GitHub Actions workflows for dedicated linting and testing
- Added `.github/dependabot.yml` for automated dependency updates
- Added new `gosec-matrix.yml` workflow for security scanning on multiple platforms

### Changed
- Improved CI/CD workflows with robust security scanning
- Fixed dependency handling in GitHub Actions workflows
- Enhanced error reporting in security scanning jobs
- Refactored code to address linter warnings
- Updated to Go 1.24 in all workflows

### Fixed
- Fixed issues with gosec security scanning on different platforms
- Resolved dependency verification problems in GitHub Actions

## [1.0.0] - 2025-06-01

### Added
- Initial implementation of Go STL library
- Set data structure with comprehensive operations
- MultiSet data structure with frequency tracking
- MultiMap data structure for one-to-many relationships
- Deque (double-ended queue) with O(1) operations
- Binary Search Tree (BST) with self-balancing
- Trie (prefix tree) for string operations
- Graph data structure with various algorithms
- TreeMap (ordered map) with logarithmic operations
- Stack data structure with LIFO operations
- Queue data structure with FIFO operations
- PriorityQueue with customizable ordering
- Comprehensive test coverage
- Extensive documentation and examples
- Functional programming support across all data structures

### Features
- Generic type support for all data structures
- Comprehensive API with standard operations
- Advanced algorithms (graph traversal, trie pattern matching, etc.)
- Memory-efficient implementations
- Cross-platform compatibility

## [0.1.0] - 2025-07-05

### Added
- Initial release of Go STL library
- 11 core data structures implemented
- Complete documentation and examples
- MIT License
- Contributing guidelines
- Code of Conduct

### Technical Details
- Go 1.21+ compatibility
- Generic implementations using Go generics
- Comprehensive test suite
- Performance-optimized algorithms
- Memory-efficient data structures

---

## Version History

- **0.1.0**: Initial release with all core data structures
- **Unreleased**: Development version with latest features

## Migration Guide

### From Unreleased to 0.1.0
- No breaking changes - this is the initial release
- All APIs are stable and ready for production use

## Deprecation Policy

- Deprecated features will be marked with `// Deprecated:` comments
- Deprecated features will be removed after 2 major versions
- Migration guides will be provided for breaking changes

## Support

For support and questions:
- Open an issue on GitHub
- Check the documentation
- Review examples in the `examples/` directory