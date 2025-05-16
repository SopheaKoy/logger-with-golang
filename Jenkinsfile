pipeline {
    agent any

    stages {
        stage('Load Configuration') {
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
                            configFileId = "221c9bb7-955e-4feb-9329-9e60b3399d33"  
                            break
                        default:
                            configFileId = "65ce3b9e-b918-4ee8-a23e-7f5239aeb2ee"  // Default to Dev config
                            break
                    }
                    echo "Using config file ID: ${configFileId} for branch: ${env.BRANCH_NAME}"

                    // Load the config file with the determined ID
                    configFileProvider([configFile(fileId: configFileId, variable: 'CONFIG_FILE')]) {
                        def config = readYaml file: env.CONFIG_FILE
                        
                        // Environment settings
                        env.ENVIRONMENT     = config.environment.name
                        env.NAMESPACE       = sh(script: 'git rev-parse --abbrev-ref HEAD | tr "[:upper:]" "[:lower:]"', returnStdout: true).trim()
                        env.DEPLOY_SERVER   = config.environment.deploy_server
                        env.HOST_USER       = config.environment.host_user
                        env.SUDO_PASSWORD   = config.environment.sudo_password
                        env.PORT            = config.environment.port.toString()
                        env.WORKERS             = config.environment.workers.toString()
                        env.URL                 = config.environment.url
                        env.SSH_KNOWN_HOSTS     = config.environment.ssh_known_hosts

                        env.ANSIBLE_PLAYBOOK    = config.environment.ansible_playbook
                        env.ANSIBLE_INVENTORY   = config.environment.ansible_inventory

                        env.APPLICATION         = config.environment.application
                        env.SERVICE_REPLICAS    = config.environment.service_replicas
                    
                        
                        // Docker configuration
                        env.DOCKER_REGISTRY = config.docker.registry
                        env.DOCKER_FOLDER   = config.docker.folder
                        env.DOCKER_IMAGE    = config.docker.image
                        env.DOCKER_REGISTRY_USER        = config.docker.registry_user
                        env.DOCKER_REGISTRY_PASSWORD    = config.docker.registry_password
                        
                        // Authentication
                        env.JWT_SECRET_KEY  = config.auth.jwt_secret_key
                        env.JWT_REFRESH_SECRET_KEY = config.auth.jwt_refresh_secret_key
                        
                        // Main database
                        env.DB_USER = config.database.user
                        env.DB_PASS = config.database.password
                        env.DB_HOST = config.database.host
                        env.DB_PORT = config.database.port.toString()
                        env.DB_NAME = config.database.name
                        env.DB_DIALECT  = config.database.dialect
                        env.DB_SSL      = config.database.ssl.toString()
                                                
                        // Telegram configuration
                        env.TELEGRAM_BOT_TOKEN  = config.telegram.bot_token
                        env.TELEGRAM_CHAT_ID    = config.telegram.chat_id
                        
                        // Display loaded configuration (except secrets)
                        echo "Loaded configuration for ${ENV} environment"
                        echo "Application: ${env.APPLICATION}"
                        echo "Namespace: ${env.NAMESPACE}"
                        echo "Deploy Server: ${env.DEPLOY_SERVER}"
                    }
                }
            }
        }
    }
}