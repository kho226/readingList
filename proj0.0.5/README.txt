Name: Kyle Ong
Building: 
    ```
        go build distsys/proj0.0.5/frontend/web/main.go
        go build distsys/proj0.0.5/backend/main.go
    ```
Testing:
    ```
        go test distsys/proj0.0.5/frontend/services
        go test distsys/proj0.0.5/frontend/utils

        go test distsys/proj0.0.5/backend/services
        go test distsys/proj0.0.5/backend/utils
    ```
Running:
    ```
        cd distsys/proj0.0.5/backend
        go run main.go --listen <portNum>

        cd ..
        cd frontend
        go run main.go --listen <portNum> --backend <hostName>
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
		cd distsys/proj0.0.5/frontend/web
		go run main.go - - listen 8899
	```
- in another terminal window (Client 2)
	```
		cd distsys/proj0.0.5/frontend/web
		go run main.go - - listen 8989

	```
- in another terminal window (backend)
	```
		cd distsys/proj0.0.5/backend
		go run main.go 
	```
- navigate to localhost:8899 and localhost:8989
- mark a book as completed in localhost:8899
- navigate to localhost:8989
- green check mark is not replicated across both clients

Todo
- [ ] implement more robust error handling / retries on server connections
- [ ] implement comprehensive suite of tests to mock backend server to test backendservices

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
