# usecase/

Application-specific business flows.

**Responsibilities:**
- Orchestrate domain objects
- ID generation
- Time handling
- Transaction boundaries

**Dependencies:**
- Can depend on: domain, repository (interfaces only)
- Must NOT depend on: controller, infrastructure, apigen
