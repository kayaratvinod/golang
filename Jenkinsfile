pipeline {
    agent any

    environment {
        GO122MODULE = 'on'
        GOPATH = "${env.WORKSPACE}/go"
        PATH = "${env.PATH}:${env.GOPATH}/bin"
	BRANCHNAME = ${env.BRANCH_NAME}
    }
    options {
        // This is required if you want to clean before build
        skipDefaultCheckout(true)
    }

    stages {
        stage('Checkout') {
            steps {
                echo "Hello ${params.branchName}"
                git url: 'https://github.com/kayaratvinod/golang.git', branch: "${BRANCHNAME}"
            }
        }
        stage('Initialize golang') {
            steps {
                sh 'go mod init golang'
            }
        }
        stage('Code Analysis') {
            parallel {
                stage('Vet') {
                    steps {
                        sh 'go vet ./...'
                    }
                }
            }
        }
        stage('Install Dependencies') {
            steps {
                sh 'go mod tidy'
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
