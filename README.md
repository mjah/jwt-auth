# JWT Authentication Microservice

[![GoDoc Badge]][GoDoc] [![GoReportCard Badge]][GoReportCard]

[GoDoc]: https://godoc.org/github.com/mjah/jwt-auth
[GoDoc Badge]: https://godoc.org/github.com/mjah/jwt-auth?status.svg
[GoReportCard]: https://goreportcard.com/report/github.com/mjah/jwt-auth
[GoReportCard Badge]: https://goreportcard.com/badge/github.com/mjah/jwt-auth

A simple JWT based authentication server.

Features:

* Token based stateless authentication.
* User sign up, sign in, sign out, update, confirm, delete, and reset password.
* Send welcome, confirm, and reset password emails.
* Issue access and refresh token on signin.
* Refresh token revocation on sign out and ability to revoke all refresh tokens on sign out everywhere.
* JWT signed using RS256 signing algorithm for asymmetric encryption.

## Quick Start

This section will guide you through getting this project up and running as quickly as possible. It will only require that [docker](https://www.docker.com/) is installed and nothing else.

**This quick start is recommended for experimenting/testing purposes only.**

### Run Postgresql

```sh
docker run -it --rm \
--name postgres \
-p 5432:5432 \
-e POSTGRES_DB=jwt-auth \
-e POSTGRES_USER=postgres \
-e POSTGRES_PASSWORD=postgres \
postgres:11-alpine
```

Note: If you want to use an existing Postgresql setup with same port, then ensure that the jwt-auth database is created.

### Run Rabbitmq

```sh
docker run -it --rm \
--name rabbitmq \
-p 5672:5672 \
-p 15672:15672 \
rabbitmq:3-management
```

### Generate an RSA keypair with a 2048 bit private key

```sh
JA_KEYS_DIR="$HOME/.jwt-auth/keys"

mkdir -p "$JA_KEYS_DIR"
openssl genpkey -algorithm RSA -out "$JA_KEYS_DIR/private_key.pem" -pkeyopt rsa_keygen_bits:2048
openssl rsa -pubout -in "$JA_KEYS_DIR/private_key.pem" -out "$JA_KEYS_DIR/public_key.pem"
```

### Run application

```sh
docker run \
--network host \
--volume "$JA_KEYS_DIR":/keys \
-e JA_TOKEN_PRIVATE_KEY_PATH=/keys/private_key.pem \
-e JA_TOKEN_PUBLIC_KEY_PATH=/keys/public_key.pem \
docker.pkg.github.com/mjah/jwt-auth/jwt-auth:latest serve
```

Go to [localhost:9096/ping](http://localhost:9096/ping), if you receive a pong then you are now up and running.

## Configuration

See the [config.example.yml](https://github.com/mjah/jwt-auth/blob/master/config.example.yml) file for an example of the configuration.

Environment variables are also supported. This will be the configuration name in all capital letters, 'JA\_' prefixed, and '.' replaced with '\_'. E.g. *email.smtp_host* becomes *JA_EMAIL_SMTP_HOST*.

## API

### Public Routes

<table>
  <thead>
    <tr>
      <th>Path</th>
      <th>Method</th>
      <th>JSON Data</th>
      <th>Shared Error Responses</th>
      <th>Further Error Responses</th>
    </tr>
  </thead>
  <tbody>
    <tr>
      <td>/v1/auth/signup</td>
      <td>POST</td>
      <td>
        email (string, required)<br>
        username (string, required)<br>
        password (string, required)<br>
        first_name (string, required)<br>
        last_name (string, required)<br>
        confirm_email_url (string, required)
      </td>
      <td rowspan=0>
        DetailsInvalid<br>
        DatabaseConnectionFailed<br>
        DatabaseQueryFailed
      </td>
      <td>
        EmailAndUsernameAlreadyExists<br>
        EmailAlreadyExists<br>
        UsernameAlreadyExists<br>
        DefaultRoleAssignFailed<br>
        PasswordGenerationFailed<br>
        MessageQueueFailed
      </td>
    </tr>
    <tr>
      <td>/v1/auth/signin</td>
      <td>POST</td>
      <td>
        email (string, required)<br>
        password (string, required)<br>
        remember_me (bool, required)
      </td>
      <td>
        EmailDoesNotExist<br>
        PasswordInvalid<br>
        AccessTokenIssueFailed<br>
        RefreshTokenIssueFailed
      </td>
    </tr>
    <tr>
      <td>/v1/auth/confirm_email</td>
      <td>POST</td>
      <td>
        email (string, required)<br>
        confirm_email_token (string, required)
      </td>
      <td>
        EmailDoesNotExist<br>
        EmailAlreadyConfirmed<br>
        UUIDTokenDoesNotMatch<br>
        UUIDTokenExpired
      </td>
    </tr>
    <tr>
      <td>/v1/auth/reset_password</td>
      <td>POST</td>
      <td>
        email (string, required)<br>
        reset_password_token (string, required)<br>
        password (string, required)
      </td>
      <td>
        EmailDoesNotExist<br>
        UUIDTokenDoesNotMatch<br>
        UUIDTokenExpired<br>
        PasswordGenerationFailed
      </td>
    </tr>
    <tr>
      <td>/v1/auth/send_confirm_email</td>
      <td>POST</td>
      <td>
        email (string, required)<br>
        confirm_email_url (string, required)
      </td>
      <td>
        EmailDoesNotExist<br>
        EmailAlreadyConfirmed<br>
        MessageQueueFailed
      </td>
    </tr>
    <tr>
      <td>/v1/auth/send_reset_password</td>
      <td>POST</td>
      <td>
        email (string, required)<br>
        reset_password_url (string, required)
      </td>
      <td>
        EmailDoesNotExist<br>
        MessageQueueFailed
      </td>
    </tr>
  </tbody>
</table>

### Private Routes

Accessing private routes will require the refresh token in the authorization bearer.

<table>
  <thead>
    <tr>
      <th>Path</th>
      <th>Method</th>
      <th>JSON Data</th>
      <th>Shared Error Responses</th>
      <th>Further Error Responses</th>
    </tr>
  </thead>
  <tbody>
    <tr>
      <td>/v1/auth/signout</td>
      <td>GET</td>
      <td></td>
      <td rowspan=0>
        AuthorizationBearerTokenEmpty<br>
        JWTTokenInvalid<br>
        DatabaseConnectionFailed<br>
        DatabaseQueryFailed<br>
        UserDoesNotExist<br>
        UserIsNotActive<br>
        RefreshTokenIsRevoked
      </td>
      <td></td>
    </tr>
    <tr>
      <td>/v1/auth/signout_all</td>
      <td>GET</td>
      <td></td>
      <td></td>
    </tr>
    <tr>
      <td>/v1/auth/refresh_token</td>
      <td>GET</td>
      <td></td>
      <td>
        UserDoesNotExist<br>
        AccessTokenIssueFailed
      </td>
    </tr>
    <tr>
      <td>/v1/auth/update</td>
      <td>PATCH</td>
      <td>
        email (string, optional)<br>
        username (string, optional)<br>
        password (string, optional)<br>
        first_name (string, optional)<br>
        last_name (string, optional)
      </td>
      <td>
        DetailsInvalid<br>
        UserDoesNotExist<br>
        EmailAndUsernameAlreadyExists<br>
        EmailAlreadyExists<br>
        UsernameAlreadyExists<br>
        PasswordGenerationFailed
      </td>
    </tr>
    <tr>
      <td>/v1/auth/delete</td>
      <td>DELETE</td>
      <td></td>
      <td></td>
    </tr>
  </tbody>
</table>

### Error Responses

Error responses and their codes can be seen in [errors/codes.go](https://github.com/mjah/jwt-auth/blob/master/errors/codes.go)

## Example Client

To see an implementation of the jwt-auth API, please see the following [example client](https://github.com/mjah/jwt-auth-client-example).

## Contributing

Any feedback and pull requests are welcome and highly appreciated. Please open an issue first if you intend to send in a larger pull request or want to add additional features.

## License

This project is licensed under the MIT License - see the [LICENSE](https://github.com/mjah/jwt-auth/blob/master/LICENSE) file for details.
