pipeline {
    agent any

    stages {
        stage('Build') {
            steps {
                echo 'Building..'
                 sh 'echo $PWD'
                 sh 'mkdir $PWD/go'
                 sh 'env.GOPATH=$PWD/go'
                 sh 'echo $GOPATH'
                 git url: 'https://github.com/sandramarta1912/admin'
                 sh 'mkdir $GOPATH/src/github.com/sandramarta1912/admin'
                 sh 'mv * $GOPATH/src/github.com/sandramarta1912/admin'

                 dir('$GOPATH/src/github.com/sandramarta1912/admin') {
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
