# Forced static body

Forced static body is a middleware plugin for [Traefik](https://github.com/traefik/traefik) which rewrites the HTTP response body
with a constant string passed in the config.

## Configuration

### Static

```toml
[pilot]
  token = "xxxx"

[experimental.plugins.forcedbody]
  modulename = "github.com/benoitg31/traefik-forced-plugin"
  version = "v1.0.2"
```


