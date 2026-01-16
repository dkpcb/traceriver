# Contributing

## Core Rules

- Do not violate Clean Architecture
- Do not introduce external dependencies into the domain layer
- Never edit generated code

## Coding Guidelines

- Keep implementations small
- Use meaningful names
- Comments should explain *why*, not *what*

## Prohibited Practices

- Calling repositories directly from controllers
- Introducing HTTP, OpenAPI, or database concerns into the domain
- Manually editing files under the apigen directory
