# Architecture

## Core Principles

This project strictly follows **Clean Architecture**.

Allowed dependency directions:

controller → usecase → domain
usecase → repository (interface)

yaml
コードをコピーする

All external concerns (HTTP, OpenAPI, databases, blockchain) must remain at the outermost layer.

---

## Directory Structure

services/api/
main.go # Application entry point (DI + startup only)

openapi.yaml # API definition (single source of truth)
oapi-codegen.yaml

apigen/ # OpenAPI generated code (DO NOT EDIT)
api.gen.go

domain/ # Business concepts and rules
repository/ # Persistence abstractions
usecase/ # Application use cases
controller/ # HTTP adapters

markdown
コードをコピーする

---

## Responsibilities by Layer

### domain
- The core of the business
- Defines entities, identities, and invariants
- Must not know about HTTP, databases, OpenAPI, or frameworks

### repository
- Persistence abstraction
- Interface definitions only

### usecase
- Application-specific business flows
- Creates and orchestrates domain objects
- Owns ID generation and time handling

### controller
- HTTP and OpenAPI-specific logic
- Request/response translation
- Calls usecases

### apigen
- Code generated from OpenAPI
- **Must never be manually edited**
- Responsible for routing and validation

---

## OpenAPI Policy

- OpenAPI is the only authoritative API specification
- Controllers must use OpenAPI-generated types
- Routing must be handled by generated code