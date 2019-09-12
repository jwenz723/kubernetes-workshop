# kubernetes-workshop

A set of modules used to learn about how to use kubernetes.

## Modules

* [skaffold](/skaffold): teaches about how to use skaffold to increase local development velocity


# Local Setup

* Create a github account (if you don't already have one)
* Install docker
    * Create docker hub account
    * authenticate local docker to your docker hub account
* Install minikube
* Install kubectl
* Install helm
* (Optional) install fluxctl



# skaffolddemo

An example of how 2 go services can be ran in parallel in kubernetes using skaffold.

# How to run

1. Install Skaffold: https://skaffold.dev/docs/getting-started/
2. Run the apps: `skaffold dev --port-forward` (this will port-forward the pod ports to your workstation)
3. Once the pods have started up you will see all of the logs for both pods flowing in your terminal, including
a log line for each app stating which port (frontend: 9000, backend: 8080) each pod was forwarded to on your workstation.
4. Browse to http://localhost:9000/hello to test the frontend app. This page does a call to the backend service/pod then
prints the message that was received from the backend (`Hello, from the underworld!`) on the page with a frontend prefix
(`Underworld says:`). The complete message you should see is the following:

```text
Underworld says: Hello, from the underworld!
```