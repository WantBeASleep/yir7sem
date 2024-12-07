package main

import (
	"github.com/WantBeASleep/goooool/brokerlib"

	pb "uzi/internal/generated/broker/produce/uzisplitted"

	"google.golang.org/protobuf/proto"
)

func main() {
	b, _ := brokerlib.NewProducer([]string{"localhost:19092"})
	msg := &pb.UziSplitted{
		UziId: "9d5dc785-720e-4866-bc07-0d21c2ee6583",
		PagesId: []string{
			"a0796f77-b1cf-41df-959b-219b4f859b66",
			"5f9238cd-3a9b-4f5d-93f6-ed6a9fbe8c0d",
			"8a5e9334-de6f-4954-b2c4-a69338f61483",
			"34e2695d-64ff-437a-a233-b39f20029735",
			"5355f85d-0bee-4fc8-a5ef-703fb687b3e9",
			"3657c92a-31e7-46f8-958c-b68c7964f2a9",
			"6041fee1-7ed7-41b6-9af9-4dffba23645c",
			"cfc60c37-1070-4d38-abcb-84e72574d8ec",
			"3ae9f1b4-ca0d-4279-bf34-6fd9e086e809",
			"7cb99129-8d7a-4e0e-a435-82d5144d19ee",
			"55c73afa-3f1a-488c-b3d2-d576503ef877",
			"aa981926-b800-48de-8d77-f616303ab4be",
			"bb138893-8fbe-498d-9101-a9c5b7797965",
			"2af688a3-ae86-4786-abff-fe55c0fda9b0",
			"d6f970fd-6b64-455b-ba1c-cb2146b5a6cf",
			"cad74494-dccb-40e8-9d57-c76c5c294321",
			"d22426d1-bf9b-43cc-bd15-07997ac388fd",
			"7ad686e4-69d1-44d8-9c0f-3d940818bf05",
			"163ee1b7-e900-4bb4-93a0-abdf2ce550de",
			"f29ff62e-ba00-483e-8c46-f08534d3ba96",
			"1e8994ec-7280-4175-b6a0-246feca930b6",
			"2aae8df6-2cdb-47de-804b-5944868a4e5a",
			"71e092f8-cc30-4a33-a2ed-d30a02ef839b",
			"6acc6d20-2ef0-4291-bc04-19ad4d6b0505",
			"46f62c9b-6c8f-465f-8e7f-6ae0bac77126",
			"574cef34-f64c-491a-a574-4e9ef260f84d",
			"f9c382d6-8208-4771-9261-9dc7e8ee4ad1",
			"d9e46401-e95b-4f9b-9776-a10fb023ac62",
			"4968e45a-f26d-4118-898e-92209299e76f",
			"cb667720-5503-4d54-966f-b45e1995d0a4",
			"5bde183c-dc7f-42a4-88e4-46b048becde2",
			"96a04ef0-4390-44c2-8d99-aceae5a16532",
			"55e11de2-b1a2-4ea9-88c2-ca895bcce55c",
			"c3d69024-b9b5-46cd-89d3-6e505ece07d7",
			"dfdfa341-06b4-4b29-b7e5-6052c329cf63",
			"bf74875b-3ab1-47e0-bb16-3f5bca5a0dc4",
			"859995d9-7884-47cd-bca5-e8e9b62ccb0f",
			"0e2fe397-7fde-4ba4-b730-ee57639d08e9",
			"f2dbda34-3f8e-4050-81af-ffb2b82b65fb",
			"09cd0725-5eb2-4f15-adc2-eb674c2e57c6",
		},
	}
	p, err := proto.Marshal(msg)
	if err != nil {
		panic("marshal")
	}
	b.Send("uzisplitted", "", p)
}
