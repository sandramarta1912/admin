pipeline {
    agent any

    stages {
        stage('Build') {
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
