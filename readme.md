## Intro:

- Everything is a File

## Basic Operations

- Create Empty File
- Truncate a File
- Get File Info
- Rename and Move a File
- Delete Files
- Open and Close Files
- Check if File Exists
- Check Read and Write Permissions
- Change Permissions, Ownership, and Timestamps
- Create Hard Links and Symlinks

## Reading and Writing

- Copy a File
- Seek Positions in File
- Write Bytes to a File
- Quick Write to File
- Use Buffered Writer
- Read up to n Bytes from File
- Read Exactly n Bytes
- Read At Least n Bytes
- Read All Bytes of File
- Quick Read Whole File to Memory
- Use Buffered Reader
- Read with a Scanner

## Archiving(Zipping)

- Archive(Zip) Files
- Extract(Unzip) Archived Files
- Compressing
- Compress a File
- Uncompress a File

## Misc

- Temporary Files and Directories
- Downloading a File Over HTTP
- Hashing and Checksums

---

# Hands-On System Programming with Go

## Chapter 1,

- An Introduction to System Programming, introduces you to Go and
  system programming and provides some basic concepts and an overview of
  Unix and its resources, including the kernel API. It also defines many
  concepts that are used throughout the rest of the book.

## Chapter 2,

- Unix OS Components, focuses on the Unix operating system and the
  components that you will interact with—files and the filesystem, processes,
  users and permissions, threads, and others. It also explains the various
  memory management techniques of the operating system, and how Unix
  handles resident and virtual memory.

## Chapter 3,

- An Overview of Go, takes a look at Go, starting with some history of
  the language and then explaining, one by one, all its basic concepts, starting with namespaces and the type system, variables, and flow control, and finishing with built-in functions and the concurrency model, while also
  offering an explanation of how Go interacts and manages its memory.

## Chapter 4,

- Working with the Filesystem, helps you to understand how the Unix
  filesystem works and how to master the Go standard library to handle file
  path operations, file reading, and file writing.

## Chapter 5,

- Handling Streams, helps you to learn about the interfaces for the
  input and output streams that Go uses to abstract data flows. It explains how they work and how to combine them and best use them without leaking
  information.

## Chapter 6,

- Building Pseudo-Terminals, helps you understand how a pseudoterminal application works and how to create one. The result will be an
  interactive application that uses standard streams just as the command line
  does.

## Chapter 7,

- Handling Processes and Daemons, provides an explanation of what
  processes are and how to handle them in Go, how to start child processes
  from a Go application, and how to create a command-line application that
  will stay in the background (a daemon) and interact with it.
  Chapter 8, Exit Codes, Signals, and Pipes, discusses Unix inter-process
  communication. It explains how to use exit codes effectively. It shows you
  how signals are handled by default inside an application, and how to manage
  them with some patterns for effective signal handling. Furthermore, it
  explains how to connect the output and input of different processes using
  pipes.

## Chapter 9,

- Network Programming, explains how to use a network to make
  processes communicate. It explains how network communication protocols
  work. It initially focuses on low-level socket communication, such as TCP
  and UDP, before moving on to web server development using the wellknown HTTP protocol. Finally, it shows how to use the Go template engine.

## Chapter 10,

- Data Encoding Using Go, explains how to leverage the Go
  standard library to encode complex data structures in order to facilitate
  process communications. It analyzes both text-based protocols, such as XML
  and JSON, and binary-based protocols, such as GOB.

## Chapter 11,

- Dealing with Channels and Goroutines, explains the basics of
  concurrency and channels and some general rules that prevent the creation of
  deadlocks and resource-leaking inside an application.

## Chapter 12,

- Synchronization with sync and atomic, discusses the
  synchronization packages of the sync and sync/atomic standard libraries, and
  how they can be used instead of channels to achieve concurrency easily. It
  also focuses on avoiding the leaking of resources and on recycling resources.

## Chapter 13,

- Coordination Using Context, discusses Context, a relatively new
  package introduced in Go that offers a simple way of handling asynchronous
  operations effectively.

## Chapter 14,

- Implementing Concurrency Patterns, uses the tools from the
  previous three chapters and demonstrates how to use and combine them to
  communicate effectively. It focuses on the most common patterns used in Go
  for concurrency.

## Chapter 15,

- Using Reflection, explains what reflection is and whether you
  should use it. It shows where it's used in the standard library and guides you in creating a practical example. It also shows how to avoid using reflection where there is no need to.

# Chapter 16,

- Using CGO, explains how CGO works and why and when you
  should use it. It explains how to use C code inside a Go application, and vice-versa.

1. ## Section 1: An Introduction to System Programming and Go
1. An Introduction to System Programming
   Technical requirements

- Beginning with system programming
- Software for software
- Languages and system evolution
- System programming and software engineering
- Application programming interfaces
- Types of APIs
- Operating systems
- Libraries and frameworks
- Remote APIs
- Web APIs
- Understanding the protection ring
- Architectural differences
- Kernel space and user space
- Diving into system calls
- Services provided
- Process control
- File management
- Device management
- Information maintenance
- Communication
- The difference between operating systems
- Understanding the POSIX standard
- POSIX standards and features
- POSIX.1 &#x2013; core services
- POSIX.1b and POSIX.1c &#x2013; real-time and thread extensions
- POSIX.2 &#x2013; shell and utilities&#xA0;
- OS adherence
- Linux and macOS
- Windows
  Summary
  Questions

2. ## Unix OS Components
   Technical requirements

- Memory management
- Techniques&#xA0;of management
- Virtual memory
- Understanding files and filesystems
- Operating systems and filesystems
- Linux
- macOS
- Windows
- Files and hard and soft links
- Unix filesystem
- Root and inodes
- Directory structure
- Navigation and interaction
- Mounting and unmounting
- Processes
- Process properties
- Process life cycle
- Foreground and background
- Killing a job
- Users, groups, and permissions
- Users and groups
- Owner, group, and others
- Read, write, and execute
- Changing permission
- Process communications
- Exit codes
- Signals
- Pipes
- Sockets
  Summary
  Questions

3. ## An Overview of Go
   Technical requirements

- Language features
- History of Go
- Strengths and weaknesses
- Namespace
- Imports and exporting symbols
- Type system
- Basic types
- Composite types
- Custom-defined types
- Variables and functions
- Handling variables
- Declaration
- Operations
- Casting
- Scope
- Constants
- Functions and methods
- Values and pointers
- Understanding flow control
- Condition
- Looping
- Exploring built-in functions
- Defer, panic, and recover
- Concurrency model
- Understanding channels and goroutines
- Understanding memory management
- Stack and heap
- The history of GC in Go
- Building and compiling programs
- Install
- Build
- Run
  Summary
  Questions

2. # Section 2: Advanced File I/O Operations

3. ## Working with the Filesystem
   Technical requirements

- Handling paths
- Working directory
- Getting and setting the&#xA0;working directory
- Path manipulation
- Reading from files
- Reader interface
- The file structure
- Using buffers
- Peeking content
- Closer and seeker
- Writing to file
- Writer interface
- Buffers and format
- Efficient writing
- File modes
- Other operations
- Create
- Truncate
- Delete
- Move
- Copy
- Stats
- Changing properties
- Third-party packages
- Virtual filesystems
- Filesystem events
  Summary
  Questions

5. ## Handling Streams
   Technical requirements

- Streams
- Input and readers
- The bytes reader
- The strings reader
- Defining a reader
- Output and writers
- The bytes writer
- The string writer
- Defining a writer
- Built-in utilities
- Copying from one stream to another
- Connected readers and writers
- Extending readers
- Writers and decorators
  Summary
  Questions

6. ## Building Pseudo-Terminals
   Technical requirements

- Understanding pseudo-terminals
- Beginning with teletypes
- Pseudo teletypes
- Creating a basic PTY
- Input management
- Selector
- Command execution
- Some refactor
- Improving the PTY
- Multiline input
- Providing color support to the pseudo-terminal
- Suggesting commands
- Extensible commands
- Commands with status
- Volatile status
- Persistent status
- Upgrading the Stack command
  Summary
  Questions

3. # Section 3: Understanding Process Communication

4. ## Handling Processes and Daemons
   Technical requirements

- Understanding processes
- Current process
- Standard input
- User and group ID
- Working directory
- Child processes
- Accessing child properties
- Standard input
- Beginning with daemons
- Operating system support
- Daemons in action
- Services
- Creating a service
- Third-party packages
  Summary
  Questions

8. ## Exit Codes, Signals, and Pipes
   Technical requirements

- Using exit codes
- Sending exit codes
- Exit codes in bash
- The exit value bit size
- Exit and deferred functions
- Panics and exit codes
- Exit codes and goroutines
- Reading child process exit codes
- Handling signals
- Handling incoming signals
- The signal package
- Graceful shutdowns
- Exit cleanup and resource release
- Configuration reload
- Sending signals to other processes
- Connecting streams
- Pipes
- Anonymous pipes
- Standard input and output pipes
  Summary
  Questions

9. ## Network Programming
   Technical requirements
   - Communicating via networks
   - OSI model
   - Layer 1 &#x2013; Physical layer
   - Layer 2 &#x2013; Data link layer
   - Layer 3 &#x2013; Network layer
   - Layer 4 &#x2013; Transport layer
   - Layer 5 &#x2013; Session layer
   - Layer 6 &#x2013; Presentation layer
   - Layer 7 &#x2013; Application layer
   - TCP/IP &#x2013; Internet protocol suite
   - Layer 1 &#x2013; Link layer
   - Layer 2 &#x2013; Internet layer
   - Layer 3 &#x2013; Transport layer
   - Layer 4 &#x2013; Application layer
   - Understanding socket programming
   - Network package
   - TCP connections
   - UDP connections
   - Encoding and checksum
   - Web servers in Go
   - Web server
   - HTTP protocol
   - HTTP/2 and Go
   - Using the standard package
   - Making a HTTP request
   - Creating a simple server
   - Serving&#xA0;filesystem
   - Navigating through routes and methods
   - Multipart request and files
   - HTTPS
   - Third-party packages
   - gorilla/mux
   - gin-gonic/gin
   - Other functionalities
   - HTTP/2 Pusher
   - WebSockets protocol
   - Beginning with the template engine
   - Syntax and basic usage
   - Creating, parsing, and executing templates
   - Conditions and loops
   - Template functions
   - RPC servers
   - Defining a service
   - Creating the server
   - Creating the client
     Summary
     Questions
10. ## Data Encoding Using Go
    Technical requirements
    - Understanding text-based encoding
    - CSV
    - Decoding values
    - Encoding values
    - Custom options
    - JSON
    - Field tags
    - Decoder
    - Encoder
    - Marshaler and unmarshaler
    - Interfaces
    - Generating structs
    - JSON schemas
    - XML
    - Structure
    - Document Type Definition
    - Decoding and encoding
    - Field tags
    - Marshaler and unmarshaler
    - Generating structs
    - YAML
    - Structure
    - Decoding and encoding
    - Learning about binary encoding
    - BSON
    - Encoding
    - Decoding
    - gob
    - Interfaces
    - Encoding
    - Decoding
    - Interfaces
    - Proto
    - Structure
    - Code generation
    - Encoding
    - Decoding
    - gRPC protocol
      Summary
      Questions
11. # Section 4: Deep Dive into Concurrency

12. ## Dealing with Channels and Goroutines
    Technical requirements
    - Understanding goroutines
    - Comparing threads and goroutines
    - Threads
    - Goroutines
    - New goroutine
    - Multiple goroutines
    - Argument evaluation
    - Synchronization
    - Exploring channels
    - Properties and operations
    - Capacity and size
    - Blocking operations
    - Closing channels
    - One-way channels
    - Waiting receiver
    - Special values
    - nil channels
    - Closed channels
    - Managing multiple operations
    - Default clause
    - Timers and tickers
    - Timers
    - AfterFunc
    - Tickers
    - Combining channels and goroutines
    - Rate limiter
    - Workers
    - Pool of workers
    - Semaphores
      Summary
      Questions
13. ## Synchronization with sync and atomic
    Technical requirements
    - Synchronization primitives
    - Concurrent access and lockers
    - Mutex
    - RWMutex
    - Write starvation
    - Locking gotchas
    - Synchronizing goroutines
    - Singleton in Go
    - Once and Reset
    - Resource recycling
    - Slices recycling issues
    - Conditions
    - Synchronized maps
    - Semaphores
    - Atomic operations
    - Integer operations
    - clicker
    - Thread-safe floats
    - Thread-safe Boolean
    - Pointer operations
    - Value
    - Under the hood
      Summary
      Questions
14. ## Coordination Using Context
    Technical requirements
    - Understanding context
    - The interface
    - Default contexts
    - Background
    - TODO
    - Cancellation, timeout, and deadline
    - Cancellation
    - Deadline
    - Timeout
    - Keys and values&#xA0;
    - Context in the standard library
    - HTTP requests
    - Passing scoped values
    - Request cancellation
    - HTTP server
    - Shutdown
    - Passing values
    - TCP dialing
    - Cancelling a connection
    - Database operations
    - Experimental packages
    - Context in your application
    - Things to avoid
    - Wrong types as keys
    - Passing parameters
    - Optional arguments
    - Globals
    - Building a service with Context
    - Main interface and usage
    - Exit and entry points
    - Exclude list
    - Handling directories
    - Checking file names and contents
      Summary
      Questions
15. ## Implementing Concurrency Patterns
    Technical requirements
    - Beginning with generators
    - Avoiding leaks
    - Sequencing with pipelines
    - Muxing and demuxing
    - Fan-out
    - Fan-in
    - Producers and consumers
    - Multiple producers (N \_ 1)
    - Multiple consumers (1 \_ M)
    - Multiple consumers and producers (N\*M)
    - Other patterns
    - Error groups
    - Leaky bucket
    - Sequencing
      Summary
      Questions
16. # Section 5: A Guide to Using Reflection and CGO
17. ## Using Reflection
    Technical requirements
    - What's reflection?
    - Type assertions
    - Interface assertion
    - Understanding basic mechanics
    - Value and Type methods
    - Kind
    - Value to interface
    - Manipulating values
    - Changing values
    - Creating new values
    - Handling complex types
    - Data structures
    - Changing fields
    - Using tags
    - Maps and slices
    - Maps
    - Slices
    - Functions
    - Analyzing a function
    - Invoking a function
    - Channels
    - Creating channels
    - Sending, receiving, and closing
    - Select statement
    - Reflecting&#xA0;on reflection
    - Performance cost
    - Usage in the standard library
    - Using reflection in a package
    - Property files
    - Using the package
      Summary
      Questions
18. ## Using CGO
    Technical requirements
    - Introduction to CGO
    - Calling C code from Go
    - Calling Go code from C
    - The&#xA0;C and Go type systems
    - Strings and byte slices
    - Integers
    - Float types
    - Unsafe conversions
    - Editing a byte slice directly
    - Numbers
    - Working with slices
    - Working with structs
    - Structures in Go
    - Manual padding
    - Structures in C
    - Unpacked structures
    - Packed structures
    - CGO recommendations
    - Compilation and speed
    - Performance
    - Dependency from C
