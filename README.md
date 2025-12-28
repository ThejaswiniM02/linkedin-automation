# LinkedIn Automation – Scoped Proof of Concept (Go + Rod)

## Overview
This repository contains a scoped proof-of-concept demonstrating browser automation, human-like behavior simulation, and basic anti-detection techniques using Go and the Rod library.
⚠️ Disclaimer (Important)

This project is a technical proof-of-concept built strictly for educational and evaluation purposes.
It demonstrates browser automation architecture, stealth techniques, and human-like interaction patterns.

Automating LinkedIn violates their Terms of Service.
This project must not be used on real accounts or in production.

## Features Implemented
- Go-based modular architecture
- Browser automation using Rod
- User-agent spoofing
- navigator.webdriver masking
- Randomized human-like delays
- Environment-based configuration

## Stealth Techniques
Implemented:
- Browser fingerprint masking (webdriver flag)
- Randomized timing patterns

Planned (documented but not implemented):
- Bézier mouse movement
- Typing simulation with errors
- Scroll behavior modeling
- Rate limiting & scheduling
- Session persistence

## Project Structure
- cmd/app: application entry point
- config: configuration loading
- stealth: anti-detection logic
- utils: human behavior utilities

## Future Improvements
- SQLite-based state persistence
- Search & targeting workflows
- Connection request tracking
- Messaging templates
- Advanced fingerprint spoofing
