package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"

	"github.com/jmcvetta/napping"
	collectionpb "github.com/thanhlam/collection-service-new/collectionpb"
	"github.com/thanhlam/collection-service-new/service"
	"google.golang.org/grpc"
)

type server struct{}

//Check Token
/*func (*server) CheckToken(ctx context.Context, req *collectionpb.TokenRequest) (*collectionpb.TokenResponse, error) {
	log.Println("Server check token")
	fmt.Println("================")
	fmt.Println(req.GetTokenReq())
	fmt.Println("================")
	token := req.GetTokenReq()
	fmt.Println(token)
	status := CheckToken(token)
	if status == "ACTIVE" {
		resp := &collectionpb.TokenResponse{
			TokenRes: "YES",
		}
		return resp, nil
	} else {
		resp := &collectionpb.TokenResponse{
			TokenRes: "NO",
		}
		return resp, nil
	}
}*/
//Check Token
//Check Token and Role
func (*server) CheckTokenAndRole(ctx context.Context, req *collectionpb.CheckRequest) (*collectionpb.CheckResponse, error) {
	log.Println("Server check token")
	token := req.GetCheckToken()
	thingId := req.GetCheckThingId()
	thingRole := req.GetCheckRole()
	/*fmt.Println("=======token=========")
	fmt.Println(token)
	fmt.Println("=======token=========")
	fmt.Println("=======thingId=========")
	fmt.Println(thingId)
	fmt.Println("=======thingId=========")
	fmt.Println("=======thingRole=========")
	fmt.Println(thingRole)
	fmt.Println("=======thingRole=========")*/
	//check user status
	status := CheckToken(token)
	////check user status
	//check thing role
	arrayRole := GetThingRole(thingId)
	/*fmt.Println("=======thingRole=========")
	fmt.Println(arrayRole)
	fmt.Println("=======thingRole=========")
	fmt.Println("=======resp=========")
	fmt.Println(arrayRole)
	fmt.Println("=======resp=========")*/
	checkRole := contains(arrayRole, thingRole)
	/*fmt.Println("=======checkRole=========")
	fmt.Println(checkRole)
	fmt.Println("=======checkRole=========")*/
	//check thing role
	if status == "ACTIVE" && checkRole == true {
		resp := &collectionpb.CheckResponse{
			CheckRes: "YES",
		}
		return resp, nil
	} else {
		resp := &collectionpb.CheckResponse{
			CheckRes: "NO",
		}
		return resp, nil
	}
	/*if status == "ACTIVE" {
		resp := &collectionpb.CheckResponse{
			CheckRes: "YES",
		}
		return resp, nil
	} else {
		resp := &collectionpb.CheckResponse{
			CheckRes: "NO",
		}
		return resp, nil
	}*/
}

//Check Token and Role
//-------------------------
//send data
func (*server) SendData(stream collectionpb.CollectionService_SendDataServer) error {
	log.Println("Send data...")
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			resp := &collectionpb.DataStreamResponse{
				DataRes: "Done",
			}
			return stream.SendAndClose(resp)
		}
		if err != nil {
			log.Fatal("Err while Recv Average %v", err)
			return err
		}
		log.Println("Receive num %v", req)
		//send to kafka
		//fmt.Printf("%T", req.GetDataReq())
		service.ProducerMessage(req.String(), "testTopic")
		//send to kafka
	}

}

//send data
func main() {
	lis, err := net.Listen("tcp", "0.0.0.0:50070")
	if err != nil {
		log.Fatal("Error while create listen %v", err)
	}
	s := grpc.NewServer()

	collectionpb.RegisterCollectionServiceServer(s, &server{}) // ham nay tu protoc go tao ra

	fmt.Println("Collection data is running...")

	err = s.Serve(lis)

	if err != nil {
		log.Fatal("Error while server %v", err)
	}
}

//check token
func CheckToken(token string) string {
	tokenInputString := `{"token":"` + token + `"}`
	//fmt.Println(tokenInputString)
	//-------------------parse string to Token
	////////var tokenInput TokenInput
	////////json.Unmarshal([]byte(tokenInputString), &tokenInput)
	////////fmt.Println(tokenInput.Token)
	//-------------------parse string to Token
	url := "http://54.254.177.239:1323/api/sso/parseToken"
	s := napping.Session{}
	h := &http.Header{}
	h.Set("X-Custom-Header", "")
	s.Header = h
	var jsonStr = []byte(tokenInputString)

	var data map[string]json.RawMessage
	err := json.Unmarshal(jsonStr, &data)
	if err != nil {
		fmt.Println(err)
	}
	resp, err := s.Post(url, &data, nil, nil)
	if err != nil {
		log.Fatal(err)
	}
	var result map[string]interface{}
	json.Unmarshal([]byte(resp.RawText()), &result)
	if result["code"].(string) == "0" {
		tokenInfo := result["data"].(map[string]interface{})
		status := (tokenInfo["userstatus"]).(string)
		return status
	} else {
		return ""
	}
	//fmt.Println(result["data"])
	//
	/*tokenInfo := result["data"].(map[string]interface{})*/
	/*for key, value := range tokenInfo {
		// Each value is an interface{} type, that is type asserted as a string
		fmt.Println(key, value.(string))
	}*/
	/*status := (tokenInfo["userstatus"]).(string)
	return status*/
}

//get thing role by thingid
type GetThingRoleBody struct {
	Thingid string `json:"thingid"`
}
type RespThingRole struct {
	Data []string
}

func GetThingRole(thingid string) (arrayRes []string) {
	url := "http://54.254.177.239:1323/api/thing/getThingRole"
	var getThingRoleBody GetThingRoleBody
	getThingRoleBody.Thingid = thingid
	//convert struct to string
	e, err := json.Marshal(getThingRoleBody)
	if err != nil {
		fmt.Println(err)
		return arrayRes
	}
	requestTokenStr := string(e)
	//convert struct to string
	resp := service.Post(url, requestTokenStr)
	fmt.Println(resp)
	var respThingRole RespThingRole
	json.Unmarshal([]byte(resp), &respThingRole)
	return respThingRole.Data
}

//check element exits in array
func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}
