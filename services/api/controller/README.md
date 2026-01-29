# controller/

HTTP adapters that translate OpenAPI requests to usecase calls.

**Responsibilities:**
- Request/response translation
- HTTP-specific logic
- Call usecases
- Must NOT call repositories directly

**Dependencies:**
- Can depend on: usecase, domain, apigen
- Must NOT depend on: infrastructure
