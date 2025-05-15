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
                                        