# rabbitmq
This repo is for learning purposes. I tried to implement in code my new knowledge as an output of my learning.
You can see every directory as a single project like Hello-World is one project and also ...

### Used technologies
- `go1.13+`
- `rabbitmq client`

### Hello World
    Hello World is a common project that sends and receives a message. It has no any fancy configuration at all

### Task Queue
    Task Queue is program that sends tasks to multiple workers. It has an ack.
    With ack sender sends the message to worker and waits to workers reply that I have done with this you can delete the queue. If worker 
    terminated or dies before proceed the task, sender will know it and try to re-send the message to next online workers and you can 
    be sure no task will lost.
    
    Even we add ack to workers, we can easily lose messages, because of server's termination. In order not to lose the tasks we can add 
    durability to our queues and even server restarts message can survive. In this case server try to write tasks to disk
    
    If we run the program sender sends the tasks to workers blindly, even the worker processed heavy tasks and the others do light tasks 
    wait for the next tasks. In order to send the tasks to free workers we can add prefetch count and set it 1. This ensures number of 
    tasks will be only one and sender delivir the taks to free workers.