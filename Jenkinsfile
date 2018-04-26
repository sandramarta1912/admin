pipeline {
    agent any
    
    environment {
        GOPATH = '$PWD/go'
    }

    stages {
        stage('Build') {
            steps {
                echo 'Building..'
                git url: 'https://github.com/sandramarta1912/admin'
                sh 'echo ${env.GOPATH}'
            }
        }
        stage('Test') {
            steps {
                echo 'Testing..'
            }
        }
        stage('Deploy') {
            steps {
                echo 'Deploying....'
            }
        }
    }
}
