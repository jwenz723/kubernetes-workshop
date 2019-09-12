# kubernetes-workshop-demohttp

An example app written in go to show some of the features of kubernetes.

## Getting Started

1. Create a github account (if you don't already have one): https://github.com/join?source=header-home
1. Install docker: https://hub.docker.com/?overlay=onboarding
    * Create a docker hub account (if you don't already have one): https://hub.docker.com/signup?next=%2F%3Fref%3Dlogin
    * Authenticate your local docker installation to your docker hub account by opening a terminal/prompt and executing 
    `docker login` then entering your credentials
1. Follow OS specific instructions for your machine:

#### Windows
1. Install chocolatey: https://chocolatey.org/docs/installation#installing-chocolatey
1. Install git: https://git-scm.com/download/win
1. Install kubectl: `choco install kubernetes-cli`
1. Install minikube: `choco install minikube`
1. Start minikube: `minikube start`
1. Install helm: `choco install kubernetes-helm`
1. Install skaffold: `choco install skaffold`

#### Mac

1. Install homebrew: `/usr/bin/ruby -e "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/master/install)"`
1. Install git: `brew install git`
1. Install kubectl: `brew install kubernetes-cli`
1. Install minikube: `brew cask install minikube`
1. Start minikube: `minikube start`
1. Install helm: `brew install kubernetes-helm`
1. Install skaffold: `brew install skaffold`


## Using skaffold

Minikube provides you with a fully functional kubernetes cluster to test with. Skaffold provides the ability to quickly
push changes that you write into an app or deployment into a cluster (like minikube).

#### How to run an app using skaffold

1. Start by forking this repository into your own github account by clicking the 'fork' button at the top right corner of this page.
1. Now clone this repository to a directory on your workstation (replace `jwenz723` with your username): `git clone https://github.com/jwenz723/kubernetes-workshop-demohttp.git`
1. cd into the cloned directory: `cd kubernetes-workshop-demohttp`
1. Start up the apps using skaffold: `skaffold dev --port-forward` (this will port-forward the pod ports to your workstation)
1. Once the pods have started up you will see all of the logs for both pods flowing in your terminal, including
a log line for each app stating which port (frontend: 9000, backend: 9001) each pod was forwarded to on your workstation.
1. Browse to http://localhost:9000/hello to test the frontend app. This page does a call to the backend service/pod then
prints the message that was received from the backend (`Hello, from the underworld!`) on the page with a frontend prefix
(`Underworld says:`). The complete message you should see is the following:

```text
Underworld says: Hello, from the underworld!
```

#### Testing local development changes

1. While the apps are running, modify the source code of either the backend or frontend app then save your changes. (If 
you're not sure what code to change, you can start by modifying the value of `HELLO_RESPONSE` in [cmd/frontend/main.go](/cmd/frontend/main.go)) 
1. Skaffold will automatically recompile the app and then deploy the changes into your minikube instance. 

#### Making a 'production' change

1. Now that you have played with development of these apps, it is time for you to commit back to the project.
1. Add yourself as a contributor by modifying the `contributors` slice/array found in 
[cmd/backend/main.go](/cmd/backend/main.go) within the `handleContributors()` function. 
1. Create a new git commit with your changes included:
 
    ```bash
    git add .
    git commit -m "Added myself to contributors"
    git push
    ```
   
1. Now create a pull request to merge your fork into the original repository: https://github.com/jwenz723/kubernetes-workshop-demohttp/compare
1. Once your changes have been merged, you will see a build kick off in Jenkins, followed by the newly created docker image deployed to EKS.
1. You have successfully followed the flow of GitOps to deploy code changes into a kubernetes cluster.