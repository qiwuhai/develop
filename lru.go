package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

var (
	input_file *string = flag.String("input", "", "use -input=<>")
	cap        *int    = flag.Int("cap", 10, "use -cap=<>")
	REQTOTAL   int     = 0
	HITCACHE   int     = 0
)

type SinglyLink struct {
	head     *Node
	capacity int
	size     int
}

type Node struct {
	Value string
	Next  *Node
}

func (s *SinglyLink) prependNode(val string) {
	currentHead := s.head
	s.size++
	if currentHead == nil {
		s.head = &Node{val, nil}
		return

	}

	s.head = &Node{val, currentHead}

}

func (s *SinglyLink) deleteLastNode() {
	if s.size <= 1 {
		s.head = nil
		s.size = 0
		return

	}

	tmp := s.head

	for tmp.Next.Next != nil {
		tmp = tmp.Next

	}
	tmp.Next = nil
	s.size--

}

func (s *SinglyLink) deleteNode(val string) {
	if s.size == 0 {
		return

	}
	currentNode := s.head

	if currentNode.Value == val {
		s.head = currentNode.Next
		s.size--
		return

	}
	for currentNode.Next != nil {
		if currentNode.Next.Value == val {
			currentNode.Next = currentNode.Next.Next
			s.size--
			return

		}
		currentNode = currentNode.Next

	}

}

func (s *SinglyLink) searchNode(val string) {
	currentNode := s.head
	REQTOTAL++
	if currentNode == nil {
		s.prependNode(val)
		//fmt.Println("-----------first insert into cache:", val)
		return
	}
	for currentNode != nil {
		if currentNode.Value == val {
			s.deleteNode(val)
			s.prependNode(val)
			HITCACHE++
			//fmt.Println("---------exist in cache:", val)
			return

		}
		nextNode := currentNode.Next
		if nextNode == nil {
			s.checkFull(val)
			//fmt.Println("-------append to cache:", val)
			return

		}
		currentNode = nextNode

	}

}

func (s *SinglyLink) checkFull(val string) {
	isFull := s.size >= s.capacity
	if isFull {
		s.deleteLastNode()

	}
	s.prependNode(val)

}

func (s *SinglyLink) print() {
	tmp := s.head
	for tmp != nil {
		fmt.Printf("%s -> ", tmp.Value)
		tmp = tmp.Next

	}
	fmt.Println()

}

func newSinglyLink(capacity int) *SinglyLink {
	return &SinglyLink{nil, capacity, 0}

}

func main() {

	flag.Parse()
	if len(*input_file) == 0 {
		fmt.Println("ERROR: input file should be set")
		return
	}
	if *cap == 0 {
		*cap = 10
	}
	lru := newSinglyLink(*cap)

	file, err := os.Open(*input_file)
	if err != nil {
		fmt.Println("open file error ", err, *input_file)
		return
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF {
			fmt.Println("read END")
			break
		}
		str = strings.TrimSpace(str)
		lru.searchNode(str)
		//	lru.print()
	}
	lru.print()
	fmt.Println(">>>>>>>>>>>>", REQTOTAL, HITCACHE)
	/*
		lru.prependNode("a")
		lru.prependNode("b")
		lru.prependNode("c")
		lru.searchNode("e")
		lru.print() // e -> c -> b ->
		lru.searchNode("g")
		lru.print() // g ->e -> c ->
		lru.searchNode("h")
		lru.searchNode("a")
		lru.searchNode("h")
		lru.print() // h -> a -> g ->
	*/
}
