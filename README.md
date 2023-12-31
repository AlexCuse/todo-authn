# todo-authn

This repo contains a sample application demonstrating use of [authn-server](https://github.com/keratin/authn-server).

The application has two parts:

1. A TODO api hosted at /api/v1/
2. A svelte-based front end hosted at /

Infrastructure is provided via postgres, redis and caddy (and authn).

# running

Set up your hosts file for a few todoauthn subdomains
```
127.0.0.1       todoauthn.com auth.todoauthn.com api.todoauthn.com swagger.todoauthn.com
```

Generate a certificate using certutil and make it available to caddy, then run through docker.
```shell
$ make ssl
$ make run
```

You may need to stop and start again to allow for postgres initialization.

Also may need to refresh the page a few time after registration/login until I get the authn client working better w svelte.
# local development

This repo uses a codegen-heavy workflow.

API is generated using https://github.com/deepmap/oapi-codegen
Database operations generated using https://sqlc.dev

`$make generate-all` will regenerate the API contract and database code.  
From there the compiler will tell you what to do.