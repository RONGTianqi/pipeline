env.PROJ_DIR='src/learningGo'
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

        try{
          sh 'cd ${PROJ_DIR} ; golint'
        } catch (err){
          sh "echo static analyis failed.  See report"
        }
    }
    }
  stage ('Test') {
     withEnv(["GOPATH=${WORKSPACE}", "PATH+GO=${root}/bin:${WORKSPACE}/bin", "GOBIN=${WORKSPACE}/bin"]){
        sh 'cd ${PROJ_DIR}'
        sh 'go test -v -coverprofile=coverage.out -covermode count > tests.out'
        sh 'go2xunit < tests.out -output tests.xml'
        junit 'tests.xml'

        sh 'gocover-cobertura < coverage.out > coverage.xml'
        step([$class: 'CoberturaPublisher', coberturaReportFile: 'coverage.xml'])
      }
    }
    stage ('Archive') {
      archiveArtifacts '**/tests.out, **/tests.xml, **/coverage.out, **/coverage.xml, **/coverage2.xml'
    }

}
