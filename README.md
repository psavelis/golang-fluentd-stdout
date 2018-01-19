# golang-fluentd-stdout

# Run docker image on your machine
> docker run gcr.io/dotzcloud-production/golang-fluentd

# Build and Push to Registry
> docker build -t gcr.io/dotzcloud-production/golang-fluentd .
> gcloud docker -- push gcr.io/dotzcloud-production/golang-fluentd:latest

# Run Deployment on Kubernetes
> kubectl apply -f deployment.yaml

# Expose "for new deployments only" to TCP 80.
> kubectl expose deployment golang-fluentd --type=LoadBalancer --port 80

# Get External IP
> kubectl get svc --selector app=golang-fluentd