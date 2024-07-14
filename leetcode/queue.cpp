#include <iostream>
#include <string>
#include <vector>
using namespace std;

const int MAX = 100000;

struct Node {
  Node* next;
  string name;

  Node()
}

int qu[MAX];
int tail = 0, head = 0;

void init() {
  head = tail = 0;
}

bool isEmpty() {
  return (head == tail);
}

bool isFull() {
  return (head == (tail + 1) % MAX);
}

void enqueue(int x) {
  if (isFull()) {
    cout << "error: queue is full" << endl;
    return;
  }

  qu[tail] = x;
  ++tail;
  if(tail == MAX) tail = 0;
}

int dequeue() {
  if(isEmpty()){
    cout << "error: queue is empty" << endl;
    return;
  }
  int res = qu[head];
  ++head;
  if(head==MAX) head = 0;
  return res;
}

int main() {
  init();

  enqueue(3);
  enqueue(5);
  enqueue(9);

  cout << qu << endl;
  cout << dequeue() << endl;

  cout << qu << endl;
  cout << dequeue() << endl;
  
  enqueue(33);
  cout << qu << endl;  
}
