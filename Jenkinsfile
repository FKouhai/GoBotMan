pipeline {
  agent any
  stages {
    stage("build"){
      steps {
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
