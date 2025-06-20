**gRPC Chat App**
A full-stack bi-directional chat application built with React, Go, and gRPC, supporting real-time streaming communication between users. This project demonstrates the use of modern backend architecture patterns (gRPC, Protocol Buffers, Concurrency Handling) with seamless REST/gRPC bridging and a clean frontend built in React.

**PROJECT OVERVIEW**:
This is a real-time chat app the tech stack used is:
-Frontend: React (Axios)
-Backend: Go (gRPC server + RESTful Gateway)
-Transport Protocols: HTTP/2 (gRPC) and HTTP/1.1 (REST via gRPC Gateway)
-Streaming: Bi-directional streaming using gRPC
-Containerization: Docker

![image](https://github.com/user-attachments/assets/cbef4b7a-d28f-4709-8cf2-476c83ae3ce1)


**ARCHITECTURE:**
Here's how the internal architechture looks like
![image](https://github.com/user-attachments/assets/821a2793-09c1-4573-bf10-c7d49ca154af)

**BACKEND DESIGN:**
I've used Golang for writing the server side code as well as the gateway. This could've been done in Java/C++ or any other language, Golang was just a personal preference as Go compiles directly to machine code, which results in faster execution speeds compared to Java.

**KEY CONCEPTS:**
-Mutex (sync.Mutex): Prevents race conditions when multiple users interact with shared structures like messageHistory or clients map.
-Channel Buffers: Used to queue messages to individual clients efficiently.
-Clean Folder Structure: Code is modularized with separation of concerns (proto, server, gateway).
-Cross-Origin Resource Sharing (CORS): Enabled using rs/cors for safe frontend communication.
-HTTP/REST to gRPC Translation via gRPC-Gateway
-Protocol Buffers
-Multiplexed Bi-directional Streaming
-RESTful Send and Receive APIs
-TLS/HTTPS ready support (for future productionization)
-Graceful Error Handling & Logging

**RUNNING LOCALLY:**
1. Start the gRPC server located in the 'cmd' folder.
2. Start the gRPC gateway located in the 'cmd' as well.
3. Run the frontend using 'npm run dev'.
#NOTE: If you have firewall protection enabled then you will be asked for permission when you run server and gateway, allow it.

**WHY THIS STACK?**
-gRPC: Efficient binary protocol suited for real-time, low-latency communication.
-Go: Fast, concurrency-friendly language for scalable servers.
-React: Flexible and interactive UI development.
-gRPC Gateway: Allows hybrid use (HTTP + gRPC) without frontend complications.

📂 **PROTO FILES AND CODE REFERENCE:**
Refer to main.proto and main.go for core logic.

**CONTACT & FEEDBACK:**
Feel free to raise an issue or reach out on LinkedIn if you want to collaborate or learn more :)
Linkedin :- https://www.linkedin.com/in/tilak-jain-521913328/
Email : tilakj0108@gmail.com
