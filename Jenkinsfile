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
                    // Set your config file ID
                    def configFileId = "221c9bb7-955e-4feb-9329-9e60b3399d33"
                    echo "Looking for config file with ID: ${configFileId}"

                    try {
                        // Attempt to load config file using Config File Provider
                        configFileProvider([configFile(fileId: configFileId, variable: 'CONFIG_FILE')]) {
                            def configFilePath = env.CONFIG_FILE
                            
                            if (fileExists(configFilePath)) {
                                echo "✅ Successfully loaded config file at: ${configFilePath}"

                                def config = readYaml file: configFilePath
                                
                                // Example usage of a config property
                                if (config.PROJECT_NAME) {
                                    echo "Project Name: ${config.PROJECT_NAME}"
                                } else {
                                    echo "⚠️ PROJECT_NAME is not defined in the config file."
                                }

                            } else {
                                error "❌ Config file not found at path: ${configFilePath}. Check file ID and folder scope."
                            }
                        }
                    } catch (Exception e) {
                        echo "❌ Exception while accessing config file: ${e}"
                        echo "➡️ Possible reasons:"
                        echo " - Wrong config file ID"
                        echo " - File is scoped to a different folder"
                        echo " - File was deleted"
                        error "🚨 Failed to load config file. See above for details."
                    }
                }
            }
        }
    }
}
