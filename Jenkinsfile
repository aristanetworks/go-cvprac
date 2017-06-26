#!/usr/bin/env groovy
pipeline {
    agent { docker 'golang:1.8.3' }
    stages {
        stage('build') {
            steps {
                sh 'env.GOPATH=$PWD'
                sh 'mkdir -p $GOPATH/src/github.com/aristanetworks/go-cvprac'
                dir('$GOPATH/src/github.com/aristanetworks/go-cvprac') {
                    sh 'go version'
                    sh 'echo $GOPATH'
                }
            }
        }
    }
}