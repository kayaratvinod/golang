pipeline {
    agent { label '10.134.135.130' }
    environment {
        GO122MODULE = 'on'
        GOPATH = "${env.WORKSPACE}/go"
        PATH = "${env.PATH}:${env.GOPATH}/bin"
	BRANCHNAME = "${env.BRANCH_NAME}"
    }
    options {
        // This is required if you want to clean before build
        skipDefaultCheckout(true)
    }
    stages {
        stage('upload') {
           steps {
		git url: 'git@github.com:kayaratvinod/golang.git', branch: "${BRANCHNAME}"		
		jf 'rt u *.txt vinod/mybuild'
		jf 'rt build-publish'
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
