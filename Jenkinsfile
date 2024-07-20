node('10.134.135.130') {
    def goVersion = '1.22.5'
    def goROOT = 'C:\\Go1225'
    def goPATH = 'C:\\Users\\Administrator\\go' 
    def goWorkspace = "${env.WORKSPACE}\\go"
    def autoCancelled = false
    parameters {
        string(name: 'NAME', defaultValue: 'World', description: 'Name to greet')
    }
    try {
	    env.PATH = "${goROOT}\\bin;${goPATH}\\bin;${env.PATH}"
            env.GO1225MODULE = 'on'
	    if (env.CHANGE_ID) {
                echo 'Change ID  ...' + env.CHANGE_ID
            	def skipBuild=env.SKIP_BUILD
            	def branchname=env.CHANGE_BRANCH
	    	echo 'entering pre-flights ...' + env.CHANGE_BRANCH
//              if (branchname == "vinod") {
		 echo 'Starting Pre-Flights'
                 echo 'Branch  ...' + env.CHANGE_BRANCH
                 echo 'Source branch of Pull Request ...' + env.CHANGE_BRANCH
	         checkout scmGit(branches: [[name: "*/${env.CHANGE_BRANCH}"]], extensions: [], userRemoteConfigs: [[credentialsId: 'root', url: 'https://github.com/kayaratvinod/golang.git']])
                 echo 'Target branch of pull request ...' + env.CHANGE_TARGET
	         echo 'Build Number...' + env.BUILD_NUMBER
            	 def BUILDNUMBER=env.BUILD_NUMBER

	        stage('Initialize golang') {
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

		 autoCancelled = true	
//	    	 error('Pre-Flight Succeded')
 //      	    } else {
  //               echo 'This is not a pull request ...' + env.BRANCH_NAME
   //    	    } 
    	} else {
        stage('Build') {
	    echo 'Pulling...' + env.BRANCH_NAME
	    checkout scmGit(branches: [[name: "*/${env.BRANCH_NAME}"]], extensions: [], userRemoteConfigs: [[credentialsId: 'root', url: 'https://github.com/kayaratvinod/golang.git']])
	    echo 'Pulling...' + env.GIT_PR_TRIGGER
	    echo 'Pulling...' + env.GITHUB_PR_STATE
	    echo 'Pulling...' + env.BRANCH_NAME
	    echo 'Pulling...' + env.CHANGE_TARGET 
	    echo 'Build Number...' + env.BUILD_NUMBER 
	    def BUILDNUMBER = env.BUILD_NUMBER 
	    echo BUILDNUMBER
            // Build steps
        }
//	stage('Build Another Job') {
//		build job: 'newpipelinebranch1/ranjith', wait: false
//	}
	stage('Powershell') {
            script {
                // Define a PowerShell script
		def build = env.BUILD_NUMBER
                def scriptContent = '''
                param (
                    [string]$Name = "Jenkins"
                )
                Write-Output "Hello, $Name! Running PowerShell from Jenkins."
                '''

                // Write the PowerShell script to a file
                writeFile file: 'script.ps1', text: scriptContent

                // Run the PowerShell script
                bat 'powershell.exe -NoProfile -ExecutionPolicy Bypass -File script.ps1 -Name "World"'
            }
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
        //      bat 'go build -o hello-world'
		bat 'go build -ldflags="-X main.Version=v1.0.0 -s -w" hello-world.go'
        }
        stage('Run Batch Command') {
        // Reference the parameter inside a batch command
                bat """
        	echo Parameter value: ${params.MY_PARAM}
        	"""
        }
        stage('Package') {
	        def BUILDNUMBER = env.BUILD_NUMBER 
		echo BUILDNUMBER
		def oldFileName = 'hello-world.go'
		def newFileName = env.BUILD_NUMBER
		bat 'echo %BUILD_TAG%'
                bat 'ren hello-world.go %BUILD_TAG%_hello-world.go'
//		bat 'jf rt u ${BUILDNUMBER}_hello-world.go "vinod/"'
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

