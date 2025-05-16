pipeline {
    agent any

    stages {
        stage('Load Configuration') {
            steps {
                script {
                    // Define the config file ID based on the branch
                    def configFileId = ""

                    switch(env.BRANCH_NAME) {
                        case 'sophea':
                            configFileId = "221c9bb7-955e-4feb-9329-9e60b3399d33"  // Config file ID for sophea branch
                            break
                        default:
                            configFileId = "221c9bb7-955e-4feb-9329-9e60b3399d33"  // Default config for other branches
                            break
                    }

                    echo "Using config file ID: ${configFileId} for branch: ${env.BRANCH_NAME}"

                    // Load the configuration file
                    configFileProvider([configFile(fileId: configFileId, variable: 'CONFIG_FILE_PATH')]) {
                        // Read and parse the YAML configuration file
                        def config = readYaml file: env.CONFIG_FILE_PATH

                        // Set environment variables from the YAML file content
                        env.PROJECT_NAME = config.project_name
                        env.APPLICATION = config.application
                        env.NAMESPACE = config.namespace
                        env.DEPLOY_SERVER = config.deploy_server
                        env.ENVIRONMENT = config.environment  // Make sure to set this from the YAML

                        // Echo the loaded environment variables for debugging purposes
                        echo "Loaded configuration for ${env.ENVIRONMENT} environment"
                        echo "Project: ${env.PROJECT_NAME}"
                        echo "Application: ${env.APPLICATION}"
                        echo "Namespace: ${env.NAMESPACE}"
                        echo "Deploy Server: ${env.DEPLOY_SERVER}"
                    }
                }
            }
        }

        // Add further stages for your pipeline, e.g., deploy, build, etc.
        stage('Deploy') {
            steps {
                script {
                    // Use the environment variables (e.g., deploy based on the environment)
                    echo "Deploying ${env.APPLICATION} to ${env.NAMESPACE} on server ${env.DEPLOY_SERVER}"
                }
            }
        }
    }
}
