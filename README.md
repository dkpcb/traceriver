# river

river is an application designed to **slowly shorten the distance between people through the exchange of expressions**.

This product deliberately prioritizes:
- expressions that are *carefully observed*
over
- expressions that are *rapidly consumed*.

Users share artworks they have created (or generated expressions) with others.
Timelines are shared only among people they have met directly, or within a maximum of **2 hops** of social distance.

## Core Ideas

- An expression is something to be *looked at*, and something to be *exchanged*
- Connections should not scale infinitely (maximum 2 hops)
- The primary KPI is **scrolling speed (slowness)**
- Everyone can become the center of a small belief system — a “founder” of meaning

## Core Features (Initial)

- Posting expressions (Trace)
- Building relationships through exchanging expressions
- Timeline sharing with a 2-hop constraint
- API defined strictly by OpenAPI

## Tech Stack

- Go
- Echo
- OpenAPI (oapi-codegen)
- Clean Architecture

## Development Principles

- OpenAPI is the single source of truth
- Generated code and human-written code must be strictly separated
- The domain layer is the most important layer
