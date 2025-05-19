pipeline {
    agent any

    stages {
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
                        case 'sophea':
                            configFileId = "dev/221c9bb7-955e-4feb-9329-9e60b3399d33"
                            break
                        default:
                            configFileId = "dev/221c9bb7-955e-4feb-9329-9e60b3399d33"
                            break
                    }
                    echo "Using config file ID: ${configFileId} for branch: ${env.BRANCH_NAME}"

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