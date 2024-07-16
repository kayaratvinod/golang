node('10.134.135.130') {
    environment {
        GIT_PR_TRIGGER = "${env.CHANGE_ID}"
	BRANCHNAME = "${env.BRANCH_NAME}"
    }
    try {
    stages {
        stage('build') {
            when {
                expression {
                    def skipBuild=env.SKIP_BUILD 
                    return skipBuild == null || skipBuild.isEmpty()
                }
            }
            steps {
                script {
                    echo 'starting build ...'
                }
            }
        }
    
        stage('build') {
            def skipBuild=env.SKIP_BUILD
            if (skipBuild == null || skipBuild.isEmpty()) {
                echo 'starting build ...'
       	    } else {
            	echo 'skipping build ...'
       	    } 
    	}
        stage('Build') {
	    echo 'Pulling...' + env.GIT_PR_TRIGGER
	    echo 'Pulling...' + env.BRANCH_NAME
            // Build steps
        }
        stage('Test') {
            echo 'Testing on a Linux node...'
            // Test steps
        }
        stage('Deploy') {
            echo 'Deploying on a Linux node...'
            // Deploy steps
        }
    }
    } catch (Exception e) {
        currentBuild.result = 'FAILURE'
        throw e
    } finally {
        // Cleanup steps
        cleanWs() // Clean up the workspace
    }
}
