# oidc-copycat

Mock server for OpenID Connect (OIDC) testing

## Setup

TODO

## Goals

1. Help developers test and/or simulate OIDC authentication flows with minimal setup and configuration. Intended to be used for local development or with integration tests.

1. Offer a configurable admin interface for advanced scenarios.

1. Support simulating error scenarios & unique IdP implementations.

## Features

### Basic OIDC Endpoints

- [ ] `/.well-known/openid-configuration` discovery
- [ ] `/token` exchange
- [ ] `/authorize` - begin the authentication & receive authorization grant
- [ ] `/token` - Obtain access & ID token in exchange for auth grant
- [ ] `/userinfo` - get claims about authenticated user

### Server Configuration

- [ ] Manage user identities via a config file or admin UI
- [ ] Force calls to return specific errors for testing edge cases
