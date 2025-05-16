pipeline {
    agent any

    stages {
        stage('Load Configuration') {
            steps {
                script {
                    def configFileId = "221c9bb7-955e-4feb-9329-9e60b3399d33"  // Config file ID

                    echo "Attempting to load config file ID: ${configFileId}"

                    // Load the configuration file using Config File Provider
                    configFileProvider([configFile(fileId: configFileId, variable: 'CONFIG_FILE_PATH')]) {
                        // Ensure the config file exists before proceeding
                        if (!fileExists(env.CONFIG_FILE_PATH)) {
                            error "Config file not found: ${env.CONFIG_FILE_PATH}"
                        }

                        echo "Config file loaded: ${env.CONFIG_FILE_PATH}"

                        // Read and parse the YAML file
                        def config = readYaml file: env.CONFIG_FILE_PATH

                        // Set the necessary environment variables
                        env.PROJECT_NAME = config.project_name
                        env.APPLICATION = config.application
                        env.NAMESPACE = config.namespace
                        env.DEPLOY_SERVER = config.deploy_server
                        env.ENVIRONMENT = config.environment

                        echo "Configuration loaded:"
                        echo "Project: ${env.PROJECT_NAME}"
                        echo "Application: ${env.APPLICATION}"
                        echo "Namespace: ${env.NAMESPACE}"
                        echo "Deploy Server: ${env.DEPLOY_SERVER}"
                    }
                }
            }
        }
    }
}
