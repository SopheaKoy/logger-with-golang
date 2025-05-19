pipeline {
    agent any

    environment {
        MY_SECRET_FILE = credentials('85da67ed-e8af-4b0b-991a-44e2a306fead')
    }

    stages {
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
                script {
                    echo "Using secret file path: ${env.MY_SECRET_FILE}"
                    sh 'cat $MY_SECRET_FILE'
                }
            }
        }

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
                configFileProvider([configFile(fileId: '85da67ed-e8af-4b0b-991a-44e2a306fead', variable: 'CONFIG_FILE')]) {
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
