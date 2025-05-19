pipeline {
    agent any

    environment {
        MY_CONFIG = credentials('85da67ed-e8af-4b0b-991a-44e2a306fead')
    }

    stages {
        when {
            anyOf {
                branch 'main'
                branch 'sophea'
                branch 'dev'
                branch 'staging'
            }
        }
        stage('Print YAML Content') {
            steps {
                script {
                    echo "Secret file is stored at: ${env.MY_CONFIG}"
                    sh 'cat $MY_CONFIG'  // Shows the file contents
                }
            }
        }
    }
}
