package main

type node struct {
	next  *node
	value int
}

// func reverse(head *node, left, right int) *node {
// 	flag := false
// 	cur := head

// 	for cur != nil {

// 	}
// 	beg.next = tail
// 	prev.next = beg

// 	return head
// }

func reversee(head *node) (head1, tail *node) {
	var behind *node
	cur := head
	for head != nil {
		next := head.next
		head.next = behind
		behind = head
		head = next
	}
	return behind, cur
}

// func recRev(head,tail *node){
// 	if ==nil{
// 		return
// 	}
// 	prev:=head
// 	cur :=head.next
// 	fur :=cur.next
// 	cur.next=prev

// }

func main() {
	head := &node{
		next:  nil,
		value: 1,
	}
	cur := head
	for i := 2; i <= 5; i++ {
		cur.next = &node{
			next:  nil,
			value: i,
		}
		cur = cur.next
	}
	// head1 := reverse(head, 2, 4)
	// for head1 != nil {
	// 	fmt.Println(head1)
	// 	head1 = head1.next

	// }

}
