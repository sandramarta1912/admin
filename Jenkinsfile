node {

   stage 'build'
   git url: ''
   sh './build.sh'
   sh 'docker build . --no-cache'
   sh 'docker build -t martasandra/admin:latest . '
   sh 'docker push martasandra/admin:latest '
}