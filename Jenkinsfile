pipeline {
    agent any
    stages {
        stage ('First') {
            steps {
                echo "First"

            if (env.BRANCH_NAME == 'vinod') {
                echo 'First stage is enought, exit 0 shoul happened here'
                currentBuild.result = 'SUCCESS'
                return
            } else {
                echo 'Second stage must be executed'
            }

        }
	}

        stage('Second') {
            steps {
                echo "Second"
            }
        }
    }
}
