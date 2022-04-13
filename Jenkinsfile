pipeline {
  agent any
  tools {
    go '1.18'
  }
  environment {
    GO111MODULE="off"
    def CURRDATE = sh(script: "echo `date +%s`", returnStdout: true).trim()
    GOPATH="${WORKSPACE}"
    BINDEST="${JENKINS_HOME}/binaries"
    AGENTBIN="agent_${BUILD_ID}_${CURRDATE}"
  }
  stages {
    stage("build"){
      steps {
        echo 'Building the agent binary'
        sh 'go build -o ${AGENTBIN} src/agent/agent.go'
      }
    }
  }
  post {
    success {
      echo 'Building has finished successfully'
      sh "cp ${AGENTBIN} ${env.BINDEST}"
    }
    failure {
      echo 'Building has failed'
    }
  }
}
