# Mountebank-record-to-pact-file

## Start Stub

```shell
docker-compose up -d
```

## Run Test

### Run Test from Newman

```shell
newman run tests/JsonPlaceholder-Todo-1.json
```

### Run Test from Pact Verifier

```shell
docker-compose -f docker-compose.test.yaml up
```
