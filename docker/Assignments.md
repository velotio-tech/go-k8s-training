# Docker Assignments for the Deep Dive

- Install Docker on your local laptop.

- Create a public docker hub repository with your name/email and login.
  
- Run nginx on the docker and expose it to the port 8080 on localhost. Check in your browser at "http://localhost:8080".
  It should show a page of Welcome to Nginx. Add the screenshot to your branch

- Now download this index.html file from [here](.extrafiles/index.html). Now run the same nginx server from previous assignment but the index.html
  should be replaced with the file provided by us. Now check the browser again and it should show a page of Welcome to Velotio
  Add the screenshot to your branch
  
- Create a mysql docker container and add some data in it. Now delete that container and recreate another mysql container with 
  the same volume and the data you added previously should persist. Add a small video of that.

- Create a simple golang application which takes one input as argument and print that input. Following would be the go binary for the program
  
  For example: ./hello --input "Prafull"
  Output: Prafull
  
  Containerize this application in such a way that you can provide argument "Welcome to Velotio" when you run "docker run" command. 
  The container log should print the argument provided by you, if no argument is provided it should print "Hello World!!!"

  <details><summary>Hint</summary>
  <p>
    #### Explore the command lint input in golang..
    #### docker run hello:velotio "--input Welcome to Velotio"
  </p>
  </details>

  Push the docker image on your repo with image tag: hello:velotio