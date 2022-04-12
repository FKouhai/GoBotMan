pipeline {
  agent any
  tools {
    go '1.18'
  }
  environment {
    GO111MODULE="off"
    GOPATH="${WORKSPACE}"
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
    }
    failure {
      echo 'Building has failed'
    }
  }
}
