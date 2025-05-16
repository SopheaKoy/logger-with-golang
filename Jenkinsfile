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
                    def configFileId = "221c9bb7-955e-4feb-9329-9e60b3399d33"
                    echo "Using config file ID: ${configFileId}"

                    configFileProvider([configFile(fileId: configFileId, variable: 'CONFIG_FILE')]) {
                        // At this point, env.CONFIG_FILE contains the full path to the temp config file
                        def configFilePath = env.CONFIG_FILE
                        echo "Config file is at: ${configFilePath}"

                        if (fileExists(configFilePath)) {
                            def config = readYaml file: configFilePath
                            echo "Loaded config. Project name: ${config.PROJECT_NAME ?: 'N/A'}"
                        } else {
                            error "Config file not found at path: ${configFilePath}"
                        }
                    }
                }
            }
        }
    }
}
