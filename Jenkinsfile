node('10.134.137.117') {
    // Define variables
    def goVersion = '1.20.3'
    def goInstaller = "go${goVersion}.windows-amd64.msi"
    def goInstallerUrl = "https://dl.google.com/go/${goInstaller}"
    def goInstallDir = "C:\\myfolder"
    def goPath = "${env.WORKSPACE}\\go"

    try {
        stage('Install Go') {
            // Download Go installer
            bat "curl -o ${goInstaller} ${goInstallerUrl}"
            
            // Run the Go installer
            bat "msiexec /i ${goInstaller} /qn INSTALLDIR=${goInstallDir}"
            
            // Clean up
            bat "del ${goInstaller}"
            
            // Add Go to PATH
            bat '''
                setx PATH "%PATH%;C:\\myfolder\\bin"
                go version
            '''
        }
    } catch (Exception e) {
        // Handle any errors that occur during the pipeline execution
        echo "An error occurred: ${e.message}"
        currentBuild.result = 'FAILURE'
    }
}
