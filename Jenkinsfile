pipeline {
    agent any

    stages {
        stage('Build') {
            steps {
                sh 'echo $PWD'
                sh 'echo $GOPATH'
                echo 'Building..'
                sh "mkdir -p $GOPATH/src/github.com/sandramarta1912/admin"
		        git url: 'https://github.com/sandramarta1912/admin'
                sh "mv ./*  $GOPATH/src/github.com/sandramarta1912/admin/"                 
                dir("$GOPATH/src/github.com/sandramarta1912/admin") {
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
