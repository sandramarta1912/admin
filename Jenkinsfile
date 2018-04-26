pipeline {
    agent any

    stages {
        stage('Build') {
            steps {
                echo 'Building..'
                sh 'echo $PWD'
                sh 'env.PWD=/var/lib/jenkins/workspace/pipeline-jenkins/go/src/github.com/conves/admin'
                sh 'echo $PWD'
                dir('/var/lib/jenkins/workspace/pipeline-jenkins/go/src/github.com/conves/admin') {
                   git url: 'https://github.com/sandramarta1912/admin'
                   sh 'go version'
                   sh 'echo $GOPATH'
                   sh './build.sh'
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
}
