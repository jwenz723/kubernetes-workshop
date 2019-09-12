import com.niceincontact.pipeline.*
@Library('cicd-jenkins-shared-libraries@master') _

node {
    try {
        properties(
                [
                        buildDiscarder(logRotator(numToKeepStr: '5'))
                ]
        )

        checkout scm

        def appBackend = new GoApp("demohttp-backend", "backend")
        def appFrontend = new GoApp("demohttp-frontend", "frontend")
        def apps = [appBackend, appFrontend]

        def params = new GoStandardPipelineParams()
        params.apps = apps
        params.directory = "${WORKSPACE}/1-skaffold"
        params.jenkinsCredential_GitAccess = "ESCWORKSTREAM_GITHUB_PERSONAL_ACCESS_TOKEN"
        params.organization = "kubernetes-workshop"

        go.standardPipeline(params)
    }
    catch (exc) {
        currentBuild.result = 'FAILURE'
        echo "${exc}"
        throw exc
    }
    finally {
        stage('Cleanup') {
            cleanWs()
        }
    }
}