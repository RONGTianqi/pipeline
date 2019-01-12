#!/usr/bin/env groovy
// https://www.cloudbees.com/blog/top-10-best-practices-jenkins-pipeline-plugin


node() {
  def root = tool name: 'go-1.11', type: 'go'

  withEnv(["GOPATH=${WORKSPACE}","GOROOT=${root}"]) {
        env.PATH="${GOPATH}/bin:$PATH"
        env.PATH="${GOROOT}/bin:$PATH"
        sh 'go version'
        sh 'rm -rf * '

  stage('Checkout') {
     checkout([
                    $class: 'GitSCM',
                    branches: [[name: '*/master']],

                    userRemoteConfigs: [[
                        credentialsId: 'e4d4cf21-2d28-4212-809c-960b68ff5c6f',
                        url: 'git@github.com:RONGTianqi/pipeline.git'
                    ]]
                ])

  }
  stage ('Compile') {
    sh 'cd ${WORKSPACE}/src/learningGo; go build '
  }
  stage ('Static Analysis'){

        try{
          sh 'cd ${WORKSPACE}/src/learningGo ; golint'
        } catch (err){
          sh "echo static analyis failed.  See report"
        }


    }



}
}
