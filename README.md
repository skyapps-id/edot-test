User
 └─> [User Service] ─────────────┐
                                 ├─> Authenticated Access
Product Viewer                   │
 └─> [Product Service] <────────┴─> Get product list (stock info via Warehouse)
                                 
Buyer
 └─> [Order Service] ──────────┐
      │                        └─> Reserve/Release stock via Warehouse Service
      └─> [Warehouse Service] ──> Manage stock, transfer, active/inactive

Admin
 └─> [Shop Service] ──────────┐
      └─> Manage warehouse map └─> [Warehouse Service]


POST /orders/checkout
POST /orders/:id/payment
POST /orders/:id/cancel
                                        
[User]
   |
   v
[Login via User Service]
   |
   v
[View Products - Product Service]
   |
   v
[Click Checkout - Order Service]
   |
   v
[Order Service: Ambil info produk & kuantitas]
   |
   v
[Order Service → Shop Service]
   |
   v
Shop Service:
- Cari Shop dari produk terkait
- Dapatkan daftar Warehouse aktif milik Shop
   |
   v
[Order Service → Warehouse Service]
   |
   v
Warehouse Service:
- Cek ketersediaan stok di gudang aktif
- Reserve (lock) stok
   |
   v
[Order Service]
   |
   v
[Create Order]
   |
   v
[Tunggu Pembayaran]
   |
   v
(Timeout?)
 ├── No  → [Order Berhasil]
 └── Yes → [Order Dibatalkan → Release Stok via Warehouse Service]

 