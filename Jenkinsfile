#!/usr/bin/env groovy
// https://www.cloudbees.com/blog/top-10-best-practices-jenkins-pipeline-plugin

env.PROJ_DIR='src/learningGo'
node() {
  def root = tool name: 'go-1.11', type: 'go'

  withEnv(["GOPATH=${WORKSPACE}","GOROOT=${root}"]) {
        env.PATH="${GOPATH}/bin:$PATH"
        env.PATH="${GOROOT}/bin:$PATH"
        sh 'go version'
  stage('Init gopath') {
              sh 'mkdir -p $GOPATH/{bin,pkg,src}'  // go运行环境目录
          }
  stage('Checkout') {
     checkout([
                    $class: 'GitSCM',
                    branches: [[name: '*/master']],
                    doGenerateSubmoduleConfigurations: false,
                    extensions: [[$class: 'RelativeTargetDirectory', relativeTargetDir: 'src/learningGo']],
                    submoduleCfg: [],
                    userRemoteConfigs: [[
                        credentialsId: 'e4d4cf21-2d28-4212-809c-960b68ff5c6f',
                        url: 'git@github.com:RONGTianqi/pipeline.git'
                    ]]
                ])

  }
  stage ('Compile') {
    sh 'cd ${PROJ_DIR} ; go build '
  }

  stage ('Static Analysis'){
       sh 'go get github.com/golang/lint/golint'
      try{
        sh 'cd ${PROJ_DIR} ; golint'
      } catch (err){
        sh "echo static analyis failed.  See report"
      }
      
      warnings canComputeNew: true, canResolveRelativePaths: true, categoriesPattern: '', consoleParsers: [[parserName: 'Go Vet'], [parserName: 'Go Lint']], defaultEncoding: '', excludePattern: '', healthy: '', includePattern: '', messagesPattern: '', unHealthy: ''

  }
    stage ('Test') {
    sh 'cd ${PROJ_DIR} ; go test'

    }

}
}
