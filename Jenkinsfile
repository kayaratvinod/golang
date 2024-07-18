pipeline {
  agent any 
  stages {
        stage('Build') { 
            steps {
                echo "This is Build step."
            }
        }
        stage('Test') { 
            steps {
                echo "This is Test step."
            }
        }
        stage('Deploy') {         
            steps {
              script{
                 echo "This is Deploy step."
                 def changedFiles = [];
                 
                 // Calling getChangedFiles method...
                 changedFiles = getChangedFiles();
                 if(changedFiles.size>0){
                    println(changedFiles)
                 }
                 else{
                    println("No changed file.")
                 }
                 
                 def branchName = "${env.BRANCH_NAME}"
                 // Calling deploy method... 
                 deploy(branchName);
              }
            }
        }
   }
}
def void deploy(String branchName){
    if(branchName == "master"){
       println("Deploying to Prod.")
    }
    else if(branchName == "test"){
       println("Deploying to Test.")
    }

}
def getChangedFiles(){
   def changes = []
   def changeLogSets = currentBuild.changeSets
   def filePath = ""
   for (int i = 0; i < changeLogSets.size(); i++) {
      def entries = changeLogSets[i].items
      for (int j = 0; j < entries.length; j++) {
         def entry = entries[j]
         def files = new ArrayList(entry.affectedFiles)
         for (int k = 0; k < files.size(); k++) {
            def file = files[k]
            filePath = "${file.path}"
            if("${file.editType.name}"!="delete"){
               changes.add(filePath);
             }
         }
      }
   }
   changes.unique();
   changes.sort();
   return changes;
}
