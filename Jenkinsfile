pipeline {
    agent any

    stages {
        stage('upload') {
           steps {
              script { 
                 def server = Artifactory.server 'vinod'
                 def uploadSpec = """{
                    "files": [{
                       "pattern": "*.txt",
                       "target": "vinod/"
                    }]
                 }"""

                 server.upload(uploadSpec) 
               }
            }
        }
    } 
}
