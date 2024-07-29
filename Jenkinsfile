node('10.134.137.117') {
    // Define variables
    def goVersion = '1.22.5'
    def goInstaller = "go${goVersion}.windows-amd64.msi"
    def goInstallerUrl = "https://dl.google.com/go/${goInstaller}"
    // def goInstallDir = "C:\\myfolder"
    def goInstallDir = "${env.WORKSPACE}\\mygofolder"
    def goPath = "${env.WORKSPACE}\\go"

    try {
        stage('GIT Repo') {
	    cleanWs()
	    echo 'Pulling...' + env.BRANCH_NAME
	    checkout scmGit(branches: [[name: "*/${env.BRANCH_NAME}"]], extensions: [], userRemoteConfigs: [[credentialsId: 'root', url: 'https://github.com/kayaratvinod/golang.git']])
        }

        stage('Install Go') {
            // Download Go installer
            bat "curl -o ${goInstaller} ${goInstallerUrl}"

            // Run the Go installer
            bat "msiexec /i ${goInstaller} /qn INSTALLDIR=${goInstallDir}"

            // Clean up
            bat "del ${goInstaller}"

        }

        stage('Set Up Go Environment') {
            // Create Go workspace
//            bat "mkdir ${goPath}"
 //           bat "mkdir ${goPath}\\src ${goPath}\\bin ${goPath}\\pkg"

            // Set Go environment variables
            bat """
                ${goInstallDir}\\bin\\go version
            """
        }
        stage('Initialize golang') {
            bat """
		${goInstallDir}\\bin\\go mod init golang
            """
        }
        stage('Code Vetting') {
            bat """
		${goInstallDir}\\bin\\go vet ./...
	    """
        }
        stage('Go code formatting FMT Stage') {
            bat """
                ${goInstallDir}\\bin\\go fmt ./...
	    """
        }
	stage('cleanup') {
		cleanWs()		
	}

    } catch (Exception e) {
        // Handle any errors that occur during the pipeline execution
        echo "An error occurred: ${e.message}"
        currentBuild.result = 'FAILURE'
    }
}
