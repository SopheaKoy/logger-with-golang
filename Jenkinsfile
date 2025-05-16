// pipeline {
//     agent any
//     stages {
//         stage('Load Configuration') {
//             steps {
//                 script {
                    
//                     try {
//                         def provider = org.jenkinsci.lib.configprovider.ConfigProvider.all()
//                         provider.each { p ->
//                             p.getAllConfigs().each { c ->
//                                 echo "Found config file: ${c.name} with ID: ${c.id}"
//                             }
//                         }
//                     } catch (Exception e) {
//                         echo "Error listing config files: ${e.message}"
//                     }

//                     echo "Using config file ID: ${configFileId} for branch: ${env.BRANCH_NAME}"

//                     // Load the config file with the determined ID
//                     configFileProvider([configFile(fileId: configFileId, variable: 'CONFIG_FILE')]) {
//                         def config = readYaml file: env.CONFIG_FILE 

//                         // project name
//                         project_name = config.PROJECT_NAME
//                         echo "Project Name: ${project_name}"   
//                     }
//                 }
//             }
//         }
//     }
// }

pipeline {
    agent any

    stages {
        stage('Stage 1 - Preparation') {
            steps {
                echo 'This is Stage 1: Preparation'
            }
        }

        stage('Stage 2 - Build') {
            steps {
                echo 'This is Stage 2: Build'
            }
        }

        stage('Stage 3 - Deploy') {
            steps {
                echo 'This is Stage 3: Deploy'
            }
        }
    }
}
