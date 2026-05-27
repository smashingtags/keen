# Security Policy

## Reporting a Vulnerability

**Do not file public GitHub issues for security vulnerabilities.**

Email **security@imogenlabs.ai** with:
- Description of the vulnerability
- Steps to reproduce
- Impact assessment

We will respond within 72 hours and credit reporters in release notes.

## Supported Versions

| Version | Supported |
|---------|-----------|
| Latest release | Yes |
| Older releases | No |

## Scope

Keen stores API credentials in environment variables only — never on disk. The local SQLite database contains cached Atlassian API responses and should be treated as sensitive data.
