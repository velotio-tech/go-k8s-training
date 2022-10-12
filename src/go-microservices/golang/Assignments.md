# Assignments
---

### Assignment 1

Topics covered:

**Go inbuilt packages like `os`,`bufio`, executing linux commands in Golang**

*Linux shell*

Create a simple linux interactive shell that supports bare minimum commands like `ls`, `pwd`, `exit`, `cd`. Also print 
the current user, hostname and present working directory before any command input like this

`abhishek@ubuntu ~/go/src/github.com/hello-world` ls

Bonus: Implement `history` command support.

---

### Assignment 2

Topics covered:

**File handling, basic encoding/encryption, cobra package, command line input handling, various inbuilt packages like 
strings, os, time etc**

*Personal Journal App*

Create a CLI application using cobra to store personal journal log with user management. For simplicity consider you could 
enter only textual content in journal data.

Terminologies:

  - User : An independent entity with access to its own journal. Data should not be shared between users.

  - Journal : A text log containing multiple entries.

  - Journal Entry : A piece for text accompanied by a timestamp

Features required:

   - User Management: On starting the application the user should be presented with either Login or Sign Up options. 
   On Sign Up do not ask for a password again.
   - Journal Management: After authentication, the user should be presented with two options to either list all his previous 
   entries or create a new entry. Maximum 50 Journal entries should be allowed per user. Newer entry after 50 should replace
   the oldest entry.
   - While showing previous entries each entry should be preceded by time in readable format and data followed by the text input
         Eg.
         25 Jun 2019 10.30pm - Some text that the user entered
         23 Jun 2019 10.00am - Some text that user entered

   - The application should store data on files locally but not in any database (hosted locally including authentication data).
     These files must be encrypted so that User should not be able to read journals by looking at the files through explorers or 
     any other software 
   - The application should support multiple account creation up to 10 accounts.

Bonus: Add support for flags also along with user prompt interface like

`journal entry --add <JOURNAL-ENTRY-TEXT> --user <USERNAME> --passwd <PASSWORD>`

---

### Assignment 3

Topics covered:

**API development and Http communication using client, Json Encoding/decoding, DB connection, Docker**

*Microservice based e-commerce backend*

Create a small e-commerce backend having two micro-services: `User` and `Order`. For now ignore the authentication part.
`User` service only will be exposed to the public and will communicate with the `Order` service for any order related queries.

You are admin and can perform all operations. Things to consider
- Handle CRUD operation of Users and Orders. There can be 1:n mapping between user and orders.
- Can only update one order at a time
- Can delete one or all orders of a user at a time
- Store the data in any DB of your choice
- Containerize your services and test it.


    Eg:
    http://localhost:80/users -> all users
   
    http://localhost:80/users/0/orders -> all orders of user 0
   
    http://localhost:80/users/0/orders/0 - order 0 of user 0
---

### Exercise 4

Topics covered:

**Helm chart, ingress, k8s in-cluster communication**

Create a helm chart for application developed in Exercise 3 and deploy it on Kubernetes cluster.
Ensure proper communication between your micro-services and use ingress to expose your application to
the outside world

---


