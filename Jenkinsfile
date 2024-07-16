node('10.134.135.130') {
def autoCancelled = false
    environment {
        GIT_PR_TRIGGER = "${env.CHANGE_ID}"
	BRANCHNAME = "${env.BRANCH_NAME}"
    }
    try {
        stage('Pre-Flight') {
            def skipBuild=env.SKIP_BUILD
            def branchname=env.BRANCH_NAME
            if ((skipBuild == null || skipBuild.isEmpty()) && branchname == "visod") {
                 echo 'starting build ...' + env.BRANCH_NAME
		 autoCancelled = true	
	    	 error('Pre-Flight Succeded')
       	    } else {
                 echo 'skipping build ...' + env.BRANCH_NAME
       	    } 
    	}
        stage('Build') {
	    echo 'Pulling...' + env.GIT_PR_TRIGGER
	    echo 'Pulling...' + env.GITHUB_PR_STATE
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
	if (autoCancelled) {
		echo 'Pre-Flights Succeeded'
		currentBuild.result = 'SUCCESS'
		return
	}
        currentBuild.result = 'FAILURE'
        throw e
    } finally {
        // Cleanup steps
        cleanWs() // Clean up the workspace
    }
}
