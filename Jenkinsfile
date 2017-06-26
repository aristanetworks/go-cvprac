#!/usr/bin/env groovy

/**
 * Jenkinsfile
 */
pipeline {
    
    agent{ label 'exec'}
    
    options {
        buildDiscarder(
            // Only keep the 10 most recent builds
            logRotator(numToKeepStr:'10'))
    }

    environment {
        REPO      = 'aristanetworks/go-cvprac'
        BUILD_DIR = '__build'
        GOPATH    = "${WORKSPACE}/${BUILD_DIR}"
        SRC_PATH  = "${GOPATH}/src/github.com/${REPO}"

        projectName = 'GoCvpRac'
        //emailTo = 'eosplus-dev@arista.com'
        emailTo = 'cwomble@arista.com'
        emailFrom = 'eosplus-dev+jenkins@arista.com'
    }

    stages {
        stage ('Checkout') {
            steps {
                sh 'printenv'
                checkout scm
            }
        }

        stage ('Install_Requirements') {
            steps {
                sh """
                make bootstrap || true
                """
                // Stub dummy file
                writeFile file: "api/cvp_node.gcfg", text: "\n[node \"10.81.110.115\"]\nusername = cvpadmin\npassword = cvp123\n"
            }
        }

        stage ('Check_style') {
            steps {
                sh """
                    make check || true
                """
            }
        }

        stage ('Unit Test') {
            steps {
                sh """
                    make unittest || true
                """
            }
        }

        stage ('System Test') {
            steps {
                sh """
                    make systest || true
                """
            }
        }

        stage ('Cleanup') {
            steps {
                sh 'echo cleanup step'
            }
        }
    }

    post {
        failure {
            mail body: "${env.JOB_NAME} (${env.BUILD_NUMBER}) ${env.projectName} build error " +
                       "is here: ${env.BUILD_URL}\nStarted by ${env.BUILD_CAUSE}" ,
                 from: env.emailFrom,
                 subject: "${env.projectName} ${env.JOB_NAME} (${env.BUILD_NUMBER}) build failed",
                 to: env.emailTo
        }
        success {
            mail body: "${env.JOB_NAME} (${env.BUILD_NUMBER}) ${env.projectName} build successful\n" +
                       "Started by ${env.BUILD_CAUSE}",
                 from: env.emailFrom,
                 subject: "${env.projectName} ${env.JOB_NAME} (${env.BUILD_NUMBER}) build successful",
                 to: env.emailTo
        }
    }
}
