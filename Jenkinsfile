pipeline {
    agent any

    environment {
        MY_CONFIG = credentials('85da67ed-e8af-4b0b-991a-44e2a306fead')
    }

    stages {
        stage('Print YAML Content') {
            when {
                anyOf {
                    branch 'sophea'
                    branch 'dev' 
                    branch 'uat'
                }
            }
            steps {
                script {
                    echo 'Reading secret config file...'
                    sh 'cat "$MY_CONFIG"'
                }
            }
        }

        stage('Read Config') {
            when {
                anyOf {
                    branch 'sophea'
                    branch 'dev'
                    branch 'uat'
                }
            }
            steps {
                script {
                    def yamlText = readFile(env.MY_CONFIG)
                    def config   = readYaml text: yamlText

                    // Access values from YAML
                    def PORT = config.env.port
                    def ENV  = config.env.env

                    echo "App Port:    ${PORT}"
                    echo "Environment: ${ENV}"
                }
            }
        }
    }
}