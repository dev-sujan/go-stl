# Contributing to Go STL

Thank you for your interest in contributing to Go STL! This document provides guidelines and information for contributors.

## Table of Contents

- [Code of Conduct](#code-of-conduct)
- [How Can I Contribute?](#how-can-i-contribute)
- [Development Setup](#development-setup)
- [Coding Standards](#coding-standards)
- [Testing](#testing)
- [Pull Request Process](#pull-request-process)
- [Release Process](#release-process)

## Code of Conduct

This project and everyone participating in it is governed by our Code of Conduct. By participating, you are expected to uphold this code.

## How Can I Contribute?

### Reporting Bugs

- Use the GitHub issue tracker
- Include a clear and descriptive title
- Provide detailed steps to reproduce the issue
- Include Go version, OS, and any relevant environment details
- Include code examples if applicable

### Suggesting Enhancements

- Use the GitHub issue tracker with the "enhancement" label
- Clearly describe the proposed feature
- Explain why this feature would be useful
- Include implementation suggestions if possible

### Contributing Code

- Fork the repository
- Create a feature branch (`git checkout -b feature/amazing-feature`)
- Make your changes following the coding standards
- Add tests for new functionality
- Ensure all tests pass
- Update documentation if needed
- Commit your changes with clear commit messages
- Push to your fork and submit a pull request

## Development Setup

### Prerequisites

- Go 1.21 or later
- Git

### Local Development

1. Fork and clone the repository:
   ```bash
   git clone https://github.com/yourusername/go-stl.git
   cd go-stl
   ```

2. Install dependencies:
   ```bash
   go mod download
   ```

3. Run tests:
   ```bash
   go test ./...
   ```

4. Run examples:
   ```bash
   go run examples/main.go
   ```

## Coding Standards

### General Guidelines

- Follow Go's official [Code Review Comments](https://github.com/golang/go/wiki/CodeReviewComments)
- Use meaningful variable and function names
- Write clear and concise comments
- Keep functions small and focused
- Use consistent formatting (run `gofmt`)

### Code Style

- Use `gofmt` for code formatting
- Follow Go naming conventions:
  - Use `CamelCase` for exported names
  - Use `camelCase` for unexported names
  - Use `UPPER_CASE` for constants
- Write comprehensive documentation for exported functions
- Use meaningful commit messages

### Documentation

- All exported functions must have Go doc comments
- Include usage examples in documentation
- Update README.md for new features
- Update IMPLEMENTATION_SUMMARY.md for new data structures

### Error Handling

- Always check for errors
- Return meaningful error messages
- Use `fmt.Errorf` for error wrapping
- Don't ignore errors

## Testing

### Test Requirements

- All new code must include tests
- Aim for at least 80% test coverage
- Include both unit tests and integration tests
- Test edge cases and error conditions

### Running Tests

```bash
# Run all tests
go test ./...

# Run tests with coverage
go test -cover ./...

# Run tests with verbose output
go test -v ./...

# Run benchmarks
go test -bench=. ./...
```

### Test Structure

- Test files should be named `*_test.go`
- Use descriptive test function names
- Group related tests using subtests
- Use table-driven tests for multiple test cases

Example:
```go
func TestSetOperations(t *testing.T) {
    tests := []struct {
        name     string
        input    []int
        expected int
    }{
        {"empty set", []int{}, 0},
        {"single element", []int{1}, 1},
        {"multiple elements", []int{1, 2, 3}, 3},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            set := NewSet[int]()
            for _, item := range tt.input {
                set.Add(item)
            }
            if set.Size() != tt.expected {
                t.Errorf("expected %d, got %d", tt.expected, set.Size())
            }
        })
    }
}
```

## Pull Request Process

### Before Submitting

1. Ensure your code follows the coding standards
2. Run all tests and ensure they pass
3. Update documentation if needed
4. Check that your changes don't break existing functionality

### Pull Request Guidelines

- Use a clear and descriptive title
- Provide a detailed description of changes
- Include any relevant issue numbers
- Add screenshots or examples if applicable
- Ensure the PR passes all CI checks

### Review Process

- All PRs require at least one review
- Address review comments promptly
- Keep discussions constructive and respectful
- Be open to feedback and suggestions

## Release Process

### Versioning

We follow [Semantic Versioning](https://semver.org/) (SemVer):

- **MAJOR** version for incompatible API changes
- **MINOR** version for backwards-compatible functionality additions
- **PATCH** version for backwards-compatible bug fixes

### Release Checklist

- [ ] All tests pass
- [ ] Documentation is up to date
- [ ] Examples work correctly
- [ ] Version is updated in `go.mod`
- [ ] Release notes are prepared
- [ ] Tag is created with version number

### Creating a Release

1. Update version in `go.mod`
2. Create a release branch
3. Update CHANGELOG.md
4. Create a pull request
5. After review, merge and create a tag
6. Publish to GitHub releases

## Getting Help

- Open an issue for bugs or feature requests
- Join our discussions in GitHub Discussions
- Check existing issues and pull requests
- Review the documentation

## Recognition

Contributors will be recognized in:
- The project README
- Release notes
- GitHub contributors page

Thank you for contributing to Go STL! 