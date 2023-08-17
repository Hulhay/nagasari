# Auth API

<details>
  <summary>Register Docs</summary>

# Register

## Decription

| Method | URL | Description |
|----------|----------|----------|
| **POST** | `/api/v1/auth/register` | Endpoint to register new user |


## Request Parameter

### Body 

| Parameter | Type | Required | Description |
|----------|----------|----------|----------|
| name | string | yes | User name |
| email | string | yes | User e-mail |
| password | string | yes | User password |
| phone_number | string | yes | User phone number |

## Example Request

```
curl --location --request POST '{base_url}/api/v1/auth/register' \
--header 'Content-Type: application/json' \
--data-raw '{
    "name": "Admin Pujas 1",
    "email": "adminpujas1@pujaskita.com",
    "password": "pujas1234",
    "phone_number": "0895401459426"
}'
```

## Response Parameter

| Parameter | Type | Description |
|----------|----------|----------|
| meta | meta object | Meta object |
| data | data object | null |

**Meta Object Contains :**
| Parameter | Type | Description |
|----------|----------|----------|
| code | number | HTTP status code |
| message | string | message |
| pagination | pagination object | null |

## Example Response

### 200

```
{
    "meta": {
        "code": 200,
        "message": "success!",
        "pagination": {}
    },
    "data": null
}
```

### 422
```
{
    "meta": {
        "code": 422,
        "message": "invalid email format",
        "pagination": {}
    },
    "data": null
}
```
```
{
    "meta": {
        "code": 422,
        "message": "invalid phone number format",
        "pagination": {}
    },
    "data": null
}
```
```
{
    "meta": {
        "code": 422,
        "message": "email cannot be empty",
        "pagination": {}
    },
    "data": null
}
```
```
{
    "meta": {
        "code": 422,
        "message": "name cannot be empty",
        "pagination": {}
    },
    "data": null
}
```
```
{
    "meta": {
        "code": 422,
        "message": "phone number cannot be empty",
        "pagination": {}
    },
    "data": null
}
```
```
{
    "meta": {
        "code": 422,
        "message": "password cannot be empty",
        "pagination": {}
    },
    "data": null
}
```
```
{
    "meta": {
        "code": 422,
        "message": "password must have a minimum of 8 characters",
        "pagination": {}
    },
    "data": null
}
```
```
{
    "meta": {
        "code": 422,
        "message": "password must contain and contain only letters and numbers",
        "pagination": {}
    },
    "data": null
}
```
```
{
    "meta": {
        "code": 422,
        "message": "email already exists",
        "pagination": {}
    },
    "data": null
}
```
```
{
    "meta": {
        "code": 422,
        "message": "phone number already exists",
        "pagination": {}
    },
    "data": null
}
```

### 500
```
{
    "meta": {
        "code": 500,
        "message": "no rows affected",
        "pagination": {}
    },
    "data": null
}
```
```
{
    "meta": {
        "code": 500,
        "message": "something wrong with encryption process",
        "pagination": {}
    },
    "data": null
}
```
```
{
    "meta": {
        "code": 500,
        "message": "failed to register",
        "pagination": {}
    },
    "data": null
}
```

</details>