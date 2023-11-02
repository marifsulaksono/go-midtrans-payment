# About
Go-midtrans-payment is my implementation of learning about payment gateway on midtrans
> [!NOTE]
> This project is using sandbox link

## ENV Example :
```sh
MIDTRANS_SERVER_KEY = "xXXxxXxxxXxx"
MIDTRANS_SANDBOX_LINK = "xXXxxXxxxXxx"
MIDTRANS_SNAP_SANDBOX_LINK = "xXXxxXxxxXxx"
```

## API Endpoint
- Snap Endpoint : ```http://localhost:8080/payments/snap```
- Core Endpoint : ```http://localhost:8080/payments/core```

## Request Body
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
    }
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