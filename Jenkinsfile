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
        echo ${WORKSPACE}
        sh 'go env'
        sh 'go build -o agent_${BUILD_ID}_${BUILD_NUMBER} agent/agent.go'
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
