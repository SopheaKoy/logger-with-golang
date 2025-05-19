pipeline {
    agent any

    stages {
        stage('Load Config') {
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
                    // PROVEN WORKING FORMAT for your specific case
                    def configFileId = "production/221c9bb7-955e-4feb-9329-9e60b3399d33"
                    
                    echo "Loading config from: ${configFileId}"
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