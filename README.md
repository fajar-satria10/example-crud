# Go Pointers Cheatsheet (`*` and `&`)

A quick reference guide to understanding pointers in Go, with real examples from this CRUD project.

---

## 🔑 Core Concepts

| Symbol | Name             | What It Does                                | Analogy                                  |
| ------ | ---------------- | ------------------------------------------- | ---------------------------------------- |
| `&`    | **Address-of**   | Gets the memory address of a variable       | Getting the **address** of a house       |
| `*`    | **Pointer type** | Declares a type that holds a memory address | A **paper** that stores a house address  |
| `*`    | **Dereference**  | Reads the value at a memory address         | **Going to** the house using the address |

```go
stock := models.Stock{Quantity: 50}  // stock = the house (actual data)
ptr := &stock                         // ptr = the address of the house
fmt.Println(*ptr)                     // go to the address → read the house
```

---

## ⚙️ `&` — Address-of Operator

**Use when passing a variable to a function that needs to MODIFY it.**

### Real examples from this project:

```go
// GORM needs to WRITE database data INTO stock
config.DB.First(&stock, c.Param("id"))

// Gin needs to WRITE JSON request body INTO stock
c.ShouldBindJSON(&stock)

// GORM needs to WRITE (save) and UPDATE timestamps in stock
config.DB.Save(&stock)

// AutoMigrate needs to READ the struct to create/update the table
config.DB.AutoMigrate(&models.Item{})
```

### ❌ When NOT to use `&`:

```go
// Gin only READS the data to build a JSON response — no & needed
c.JSON(http.StatusOK, stock)
```

---

## ⚙️ `*` — Pointer Type

**Use in function signatures when the function needs to modify the caller's data.**

### Real examples from this project:

```go
// c is a POINTER to gin.Context
// Gin passes &ctx internally, your function receives it as *gin.Context
func UpdateStock(c *gin.Context) {
    // c can read AND modify the original Context
    c.JSON(200, data)   // writes a response to the original context
}
```

### How the caller and receiver connect:

```go
// CALLER side (inside Gin framework):
ctx := gin.Context{...}
UpdateStock(&ctx)          // passes ADDRESS using &

// RECEIVER side (your handler):
func UpdateStock(c *gin.Context)  // receives a POINTER using *
```

---

## 📋 Decision Flowchart

```
Calling a function?
├── Does it need to FILL or MODIFY my variable?
│   ├── YES → pass with &
│   │         c.ShouldBindJSON(&stock)
│   │         config.DB.First(&stock, id)
│   │
│   └── NO (just reading) → pass the value
│             c.JSON(200, stock)
│             fmt.Println(stock)
│
Writing a function?
├── Does it modify the caller's data?
│   ├── YES → use *Type in parameter
│   │         func Update(s *models.Stock)
│   │
│   └── NO → use value type
│             func Print(s models.Stock)
│
└── Is the struct large?
    ├── YES → use *Type (avoids copying)
    └── NO → either is fine
```

---

## 📊 Quick Reference Table

| Situation             | Use   | Example                    | Why                         |
| --------------------- | ----- | -------------------------- | --------------------------- |
| Fill struct from DB   | `&`   | `db.First(&stock, id)`     | Function writes into it     |
| Fill struct from JSON | `&`   | `c.ShouldBindJSON(&stock)` | Function writes into it     |
| Save to DB            | `&`   | `db.Save(&stock)`          | Function updates timestamps |
| Send JSON response    | value | `c.JSON(200, stock)`       | Function only reads         |
| Handler parameter     | `*`   | `func Get(c *gin.Context)` | Receive pointer from Gin    |
| Modify caller's data  | `*`   | `func Update(s *Stock)`    | Changes affect original     |
| Read-only parameter   | value | `func Print(s Stock)`      | Copy is fine                |

---

## 🧠 Memory Model

```
STACK MEMORY
┌─────────────────────────┐
│ stock := Stock{          │ ← actual data lives here
│   ID:       1,           │    (address: 0xC0000B4000)
│   Quantity: 50,          │
│   ItemId:   3,           │
│ }                        │
├─────────────────────────┤
│ ptr := &stock            │ ← holds address 0xC0000B4000
│ (type: *Stock)           │    (just 8 bytes, not a copy)
└─────────────────────────┘

db.First(&stock, id)    → goes to 0xC0000B4000, writes data
c.ShouldBindJSON(&stock)→ goes to 0xC0000B4000, overwrites fields
c.JSON(200, stock)      → copies data from 0xC0000B4000, sends as JSON
```

---

## ⚠️ Common Mistakes

### 1. Forgetting `&` when filling a struct

```go
// ❌ WRONG — stock stays empty, a copy gets filled and thrown away
c.ShouldBindJSON(stock)

// ✅ CORRECT — stock gets filled with JSON data
c.ShouldBindJSON(&stock)
```

### 2. Nil pointer dereference

```go
// ❌ CRASH — ptr is nil, can't dereference
var ptr *models.Stock     // ptr = nil
fmt.Println(ptr.Quantity) // panic: nil pointer dereference

// ✅ SAFE — initialize first
stock := models.Stock{Quantity: 50}
ptr := &stock
fmt.Println(ptr.Quantity) // 50
```

### 3. Pointer vs Value in loops

```go
// ❌ BUG — all pointers point to the same loop variable
var ptrs []*Stock
for _, s := range stocks {
    ptrs = append(ptrs, &s) // all point to same address!
}

// ✅ FIX — create a local copy
for _, s := range stocks {
    s := s                   // shadow with local copy
    ptrs = append(ptrs, &s)  // each points to different address
}
```

---

## 🏗️ Project Structure

```
example-crud/
├── main.go              # Entry point, AutoMigrate, route setup
├── config/
│   └── database.go      # DB connection (GORM + MariaDB)
├── models/
│   ├── item.go          # Item model
│   └── stock.go         # Stock model
├── handlers/
│   ├── item_handler.go  # CRUD handlers for items
│   └── stock_handler.go # CRUD handlers for stocks
└── routes/
    └── routes.go        # Route registration
```

---

## 🚀 Running the App

```bash
# Start the server
go run main.go

# Or with live-reload (requires Air)
air

# Test endpoints
curl http://localhost:8080/items
curl -X POST http://localhost:8080/items -H "Content-Type: application/json" \
  -d '{"name":"Book","description":"A good book","price":9.99}'
```
