pipeline {
    agent { label '10.134.135.130' }

    environment {
        GOPATH = 'C:\\Users\\Administrator\\go'
        GOROOT = 'C:\\Go1225'
        PATH = "${GOROOT}\\bin;${GOPATH}\\bin;${env.PATH}"
	BRANCHNAME = "${env.BRANCH_NAME}"
	BUILDNUMBER = "${env.BUILD_NUMBER}"
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
                        bat 'go vet  hello-world.go'
                    }
                }
            }
        }
        stage('Install Dependencies') {
            steps {
                bat 'go mod tidy'
            }
        }
        stage('Formatting') {
            steps {
                bat 'go fmt  hello-world.go'
            }
        }
        stage('Linting') {
            steps {
                bat 'golint  hello-world.go'
            }
        }
	stage('Build') {
            steps {
                bat 'go build -o hello-world'
            }
        }
        stage('Package') {
            steps {
                bat 'ren hello-world.go "${BUILDNUMBER}"_hello-world.go'
		bat 'jf rt u "${BUILDNUMBER}_hello-world.go" "vinod/"'
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
