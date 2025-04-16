package generic

import (
    "fmt"
    "strconv"
    "cmp"
)

type Intfloat interface {
    int | float64
}

func Double[T Intfloat] (in T) T {
   return in * 2
}


type Printable1 interface {
    ~int | ~float64
}
type Printable2 interface {
    fmt.Stringer
}
type Printable interface {
    Printable1
    Printable2
}

type PrintableInt int

func (p PrintableInt) String() string {
    return strconv.Itoa(int(p))
}

type Printablefloat64 float64

func (p Printablefloat64) String() string {
    return strconv.FormatFloat(float64(p), 'f', 2, 64)
}

func PrintToScreen[T Printable] (in T) {
    fmt.Println(in.String())
}

// generic singly linked list

type Node[T cmp.Ordered] struct {
    Val T
    Next *Node[T]
}

func (n *Node[T]) Add (val T) {
    curr := n
    for curr != nil {
        // check if the curr node is the last node in the list
        if next_node := curr.Next; next_node == nil {
            // we are at the last node!
            new_node := Node[T]{
                Val: val,
                Next: nil,
            }
            curr.Next = &new_node
            break
        }
        curr = curr.Next
    }
}

func (n Node[T]) Contains (val T) bool {
    curr := &n
    for curr != nil {
        val_ := (*curr).Val
        if val_ == val {
            return true
        }
        curr = (*curr).Next
    }
    return false
}

func (n *Node[T]) Insert (index int, val T) bool {
    // returns false if the index is larger than the length of the list
    // index has to be starting from 0
    var running_idx int = 0
    curr := n
    for running_idx < index-1 {
        curr = (*curr).Next
        if curr == nil && running_idx < index - 1 {
            return false
        }
        running_idx++
    }
    rest_list := (*curr).Next
    new_node := Node[T]{
        Val: val,
        Next: rest_list,
    }
    (*curr).Next = &new_node
    return true

}


func (n Node[T]) PrettyPrinter() {
    curr := &n
    for curr != nil {
        fmt.Printf("%v\n", (*curr).Val)
        if next_node := (*curr).Next; next_node != nil {
            // Not at the last node
            fmt.Printf("   |\n   |\n   |\n")
        }
        curr = (*curr).Next
    }
}

