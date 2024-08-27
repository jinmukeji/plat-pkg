pipeline{
    agent any

    environment {
        GOROOT="/var/jenkins_home/tools/org.jenkinsci.plugins.golang.GolangInstallation/go_1.22/go"
        GOPATH="/var/go"
        GOBIN="${env.GOPATH}/bin"
        PATH="${env.GOBIN}:${env.GOROOT}/bin:${env.PATH}"
        GOPROXY="https://goproxy.io,direct"
        GOPRIVATE="github.com/jinmukeji/*"
        AWS_ACCESS_KEY_ID     = credentials('jenkins-aws-secret-key-id')
        AWS_SECRET_ACCESS_KEY = credentials('jenkins-aws-secret-access-key')
        AWS_DEFAULT_REGION    = credentials('jenkins-aws-secret-region')
        PROJECT = "$WORKSPACE"
        DOCKER_HOST_IP ="${DOCKER_HOST_IP}"
        BUILD_NUMBER ="${env.BUILD_NUMBER}"
    }

    stages {    

        stage('初始化环境'){
            steps {
                sh label: '', script: '$PROJECT/jenkins_ci/ci_init.sh'
             }
        }
        stage('代码格式化'){
            steps{
                sh label: '', script: '$PROJECT/jenkins_ci/ci_format.sh'
            }
        }
        stage('静态代码检查'){
            steps{
                sh label: '', script: '$PROJECT/jenkins_ci/ci_lint.sh'
            }
        }
        stage('编译'){
            steps{
                sh label: '', script: '$PROJECT/jenkins_ci/ci_build.sh'
            }
        }
       
     }

   
        success {
            emailext (
              subject: "SUCCESSFUL: Job '${env.JOB_NAME} [${env.BUILD_NUMBER}]'",
              to: "tech@jinmuhealth.com",
              body: """SUCCESSFUL: Job '${env.JOB_NAME} [${env.BUILD_NUMBER}]':</p>
    <p>Check console output at &QUOT;<a href='${env.BUILD_URL}'>${env.JOB_NAME} [${env.BUILD_NUMBER}]</a>&QUOT;</p>""",
              recipientProviders: [[$class: 'DevelopersRecipientProvider']]
            )
        }
        failure {
            emailext (
              subject: "FAILED: Job '${env.JOB_NAME} [${env.BUILD_NUMBER}]'",
              to: "tech@jinmuhealth.com",
              body: """<p>FAILED: Job '${env.JOB_NAME} [${env.BUILD_NUMBER}]':</p>
                <p>Check console output at &QUOT;<a href='${env.BUILD_URL}'>${env.JOB_NAME} [${env.BUILD_NUMBER}]</a>&QUOT;</p>""",
              recipientProviders: [[$class: 'DevelopersRecipientProvider']]
            )
        }
     }
}
