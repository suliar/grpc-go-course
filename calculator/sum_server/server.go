package main

import "context"

type server struct {}

func (s *server) DoSum(ctx context.Context, req *sumpb.DoSumRequest) (*sumpb.DoSumResponse, error) {

}