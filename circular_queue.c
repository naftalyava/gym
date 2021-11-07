typedef enum {
    EMPTY = 0,
    FULL,
    OTHER
} MyCircularQueueState;

typedef struct {
    int *arr;
    int size;
    int head;
    int tail;
    MyCircularQueueState state;
} MyCircularQueue;


MyCircularQueue* myCircularQueueCreate(int k) {
    MyCircularQueue *cq = (MyCircularQueue*)malloc(sizeof(*cq));
    if (!cq) return NULL;

    cq->arr = (int*)malloc(sizeof(int) * k);

    if (!cq->arr) {
        free(cq);
        return NULL;
    }

    cq->size = k;
    cq->head = 0;
    cq->tail = 0;
    cq->state = EMPTY;

    return cq;
}

bool myCircularQueueEnQueue(MyCircularQueue* obj, int value) {
    if (!obj || !obj->arr || obj->state == FULL) return false;
    obj->arr[obj->tail] = value;

    obj->tail = (obj->tail + 1) % obj->size;
    
    if (obj->head == obj->tail)
        obj->state = FULL;
    else 
        obj->state = OTHER;
    
    return true;
}

bool myCircularQueueDeQueue(MyCircularQueue* obj) {
    if (!obj || !obj->arr || obj->state == EMPTY) return false;

    obj->head = (obj->head + 1) % obj->size;


    if (obj->tail == obj->head) {
        obj->state = EMPTY;
    }
    if (obj->state == FULL) {
        obj->state = OTHER;
    }

    return true;
}

int myCircularQueueFront(MyCircularQueue* obj) {
    if (!obj || !obj->arr || obj->state == EMPTY) return -1;

    return obj->arr[obj->head];

}

int myCircularQueueRear(MyCircularQueue* obj) {
    if (!obj || !obj->arr || obj->state == EMPTY) return -1;

    if (obj->tail == 0) {
        return obj->arr[obj->size - 1];
    } else {
        return obj->arr[obj->tail - 1];
    }
    
}

bool myCircularQueueIsEmpty(MyCircularQueue* obj) {
    if (!obj || !obj->arr || obj->state == EMPTY) return true;
    else return false;
}

bool myCircularQueueIsFull(MyCircularQueue* obj) {
    if (!obj || !obj->arr || obj->state == FULL) return true;
    else return false;
}

void myCircularQueueFree(MyCircularQueue* obj) {
    if (!obj || !obj->arr) return NULL;

    free(obj->arr);
    free(obj);
}

/**
 * Your MyCircularQueue struct will be instantiated and called as such:
 * MyCircularQueue* obj = myCircularQueueCreate(k);
 * bool param_1 = myCircularQueueEnQueue(obj, value);
 
 * bool param_2 = myCircularQueueDeQueue(obj);
 
 * int param_3 = myCircularQueueFront(obj);
 
 * int param_4 = myCircularQueueRear(obj);
 
 * bool param_5 = myCircularQueueIsEmpty(obj);
 
 * bool param_6 = myCircularQueueIsFull(obj);
 
 * myCircularQueueFree(obj);
*/
