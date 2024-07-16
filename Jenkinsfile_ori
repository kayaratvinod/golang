node('10.134.135.130') {
    environment {
        GIT_PR_TRIGGER = "${env.CHANGE_ID}"
	BRANCHNAME = "${env.BRANCH_NAME}"
    }
    def autoCancelled = false
    try {
        stage('Pre-Flight') {
                 def skipBuild=env.SKIP_BUILD
                 def branchname=env.BRANCH_NAME
                 if ((skipBuild == null || skipBuild.isEmpty()) && branchname == "vinod") {
                      echo 'starting build ...' + env.BRANCH_NAME
		      currentBuild.result = 'SUCCESS'
		      autoCancelled = true	
	    	      error('Aborting the build.')
       	         } else {
                      echo 'skipping build ...' + env.BRANCH_NAME
       	         } 
	}
   }catch (Exception e) {
	if (autoCancelled) {
		echo 'PipeLine Succeeded'
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
