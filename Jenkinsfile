pipeline {
    agent any

    environment {
        MY_CONFIG = credentials('85da67ed-e8af-4b0b-991a-44e2a306fead')
    }

    stages {
        stage('Print YAML Content') {
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
                    echo "Secret file is stored at: ${env.MY_CONFIG}"
                    sh 'cat $MY_CONFIG'  // Shows the file contents
                }
            }
        }
    }

    stage('Read Config') {
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
                def yamlText = readFile("${CONFIG_FILE}")
                def config   = readYaml text: yamlText
                def port     = config.env.port
                def envName  = config.env.env

                echo "App Port: ${port}"
                echo "Environment: ${envName}"  // Be careful not to echo secrets
            }
        }
    }
}
