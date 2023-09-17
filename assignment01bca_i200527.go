package assignment01bcaa

import (
	"bufio"
	"crypto/sha256"
	"fmt"
	"os"
	"strings"
)

type Block struct {
	Transaction   string
	Nonce         int
	PreviousHash  string
	Hash          string
}

func NewBlock(transaction string, nonce int, previousHash string) *Block {
	block := &Block{
		Transaction:  transaction,
		Nonce:        nonce,
		PreviousHash: previousHash,
	}
	block.Hash = block.CreateHash()
	return block
}

func (b *Block) CreateHash() string {
	data := fmt.Sprintf("%s%d%s", b.Transaction, b.Nonce, b.PreviousHash)
	hash := sha256.Sum256([]byte(data))
	return fmt.Sprintf("%x", hash)
}

type Blockchain struct {
	Blocks []*Block
}

func (bc *Blockchain) AddBlock(transaction string, nonce int) {
	var previousHash string
	if len(bc.Blocks) > 0 {
		previousHash = bc.Blocks[len(bc.Blocks)-1].Hash
	}
	block := NewBlock(transaction, nonce, previousHash)
	bc.Blocks = append(bc.Blocks, block)
}

func (bc *Blockchain) ListBlocks() {
	for i, block := range bc.Blocks {
		fmt.Printf("Block %d:\n", i)
		fmt.Printf("Transaction: %s\n", block.Transaction)
		fmt.Printf("Nonce: %d\n", block.Nonce)
		fmt.Printf("Previous Hash: %s\n", block.PreviousHash)
		fmt.Printf("Current Hash: %s\n", block.Hash)
		print("\n")
		fmt.Println("==========================================================================")
	}
}

func (bc *Blockchain) ChangeBlock(blockIndex int, newTransaction string, nonce int) {
	if blockIndex >= 0 && blockIndex < len(bc.Blocks) {
		bc.Blocks[blockIndex].Transaction = newTransaction
		bc.Blocks[blockIndex].Nonce = nonce
		bc.Blocks[blockIndex].Hash = bc.Blocks[blockIndex].CreateHash()
	}
}

func (bc *Blockchain) VerifyChain() bool {
	for i := 1; i < len(bc.Blocks); i++ {
		currentBlock := bc.Blocks[i]
		previousBlock := bc.Blocks[i-1]

		if currentBlock.Hash != currentBlock.CreateHash() ||
			currentBlock.PreviousHash != previousBlock.Hash {
			return false
		}
	}
	return true
}

func main() {

	blockchain := &Blockchain{}

	for {
		fmt.Print("\n\t\t******** Select from the optoins ********\n\nAdd new Block: (add)\nChange an existing block: (change)\nExit: (no) ")
		print("\n")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		response := strings.ToLower(scanner.Text())

		if response == "add" {
			fmt.Print("Enter the new transaction: ")
			scanner.Scan()
			transaction := scanner.Text()

			fmt.Print("Enter the nonce: ")
			scanner.Scan()
			var nonce int
			_, err := fmt.Sscanf(scanner.Text(), "%d", &nonce)
			if err != nil {
				fmt.Println("Invalid nonce. Please enter a valid integer.")
				continue
			}

			// Add a new block to the blockchain
			blockchain.AddBlock(transaction, nonce)

			// Verify the updated blockchain
			isValid := blockchain.VerifyChain()
			fmt.Printf("Is Blockchain Valid: %v\n", isValid)
		} else if response == "change" {
			// Prompt the user to change an existing block
			fmt.Print("Enter the index of the block you want to change: ")
			scanner.Scan()
			var blockIndex int
			_, err := fmt.Sscanf(scanner.Text(), "%d", &blockIndex)
			if err != nil || blockIndex < 0 || blockIndex >= len(blockchain.Blocks) {
				fmt.Println("Invalid block index. Please enter a valid index.")
				continue
			}

			// Prompt the user for new transaction and nonce
			fmt.Print("Enter the new transaction: ")
			scanner.Scan()
			newTransaction := scanner.Text()

			fmt.Print("Enter the new nonce: ")
			scanner.Scan()
			var newNonce int
			_, err = fmt.Sscanf(scanner.Text(), "%d", &newNonce)
			if err != nil {
				fmt.Println("Invalid nonce. Please enter a valid integer.")
				continue
			}

			// Change the selected block
			blockchain.ChangeBlock(blockIndex, newTransaction, newNonce)

			// Verify the updated blockchain
			isValid := blockchain.VerifyChain()
			fmt.Printf("Is Blockchain Valid: %v\n", isValid)
		} else if response == "no" {
			// Display all blocks and exit the loop if the user chooses not to add more blocks or change existing blocks
			print("\n")
			fmt.Println("All Blocks in the Blockchain: ")
			print("\n")
			blockchain.ListBlocks()
			break
		} else {
			fmt.Println("Invalid response. Please enter 'add', 'change', or 'no'.")
		}
	}
}
