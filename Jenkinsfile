#!/usr/bin/env groovy

/**
 * Jenkinsfile
 */
node ('exec') {
  env.REPO      = 'aristanetworks/go-cvprac'
  env.BUILD_DIR = '__build'
  env.GOPATH    = "${WORKSPACE}/${BUILD_DIR}"
  env.SRC_PATH  = "${env.GOPATH}/src/github.com/${REPO}"

    // Install the desired Go version
    def root = tool name: 'Go 1.8', type: 'go'
 
    // Export environment variables pointing to the directory where Go was installed
    //withEnv(["GOROOT=${root}", "PATH+GO=${root}/bin"]) {
    withEnv(["GOPATH=${root}", "PATH+GO=${root}/bin"]) {

  stage ('Checkout') {
          sh "mkdir -p ${env.SRC_PATH}"
          dir(env.SRC_PATH) {
             sh 'go version'
             sh 'printenv'
             checkout scm
          }
  }

  stage ('Install_Requirements') {
          sh """
          make bootstrap || true
          """
          // Stub dummy file
          writeFile file: "api/cvp_node.gcfg", text: "\n[node \"10.81.110.115\"]\nusername = cvpadmin\npassword = cvp123\n"
  }

  stage ('Check_style') {
          sh """
              make check || true
          """
  }

  stage ('Unit Test') {
          sh """
              make unittest || true
          """
  }

  stage ('System Test') {
          sh """
              make systest || true
          """
  }

  stage ('Cleanup') {
          sh 'echo cleanup step'
  }
    }
}
