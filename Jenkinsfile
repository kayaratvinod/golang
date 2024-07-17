pipeline {
    agent { label '10.134.135.130' }

    environment {
        GOROOT = 'C:\\Users\\Administrator\\go'
        GOPATH = 'C:\\Go1225'
        PATH = "${GOROOT}\\bin;${GOPATH}\\bin;${env.PATH}"
	BRANCHNAME = "${env.BRANCH_NAME}"
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
                bat 'go mod init golang'
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
	stage('Build') {
            steps {
                sh 'go build -o hello-world'
            }
        }
        stage('Package') {
            steps {
                sh 'mv hello-world.go /tmp/"${BRANCHNAME}"_hello-world.go'
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
