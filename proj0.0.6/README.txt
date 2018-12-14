Name: Kyle Ong
Building: 
    ```
        go build distsys/proj0.0.6/frontend/web/main.go
        go build distsys/proj0.0.6/backend/main.go
    ```
Running:
    ```
        cd distsys/proj0.0.6/backend
        go run main.go --listen <portNum> --backend <hostName>:<portNum>,<hostName>:<portNum>

        cd ..
        cd frontend
        go run main.go --listen <portNum> --backend <hostName>:<portNum>,<hostName>:<portNum>,<hostName>:<portNum>
    ```
Testing:
    ```
        go test distsys/proj0.0.6/frontend/services
        go test distsys/proj0.0.6/frontend/utils

        go test distsys/proj0.0.6/backend/services
        go test distsys/proj0.0.6/backend/utils
    ```
Stress Testing:
    ```
        cd ~/distsys/proj0.0.6/stressTest
        chmod +x ./runTests.sh
        ./runTests.sh
    ```

Application state becomes out of synch when you…

(1) Completed Books loose green check mark
- start the app fresh
- mark books as ‘completed’ in the ‘active’ tab
- select ‘All’ tab

(2) Completed Books loose green check mark
- select ‘All’
- mark a book as ‘completed’
- select ‘Finished’
- select ‘All’

(3) Completed Books loose green check mark
- select ‘All’
- mark a book as ‘completed’
- select ‘Active’
- select ‘Finished’
- select ‘All’

(4) Completed books loose green check across multiple clients
- open multiple clients 
- in one terminal window (Client 1)
	```
		cd distsys/proj0.0.6/frontend/web
		go run main.go - - listen 8899
	```
- in another terminal window (Client 2)
	```
		cd distsys/proj0.0.6/frontend/web
		go run main.go - - listen 8989

	```
- in another terminal window (backend)
	```
		cd distsys/proj0.0.6/backend
		go run main.go 
	```
- navigate to localhost:8899 and localhost:8989
- mark a book as completed in localhost:8899
- navigate to localhost:8989
- green check mark is not replicated across both clients

Todo
- [ ] implement more robust error handling / retries on server connections
- [ ] implement comprehensive suite of tests to mock backend server to test backendservices


Part 2

Application follows a MVC architecture. 
Models and associated logic sit on the backend server.
View controllers sit on the frontend server.
Frontend processes HTTP requests from Client.
Frontend then serializes a message to the backend.
Backend performs necessary CRUD operations.
Backend serializes updated models to the frontend.
All happens over TCP.
Frontend renders updated models.

Pros:
    - MVC architecture is modular
    - Service oriented
    - Front-end server is stateless

Cons:
    - View Controller will become bloated
    - Backend server is stateful and contains most of the logic
    - A lot of code


Part 3

I completed the assignment.

Todo
- [ ] exit gracefully after two deteced failures

1.) I provided safe, concurrent, performant access to the data store by adding a mutex. The con of this
    approach, obviously, is that it locks the entire datastore. Thus, only one process can access it at
    a time. The pros, however; are that the datastore is thread-safe and their are absolutely zero chances
    of a deadlock, race-condition, or live-lock.

    This approach is optimal for the use of a reading list. If we consider the use case of a reading list, it is
    safe to assume that it will not be accessed or modified concurrently by many workers. Furthermore, the app
    does not lend itself to heavy amounts of concurrent reads or writes.

    However, additional approaches would need to be considered if this application were to be scaled to a to-do
    list for a group of people in a collaborative setting.

    In terms of performance metrics mean latency was the metric of most interest. Refer to ~/distsys/stressTest
    for test.html files. Here are the relevant terminal outputs from vegeta...
    
    Runing GET stress tests
    mkdir: completed: File exists
    mkdir: get: File exists
    Requests      [total, rate]            1500, 50.02
    Duration      [total, attack, wait]    29.98595123s, 29.985086s, 865.23µs
    Latencies     [mean, 50, 95, 99, max]  880.218µs, 780.673µs, 1.031084ms, 2.205681ms, 26.483977ms
    Bytes In      [total, mean]            3804000, 2536.00
    Bytes Out     [total, mean]            0, 0.00
    Success       [ratio]                  100.00%
    Status Codes  [code:count]             200:1500  
    
    mkdir: active: File exists
    mkdir: get: File exists
    Requests      [total, rate]            1500, 50.03
    Duration      [total, attack, wait]    29.985135061s, 29.984422s, 713.061µs
    Latencies     [mean, 50, 95, 99, max]  894.116µs, 785.09µs, 1.030568ms, 2.260807ms, 48.728809ms
    Bytes In      [total, mean]            3804000, 2536.00
    Bytes Out     [total, mean]            0, 0.00
    Success       [ratio]                  100.00%
    Status Codes  [code:count]             200:1500  
    
    mkdir: all: File exists
    mkdir: get: File exists
    Requests      [total, rate]            1500, 50.03
    Duration      [total, attack, wait]    29.984451115s, 29.983708s, 743.115µs
    Latencies     [mean, 50, 95, 99, max]  932.126µs, 760.091µs, 1.540997ms, 5.690506ms, 30.325703ms
    Bytes In      [total, mean]            3804000, 2536.00
    Bytes Out     [total, mean]            0, 0.00
    Success       [ratio]                  100.00%
    Status Codes  [code:count]             200:1500  
    
    mkdir: readingList: File exists
    mkdir: get: File exists
    Requests      [total, rate]            1500, 50.03
    Duration      [total, attack, wait]    29.985624481s, 29.98218s, 3.444481ms
    Latencies     [mean, 50, 95, 99, max]  3.056803ms, 2.712998ms, 5.810233ms, 9.488214ms, 80.438378ms
    Bytes In      [total, mean]            4896, 3.26
    Bytes Out     [total, mean]            0, 0.00
    Success       [ratio]                  100.00%
    Status Codes  [code:count]             200:1500  
    
    Running POST stress tests
    mkdir: readingList: File exists
    mkdir: post: File exists
    Requests      [total, rate]            1500, 50.03
    Duration      [total, attack, wait]    29.982681349s, 29.981167s, 1.514349ms
    Latencies     [mean, 50, 95, 99, max]  2.220557ms, 2.21154ms, 2.622434ms, 3.997082ms, 5.234903ms
    Bytes In      [total, mean]            31500, 21.00
    Bytes Out     [total, mean]            0, 0.00
    Success       [ratio]                  100.00%
    Status Codes  [code:count]             200:1500  
    
2.) The failure detector follows a ping-ack implementation. Every 10 seconds the front-end server will ask if the
    backend is alive. If the backend is not alive, print the necessary output and try again. After two failures,
    the system exits.

    A heartbeat would also be a viable implementation of a failure detector, however; the general flow of 
    data in the app lends to a ping ack. 

Part 4

Status: Incomplete

Done:
- [x] many to many client-server architecture
- [x] primary and backups
- [x] requests from client to replicas
- [x] rpc calls necessary to implement requests, prepare, prepareok, commit
- [x] clientTable
- [x] in memory log

To-do:
- [ ] view change
- [ ] recovery
- [ ] tests

ViewStamped replication assumes a non-byzantine asynchronous network. It differs from a consensus algorithm such as PAXOS or raft
in that it is a replication strategy with built in consensus rather than a pure consensus algorithm.

Pros:
- Does not need to elect a new leader
- In memory log does not require disk I/O
- Recover is easily achieved
- Merkle Tree can optimize recovery
- Epoch or reconfigurations are easy

Cons:
- Client makes network calls proportional to 2f + 1