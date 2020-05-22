/*
Given a linked list, remove all consecutive nodes that sum to zero. Print out the remaining nodes.

For example, suppose you are given the input 3 -> 4 -> -7 -> 5 -> -6 -> 6. 

In this case, you should first remove 3 -> 4 -> -7, then -6 -> 6, leaving only 5

*/

class LinkedList {
    constructor(value) {
        this.head = null;
        this.length = 0;
        this.addToHead(value);
    }

    addToHead(value) {
        const newNode = { value };
        newNode.next = this.head;
        this.head = newNode;
        this.length++;
        return this;
    }

    removeFromHead() {
        if (this.length === 0) {
            return undefined;
        }

        const value = this.head.value;
        this.head = this.head.next;
        this.length--;

        return value;
    }

    find(val) {
        let thisNode = this.head;

        while (thisNode) {
            if (thisNode.value === val) {
                return thisNode;
            }

            thisNode = thisNode.next;
        }

        return thisNode;
    }

    print() {
        let thisNode = this.head;
        for (let i = 0; i < this.length; i++) {
            console.log(i,':',thisNode.value);    
            thisNode = thisNode.next;
        }
    }

    remove(val) {
        if (this.length === 0) {
            return undefined;
        }

        if (this.head.value === val) {
            return this.removeFromHead();
        }

        let previousNode = this.head;
        let thisNode = previousNode.next;

        while (thisNode) {
            if (thisNode.value === val) {
                break;
            }

            previousNode = thisNode;
            thisNode = thisNode.next;
        }

        if (thisNode === null) {
            return undefined;
        }

        previousNode.next = thisNode.next;
        this.length--;
        return this;
    }
}


const list = new LinkedList(9);
list.addToHead(-6);
list.addToHead(6);
list.addToHead(5);
list.addToHead(-7);
list.addToHead(4);
list.addToHead(3);
list.addToHead(-5);

console.log(list);

// 3,7,0,

let map = new Map();
console.log('length of the list : ',list.length);
let a = [];
let node = list.head;

for (let i =0; i < list.length;i++){
    if (i > 0){
        a[i] = a[i-1] + node.value;
    } else {
        a[i] = node.value;
    }
    if (map.has(a[i])){
        // remove the node from [map.get(a[i])+1] to i
        console.log('remove from:', map.get(a[i])+1, ' - ',i);
        /*
        for (let j = map.get(a[i])+1 ; j <= i ;j++) {
            console.log('index:',j,' value:',node.value)
        }*/
        map.get(a[i])
    } else {
        map.set(a[i], i);
    }
    console.log(node.value);
    node = node.next;
    //a.push()
}
console.log(a);
list.print();
