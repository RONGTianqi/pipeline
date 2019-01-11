#!/usr/bin/env groovy
// https://www.cloudbees.com/blog/top-10-best-practices-jenkins-pipeline-plugin

node() {
  def root = tool name: 'go-1.11', type: 'go'

  stage('Preparation') {
    checkout scm
  }

  stage ('Static Analysis'){
    withEnv(["GOPATH=${WORKSPACE}", "PATH+GO=${root}/bin:${WORKSPACE}/bin", "GOBIN=${WORKSPACE}/bin"]){

      try{
        sh 'cp golint ${WORKSPACE}/bin'
        sh "golint ."
        sh "go vet ."
      } catch (err){
        sh "echo static analyis failed.  See report"
      }
      
      warnings canComputeNew: true, canResolveRelativePaths: true, categoriesPattern: '', consoleParsers: [[parserName: 'Go Vet'], [parserName: 'Go Lint']], defaultEncoding: '', excludePattern: '', healthy: '', includePattern: '', messagesPattern: '', unHealthy: ''
    }
  }

}
