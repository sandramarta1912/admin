pipeline {
    agent any

    stages {
        stage('Build') {
            steps {
                echo 'Building..'
                 sh 'echo $PWD'
                 sh 'mkdir -p ${env.PWD}/go'
                sh 'GOPATH=${env.PWD}/go'
                sh 'echo ${env.GOPATH}'
                 git url: 'https://github.com/sandramarta1912/admin'
                 sh 'mkdir -p ${env.GOPATH}/src/github.com/sandramarta1912/admin'
                 sh 'mv * ${env.GOPATH}/src/github.com/sandramarta1912/admin'

                 dir('${env.GOPATH}/src/github.com/sandramarta1912/admin') {
                       sh 'echo $PWD'
                       sh 'go version'
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
