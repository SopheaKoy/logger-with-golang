pipeline {
    agent any

    stages {

        stage('Use Config File') {
            when {
                anyOf {
                    branch 'main'
                    branch 'sophea'
                    branch 'dev'
                    branch 'staging'
                }
            }
            steps {
                configFileProvider([configFile(fileId: 'b3a77caf-e908-4e73-a342-1ba1b8621edf', variable: 'CONFIG_FILE')]) {
                    sh 'cat $CONFIG_FILE'
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