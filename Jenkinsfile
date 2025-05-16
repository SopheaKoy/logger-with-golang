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
        stage('Load Config') {
            steps {
                script {
                    def configFileId = '221c9bb7-955e-4feb-9329-9e60b3399d33'  // folder-scoped ID

                    echo "Loading config file with ID: ${configFileId}"

                    try {
                        configFileProvider([configFile(fileId: configFileId, variable: 'CONFIG_FILE')]) {
                            if (fileExists(env.CONFIG_FILE)) {
                                echo "Config file found at ${env.CONFIG_FILE}"
                                def config = readYaml file: env.CONFIG_FILE
                                echo "Project name: ${config.PROJECT_NAME ?: 'NOT SET'}"
                            } else {
                                error "Config file not found at path: ${env.CONFIG_FILE}"
                            }
                        }
                    } catch (err) {
                        echo "Error loading config file: ${err}"
                        error "Failed to load folder-scoped config file. Make sure the job is inside the folder and fileId is correct."
                    }
                }
            }
        }
    }
}
