pipeline {
    agent any
    stages {
        stage ('First') {
            steps {
                echo "First"
                currentBuild.result = 'SUCCESS'
                return
            }
	}

        stage('Second') {
            steps {
                echo "Second"
            }
        }
    }
}
