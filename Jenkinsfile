node('10.134.135.130') {
    def goVersion = '1.22.5'
    def goROOT = 'C:\\Go1225'
    def goPATH = 'C:\\Users\\Administrator\\go' 
    def goWorkspace = "${env.WORKSPACE}\\go"
    def autoCancelled = false

    try {
        stage('Pre-Flight') {
	    env.PATH = "${goROOT}\\bin;${goPATH}\\bin;${env.PATH}"
            env.GO1225MODULE = 'on'
	    if (env.CHANGE_ID) {
            def skipBuild=env.SKIP_BUILD
            def branchname=env.BRANCH_NAME
	    echo 'starting pre-flights ...' + env.CHANGE_BRANCH
     //       if (branchname == "ranjith") {
                 echo 'Branch  ...' + env.BRANCH_NAME
                 echo 'Source branch of Pull Request ...' + env.CHANGE_BRANCH
	         checkout scmGit(branches: [[name: "*/${env.CHANGE_BRANCH}"]], extensions: [], userRemoteConfigs: [[credentialsId: 'root', url: 'https://github.com/kayaratvinod/golang.git']])
                 echo 'Target branch of pull request ...' + env.CHANGE_TARGET
	         echo 'Build Number...' + env.BUILD_NUMBER

	         stage('Initialize golang') {
                	echo 'came here'
                	bat 'go mod init golang'
        	 }
//		 bat 'go mod init golang'
		 bat 'go vet ./...'
		 bat 'go mod tidy'
		 bat 'go fmt ./...'
		 bat 'golint ./...'
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
                bat 'go vet ./...'
        }
        stage('Install Dependencies') {
                bat 'go mod tidy'
        }
        stage('FMT Stage') {
                bat 'go fmt ./...'
        }
        stage('Linting') {
                bat 'golint  ./...'
        }
	stage('Build') {
                bat 'go build -o hello-world'
        }
        stage('Package') {
                bat 'ren hello-world.go "${BUILDNUMBER}"_hello-world.go'
		bat 'jf rt u "${BUILDNUMBER}_hello-world.go" "vinod/"'
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

