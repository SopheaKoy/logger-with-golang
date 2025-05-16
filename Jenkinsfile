pipeline {
    agent any

    stages {
        stage('Load Configuration') {
            steps {
                script {
                    def configFileId = '221c9bb7-955e-4feb-9329-9e60b3399d33'
                    
                    echo "Starting to search for config file ID: ${configFileId}"
                    def found = false
                    def allFiles = []
                    
                    try {
                        def provider = org.jenkinsci.lib.configprovider.ConfigProvider.all()
                        echo "Found ${provider.size()} config providers"
                        
                        provider.each { p ->
                            echo "Provider: ${p.getClass().getName()}"
                            def configs = p.getAllConfigs()
                            echo "Provider has ${configs.size()} configs"
                            
                            configs.each { c ->
                                allFiles.add("${c.name} (${c.id})")
                                echo "Found config file: ${c.name} with ID: ${c.id}"
                                if (c.id == configFileId) {
                                    echo "MATCH FOUND: ${c.name} with ID: ${c.id}"
                                    found = true
                                }
                            }
                        }
                    } catch (Exception e) {
                        echo "Error listing config files: ${e.message}"
                        echo "Exception stack trace: ${e.printStackTrace()}"
                    }

                    if (found) {
                        echo "Using config file ID: ${configFileId} for branch: ${env.BRANCH_NAME}"
                        try {
                            configFileProvider([configFile(fileId: configFileId, variable: 'CONFIG_FILE')]) {
                                echo "Config file loaded to: ${env.CONFIG_FILE}"
                                echo "Config file exists: ${fileExists(env.CONFIG_FILE)}"
                                
                                if (fileExists(env.CONFIG_FILE)) {
                                    echo "Config file content: ${readFile(env.CONFIG_FILE).take(100)}..." // Show first 100 chars
                                    def config = readYaml file: env.CONFIG_FILE
                                    def project_name = config.PROJECT_NAME
                                    echo "Project Name: ${project_name}"
                                } else {
                                    error "Config file resolved but doesn't exist at path: ${env.CONFIG_FILE}"
                                }
                            }
                        } catch (Exception e) {
                            echo "Error in configFileProvider: ${e.message}"
                            echo "Exception stack trace: ${e.printStackTrace()}"
                            echo "Available config files: ${allFiles.join('\n')}"
                            error "Failed to use config file"
                        }
                    } else {
                        echo "Available config files: ${allFiles.join('\n')}"
                        error "Config file ID ${configFileId} not found! Please verify it exists in Jenkins."
                    }
                }
            }
        }
    }
}