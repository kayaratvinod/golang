node('10.134.137.117') {
    // Define variables
    def goVersion = '1.22.5'
    def goInstaller = "go${goVersion}.windows-amd64.msi"
    def goInstallerUrl = "https://dl.google.com/go/${goInstaller}"
    def newGoInstallDir = "${env.WORKSPACE}\\godir" // Change this to your desired installation path
    def goPath = "${env.WORKSPACE}\\go"

    try {
        stage('Clean Up Existing Go Installation') {
            // Uninstall existing Go installation if it exists
            bat """
                IF EXIST "C:\\Go" (
                    rmdir /S /Q "C:\\Go"
                )
                IF EXIST "${newGoInstallDir}" (
                    rmdir /S /Q "${newGoInstallDir}"
                )
            """
            
            // Reset PATH to default
            bat """
                set "PATH=%SystemRoot%\\system32;%SystemRoot%;%SystemRoot%\\System32\\Wbem;%SystemRoot%\\System32\\WindowsPowerShell\\v1.0\\"
            """
        }
        
        stage('Install Go') {
            // Download Go installer
            bat "curl -o ${goInstaller} ${goInstallerUrl}"
            
            // Run the Go installer with custom installation path
            bat "msiexec /i ${goInstaller} /qn INSTALLDIR=${newGoInstallDir}"
            
            // Clean up installer file
            bat "del ${goInstaller}"
            
            // Add Go to PATH
            bat """
                setx PATH "%PATH%;${newGoInstallDir}\\bin"
            """
            
            // Verify Go installation
            bat """
                @echo off
                setlocal
                set "PATH=%PATH%;${newGoInstallDir}\\bin"
                go version
                endlocal
            """
        }
        
        stage('Set Up Go Environment') {
            // Create Go workspace
            bat "mkdir ${goPath}\\src ${goPath}\\bin ${goPath}\\pkg"
            
            // Set Go environment variables
            bat """
                setx GOPATH ${goPath}
                setx PATH "%PATH%;${newGoInstallDir}\\bin;${goPath}\\bin"
            """
            
            // Verify Go environment
            bat """
                @echo off
                setlocal
                set "PATH=%PATH%;${newGoInstallDir}\\bin;${goPath}\\bin"
                go version
                endlocal
            """
        }
        
        stage('Run Go Code') {
            // Example: Print Go version to ensure everything is set up correctly
            bat """
                @echo off
                setlocal
                set "PATH=%PATH%;${newGoInstallDir}\\bin;${goPath}\\bin"
                go version
                endlocal
            """
            
            // Here you can add commands to run your Go code
            // bat "go run your_go_file.go"
        }
    } catch (Exception e) {
        // Handle any errors that occur during the pipeline execution
        echo "An error occurred: ${e.message}"
        currentBuild.result = 'FAILURE'
    }
}
