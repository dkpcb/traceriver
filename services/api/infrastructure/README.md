# infrastructure/

Database implementations and external service adapters.

**Responsibilities:**
- Implement repository interfaces
- Database-specific logic (MySQL)
- External service clients

**Dependencies:**
- Can depend on: domain, repository
- Must NOT depend on: usecase, controller, apigen
