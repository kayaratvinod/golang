node('10.134.137.117') {
    // Define variables
    def goVersion = '1.22.5'
    def goInstaller = "go${goVersion}.windows-amd64.msi"
    def goInstallerUrl = "https://dl.google.com/go/${goInstaller}"
    def goInstallDir = "C:\\Go"
    def goPath = "${env.WORKSPACE}\\go"

    try {
        stage('Clean Up Existing Go Installation') {
            // Uninstall existing Go installation if it exists
            bat """
                IF EXIST "${goInstallDir}" (
                    rmdir /S /Q ${goInstallDir}
                )
            """
            
            // Clean Go environment variables
//            bat """
 //               reg delete "HKCU\\Environment" /v GOPATH /f
 //               reg delete "HKCU\\Environment" /v Path /f
 //           """
            
            // Reset PATH to default
            bat """
                set "PATH=%SystemRoot%\\system32;%SystemRoot%;%SystemRoot%\\System32\\Wbem;%SystemRoot%\\System32\\WindowsPowerShell\\v1.0\\"
            """
        }
        
        stage('Install Go') {
            // Download Go installer
            bat "curl -o ${goInstaller} ${goInstallerUrl}"
            
            // Run the Go installer
            bat "msiexec /i ${goInstaller} /qn INSTALLDIR=${goInstallDir}"
            
            // Clean up
            bat "del ${goInstaller}"
            
            // Add Go to PATH
            bat '''
                setx PATH "%PATH%;C:\\Go\\bin"
            '''
            
            // Verify Go installation
            bat '''
                @echo off
                setlocal
                set "PATH=%PATH%;C:\\Go\\bin"
                go version
		go fmt ./...
                endlocal
            '''
        }

        stage('Clean Up Existing Go Installation') {
            // Uninstall existing Go installation if it exists
            bat """
                IF EXIST "${goInstallDir}" (
                    rmdir /S /Q ${goInstallDir}
                )
            """

            // Clean Go environment variables
            bat """
                reg delete "HKCU\\Environment" /v GOPATH /f
                reg delete "HKCU\\Environment" /v Path /f
            """

            // Reset PATH to default
            bat """
                set "PATH=%SystemRoot%\\system32;%SystemRoot%;%SystemRoot%\\System32\\Wbem;%SystemRoot%\\System32\\WindowsPowerShell\\v1.0\\"
            """
        }

        
    } catch (Exception e) {
        // Handle any errors that occur during the pipeline execution
        echo "An error occurred: ${e.message}"
        currentBuild.result = 'FAILURE'
    }
}
