// pipeline {
//     agent any

//     stages {
//         stage('Load Configuration') {
//             steps {
//                 script {
//                     def configFileId = ""

//                     switch(env.BRANCH_NAME) {
//                         case 'sophea':
//                             configFileId = "221c9bb7-955e-4feb-9329-9e60b3399d33"  
//                             break
//                         default:
//                             configFileId = "221c9bb7-955e-4feb-9329-9e60b3399d33"  // Default to Dev config
//                             break
//                     }

//                     echo "Using config file ID: ${configFileId} for branch: ${env.BRANCH_NAME}"

//                     configFileProvider([configFile(fileId: configFileId, variable: 'CONFIG_FILE')]) {
//                         def config = readYaml file: env.CONFIG_FILE

//                         env.PROJECT_NAME     = config.project_name

//                         echo "Loaded configuration for ${env.ENVIRONMENT} environment"
//                         echo "Application: ${env.APPLICATION}"
//                         echo "Namespace: ${env.NAMESPACE}"
//                         echo "Deploy Server: ${env.DEPLOY_SERVER}"
//                     }
//                 }
//             }
//         }
//     }
// }

pipeline {
    agent any

    stages {
        stage('Disable prod job in folder') {
            steps {
                script {
                    // Replace 'folder1/prod' with your actual folder and job name
                    def jobPath = '/job/production/'
                    def prodJob = Jenkins.instance.getItemByFullName(jobPath)

                    if (prodJob != null) {
                        if (!prodJob.isDisabled()) {
                            prodJob.disable()
                            echo "Job '${jobPath}' has been disabled."
                        } else {
                            echo "Job '${jobPath}' is already disabled."
                        }
                    } else {
                        error("Job '${jobPath}' not found.")
                    }
                }
            }
        }
    }
}
