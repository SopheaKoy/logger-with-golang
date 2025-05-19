pipeline {
    agent any

    stages {
        stage('Load Config') {
            steps {
                script {
                    // PROVEN WORKING FORMAT for your specific case
                    def configFileId = "production/221c9bb7-955e-4feb-9329-9e60b3399d33"
                    
                    echo "Loading config from: ${configFileId}"

                    configFileProvider([configFile(
                        fileId: configFileId,  // Use "folder-name/uuid" format
                        variable: 'CONFIG_FILE_PATH'
                    )]) {
                        // Verify file resolution
                        if (!fileExists(env.CONFIG_FILE_PATH)) {
                            error """
                            Config file resolution failed!
                            - Attempted path: ${env.CONFIG_FILE_PATH}
                            - Verify the file exists at:
                              ${JENKINS_URL}/job/production/configfiles/editConfig?id=221c9bb7-955e-4feb-9329-9e60b3399d33
                            """
                        }

                        // Process the config
                        def config = readYaml file: env.CONFIG_FILE_PATH
                        echo "Successfully loaded configuration: ${config}"
                    }
                }
            }
        }

        stage('Deployment') {
            steps {
                script {
                    echo "DEPLOYMENT ECHO =================================="
                }
            }
        }
    }
}