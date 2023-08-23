# Stores API

<details>
  <summary>Get Stores</summary>

# Get Stores

## Decription

| Method | URL | Description |
|----------|----------|----------|
| **GET** | `/api/v1/stores` | Endpoint to get list of stores data |


## Request Parameter

### Query Param

| Parameter | Type | Required | Description |
|----------|----------|----------|----------|
| page | number | no | Page number. Default 1 |
| size | number | no | Amount of data displayed per page. Default 20 |
| keyword | string | no | Keyoword for store name |

## Example Request

```
curl --location --request GET '{base_url}/api/v1/stores?page=1&size=2&keyword=ayam'
```

## Response Parameter

| Parameter | Type | Description |
|----------|----------|----------|
| meta | meta object | Meta object |
| data | data object | Data object |

**Meta Object Contains :**
| Parameter | Type | Description |
|----------|----------|----------|
| code | number | HTTP status code |
| message | string | message |
| pagination | pagination object | Pagination object |

**Pagination Object Contains :**
| Parameter | Type | Description |
|----------|----------|----------|
| page | number | Page number |
| size | number | Maximum amount of data displayed. Related to limit |
| total_data | number | Total data |
| total_page | number | Total page |

**Data Object Contains :**
| Parameter | Type | Description |
|----------|----------|----------|
| uuid | string | Store UUID |
| store_name | string | Store Name |
| store_photo_url | string | Store Photo |

## Example Response

### 200

```
{
    "meta": {
        "code": 200,
        "message": "success",
        "pagination": {
            "page": 1,
            "size": 2,
            "total_data": 2,
            "total_page": 1
        }
    },
    "data": [
        {
            "uuid": "98f755a7-0714-11ed-8123-1aab392cf837",
            "store_name": "Ayam Geprek Laris Manis",
            "store_photo_url": "",
        },
        {
            "uuid": "03e189c2-d84d-45ee-8b9d-556e83a9c645",
            "store_name": "Ayam Geprek Sambel Ijo",
            "store_photo_url": "",
        }
    ]
}
```

### 400
```
{
    "meta": {
        "code": 400,
        "message": "at least 3 characters",
        "pagination": {}
    },
    "data": null
}
```
```
{
    "meta": {
        "code": 400,
        "message": "malformat request",
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
        "message": "an error occured when query db",
        "pagination": {}
    },
    "data": null
}
```

</details>

