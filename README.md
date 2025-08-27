# Backend Starter CLI

[![Go](https://img.shields.io/badge/Go-1.24.4-blue)](https://golang.org/)
[![Cobra](https://img.shields.io/badge/Cobra-v1.9.1-green)](https://github.com/spf13/cobra)
![GitHub](https://img.shields.io/badge/Status-Production%20Ready-brightgreen)

**Accelerate your backend development with terminal-driven project scaffolding and setup.**

**🚀 Production Ready** - This tool is stable and actively maintained.

## Introduction

Backend Starter CLI is your command-line companion for rapidly scaffolding backend projects across multiple frameworks. Designed for developers who value efficiency, it brings the power of automated project setup directly to your terminal, eliminating the repetitive tasks of initial project configuration.

With a single command, generate production-ready boilerplate code, install dependencies, and start building your application immediately.

## Features

* **Multi-Framework Support**: Express.js, Express TypeScript, FastAPI, Fiber, and Gin
* **Zero-Config Setup**: Automatic dependency installation and environment configuration
* **Production-Ready Templates**: Industry-standard project structure and best practices
* **Cross-Platform Compatibility**: Works seamlessly on macOS, Linux, and Windows

## Installation

### macOS

```bash
brew tap Walon-Foundation/backend-starter-cli
brew install --cask backend-starter
```

### Linux

1. Download the latest Linux binary from our [Releases page](https://github.com/Walon-Foundation/backend-starter-cli/releases)
2. Extract the archive:
   ```bash
   tar -xvf backend-starter_x.x.x_linux_amd64.tar.gz
   ```
3. Install to your PATH:
   ```bash
   sudo mv ./backend-starter /usr/local/bin/
   ```

### Windows

1. Download the latest Windows .tar.gz from our [Releases page](https://github.com/Walon-Foundation/backend-starter-cli/releases)
2. Extract the archive:
    ```
    tar -xvf backend-starter_x_x_x_windows_amd64.tar.gz
    ```
3. Add the executable to your system PATH

### Go Install (Universal)

```bash
go install github.com/Walon-Foundation/backend-starter@latest
```

## Usage

```bash
# Create a new project
backend-starter init

# Show help information
backend-starter-cli help

# Display version information
backend-starter-cli version
```

## Supported Frameworks

- **Express.js** - Minimal JavaScript setup with essential middleware
- **Express TypeScript** - Type-safe Express with pre-configured TypeScript
- **FastAPI** - Python framework with modern Python features
- **Fiber** - Express-inspired web framework built on Go's Fasthttp
- **Gin** - High-performance HTTP web framework for Go

## Project Structure

Each generated project includes:
- Production-ready boilerplate code
- Framework-specific configuration files
- Pre-concluded dependency management (package.json, go.mod, requirements.txt)
- Environment setup and gitignore files
- Example entry point with basic server setup

## Technologies

* **Go 1.24.4** - Core CLI implementation language
* **Cobra 1.9.1** - Powerful CLI framework
* **Framework-specific stacks** - Each template uses industry-standard dependencies

## Contributing

We welcome contributions! Please feel free to submit pull requests, report bugs, or suggest new features.

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## Authors

* **[walonCode](https://github.com/walonCode)** - Creator and maintainer

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Support

If you encounter any issues or have questions, please file an issue on our [GitHub Issues page](https://github.com/Walon-Foundation/backend-starter-cli/issues).

---

**Thank you for using Backend Starter CLI!** Start building amazing backend applications faster than ever before.