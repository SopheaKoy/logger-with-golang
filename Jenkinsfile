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
        stage('Load Shared Managed Config') {
            steps {
                script {
                    configFileProvider([configFile(fileId: 'baefae44-c166-44dc-88b2-d9491ca9ace4', variable: 'CONFIG_FILE')]) {
                        echo "Config file path: ${env.CONFIG_FILE}"

                        if (fileExists(env.CONFIG_FILE)) {
                            def config = readYaml file: env.CONFIG_FILE
                            echo "Loaded config. Project name: ${config.environment?.project_name ?: 'N/A'}"
                        } else {
                            error "Config file not found at path: ${env.CONFIG_FILE}"
                        }
                    }
                }
            }
        }
    }
}
