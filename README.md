# JWT Authentication Microservice

[![GoDoc Badge]][GoDoc] [![GoReportCard Badge]][GoReportCard]

[GoDoc]: https://godoc.org/github.com/mjah/jwt-auth
[GoDoc Badge]: https://godoc.org/github.com/mjah/jwt-auth?status.svg
[GoReportCard]: https://goreportcard.com/report/github.com/mjah/jwt-auth
[GoReportCard Badge]: https://goreportcard.com/badge/github.com/mjah/jwt-auth

## Features

## Prerequisite

### Generate an RSA keypair with a 2048 bit private key

Private key:

```sh
openssl genpkey -algorithm RSA -out private_key.pem -pkeyopt rsa_keygen_bits:2048
```

Public key:

```sh
openssl rsa -pubout -in private_key.pem -out public_key.pem
```

### Configuration

Configuration Name | Environment Name | Type | Default | Description
---|---|---|---|---
environment | JA_ENVIRONMENT | string | development | ...
log_level | JA_LOG_LEVEL | string | debug | ...
log_email | JA_LOG_EMAIL | bool | true | ...
account.password_cost | JA_ACCOUNT_PASSWORD_COST | int | 11 | ...
account.confirm_token_expires | JA_ACCOUNT_CONFIRM_TOKEN_EXPIRES | time.Duration | 24h00m | ...
account.confirm_token_endpoint | JA_ACCOUNT_CONFIRM_TOKEN_ENDPOINT | string | | ...
account.reset_password_token_expires | JA_ACCOUNT_RESET_PASSWORD_TOKEN_EXPIRES | time.Duration | 1h00m | ...
account.reset_password_token_endpoint | JA_ACCOUNT_RESET_PASSWORD_TOKEN_ENDPOINT | string | | ...
roles.define | JA_ROLES_DEFINE | []string | [admin member guest] | ...
roles.default | JA_ROLES_DEFAULT | string | guest | ...
serve.host | JA_SERVE_HOST | string | localhost | ...
serve.port | JA_SERVE_PORT | int | 9096 | ...
token.public_key_path | JA_TOKEN_PUBLIC_KEY_PATH | string || ...
token.private_key_path | JA_TOKEN_PRIVATE_KEY_PATH | string || ...
token.issuer | JA_TOKEN_ISSUER | string | jwt-auth | ...
token.access_token_expires | JA_TOKEN_ACCESS_TOKEN_EXPIRES | time.Duration | 5m | ...
token.refresh_token_expires | JA_TOKEN_REFRESH_TOKEN_EXPIRES | time.Duration | 8h00m | ...
token.refresh_token_expires_extended | JA_TOKEN_REFRESH_TOKEN_EXPIRES_EXTENDED | time.Duration | 8760h00m | ...
postgres.host | JA_POSTGRES_HOST | string | localhost | ...
postgres.port | JA_POSTGRES_PORT | int | 5432 | ...
postgres.username | JA_POSTGRES_USERNAME | string | postgres | ...
postgres.password | JA_POSTGRES_PASSWORD | string | psotgres | ...
postgres.database | JA_POSTGRES_DATABASE | string | jwt-auth | ...
amqp.host | JA_AMQP_HOST | string | localhost | ...
amqp.port | JA_AMQP_PORT | int | 5672 | ...
amqp.username | JA_AMQP_USERNAME | string | guest | ...
amqp.password | JA_AMQP_PASSWORD | string | guest | ...
email.smtp_host | JA_EMAIL_SMTP_HOST | string || ...
email.smtp_port | JA_EMAIL_SMTP_PORT | int || ...
email.smtp_username | JA_EMAIL_SMTP_USERNAME | string || ...
email.smtp_password | JA_EMAIL_SMTP_PASSWORD | string || ...
email.from_address | JA_EMAIL_FROM_ADDRESS | string || ...
email.from_name | JA_EMAIL_FROM_NAME | string || ...
email.test_receipient | JA_EMAIL_TEST_RECEIPIENT | string || ...

## API Routes

### Public Routes

Path | Method | JSON Data | Error Codes | Description
---|---|---|---|---
/v1/auth/signup | POST | email (string, required)<br>username (string, required)<br>password (string, required)<br>first_name (string, required)<br>last_name (string, required) || ...
/v1/auth/signin | POST | email (string, required)<br>password (string, required)<br>remember_me (bool, required) || ...
/v1/auth/confirm | POST | email (string, required)<br>confirm_token (string, required) || ...
/v1/auth/resetpassword | POST | email (string, required)<br>reset_password_token (string, required)<br>password (string, required) || ...
/v1/auth/send_confirm_email | POST | email (string, required) || ...
/v1/auth/send_resetpassword_email | POST | email (string, required) || ...

### Private Routes

Requires refresh token in authorization bearer.

Path | Method | JSON Data | Error Codes | Description
---|---|---|---|---
/v1/auth/signout | GET ||| ...
/v1/auth/signout_all | GET ||| ...
/v1/auth/refreshtoken | GET ||| ...
/v1/auth/update | PATCH | email (string, optional)<br>username (string, optional)<br>password (string, optional)<br>first_name (string, optional)<br>last_name (string, optional) || ...
/v1/auth/delete | DELETE ||| ...

### Error Codes

Error codes can be seen in [errors/codes.go](https://github.com/mjah/jwt-auth/blob/master/errors/codes.go)

## Example Client

To see an implementation of the jwt-auth API, please see the following [example client](https://github.com/mjah/jwt-auth-client-example).

## Contributing

Any feedback and pull requests are welcome and highly appreciated. Please open an issue first if you intend to send in a larger pull request or want to add additional features.
