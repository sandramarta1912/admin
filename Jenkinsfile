pipeline {
    agent any

    stages {
        stage('Build') {
            steps {
                echo 'Building..'
                sh 'cd /var/lib/jenkins/go'
                sh 'echo $PWD'
                git url: 'https://github.com/sandramarta1912/admin'
                sh 'echo $PWD'
                sh 'echo $PATH'
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