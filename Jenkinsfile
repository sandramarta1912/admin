pipeline {
    agent any

    stages {
        stage('Build') {
            steps {
                echo 'Building..'
                sh 'env.GOPATH=$PWD'
                dir('$GOPATH/go/src/github.com/conves/admin') {
                   sh 'go version'
                   sh 'echo $GOPATH'
                   sh './build.sh'
                   }
                }
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
