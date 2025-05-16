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
        stage('Load Configuration') {
            steps {
                script {
                    // The full path might include folder information
                    def configFileId = "221c9bb7-955e-4feb-9329-9e60b3399d33"
                    
                    echo "Attempting to load config file from another job's folder, ID: ${configFileId}"
                    
                    try {
                        configFileProvider([configFile(fileId: configFileId, variable: 'CONFIG_FILE')]) {
                            if (fileExists(env.CONFIG_FILE)) {
                                echo "Successfully loaded config file at: ${env.CONFIG_FILE}"
                                def config = readYaml file: env.CONFIG_FILE
                                def project_name = config.PROJECT_NAME
                                echo "Project Name: ${project_name}"
                            } else {
                                error "Config file was not found at path: ${env.CONFIG_FILE}"
                            }
                        }
                    } catch (Exception e) {
                        echo "Error accessing config file: ${e.message}"
                        error "Failed to access config file from other job's folder. Check permissions and file ID."
                    }
                }
            }
        }
    }
}