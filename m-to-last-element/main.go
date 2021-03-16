package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type linkedListNode struct {
	prev    *linkedListNode
	next    *linkedListNode
	content int64
}

func newLinkedListNode(content int64, prev, next *linkedListNode) *linkedListNode {
	return &linkedListNode{prev, next, content}
}

func (n *linkedListNode) Value() int64 {
	return n.content
}

func (n *linkedListNode) Next() *linkedListNode {
	return n.next
}

func (n *linkedListNode) Prev() *linkedListNode {
	return n.prev
}

type LinkedList struct {
	head     *linkedListNode
	tail     *linkedListNode
	elements []*linkedListNode
}

func NewLinkedList() *LinkedList {
	return &LinkedList{
		elements: []*linkedListNode{},
	}
}

func (l *LinkedList) Head() *linkedListNode {
	return l.head
}

func (l *LinkedList) Tail() *linkedListNode {
	return l.tail
}

func (l *LinkedList) Size() int {
	return len(l.elements)
}

func (l *LinkedList) Insert(n int64) {
	var prevElement *linkedListNode
	if len(l.elements) > 0 {
		prevElement = l.elements[len(l.elements)-1]
	}

	newElement := newLinkedListNode(n, prevElement, nil)

	if len(l.elements) == 0 {
		l.head = newElement
	}

	l.elements = append(l.elements, newElement)
	l.tail = newElement
}

func getMthToLastElement(n int, c []int64) string {
	list := NewLinkedList()
	for _, number := range c {
		list.Insert(number)
	}

	if n > list.Size() {
		return "NIL"
	}

	tail := list.Tail()
	for i := 1; i < n; i++ {
		tail = tail.Prev()
	}

	return fmt.Sprintf("%d", tail.Value())
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	firstMultipleInput := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	nTemp, err := strconv.ParseInt(firstMultipleInput[0], 10, 64)
	checkError(err)
	n := int32(nTemp)

	cTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var c []int64

	for i := 0; i < len(cTemp); i++ {
		cItem, err := strconv.ParseInt(cTemp[i], 10, 64)
		checkError(err)
		c = append(c, cItem)
	}

	// Print the m elements to last elment
	element := getMthToLastElement(int(n), c)

	fmt.Fprintf(writer, "%s\n", element)

	writer.Flush()
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
