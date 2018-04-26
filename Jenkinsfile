pipeline {
    agent any

    stages {
        stage('Build') {
            environment {
                GOPATH = "$PWD/go"
            }
            steps {
                sh 'echo $PWD'
                sh 'echo $GOPATH'

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
