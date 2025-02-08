# 🤝 Contributing to CLI System Status

Thank you for your interest in contributing to CLI System Status! We're excited to have you join our community.

## 🌟 Ways to Contribute

### 🐛 Bug Reports
- Use the GitHub issue tracker
- Include detailed steps to reproduce
- Share system information and logs
- Attach screenshots if applicable

### 💡 Feature Requests
- Check existing issues first
- Describe the feature in detail
- Explain why it would be useful
- Consider implementation details

### 🔧 Code Contributions

#### 📋 Prerequisites
- Go 1.16 or higher
- Git
- Make (optional, for running scripts)

#### 🚀 Getting Started
1. Fork the repository
2. Clone your fork:
   ```bash
   git clone https://github.com/Amul-Thantharate/cli-systemstatus.git
   ```
3. Create a new branch:
   ```bash
   git checkout -b feature/your-feature-name
   ```

#### 🏗️ Development Workflow
1. Write your code
2. Add tests for new features
3. Ensure all tests pass:
   ```bash
   go test ./...
   ```
4. Run linting:
   ```bash
   golangci-lint run
   ```

#### 📝 Commit Guidelines
- Use conventional commits
- Format: `type(scope): message`
- Types: feat, fix, docs, style, refactor, test, chore
- Keep commits focused and atomic

Example:
```bash
feat(metrics): add CPU temperature monitoring
fix(display): correct memory usage calculation
docs(readme): update installation instructions
```

#### 🔍 Pull Request Process
1. Update documentation if needed
2. Add yourself to CONTRIBUTORS.md
3. Reference any related issues
4. Wait for CI checks to pass
5. Address review comments

### 📚 Documentation
- Fix typos and unclear sections
- Add examples and use cases
- Improve error messages
- Update API documentation

### 🌐 Translations
- Add new language translations
- Review existing translations
- Keep translations up to date

## 🎨 Code Style

### Go Guidelines
- Follow standard Go conventions
- Use `gofmt` for formatting
- Add comments for exported functions
- Keep functions focused and small

### Testing
- Write unit tests for new code
- Maintain test coverage
- Add integration tests when needed
- Use table-driven tests

## 🚀 Future Additions

### 📊 Metrics Enhancement
- [ ] GPU monitoring support
- [ ] Container metrics (Docker, K8s)
- [ ] Custom metric plugins
- [ ] Advanced alerting system

### 🎯 Performance Optimization
- [ ] Reduced memory footprint
- [ ] Faster data collection
- [ ] Better concurrency handling
- [ ] Optimized data storage

### 🔧 Tools and Integration
- [ ] CLI configuration wizard
- [ ] Integration with monitoring services
- [ ] Export to various formats
- [ ] Remote monitoring capabilities

### 📱 User Interface
- [ ] Interactive TUI mode
- [ ] Custom dashboard layouts
- [ ] Real-time graphs
- [ ] Keyboard shortcuts

### 🔒 Security Features
- [ ] Encrypted data storage
- [ ] Authentication support
- [ ] Role-based access control
- [ ] Audit logging

## ⚖️ Code of Conduct

Please note that this project follows our [Code of Conduct](CODE_OF_CONDUCT.md). By participating, you are expected to uphold this code.

## 📜 License

By contributing, you agree that your contributions will be licensed under the same MIT License that covers the project.
