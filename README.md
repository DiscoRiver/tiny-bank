# Tiny Bank Ledger - Design Definition

A lightweight, in-memory ledger API for tracking deposits, withdrawals, and balances.

For simplicity, the design and documentation are both represented in this README. In an ideal world I would create design docs, most likely an ERD/PRD with a supplemental ADR doc if needed.

For a micro exercise like this, I assume you're looking for fundamentals, and ability to create something to spec, so that's where I've focused. I've built this as I would do a prototype or proof-of-concept, with maximum expediency as I have very limited time (overall I spent about 90 minutes on all of this). I am aware of double entry accounting databases such as the one [described by Square](https://developer.squareup.com/blog/books-an-immutable-double-entry-accounting-database-service/), and consensus algorithms such as Paxos and Raft, but have never implemented them. 

Additionally, there are features that I have not implemented but are needed, and it's important I highlight them, but they're too fundamental for me to include in the future enhancements at the bottom of this page;

- Transaction timestamps.
- Client/Merchant definitions and resolution.
- Model validation. For this I would use https://github.com/go-playground/validator
- Layering of logic. I kept things as basic as possible while also showing  a consideration for separation.
- Creation of test objects (comes with layering).
- Error handling with https://github.com/hedhyw/semerr to be transport-independant for various delivery methods.
- I would default to OpenAPI definitions for models. It allows for tools such as the typescript generator to make it easy for front end integrations.

I'm happy to discuss all this, of course, if you have any questions.

# Overview

The Tiny Bank Ledger is a minimalistic financial transaction system that provides a RESTful API to:

* Record deposits and withdrawals
* Retrieve current balance
* View transaction history

This system is designed to be simple, fast, and stateless, running entirely in memory.

# Functional Requirements

* Deposit Money: Add funds to the ledger.
* Withdraw Money: Deduct funds from the ledger, ensuring sufficient balance.
* View Balance: Fetch the current account balance.
* Transaction History: Retrieve a list of past deposits and withdrawals.

# Non-Functional Requirements

* In-Memory Storage: No database, uses Go’s data structures (maps/slices).
* Concurrency-Safe: Uses mutex locks to prevent race conditions.
* RESTful API: Uses HTTP endpoints for interaction.
* Stateless Execution: Data is lost when the server stops.
* Minimal Dependencies: Only uses the standard Go library.

# System Design

## Architecture

```User ➝ HTTP API ➝ Ledger (In-memory)```

* The API is exposed via Go’s net/http package.
* Ledger is an internal component that handles transactions safely.

# Data Model

## Transaction
| Field   | Type   | Description                                                   |
|---------|--------|---------------------------------------------------------------|
| Amount  | int    | Amount transacted (positive for deposit, negative for withdrawal) |
| Type    | string | `"deposit"` or `"withdrawal"`                                  |

## Ledger
| Field        | Type           | Description                  |
|-------------|--------------|------------------------------|
| Balance     | int          | Current account balance     |
| Transactions | []Transaction | List of recorded transactions |

# API Endpoints

| Method | Endpoint         | Description                          | Query Params       |
|--------|-----------------|--------------------------------------|--------------------|
| POST   | /deposit        | Deposits money into the ledger      | amount=int        |
| POST   | /withdraw       | Withdraws money (if balance allows) | amount=int        |
| GET    | /balance        | Retrieves current balance           | -                 |
| GET    | /transactions   | Fetches transaction history         | -                 |

# Concurrency & Atomicity

✔ Mutex Locking (sync.Mutex): Ensures atomicity of transactions.  
✔ Validation Before State Change: Prevents partial transactions.  
✔ Thread-Safe Access to Data: No race conditions in multi-threaded execution.  

# Trade-offs & Assumptions

✅ Simple & Fast: No external database overhead.  
✅ No Authentication: Anyone can access the API.  
❌ No Persistence: Data is lost if the server restarts.  
❌ No Multi-Account Support: Only one account is managed.  

# How to Run
```shell
go run main.go
```

# Usage

1. Deposit Money
```shell
curl -X POST "http://localhost:8080/deposit?amount=100"
```

2. Withdraw Money
```shell
curl -X POST "http://localhost:8080/withdraw?amount=50"
```

3. Check Balance
```shell
curl -X GET "http://localhost:8080/balance"
```

4. Get Transaction History
```shell
curl -X GET "http://localhost:8080/transactions"
```

# Testing

Run the tests using:

```shell
go test ./ledger -v
```
# Future Enhancements

🚀 Add persistent storage (SQLite, PostgreSQL)  
🚀 Implement authentication & authorization.  
🚀 Add multi-account support.  
