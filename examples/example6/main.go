// Copyright (c) 2016-2017, Arista Networks, Inc. All rights reserved.
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions are
// met:
//
//   * Redistributions of source code must retain the above copyright notice,
//   this list of conditions and the following disclaimer.
//
//   * Redistributions in binary form must reproduce the above copyright
//   notice, this list of conditions and the following disclaimer in the
//   documentation and/or other materials provided with the distribution.
//
//   * Neither the name of Arista Networks nor the names of its
//   contributors may be used to endorse or promote products derived from
//   this software without specific prior written permission.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS
// "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT
// LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR
// A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL ARISTA NETWORKS
// BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR
// CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF
// SUBSTITUTE GOODS OR SERVICES; LOSS OF USE, DATA, OR PROFITS; OR
// BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY,
// WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING NEGLIGENCE
// OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN
// IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
//

package main

import (
	"fmt"
	"log"

	"github.com/aristanetworks/go-cvprac/v3/client"
)

func main() {
	TokenCvp := "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJkaWQiOjY5ODI3MzcwMTY0MjUyODM3NDgsImRzbiI6ImdvY3ZwcmFjIiwiZHN0IjoiYWNjb3VudCIsInNpZCI6ImExMzYwZmNkOWQ4N2RlMmExMGRjZDAyYThiN2IwMmNmY2UzYjBiYjBjMDY1ZTI4YzUxODRiNDUxNzM5NTVjZjUtZ0dvWk1Ld0ZtcjJtSkc1SElEUTdWRi1INFc3YUQwTmZaa3Fuc0NpRyJ9.nI9ohekDJqWpV6r1JZRxGs1W1ZV95oJeE0Jx7Ekh8BiBWTwTYf3PChOcc7xRJDxYEQ2CjFPiIs1NebRsUg03uarTFfay8rWM5gPBH3SgaekrCGv0piCaQV5lvrlFS5Ooh0x8T0eLY4NTK1bFuUCmrXBwcI618poBZZ8Byz7ocJCm2C2Y596293mp8mnEEJKaGCCIzG7Yf406AVgzCP9NkZdH33tHNKmdYvqTSwq8d2w1hrko4lOb6_K5rJiNki73iQecFnsv_Si1tJ1uPp2srgr29v_I5aN6-0RXi21oILIlYCCSjfQM442IQP4TJd1LKTOUNVU1T9yEHdA6cLuRLGc_1QR2BV_mMlp62u6aYHXXcpQFAvBTypUXC_QCvua5sKThZhzfGiHVPI93scGbuRsZ6kZV9LpZGRyJqyFK9QEgi6s_OmCFh5JfprNtz6oYCc9B32cgKHBZPnyAtgrkyBg-Qo3tsz5ci5_fN5Ip8XFACawPBiAEQlVXXUIVf6kjGimXCDYZ2OC6ldj45UmNnNtzv0Jm75skeSeT0H7Qnbe5k2waKM1kZrpT46fAFeTptKz6momMvwaFPdDeKMJY6miPNum-LsRfx2TWQtW3GN4FJbLL5rqqkQH1tr2TAyM_7ay4f21TmVzKWDjH2iK_H-5OMR_9FShlDG1j9knn9wU"
	hosts := []string{"10.20.30.186"}
	cvpClient, _ := client.NewCvpClient(
		client.Protocol("https"),
		client.Port(443),
		client.Hosts(hosts...),
		client.Token(TokenCvp),
		client.Debug(true))

	if err := cvpClient.ConnectWithToken(); err != nil {
		log.Fatalf("ERROR: %s", err)
	}

	// verify we have at least one device in inventory
	data, err := cvpClient.API.GetCvpInfo()
	if err != nil {
		log.Fatalf("ERROR: %s", err)
	}
	fmt.Printf("Data: %v\n", data)

	//configletList, err := cvpClient.API.SearchConfiglets("ConfigletName")
	//if err != nil {
	//	log.Fatalf("ERROR: %s", err)
	//}
	//fmt.Printf("Configlets: %v\n", configletList)

}
