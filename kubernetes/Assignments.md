1. Create a namespace velotio-<your_name> and now everything you create will be created in this namespace

2. Create a pod with nginx image and expose it to port 80

3. Add the label "app: velotio" to the pod.

4. Create a "nginx-deployment" deployment with image 1.7.8 having 2 replicas, defining 80 port as the port that container exposes. (Don't create service)

5. Upgrade the previous deployment image to nginx:1.7.9 and check the rollout history.

6. Create a configmap named "velotio-cm" with key value pair "env1: value1"

7. Create a "busybox-cm" deployment with image busybox and mount the configmap in previous step as an environment variable into the deployment.

8. Create a configmap 'cmvolume' with values 'key1: value1, key2: value2'. Load this as a volume inside another deployment with "busybox-cm2" on path /etc/velotio. Once the pod is created 'ls' into the /etc/velotio

9. Create a secret named "velotio-secret" with key value pairs: "username": "velotiotech" and "password": "Test@123"

10. Create a deployment "busybox-secret" with secrets mounted in such a way that two environment variable is exposed "USERNAME" as "velotiotech" and "PASSWORD" as "Test@123"

11. Create a job "velotio-job" with image busybox that executes the command "echo hello; sleep 30; echo world". Make it run 5 times, one after other.

12. Create a cronjob "velotio-cronjob" with image busybox that runs on a schedule of "*/2 * * * *" and writes "date; echo Hello world" to the logs.

13. Configure a security context for a busy box deployment.
 


