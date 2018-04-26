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
                sh 'mkdir -p $GOPATH'
                sh 'mkdir -p $GOPATH/src/github.com/sandramarta1912/admin'
                sh 'mv ./* $GOPATH/src/github.com/sandramarta1912/admin'
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
