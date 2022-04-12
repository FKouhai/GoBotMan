pipeline {
  agent any
  tools {
    go '1.18'
  }
  environment {
    GO111MODULE="off"
    GOPATH="${WORKSPACE}"
    BINDEST="${JENKINS_HOME}binaries"
  }
  stages {
    stage("build"){
      steps {
        echo 'Building the agent binary'
        sh 'go build -o agent_${BUILD_ID}_${BUILD_NUMBER} src/agent/agent.go'
      }
    }
  }
  post {
    success {
      echo 'Building has finished successfully'
      sh 'cp agent_${BUILD_ID}_${BUILD_NUMBER} ${env.BINDEST}'
    }
    failure {
      echo 'Building has failed'
    }
  }
}
