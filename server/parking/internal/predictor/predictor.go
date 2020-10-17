package predictor

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"

	pb "parking/internal/api"
	"parking/model"
)

type Predictor struct {
	host string
	port int
	Client *pb.CarDetectorClient
}

func NewPredictor(host string, port int) *Predictor {
	predictor := &Predictor{
		host:   host,
		port:   port,
	}

	return predictor
}

func (p *Predictor) CarDetector(image string) (model.Prediction, error) {
	address := fmt.Sprintf("%s:%v", p.host, p.port)

	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
		return model.Prediction{}, err
	}

	defer conn.Close()

	client := pb.NewCarDetectorClient(conn)

	imageGrpc := &pb.Image{Image: image}
	classes, err := client.Predict(context.Background(), imageGrpc)
	if err != nil {
		log.Printf("Error to get predict from best model: %s", err)
		return model.Prediction{}, err
	}

	return model.PredictResponseToPrediction(classes), nil
}
