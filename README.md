# Chaingo — Blockchain From Scratch

A full-stack blockchain implementation built with **Go** (backend) and **React + TypeScript** (frontend). Chaingo demonstrates core blockchain concepts — Proof of Work mining, ECDSA cryptography, P2P networking, and persistent storage — wrapped in a modern glassmorphism dashboard.

<p align="center">
  <!-- Backend -->
  <img src="https://img.shields.io/badge/Go-1.24+-00ADD8?style=for-the-badge&logo=go&logoColor=white" alt="Go" />
  <img src="https://img.shields.io/badge/Fiber-v2-00ACD7?style=for-the-badge&logo=go&logoColor=white" alt="Fiber" />
  <img src="https://img.shields.io/badge/BoltDB-embedded-4A90D9?style=for-the-badge&logo=databricks&logoColor=white" alt="BoltDB" />
  <img src="https://img.shields.io/badge/ECDSA-P--256-6C3483?style=for-the-badge&logo=gnuprivacyguard&logoColor=white" alt="ECDSA" />
  <!-- Frontend -->
  <img src="https://img.shields.io/badge/React-18.3-61DAFB?style=for-the-badge&logo=react&logoColor=black" alt="React" />
  <img src="https://img.shields.io/badge/TypeScript-5.x-3178C6?style=for-the-badge&logo=typescript&logoColor=white" alt="TypeScript" />
  <img src="https://img.shields.io/badge/Vite-5.4-646CFF?style=for-the-badge&logo=vite&logoColor=white" alt="Vite" />
  <img src="https://img.shields.io/badge/Tailwind_CSS-3.4-38B2AC?style=for-the-badge&logo=tailwind-css&logoColor=white" alt="Tailwind CSS" />
  <img src="https://img.shields.io/badge/shadcn%2Fui-latest-000000?style=for-the-badge&logo=shadcnui&logoColor=white" alt="shadcn/ui" />
  <img src="https://img.shields.io/badge/Framer_Motion-12-FF0055?style=for-the-badge&logo=framer&logoColor=white" alt="Framer Motion" />
</p>

---

## Table of Contents

- [Overview](#overview)
- [Features](#features)
- [Tech Stack](#tech-stack)
- [Project Structure](#project-structure)
- [Architecture](#architecture)
  - [Backend (ChainGo)](#backend-chaingo)
  - [Frontend (chainsparkle-dashboard)](#frontend-chainsparkle-dashboard)
- [API Reference](#api-reference)
- [Getting Started](#getting-started)
- [Configuration](#configuration)
- [How It Works](#how-it-works)
- [Screenshots](#screenshots)

---

## Overview

Chaingo is a complete blockchain system with two independently runnable parts:

| Part | Directory | Description |
|------|-----------|-------------|
| Backend | `ChainGo/` | Go node: blockchain engine, REST API, P2P networking |
| Frontend | `chainsparkle-dashboard/` | React dashboard: wallet manager, explorer, miner, network viewer |

The backend runs as a single binary that simultaneously serves a REST API and a TCP P2P server. Multiple nodes can discover each other, sync their chains, and reach consensus via the longest-chain rule.

---

## Features

### Blockchain Core
- Full blockchain with genesis block, block hashing, and chain validation
- **Proof of Work** mining with 16-bit difficulty target (~10 second block time)
- **ECDSA P-256** cryptography for wallet generation and transaction signing
- **Coinbase transactions** — 50 ChainGo block reward per mined block
- UTXO-based balance calculation

### Networking
- **P2P TCP layer** — nodes connect, broadcast blocks/transactions, and sync chains
- Longest-chain consensus rule during sync
- REST API with 16 endpoints (wallets, transactions, mining, chain, peers)

### Storage
- **BoltDB** embedded key/value store — no external database required
- Two buckets: `chaingo_blocks` (hash → block) and `chaingo_wallets` (address → wallet)
- Chain persists across restarts

### Dashboard
- **Wallet Manager** — create wallets, view balances, export keys as JSON
- **Transaction Center** — create and sign transactions, view mempool and history
- **Block Explorer** — browse every block, search by height or hash, inspect transactions
- **Mining Dashboard** — start/stop mining, track rewards, hash rate, and block history
- **Network View** — visual peer graph, add peers, trigger chain sync
- Dark mode with glassmorphism UI, responsive to mobile

---

## Tech Stack

**Backend**

| Technology | Purpose |
|-----------|---------|
| Go 1.24+ | Core language |
| [Fiber v2](https://github.com/gofiber/fiber) | HTTP REST API framework |
| BoltDB (bbolt) | Embedded persistent storage |
| ECDSA (P-256) | Wallet cryptography |
| TCP (stdlib) | P2P networking |
| `encoding/gob` | Binary message encoding for P2P |

**Frontend**

| Technology | Purpose |
|-----------|---------|
| React 18.3 + TypeScript | UI framework |
| Vite 5.4 | Build tool / dev server |
| Tailwind CSS 3.4 | Styling |
| shadcn/ui + Radix UI | Component library |
| Framer Motion | Animations |
| React Router v6 | Client-side routing |
| TanStack Query | HTTP caching and background refetch |
| Recharts | Charts |
| React Hook Form + Zod | Form validation |
| Lucide React | Icons |
| Sonner | Toast notifications |

---

## Project Structure

```
Chaingo/
├── ChainGo/                        # Go backend
│   ├── main.go                     # Entry point — CLI flags, starts API + P2P
│   ├── go.mod / go.sum             # Go module and dependencies
│   ├── chaingo.db                  # BoltDB file (auto-created at runtime)
│   │
│   ├── blockchain/                 # Core blockchain logic
│   │   ├── block.go                # Block struct, hashing, serialization
│   │   ├── blockchain.go           # Chain management — AddBlock, Validate, Load
│   │   ├── pow.go                  # Proof of Work algorithm
│   │   ├── transaction.go          # Transaction struct — sign, verify
│   │   ├── wallet.go               # ECDSA wallet generation, address derivation
│   │   └── utils.go                # IntToHex helper
│   │
│   ├── internal/                   # HTTP layer (Fiber)
│   │   ├── server.go               # Server setup, CORS, logger middleware
│   │   └── handlers.go             # All 16 endpoint handlers
│   │
│   ├── network/                    # P2P layer
│   │   ├── node.go                 # TCP listener, goroutine per connection
│   │   ├── peer.go                 # PeerManager — mutex-protected peer list
│   │   └── protocol.go             # gob-encoded message encode/decode
│   │
│   └── pkg/                        # Storage layer
│       ├── db.go                   # DB interface abstraction
│       └── bolt.go                 # BoltDB implementation
│
└── chainsparkle-dashboard/         # React frontend
    ├── index.html
    ├── package.json
    ├── vite.config.ts
    ├── tailwind.config.ts
    ├── tsconfig.json
    │
    └── src/
        ├── main.tsx                # React entry point
        ├── App.tsx                 # Router + route definitions
        ├── index.css               # Global styles + custom animations
        │
        ├── lib/
        │   ├── api.ts              # ApiClient class — all backend calls
        │   └── utils.ts            # cn() utility function
        │
        ├── pages/                  # One component per route
        │   ├── Dashboard.tsx       # Home — stats, chain view, activity feed
        │   ├── Wallet.tsx          # Wallet creation and balance
        │   ├── Transactions.tsx    # Send, history, mempool
        │   ├── Blocks.tsx          # Block explorer
        │   ├── Mining.tsx          # Mining controls and stats
        │   ├── Network.tsx         # Peer management and visualization
        │   └── NotFound.tsx        # 404
        │
        ├── components/
        │   ├── layout/
        │   │   ├── Layout.tsx      # App shell
        │   │   └── Header.tsx      # Sticky nav (desktop + mobile)
        │   └── ui/                 # Reusable UI components
        │       ├── StatCard.tsx
        │       ├── GlassCard.tsx
        │       ├── BlockChainVisual.tsx
        │       ├── ActivityFeed.tsx
        │       ├── StatusBadge.tsx
        │       └── [50+ shadcn components]
        │
        └── hooks/
            ├── use-toast.ts
            └── use-mobile.tsx
```

---

## Architecture

### Backend (ChainGo)

```
                    ┌─────────────────────────────────┐
                    │           main.go                │
                    │  - Parses CLI flags               │
                    │  - Opens BoltDB                  │
                    │  - Starts P2P node (goroutine)   │
                    │  - Starts Fiber REST API         │
                    └────────────┬────────────────────┘
                                 │
               ┌─────────────────┴──────────────────┐
               │                                     │
    ┌──────────▼──────────┐             ┌────────────▼────────────┐
    │     REST API :8080   │             │    P2P Network :9000    │
    │  (Fiber + handlers)  │             │  (TCP + gob messages)   │
    └──────────┬──────────┘             └────────────┬────────────┘
               │                                     │
    ┌──────────▼─────────────────────────────────────▼────────────┐
    │                    Blockchain Engine                         │
    │  ┌──────────┐  ┌─────────────┐  ┌──────────┐  ┌─────────┐  │
    │  │  Block   │  │    Chain    │  │   PoW    │  │ Wallet  │  │
    │  │  struct  │  │  management │  │ mining   │  │  ECDSA  │  │
    │  └──────────┘  └─────────────┘  └──────────┘  └─────────┘  │
    └──────────────────────────────┬───────────────────────────────┘
                                   │
                       ┌───────────▼───────────┐
                       │       BoltDB           │
                       │  chaingo_blocks bucket │
                       │  chaingo_wallets bucket│
                       └───────────────────────┘
```

#### Data Structures

**Block**
```go
type Block struct {
    Timestamp    int64           // Unix timestamp
    Transactions []*Transaction  // List of transactions
    PrevHash     []byte          // Previous block's hash (chain link)
    Hash         []byte          // SHA-256 hash of this block
    Nonce        int             // PoW solution nonce
}
```

**Transaction**
```go
type Transaction struct {
    From      string  // Sender address (SHA-256 of public key)
    To        string  // Recipient address
    Amount    int     // Transfer amount in ChainGo
    R, S      string  // ECDSA signature components (hex)
    PublicKey []byte  // 64-byte sender public key
}
```

**Wallet**
```go
type Wallet struct {
    PrivateKey *ecdsa.PrivateKey  // P-256 elliptic curve private key
    PublicKey  []byte             // 64 bytes: 32 (X) + 32 (Y) coordinates
    // Address = hex( SHA-256(PublicKey) )
}
```

#### Proof of Work

The PoW algorithm requires the block hash to start with a number of zero bits equal to the difficulty target (16 bits by default):

```
target = 1 << (256 - 16)
loop:
    nonce++
    hash = SHA-256(timestamp + transactions + prevHash + nonce)
    if hash < target → block is valid
```

#### P2P Protocol

Nodes communicate over TCP using gob-encoded messages:

| Message Type | Direction | Description |
|-------------|-----------|-------------|
| `BLOCK` | broadcast | Newly mined block |
| `TRANSACTION` | broadcast | New unconfirmed transaction |
| `CHAIN_REQUEST` | outbound | Request peer's full chain |
| `CHAIN_RESPONSE` | inbound | Peer responds with its chain |

On receiving a `CHAIN_RESPONSE`, the node adopts the peer's chain only if it is longer and valid.

---

### Frontend (chainsparkle-dashboard)

```
App.tsx (React Router)
│
├── /                → Dashboard     — auto-refreshes every 5s
├── /wallet          → Wallet        — auto-refreshes balance every 10s
├── /transactions    → Transactions  — mempool + history
├── /blocks          → Block Explorer
├── /mining          → Mining
└── /network         → P2P Network
```

**Data flow:**  
Each page uses `ApiClient` (a typed `fetch` wrapper in `src/lib/api.ts`) to call the Go backend at `http://localhost:8080`. TanStack Query manages caching and background refetching. State that must survive page reloads (wallets, mining stats) is persisted to `localStorage`.

---

## API Reference

Base URL: `http://localhost:8080`

### Wallet

| Method | Endpoint | Description |
|--------|----------|-------------|
| `POST` | `/api/wallet/create` | Generate a new ECDSA wallet |
| `GET` | `/api/wallet/balance/:address` | Get UTXO balance for address |

### Transactions

| Method | Endpoint | Description |
|--------|----------|-------------|
| `POST` | `/api/transaction/create` | Create and sign a new transaction |
| `GET` | `/api/transaction/pending` | List mempool (unconfirmed) transactions |
| `GET` | `/api/transaction/:hash` | Look up a transaction by hash |

### Mining

| Method | Endpoint | Body | Description |
|--------|----------|------|-------------|
| `GET` | `/api/mine` | — | Mine block, reward sent to genesis address |
| `POST` | `/api/mine` | `{minerAddress}` | Mine block, reward sent to specified address |

### Blockchain

| Method | Endpoint | Description |
|--------|----------|-------------|
| `GET` | `/api/chain` | Get full chain |
| `GET` | `/api/chain/:index` | Get block by height |
| `GET` | `/api/block/latest` | Get latest block |
| `GET` | `/api/validate` | Validate entire chain integrity |
| `DELETE` | `/api/block/delete/:index` | Delete block at height and all after it |

### Node & Network

| Method | Endpoint | Body | Description |
|--------|----------|------|-------------|
| `GET` | `/api/info` | — | Node name, version, ports |
| `GET` | `/api/stats` | — | Block count, pending TXs, total TXs |
| `POST` | `/api/peer/add` | `{address}` | Connect to a peer node |
| `GET` | `/api/peer/list` | — | List all connected peers |
| `GET` | `/api/sync` | — | Sync chain from all peers |

---

## Getting Started

### Prerequisites

- **Go** 1.24 or later
- **Node.js** 18 or later + npm

### 1. Clone the Repository

```bash
git clone https://github.com/Vishal-2029/Chaingo.git
cd Chaingo
```

### 2. Start the Backend

```bash
cd ChainGo
go mod tidy
go run main.go
```

The node starts with default settings:
- REST API → `http://localhost:8080`
- P2P server → `tcp://localhost:9000`
- Database → `chaingo.db` (created automatically)

### 3. Start the Frontend

Open a new terminal:

```bash
cd chainsparkle-dashboard
npm install
npm run dev
```

Frontend dev server → `http://localhost:3000`

### 4. Open the Dashboard

Navigate to [http://localhost:3000](http://localhost:3000) in your browser.

---

## Configuration

### Backend CLI Flags

```bash
go run main.go -api 8080 -p2p 9000 -db chaingo.db
```

| Flag | Default | Description |
|------|---------|-------------|
| `-api` | `8080` | HTTP REST API port |
| `-p2p` | `9000` | TCP P2P listening port |
| `-db` | `chaingo.db` | BoltDB file path |

### Running Multiple Nodes (Local P2P Test)

```bash
# Terminal 1 — Node A
cd ChainGo
go run main.go -api 8080 -p2p 9000 -db node_a.db

# Terminal 2 — Node B
cd ChainGo
go run main.go -api 8081 -p2p 9001 -db node_b.db

# Connect Node B to Node A (via API or dashboard)
curl -X POST http://localhost:8081/api/peer/add \
  -H "Content-Type: application/json" \
  -d '{"address":":9000"}'

# Trigger chain sync
curl http://localhost:8081/api/sync
```

### Frontend API URL

The frontend points to `http://localhost:8080` by default. To change this, edit `chainsparkle-dashboard/src/lib/api.ts`:

```typescript
const BASE_URL = 'http://localhost:8080';
```

---

## How It Works

### Transaction Lifecycle

```
1. User creates wallet       → ECDSA P-256 keypair generated
                                Address = hex(SHA-256(publicKey))

2. User receives coins       → Mine a block first (50 ChainGo coinbase reward)

3. User creates transaction  → Frontend signs with private key via ECDSA
                                Signed TX added to mempool

4. Miner mines next block    → PoW finds valid nonce (~10 seconds)
                                Block includes mempool transactions + new coinbase

5. Block saved to BoltDB     → Hash → Block in chaingo_blocks bucket

6. Block broadcast to peers  → P2P layer sends BLOCK message to all connections

7. Peers validate & append   → Each peer checks hash, PoW, signatures
                                Adopts block if chain is valid

8. Frontend polls every 5s   → Dashboard reflects confirmed transaction
```

### Balance Calculation (UTXO)

Balance is computed by scanning every block in the chain:
- Add `amount` for every transaction where `tx.To == address`
- Subtract `amount` for every transaction where `tx.From == address`

### Chain Validation

`GET /api/validate` checks the entire chain by verifying:
1. Each block's hash is correctly computed from its content
2. Each block's `PrevHash` matches the previous block's `Hash`
3. The PoW nonce satisfies the difficulty target

---

## Dashboard Pages

### Dashboard (/)
Real-time overview of the node. Stats cards (blocks, pending transactions, peers, total transactions) auto-refresh every 5 seconds. Includes an animated block chain visualization showing the last 5 blocks and a live activity feed.

### Wallet (/wallet)
Create unlimited wallets locally. Each wallet shows its address, public key, and private key (toggle visibility). Balance refreshes every 10 seconds. Wallets persist in `localStorage`. Export any wallet as a JSON file.

### Transactions (/transactions)
Three tabs:
- **Create** — fill sender/recipient/amount/private key, click Send
- **History** — all confirmed transactions, searchable by address or hash
- **Pending** — current mempool with transaction priority indicators

### Block Explorer (/blocks)
Browse the full chain newest-first. Search by block index or hash. Expand any block to see its transactions. Delete a block and all blocks after it (with confirmation dialog).

### Mining (/mining)
Select a wallet to receive mining rewards, then click **Start Mining**. The dashboard shows live stats: blocks mined, total rewards, elapsed time, hash rate, and a history of the last 10 mined blocks. Stats persist across page refreshes via `localStorage`.

### Network (/network)
Visual graph of the P2P network with the local node at center and connected peers orbiting it. Add a peer by entering its address (e.g. `:9001`). Each peer entry shows sync status, block height, latency, and a sync progress bar. Click **Sync Now** to pull the latest chain from all peers.
