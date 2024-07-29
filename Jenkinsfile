node('10.134.137.117') {
    def root = tool type: 'go', name: '1.22.5'

    withEnv(["GOROOT=${root}", "PATH+GO=${root}/bin"]) {
        // Output will be something like "go version go1.19 darwin/arm64"
        sh 'go version'
    }
}
