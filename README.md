# kubernetes-workshop-demohttp

An example app written in go to show some of the features of kubernetes.

## Getting Started

1. [Create a github account](https://github.com/join?source=header-home) (if you don't already have one)
1. [Install docker](https://hub.docker.com/?overlay=onboarding)
    * [Create a docker hub account](https://hub.docker.com/signup?next=%2F%3Fref%3Dlogin) (if you don't already have one)
    * Authenticate your local docker installation to your docker hub account by opening a terminal/prompt and executing 
    `docker login` then entering your credentials
1. Install a package manager
    * Windows: [Install chocolatey](https://chocolatey.org/docs/installation#installing-chocolatey)
    * Mac: Install homebrew: `/usr/bin/ruby -e "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/master/install)"`
1. Install git
    * [Windows](https://git-scm.com/download/win)
    * Mac: `brew install git`

## Installing Kubernetes

You can choose to install a local kubernetes cluster using either Minikube or Docker For Desktop. Both are good solutions.

#### Minikube

1. Install minikube: 
    * Windows: `choco install minikube`
        > Windows: If you have hyperv enabled, you will need to start minikube using hyperv rather than the default of virtualbox. You can follow these instructions to get going with hyperv: https://minikube.sigs.k8s.io/docs/reference/drivers/hyperv/
    * Mac: `brew cask install minikube`
1. Start minikube: `minikube start`

#### Docker For Desktop

1. Install Docker For Desktop
    * [Windows](https://docs.docker.com/docker-for-windows/install/)
    * [Mac](https://docs.docker.com/docker-for-mac/install/)
1. Enable kubernetes within settings:
    * [Windows](https://docs.docker.com/docker-for-windows/#kubernetes)
    * [Mac](https://docs.docker.com/docker-for-mac/#kubernetes)

## Additional Tools

Install the following tools to make use of the power of Kubernetes:

1. Install kubectl
    * Windows: `choco install kubernetes-cli`
    * Mac: `brew install kubernetes-cli`
1. Install helm
    * Windows: `choco install kubernetes-helm`
    * Mac: `brew install kubernetes-helm`
1. Install skaffold
    * Windows: `choco install skaffold`
    * Mac: `brew install skaffold`

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

1. When you are done playing with skaffold simply enter `ctrl+c` to stop skaffold and cleanup what was installed into the cluster.

#### Testing local development changes

1. While the apps are running, modify the source code of either the backend or frontend app then save your changes. (If 
you're not sure what code to change, you can start by modifying the value of `HELLO_RESPONSE` in [cmd/frontend/main.go](/cmd/frontend/main.go)) 
1. Skaffold will automatically recompile the app and then deploy the changes into your minikube instance. 

#### Testing local deployment changes

1. Follow the instructions here to get your minikube instance setup to mirror what is deployed into our 'production' kubernetes
cluster: https://github.com/jwenz723/kubernetes-workshop-flux
1. After you have finished those steps, make sure to come back to this repo and try out the steps below.

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
