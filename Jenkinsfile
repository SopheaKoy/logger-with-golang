pipeline {
    agent any

    stages {
        stage('Load Configuration') {
            steps {
                script {
                    // Use "folder-name/file-id" format
                    def configFileId = "job/prod/221c9bb7-955e-4feb-9329-9e60b3399d33"

                    echo "Loading config from: ${configFileId}"

                    configFileProvider([configFile(fileId: configFileId, variable: 'CONFIG_FILE_PATH')]) {
                        echo "Config loaded at: ${env.CONFIG_FILE_PATH}"

                        if (!fileExists(env.CONFIG_FILE_PATH)) {
                            error "Config file missing! Check: ${env.CONFIG_FILE_PATH}"
                        }

                        def config = readYaml file: env.CONFIG_FILE_PATH
                        env.PROJECT_NAME = config.project_name ?: "default"
                        // ... (set other vars)
                    }
                }
            }
        }
    }
}