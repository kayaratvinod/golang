node('10.134.135.130') {
    environment {
        GIT_PR_TRIGGER = "${env.CHANGE_ID}"
	BRANCHNAME = "${env.BRANCH_NAME}"
    }
    try {
        stage('build') {
            def skipBuild=env.SKIP_BUILD
            def branchname=env.BRANCH_NAME
            if ((skipBuild == null || skipBuild.isEmpty()) && branchname == "vinXd") {
                echo 'starting build ...' + env.BRANCH_NAME
       	    } else {
                echo 'skipping build ...' + env.BRANCH_NAME
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
    } catch (Exception e) {
        currentBuild.result = 'FAILURE'
        throw e
    } finally {
        // Cleanup steps
        cleanWs() // Clean up the workspace
    }
}
