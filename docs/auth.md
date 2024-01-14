## Overview

You can find how the auth endpoints works here.

Url prefix: `${apiUrl}/api/v1/auth

## Register

Path: `/register`

Method: `POST`

### Success Example Requests

- Request

```json
{
  "username": "safderun",
  "email": "burakberkkeskin@gmail.com",
  "password": "Test1234",
  "passwordVerify": "Test1234"
}
```

Response:

```json
{
  "message": "user created",
  "data": {
    "message": "user created."
  },
  "error": ""
}
```

### Request Field Details

Request Body Details:

- `username`: String value which should be max 16 character.

Long username response:

```json
[
  {
    "Field": "Username",
    "Tag": "max",
    "Value": "16"
  }
]
```

- `email`: Email value with valid email

Bad email response:

```json
[
  {
    "Field": "Email",
    "Tag": "email",
    "Value": ""
  }
]
```

- `password`: Password with min 8 character, max 32 character.

Bad password response:

```json
{
  "message": "failed to create user",
  "data": null,
  "error": "password doesn't meet the requirements"
}
```

- `passwordVerify`: Same with the password field to verify.

```json
{
  "message": "failed to create user",
  "data": null,
  "error": "password and verify password is not same"
}
```

- Multiple fail response:

```json
[
  {
    "Field": "Username",
    "Tag": "max",
    "Value": "16"
  },
  {
    "Field": "Email",
    "Tag": "email",
    "Value": ""
  }
]
```

## Login

Path: `/login`

Method: `POST`

### Success Example Requests

- Request

```json
{
  "email": "burakberkkeskin@gmail.com",
  "password": "Test1234"
}
```

Response:

```json
{
  "message": "user login success",
  "data": {
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImJ1cmFrQGdtYWlsLmNvbSIsImV4cCI6MTcwNTUwNTE2NCwiaWQiOiJiYjQxOTUyMy0xOGZmLTRjYmEtOTExMC1kNzJiZmI3NDQzNWYiLCJyb2xlIjoiYWRtaW4iLCJ1c2VybmFtZSI6InNhZmRlcnVuIn0.j51W9krrpNseOL8XMa1NRTDpwn-0woMCKWXiJw1L1As"
  },
  "error": ""
}
```

### Request Field Details

- `email`: Email value with valid email

Bad email response:

```json
[
  {
    "Field": "Email",
    "Tag": "email",
    "Value": ""
  }
]
```

- `password` Password string.

Bad password response:

```json
[
  {
    "Field": "Password",
    "Tag": "required",
    "Value": ""
  }
]
```

- Invalid credentials:

```json
{
  "message": "failed to authenticate user",
  "data": null,
  "error": "credentials are not valid"
}
```
