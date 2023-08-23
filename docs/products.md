# Products API

<details>
  <summary>Get Store Products</summary>

# Get Store Products

## Decription

| Method | URL | Description |
|----------|----------|----------|
| **GET** | `/api/v1/stores/:storeUUID/products` | Endpoint to get list of products data from spesific store |


## Request Parameter

### Path Param

| Parameter | Type | Required | Description |
|----------|----------|----------|----------|
| storeUUID | string | yes | Store UUID |

## Example Request

```
curl --location --request GET '{base_url}/api/v1/stores/3f53205f-01c7-11ed-9140-6a8c8fb6de13/products'
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
| product | list of menu object | Menu object |

**Store Object Contains :**
| Parameter | Type | Description |
|----------|----------|----------|
| uuid | string | Store UUID |
| store_name | string | Store Name |
| store_photo_url | string | Store Photo |
| owner_uuid | string | Owner UUID |
| owner_name | string | Owner Name |
| owner_phone_number | string | Owner Phone Number |

**Product Object Contains :**
| Parameter | Type | Description |
|----------|----------|----------|
| uuid | string | Product UUID |
| name | string | Product Name |
| price | number | Price |
| is_sold_out | boolean | Is Sold Out |
| product_photo_url | string | Product Photo |

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
                "is_sold_out": false,
                "product_photo_url": ""
            },
            {
                "uuid": "2440a831-8ad4-4d25-8bc1-d19688c74f0e",
                "name": "STMJ",
                "price": 6000,
                "is_sold_out": false,
                "product_photo_url": ""
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
        "code": 500,
        "message": "an error occured when query db",
        "pagination": {}
    },
    "data": null
}
```

</details>

