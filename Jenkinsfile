node('10.134.137.117') {
    // Define variables
    def goVersion = '1.22.5'
    def goInstaller = "go${goVersion}.windows-amd64.msi"
    def goInstallerUrl = "https://dl.google.com/go/${goInstaller}"
    // def goInstallDir = "C:\\myfolder"
    def goInstallDir = "${env.WORKSPACE}\\mygofolder"
    def goPath = "${env.WORKSPACE}\\go"

    try {
        stage('Install Go') {
	    cleanWs()		
            // Download Go installer
            bat "curl -o ${goInstaller} ${goInstallerUrl}"
            
            // Run the Go installer
            bat "msiexec /i ${goInstaller} /qn INSTALLDIR=${goInstallDir}"
            
            // Clean up
            bat "del ${goInstaller}"
            
            // Add Go to PATH
  //          bat '''
   //             setx PATH "%PATH%;${goInstallDir}\\bin"
    //        '''
        }
        stage('Set Up Go Environment') {
            // Create Go workspace
            bat "mkdir ${goPath}"
            bat "mkdir ${goPath}\\src ${goPath}\\bin ${goPath}\\pkg"
            
            // Set Go environment variables
            bat """
     //           setx GOPATH ${goPath}
      //          setx PATH "%PATH%;${goInstallDir}\\bin;${goPath}\\bin"
                ${goInstallDir}\\bin\\go version
            """
	    cleanWs()
        }
    } catch (Exception e) {
        // Handle any errors that occur during the pipeline execution
        echo "An error occurred: ${e.message}"
        currentBuild.result = 'FAILURE'
    }
}
