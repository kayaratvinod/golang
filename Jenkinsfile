node('10.134.135.130') {
def autoCancelled = false
    environment {
        GIT_PR_TRIGGER = "${env.CHANGE_ID}"
	BRANCHNAME = "${env.BRANCH_NAME}"
    }
    try {
        stage('Pre-Flight') {
	    if (env.CHANGE_ID) {
            def skipBuild=env.SKIP_BUILD
            def branchname=env.BRANCH_NAME
            if (branchname == "ranjith") {
                 echo 'starting pre-flights ...' + env.BRANCH_NAME
                 echo 'starting pre-flights ...' + env.CHANGE_BRANCH
		 autoCancelled = true	
	    	 error('Pre-Flight Succeded')
       	    } else {
                 echo 'This is not a pull request ...' + env.BRANCH_NAME
       	    } 
	    }
    	}
        stage('Build') {
	    checkout scmGit(branches: [[name: '*/ranjith']], extensions: [], userRemoteConfigs: [[credentialsId: 'root', url: 'https://github.com/kayaratvinod/golang.git']])
	    echo 'Pulling...' + env.GIT_PR_TRIGGER
	    echo 'Pulling...' + env.GITHUB_PR_STATE
	    echo 'Pulling...' + env.BRANCH_NAME
	    echo 'Pulling...' + env.CHANGE_TARGET 
	    echo 'Build Number...' + env.BUILD_NUMBER 
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

