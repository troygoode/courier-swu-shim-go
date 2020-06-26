# Contributing

PR's welcome.

## Testing

Create `env.sh` locally:

```bash
export COURIER_AUTH_TOKEN=TODO
export SHIM_TEMPLATE_ID=TODO
export SHIM_EMAIL_TO=TODO
```

Then run:

```bash
source ./env.sh
go test
```

## Releasing New Versions

```bash
git tag -a v<VERSION> -m v<VERSION>
git push origin v<VERSION>
```
