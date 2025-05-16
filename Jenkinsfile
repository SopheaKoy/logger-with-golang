pipeline {
    agent any

    stages {
        stage('Load Configuration') {
            steps {
                script {
                    def configFileId = '33a57b3f-a0f7-40d9-91ec-406743f45d63'  // <<---- Add this line

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

                    echo "Using config file ID: ${configFileId} for branch: ${env.BRANCH_NAME}"

                    // Load the config file with the determined ID
                    configFileProvider([configFile(fileId: configFileId, variable: 'CONFIG_FILE')]) {
                        def config = readYaml file: env.CONFIG_FILE 

                        // project name
                        def project_name = config.PROJECT_NAME
                        echo "Project Name: ${project_name}"   
                    }
                }
            }
        }
    }
}
