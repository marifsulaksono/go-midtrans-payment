# About
Go-midtrans-payment is implementation of learning about payment gateway on midtrans.
> [!NOTE]
> This project is using sandbox link

## ENV Example :
```sh
MIDTRANS_SERVER_KEY = "SB-Mid-server-xXXxxXxxxXxx"
MIDTRANS_SANDBOX_LINK = "https://api.sandbox.midtrans.com/v2/charge"
MIDTRANS_SNAP_SANDBOX_LINK = "https://app.sandbox.midtrans.com/snap/v1/transactions"

MONGODB_USERNAME = "username"
MONGODB_PASSWORD = "password"
MONGODB_CLUSTER = "xXXxxXx.xXXxxXx"
```

## API Endpoint
- Payment Link Endpoint : ```http://localhost:8080/payments/link```
- Snap Endpoint : ```http://localhost:8080/payments/snap```
- Core Endpoint : ```http://localhost:8080/payments/core```

## Request Body
### Payment Link
```sh
{
    "total": 300000,
    "customer_details":
    {
        "first_name": "Muhammad Arif",
        "last_name": "Sulaksono",
        "email": "marfs@gmail.com",
        "phone": "08111222333"
    },
    "item_details": [
        {
            "id": "tix-101",
            "name": "Tiket Konsep Dewa",
            "price": 100000,
            "quantity": 2,
            "merchant_name": "popuni"
        },
        {
            "id": "tix-102",
            "name": "Tiket Konsep Nidji",
            "price": 100000,
            "quantity": 1,
            "merchant_name": "popuni"
        }
    ]
}
```

### SNAP
```sh
{
    "total": 50000,
    "customer_details":
    {
        "first_name": "Muhammad Arif",
        "last_name": "Sulaksono",
        "email": "marfs@gmail.com",
        "phone": "08111222333"
    },
    "item_details": [
        {
            "id": "tix-101",
            "name": "Tiket Konsep Dangdut",
            "price":50000,
            "quantity": 1,
            "merchant_name": "popuni"
        }
    ]
}
```

### Bank BCA/BRI/BNI/CIMB (Core)
```sh
{
    "total": 50000,
    "payment_type": "bank_transfer",
    "payment_bank": "bca/bri/bni/cimb"
}
```

### Permata (Core)
```sh
{
    "total": 50000,
    "payment_type": "permata"
}
```

### Mandiri (Core)
```sh
{
    "total": 50000,
    "payment_type": "echannel",
    "echannel":
    {
        "bill_info1" : "Payment:",
        "bill_info2" : "Online purchase"
    }
}
```

### Akulaku/Kredivo (Core)
```sh
{
    "total": 50000,
    "payment_type": "akulaku/kredivo"
}
```

### Indomaret/Alfamart (Core)
```sh
{
    "total": 50000,
    "payment_type": "cstore",
    "store":
    {
        "store": "indomaret"/"alfamart",
        "message": "Thank you for order"
    }
}
```

### Gopay/QRIS (Core)
```sh
{
    "total": 50000,
    "payment_type": "qris"/"gopay"
}
```