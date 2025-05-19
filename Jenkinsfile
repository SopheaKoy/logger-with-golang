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

        stage('Use Config File') {
            steps {
                configFileProvider([configFile(fileId: '221c9bb7-955e-4feb-9329-9e60b3399d33', variable: 'CONFIG_FILE')]) {
                    sh 'cat $CONFIG_FILE'
                    // use your config file here
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