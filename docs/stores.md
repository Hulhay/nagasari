# Stores API

<details>
  <summary>Get Stores Docs</summary>

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
        "code": 400,
        "message": "an error occured when query db",
        "pagination": {}
    },
    "data": null
}
```

</details>

---


<details>
  <summary>Get Store Menus Docs</summary>

# Get Store Menus

## Decription

| Method | URL | Description |
|----------|----------|----------|
| **GET** | `/api/v1/stores/:storeUUID/menus` | Endpoint to get list of menus data from spesific store |


## Request Parameter

### Path Param

| Parameter | Type | Required | Description |
|----------|----------|----------|----------|
| storeUUID | string | yes | Store UUID |

## Example Request

```
curl --location --request GET '{base_url}/api/v1/stores/3f53205f-01c7-11ed-9140-6a8c8fb6de13/menus'
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
| pagination | - | - |

**Data Object Contains :**
| Parameter | Type | Description |
|----------|----------|----------|
| store | store object | Store object |
| menu | list of menu object | Menu object |

**Store Object Contains :**
| Parameter | Type | Description |
|----------|----------|----------|
| uuid | string | Store UUID |
| store_name | string | Store Name |
| store_photo_url | string | Store Photo |
| owner_uuid | string | Owner UUID |
| owner_name | string | Owner Name |
| owner_phone_number | string | Owner Phone Number |

**Menu Object Contains :**
| Parameter | Type | Description |
|----------|----------|----------|
| uuid | string | Menu UUID |
| name | string | Menu Name |
| price | number | Price |
| is_sold_out | boolean | Is Sold Out |

## Example Response

### 200

```
{
    "meta": {
        "code": 200,
        "message": "success",
        "pagination": {}
    },
    "data": {
        "store": {
            "uuid": "3f53205f-01c7-11ed-9140-6a8c8fb6de13",
            "store_name": "Kedai Minum Sehat Jono",
            "store_photo_url": "",
            "owner_uuid": "48e27677-cdd9-48f4-bcb5-9b8d1f819c83",
            "owner_name": "Jono Mangunsong",
            "owner_phone_number": "085245393529"
        },
        "menu": [
            {
                "uuid": "a216e21e-6196-4508-b90c-e05cb03f2ae3",
                "name": "Wedang Jahe",
                "price": 5500,
                "is_sold_out": false
            },
            {
                "uuid": "2440a831-8ad4-4d25-8bc1-d19688c74f0e",
                "name": "STMJ",
                "price": 6000,
                "is_sold_out": false
            }
        ]
    }
}
```

### 400
```
{
    "meta": {
        "code": 400,
        "message": "invalid uuid",
        "pagination": {}
    },
    "data": null
}
```
```
{
    "meta": {
        "code": 400,
        "message": "store uuid cannot be empty",
        "pagination": {}
    },
    "data": null
}
```

### 404
```
{
    "meta": {
        "code": 400,
        "message": "store not found",
        "pagination": {}
    },
    "data": null
}
```

### 500
```
{
    "meta": {
        "code": 400,
        "message": "an error occured when query db",
        "pagination": {}
    },
    "data": null
}
```

</details>

