pipeline {
    agent any

    environment {
        GO122MODULE = 'on'
        GOPATH = "${env.WORKSPACE}/go"
        PATH = "${env.PATH}:${env.GOPATH}/bin"
    }
    options {
        // This is required if you want to clean before build
        skipDefaultCheckout(true)
    }

    stages {
        stage('Checkout') {
            steps {
                echo "Hello ${params.branchName}"
            }
        }
        stage('Package') {
            steps {
                echo "Packaging" 
            }
        }
    }
    post {
        always {
            cleanWs()
        }
        success {
            echo 'Pipeline completed successfully!'
        }
        failure {
            echo 'Pipeline failed!'
        }
    }
}
