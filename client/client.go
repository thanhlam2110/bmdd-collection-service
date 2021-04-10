package main

import (
	"context"
	"fmt"
	"log"

	collectionpb "github.com/thanhlam/collection-service-new/collectionpb"
	"google.golang.org/grpc"
)

func main() {
	cc, err := grpc.Dial("0.0.0.0:50070", grpc.WithInsecure())
	if err != nil {
		log.Fatal("Err while dial $v", err)
	}
	defer cc.Close()
	client := collectionpb.NewCollectionServiceClient(cc)
	tokenResult := callCheckToken(client)
	if tokenResult == "YES" {
		callSendData(client)
	}
}

//check token
func callCheckToken(c collectionpb.CollectionServiceClient) string {
	log.Println("Call Check Token")
	resp, err := c.CheckTokenAndRole(context.Background(), &collectionpb.CheckRequest{
		CheckToken:   "eyJ6aXAiOiJERUYiLCJhbGciOiJkaXIiLCJlbmMiOiJBMTI4Q0JDLUhTMjU2IiwiY3R5IjoiSldUIiwidHlwIjoiSldUIiwib3JnLmFwZXJlby5jYXMuc2VydmljZXMuUmVnaXN0ZXJlZFNlcnZpY2UiOiIzIn0.._FTsJVQuj76UyQh3KSNptw.PTUnIsDA6gJzwmY9Gc_7tqiC88KoVkS4RbXvvsdLxlpLwNxYNnQuzgff34IwfbVGcDcA7Daw7K_h9Zignoij545YoZQkHHzSon9efzRF2yDXMcBSK5oHDHwUDQKlBs0TBx7lyV4J3mNjae9UY0ZiU-XKpWdL5FQlkinnha8QZm6JYMVfy8Ihxew98Mtml4glyGtbFj2ZgBRNOhEpKVBQ6nmw4Peo5SnEye97AU4xNnFdijrxES2rlsQuRcRf96ilLNCD7XxjaKVGvONg7GXETPMNred--DxQzkznQ0I9OEgqD6aMW_Qns5RF71U5IOYburxvt79uCQIOEqTDSB2ZaraBejyWgwuT5D-Dg0uuiJXXAGQk_9T94UDs-uUNwni6Ji_lf53hBbGvQUvcg-sHq4_xAYh3OdDqn0xdeGnI9_rRFttihiAol0rhCP1G03Z9pYWIpGJJdkE8shluD5VWLWWos2c7DdJ-lm3aY0aJuKuctxzx0w2kyIb_gzOgDD1G5TrVSqoryWkfqMR_OFaWGxLMSIGQZf4RXoym0oUH2onF6cmaKm_pQCiJ4aPY1v3Cya_1mM8kdPt3K1-p5iL69XZ_tWUbEmQIwjBEXYBPIFOzozwnJQ9DCrnkBbPt8CnLfW06VFZuTtxgojELRT8-ulIHQ82ShrfYr_tr3Cfq3Whu2HmISP3eSiVMNmJhp4Ii5BN9-7QQsgoqXWqk3BXLR1aZmjw4Hj31Uu96mPsyklI.VBNJGqyFHytkv7cuP5ofDA",
		CheckThingId: "8fcec30c-0f8f-4cec-9d89-692929cfaff9",
		CheckRole:    "HEART-BEAT",
	})
	if err != nil {
		log.Fatal("Call check token wrong %v", err)
	}
	log.Println("Call check token response", resp.GetCheckRes())
	return resp.GetCheckRes()
}

//send data
func callSendData(c collectionpb.CollectionServiceClient) {
	log.Println("Call send data")
	stream, err := c.SendData(context.Background())
	if err != nil {
		log.Fatal("Call send data err %v", err)
	}
	//listReq := []collectionpb.DataStreamRequest{}
	var listReq []collectionpb.DataStreamRequest
	var i int
	for i = 1; i <= 10; i++ {
		j := collectionpb.DataStreamRequest{
			DataReq: int32(i),
		}
		listReq = append(listReq, j)
	}
	//gửi đi
	for _, req := range listReq {
		err := stream.Send(&req)
		if err != nil {
			log.Fatal("Send data request err %v", err)
		}
	}
	//gửi hết thì nhận response
	resp, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatal("Receive send response err %v", err)
	}
	fmt.Printf("Response %v", resp)

}
