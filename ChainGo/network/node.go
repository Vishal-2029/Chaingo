package network

import (
	"fmt"
	"net"

	"github.com/Vishal-2029/blockchain"
)

type Node struct {
	Address    string
	Peers      *PeerManager
	Blockchain *blockchain.Blockchain
}

func NewNode(address string) *Node {
	return &Node{
		Address: address,
		Peers:   NewPeerManager(),
	}
}

func (n *Node) SetBlockchain(bc *blockchain.Blockchain) {
	n.Blockchain = bc
}

func (n *Node) Start() {
	listener, err := net.Listen("tcp", n.Address)
	if err != nil {
		panic(err)
	}
	fmt.Printf("🌐 Node listening on %s\n", n.Address)

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		go n.handleConnection(conn)
	}
}

func (n *Node) handleConnection(conn net.Conn) {
	defer conn.Close()

	// Get remote address
	remoteAddr := conn.RemoteAddr().String()
	// fmt.Printf("🔗 New P2P connection from: %s\n", remoteAddr)

	// Add this peer to our list if valid
	// n.Peers.AddPeer(remoteAddr)

	buf := make([]byte, 409600) // Increase buffer for large chains
	nBytes, err := conn.Read(buf)
	if err != nil {
		// fmt.Printf("Error reading from %s: %v\n", remoteAddr, err)
		return
	}

	// fmt.Printf("📨 Received %d bytes from %s\n", nBytes, remoteAddr)

	msg, err := DecodeMessage(buf[:nBytes])
	if err != nil {
		return
	}

	switch msg.Type {
	case "BLOCK":
		// fmt.Printf("Received new BLOCK from peer: %s\n", remoteAddr)
		// Handle new block logic here
	case "TRANSACTION":
		// fmt.Printf("Received new TRANSACTION from peer: %s\n", remoteAddr)
		// Handle new tx logic here
	case "CHAIN_REQUEST":
		targetAddr, ok := msg.Data.(string)
		if ok {
			fmt.Printf("Peer %s requested chain sync. Replying to %s\n", remoteAddr, targetAddr)
			go n.sendChainTo(targetAddr)
		} else {
			// Fallback if data is missing (legacy support?), use connection if we could (but we can't).
			// Just log error.
			fmt.Printf("Peer %s requested chain sync but provided invalid address\n", remoteAddr)
		}
	case "CHAIN_RESPONSE":
		n.handleChainResponse(msg.Data)
	default:
		fmt.Printf("Unknown message type from %s: %s\n", remoteAddr, msg.Type)
	}
}

func (n *Node) sendChainTo(addr string) {
	if n.Blockchain == nil {
		return
	}

	msg := Message{
		Type: "CHAIN_RESPONSE",
		Data: n.Blockchain.Block,
	}

	fmt.Printf("Sending chain to %s\n", addr)
	n.sendMessage(addr, msg)
}

func (n *Node) unused_sendChain(conn net.Conn) {
	// Deprecated
	if n.Blockchain == nil {
		return
	}

	msg := Message{
		Type: "CHAIN_RESPONSE",
		Data: n.Blockchain.Block,
	}

	data, _ := EncodeMessage(msg)
	conn.Write(data)
	fmt.Printf("Sent blockchain (height: %d) to peer\n", len(n.Blockchain.Block))
}

func (n *Node) handleChainResponse(data interface{}) {
	// Attempt to cast to slice of blocks
	newChain, ok := data.([]*blockchain.Block)
	if !ok {
		fmt.Printf("Received invalid chain data type: %T\n", data)
		return
	}

	fmt.Printf("📦 Received valid blockchain from peer. Height: %d\n", len(newChain))

	// In a real implementation:
	// if len(newChain) > len(n.Blockchain.Block) && n.Blockchain.ValidateChain(newChain) {
	//     n.Blockchain.ReplaceChain(newChain)
	// }
}

func (n *Node) Broadcast(msg Message) {
	peers := n.Peers.ListPeers()
	// fmt.Printf("Broadcasting '%s' to %d peers\n", msg.Type, len(peers))

	for _, addr := range peers {
		go n.sendMessage(addr, msg)
	}
}

func (n *Node) sendMessage(addr string, msg Message) {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		// fmt.Printf("Failed to connect to peer %s: %v\n", addr, err)
		n.Peers.mu.Lock()
		delete(n.Peers.Peers, addr)
		n.Peers.mu.Unlock()
		return
	}
	defer conn.Close()

	data, err := EncodeMessage(msg)
	if err != nil {
		return
	}

	conn.Write(data)
}
