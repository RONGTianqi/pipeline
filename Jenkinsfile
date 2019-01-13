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
        try{
            sh 'cd ${PROJ_DIR}/test ; golint'
        } catch (err){
            sh "echo static analyis failed.  See report"
        }
    }
    }
    stage ('Test') {
        withEnv(["GOPATH=${WORKSPACE}", "PATH+GO=${root}/bin:${WORKSPACE}/bin", "GOBIN=${WORKSPACE}/bin"]){
            sh 'cd ${PROJ_DIR}/test ; go test -v -coverprofile=coverage.out -covermode count > tests.out'
            sh "go get github.com/tebeka/go2xunit"
            sh 'cd ${PROJ_DIR}/test ; go2xunit < tests.out -output tests.xml || true'
            sh 'cd ${PROJ_DIR}/test '
            junit 'tests.xml'
            sh "go get github.com/t-yuki/gocover-cobertura"
            sh 'cd ${PROJ_DIR}/test ; gocover-cobertura < coverage.out > coverage.xml'

        }
    }
    stage ('Archive') {
      archiveArtifacts '**/tests.out, **/tests.xml, **/coverage.out, **/coverage.xml, **/coverage2.xml'
    }
}
