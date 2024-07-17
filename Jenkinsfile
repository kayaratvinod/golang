node('10.134.135.130') {
    def goVersion = '1.22.5'
    def goInstallDir = 'C:\\Go1225'
    def goWorkspace = "${env.WORKSPACE}\\go"
    def autoCancelled = false

    try {
        stage('Pre-Flight') {
	    env.PATH = "${goInstallDir}\\bin;${env.PATH}"
            env.GO1225MODULE = 'on'
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
		 bat 'go mod init golang'
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
		echo 'came here'
                bat 'go mod init golang'
        }
        stage('Code Analysis') {
		echo 'came here'
                bat 'go vet hello-world.go'
        }
        stage('Install Dependencies') {
                bat 'go mod tidy'
        }
	stage('Build') {
                bat 'go build -o hello-world'
        }
//        stage('Package') {
 //               sh 'mv hello-world.go /tmp/"${BRANCHNAME}"_hello-world.go'
  //      }
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

