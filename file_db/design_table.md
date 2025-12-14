```mermaid
erDiagram
 Categories {
    SERIAL category_id PK
    VARCHAR(50) name_category
    VARCHAR(250) description
 }

 Items {
    SERIAL item_id PK
    SERIAL category_id FK
    VARCHAR(100) name_item
    NUMERIC price
    TIMESTAMP purchase_date
    BOOLEAN replaced
 }

 Categories ||--o{ Items : "has"

 ```