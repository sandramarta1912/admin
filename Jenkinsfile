pipeline {
    agent any

    stages {
        stage('Build') {
            steps {
                echo 'Building..'
                git url: 'https://github.com/sandramarta1912/admin'
                sh 'echo $PWD'
                sh 'which go'
                sh 'echo $USER'
                sh './build.sh'
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