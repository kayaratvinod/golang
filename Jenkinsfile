pipeline {
    agent { label '10.134.135.130' }
    stages {
        stage('upload') {
           steps {
		jf 'rt u *.txt vinod/'
		jf 'rt build-publish'
               }
            }
        }
} 
