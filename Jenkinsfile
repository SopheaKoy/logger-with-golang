pipeline {
    agent any

    stages {
        stage('Load Configuration') {
            steps {
                script {
                    def configFileId = "/job/production/configfiles/c5560f0a-c718-4a62-bfbc-736d8ca54674"

                    echo "Attempting to load config file ID: ${configFileId}"

                    // Use the Config File Provider plugin to load the configuration
                    configFileProvider([configFile(fileId: configFileId, variable: 'CONFIG_FILE_PATH')]) {
                        echo "Configuration file loaded successfully."

                        // Ensure the config file exists before proceeding
                        if (!fileExists(env.CONFIG_FILE_PATH)) {
                            error "Config file not found: ${env.CONFIG_FILE_PATH}. Please verify the config file ID and ensure the file exists."
                        }

                        echo "Config file exists: ${env.CONFIG_FILE_PATH}"

                        // Read and parse the YAML file
                        def config = readYaml file: env.CONFIG_FILE_PATH

                        // Set the necessary environment variables from the YAML
                        env.PROJECT_NAME = config.project_name
                        env.APPLICATION = config.application
                        env.NAMESPACE = config.namespace
                        env.DEPLOY_SERVER = config.deploy_server
                        env.ENVIRONMENT = config.environment

                        echo "Loaded configuration:"
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
