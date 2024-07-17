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
	    echo 'starting pre-flights ...' + env.CHANGE_BRANCH
     //       if (branchname == "ranjith") {
                 echo 'Branch  ...' + env.BRANCH_NAME
                 echo 'Source branch of Pull Request ...' + env.CHANGE_BRANCH

	    	// checkout scmGit(branches: [[name: '*/branchname']], extensions: [], userRemoteConfigs: [[credentialsId: 'root', url: 'https://github.com/kayaratvinod/golang.git']])
	         checkout scmGit(branches: [[name: "*/${env.CHANGE_BRANCH}"]], extensions: [], userRemoteConfigs: [[credentialsId: 'root', url: 'https://github.com/kayaratvinod/golang.git']])
                 echo 'Target branch of pull request ...' + env.CHANGE_TARGET
	         echo 'Build Number...' + env.BUILD_NUMBER
		 autoCancelled = true	
	    	 error('Pre-Flight Succeded')
      // 	    } else {
       //          echo 'This is not a pull request ...' + env.BRANCH_NAME
       //	    } 
	    }
    	}
        stage('Build') {
	    echo 'Pulling...' + env.BRANCH_NAME
	    checkout scmGit(branches: [[name: "*/${env.BRANCH_NAME}"]], extensions: [], userRemoteConfigs: [[credentialsId: 'root', url: 'https://github.com/kayaratvinod/golang.git']])
	    echo 'Pulling...' + env.GIT_PR_TRIGGER
	    echo 'Pulling...' + env.GITHUB_PR_STATE
	    echo 'Pulling...' + env.BRANCH_NAME
	    echo 'Pulling...' + env.CHANGE_TARGET 
	    echo 'Build Number...' + env.BUILD_NUMBER 
            // Build steps
        }
        stage('Initialize golang') {
            steps {
                sh 'go mod init golang'
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
    } catch (Exception e) {
	if (autoCancelled) {
		cleanWs()
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

