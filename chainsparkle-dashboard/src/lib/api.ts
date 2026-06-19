const API_BASE_URL = import.meta.env.VITE_API_URL || 'http://localhost:8080/api';

// Types for API responses
export interface Block {
  index: number;
  hash: string;
  previousHash: string;
  timestamp: number;
  transactions: Transaction[];
  nonce: number;
}

export interface Transaction {
  from: string;
  to: string;
  amount: number;
  publicKey?: string;
  signed?: boolean;
  verified?: boolean;
}

export interface Wallet {
  address: string;
  publicKey: string;
  privateKey: string;
  balance?: number;
}

export interface ChainStats {
  blocks: number;
  pendingTransactions: number;
  totalTransactions: number;
  latestBlockHash: string;
}

export interface NodeInfo {
  name: string;
  version: string;
  status: string;
  networkPort: number;
  apiPort: number;
}

export interface Peer {
  address: string;
}

// API Client
class ApiClient {
  private async fetch<T>(endpoint: string, options?: RequestInit): Promise<T> {
    const response = await fetch(`${API_BASE_URL}${endpoint}`, {
      ...options,
      headers: {
        'Content-Type': 'application/json',
        ...options?.headers,
      },
    });

    if (!response.ok) {
      const error = await response.json().catch(() => ({ error: 'Network error' }));
      throw new Error(error.error || 'API request failed');
    }

    return response.json();
  }

  // Wallet APIs
  async createWallet(): Promise<Wallet> {
    return this.fetch<Wallet>('/wallet/create', { method: 'POST' });
  }

  async getWalletBalance(address: string): Promise<{ address: string; balance: number }> {
    return this.fetch(`/wallet/balance/${address}`);
  }

  // Transaction APIs
  async createTransaction(from: string, to: string, amount: number, privateKey: string): Promise<any> {
    return this.fetch('/transaction/create', {
      method: 'POST',
      body: JSON.stringify({ from, to, amount, privateKey }),
    });
  }

  async getPendingTransactions(): Promise<{ count: number; transactions: Transaction[] }> {
    return this.fetch('/transaction/pending');
  }

  async getTransaction(hash: string): Promise<any> {
    return this.fetch(`/transaction/${hash}`);
  }

  // Blockchain APIs
  async getChain(): Promise<{ length: number; chain: Block[] }> {
    return this.fetch('/chain');
  }

  async getBlockByIndex(index: number): Promise<Block> {
    return this.fetch(`/chain/${index}`);
  }

  async getLatestBlock(): Promise<Block> {
    return this.fetch('/block/latest');
  }

  async validateChain(): Promise<{ status: string; message: string; blocks: number }> {
    return this.fetch('/validate');
  }

  // Mining APIs
  async mine(minerAddress?: string): Promise<any> {
    return this.fetch('/mine', {
      method: 'POST',
      body: JSON.stringify({ minerAddress }),
    });
  }

  // Node Stats APIs
  async getNodeInfo(): Promise<NodeInfo> {
    return this.fetch('/info');
  }

  async getChainStats(): Promise<ChainStats> {
    return this.fetch('/stats');
  }

  // Networking APIs
  async addPeer(address: string): Promise<{ message: string; peer: string }> {
    return this.fetch('/peer/add', {
      method: 'POST',
      body: JSON.stringify({ address }),
    });
  }

  async listPeers(): Promise<{ count: number; peers: string[] }> {
    return this.fetch('/peer/list');
  }

  async syncChain(): Promise<{ message: string }> {
    return this.fetch('/sync');
  }

  // Block deletion API (will be added to backend)
  async deleteBlock(index: number): Promise<{ message: string }> {
    return this.fetch(`/block/delete/${index}`, { method: 'DELETE' });
  }
}

export const api = new ApiClient();
