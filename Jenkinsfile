pipeline {
    agent any

    stages {
        stage('Build') {
            steps {
                echo 'Building..'
                git url: 'https://github.com/sandramarta1912/admin'
                sh git pull
                sh cd admin
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