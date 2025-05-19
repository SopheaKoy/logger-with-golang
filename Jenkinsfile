pipeline {
    agent any

    stages {
        stage('List Credentials') {
            when {
                anyOf {
                    branch 'main'
                    branch 'sophea'
                    branch 'dev'
                    branch 'staging'
                }
            }
            steps {
                script {
                    def creds = com.cloudbees.plugins.credentials.CredentialsProvider.lookupCredentials(
                        com.cloudbees.plugins.credentials.common.StandardCredentials.class,
                        Jenkins.instance,
                        null,
                        null)
                    echo "Found credentials:"
                    creds.each { c ->
                        echo "- ${c.id}"
                    }
                }
            }
        }
        stage('Use Secret File') {
            when {
                anyOf {
                    branch 'main'
                    branch 'sophea'
                    branch 'dev'
                    branch 'staging'
                }
            }
            steps {
                // withCredentials exposes the secret file path as an environment variable
                withCredentials([file(credentialsId: 'dev_credetial', variable: 'SECRET_FILE_PATH')]) {
                    script {
                        echo "Secret file path: ${env.SECRET_FILE_PATH}"

                         // Read YAML content from the secret file
                        def secretConfig = readYaml file: env.SECRET_YAML
                        
                        // Example: print first 5 lines of the secret file (do NOT print secrets in real logs!)
                        sh 'head -n 5 "$SECRET_FILE_PATH"'
                    }
                }
            }
        }

        stage('Load Configuration') {
            when {
                anyOf {
                    branch 'main'
                    branch 'sophea'
                    branch 'dev'
                    branch 'staging'
                }
            }
            steps {
                script {
                    
                    try {
                        def provider = org.jenkinsci.lib.configprovider.ConfigProvider.all()
                        provider.each { p ->
                            p.getAllConfigs().each { c ->
                                echo "Found config file: ${c.name} with ID: ${c.id}"
                            }
                        }
                    } catch (Exception e) {
                        echo "Error listing config files: ${e.message}"
                    }
                    def configFileId = ""

                    switch(env.BRANCH_NAME) {
                        case 'dev':
                            configFileId = "dev/221c9bb7-955e-4feb-9329-9e60b3399d33"
                            break
                        case 'prod':
                            // Optional: Skip or assign different config
                            configFileId = "prod/b3a77caf-e908-4e73-a342-1ba1b8621edf"
                            break
                        default:
                            configFileId = "dev/221c9bb7-955e-4feb-9329-9e60b3399d33"
                            break
                    }

                    echo "Using config file ID: ${configFileId} for branch: ${env.BRANCH_NAME}"

                    if (configFileId?.trim()) {
                        configFileProvider([configFile(fileId: configFileId, variable: 'CONFIG_FILE')]) {
                            def config = readYaml file: CONFIG_FILE
                            echo "Loaded config: ${config}"
                        }
                    } else {
                        echo "No config file ID defined for branch: ${env.BRANCH_NAME}. Skipping config loading."
                    }

                    // Load the config file with the determined ID
                    configFileProvider([configFile(fileId: configFileId, variable: 'CONFIG_FILE')]) {
                        def config = readYaml file: env.CONFIG_FILE
                    

                        // Display loaded configuration (except secrets)
                        echo "Loaded configuration for ${ENV} environment"
                        echo "Application: ${env.APPLICATION}"
                        echo "Namespace: ${env.NAMESPACE}"
                        echo "Deploy Server: ${env.DEPLOY_SERVER}"
                    }
                }
            }
        }

        stage('Deployment') {
            when {
                anyOf {
                    branch 'main'
                    branch 'sophea'
                    branch 'dev'
                    branch 'staging'
                }
            }
            steps {
                echo "Deploying branch ${env.BRANCH_NAME}"
            }
        }
    }
}